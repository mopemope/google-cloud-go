// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and

package pscompat

import (
	"context"
	"errors"
	"sync"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsublite/internal/wire"
	"google.golang.org/api/option"

	ipubsub "cloud.google.com/go/internal/pubsub"
)

var (
	errNackCalled       = errors.New("pubsublite: subscriber client does not support nack. See NackHandler for how to customize nack handling")
	errDuplicateReceive = errors.New("pubsublite: receive is already in progress for this subscriber client")
	errMessageIDSet     = errors.New("pubsublite: pubsub.Message.ID must not be set")
)

// handleNack is the default NackHandler implementation.
func handleNack(_ *pubsub.Message) error {
	return errNackCalled
}

// pslAckHandler is the AckHandler for Pub/Sub Lite.
type pslAckHandler struct {
	ackh        wire.AckConsumer
	msg         *pubsub.Message
	partition   int
	nackh       NackHandler
	subInstance *subscriberInstance
}

func (ah *pslAckHandler) OnAck() {
	if ah.subInstance == nil {
		return
	}

	ah.ackh.Ack()
	ah.subInstance = nil
}

func (ah *pslAckHandler) OnNack() {
	if ah.subInstance == nil {
		return
	}

	// Ignore nacks for partitions that have been assigned away.
	if !ah.subInstance.wireSub.PartitionActive(ah.partition) {
		return
	}

	err := ah.nackh(ah.msg)
	if err != nil {
		// If the NackHandler returns an error, shut down the subscriber client.
		ah.subInstance.Terminate(err)
	} else {
		// If the NackHandler succeeds, just ack the message.
		ah.ackh.Ack()
	}
	ah.subInstance = nil
}

// OnAckWithResult is required implementation for the ack handler
// for Cloud Pub/Sub's exactly once delivery feature. This will
// ack the message and return an AckResult that always resolves to success.
func (ah *pslAckHandler) OnAckWithResult() *ipubsub.AckResult {
	ah.OnAck()
	ar := ipubsub.NewAckResult()
	ipubsub.SetAckResult(ar, ipubsub.AcknowledgeStatusSuccess, nil)
	return ar
}

// OnNackWithResult is required implementation for the ack handler
// for Cloud Pub/Sub's exactly once delivery feature. This will
// nack the message and return an AckResult that always resolves to success.
func (ah *pslAckHandler) OnNackWithResult() *ipubsub.AckResult {
	ah.OnNack()
	ar := ipubsub.NewAckResult()
	ipubsub.SetAckResult(ar, ipubsub.AcknowledgeStatusSuccess, nil)
	return ar
}

// wireSubscriberFactory is a factory for creating wire subscribers, which can
// be overridden with a mock in unit tests.
type wireSubscriberFactory interface {
	New(context.Context, wire.MessageReceiverFunc, wire.ReassignmentHandlerFunc) (wire.Subscriber, error)
}

type wireSubscriberFactoryImpl struct {
	settings     wire.ReceiveSettings
	region       string
	subscription wire.SubscriptionPath
	options      []option.ClientOption
}

func (f *wireSubscriberFactoryImpl) New(ctx context.Context, receiver wire.MessageReceiverFunc, onReassignment wire.ReassignmentHandlerFunc) (wire.Subscriber, error) {
	return wire.NewSubscriber(ctx, f.settings, receiver, onReassignment, f.region, f.subscription.String(), f.options...)
}

type messageReceiverFunc = func(context.Context, *pubsub.Message)

// subscriberInstance wraps an instance of a wire.Subscriber. A new instance is
// created for each invocation of SubscriberClient.Receive().
type subscriberInstance struct {
	settings        ReceiveSettings
	receiver        messageReceiverFunc
	recvCtx         context.Context    // Context passed to the receiver
	recvCancel      context.CancelFunc // Corresponding cancel func for recvCtx
	wireSub         wire.Subscriber
	activeReceivers sync.WaitGroup

	// Fields below must be guarded with mu.
	mu  sync.Mutex
	err error
}

func newSubscriberInstance(recvCtx, clientCtx context.Context, factory wireSubscriberFactory, settings ReceiveSettings, receiver messageReceiverFunc) (*subscriberInstance, error) {
	recvCtx, recvCancel := context.WithCancel(recvCtx)
	subInstance := &subscriberInstance{
		settings:   settings,
		recvCtx:    recvCtx,
		recvCancel: recvCancel,
		receiver:   receiver,
	}

	// Note: The context from Receive (recvCtx) should not be used, as when it is
	// cancelled, the gRPC streams will be disconnected and the subscriber will
	// not be able to process acks and commit the final cursor offset. Use the
	// context from NewSubscriberClient (clientCtx) instead.
	wireSub, err := factory.New(clientCtx, subInstance.onMessage, subInstance.onReassignment)
	if err != nil {
		return nil, err
	}

	subInstance.wireSub = wireSub
	if subInstance.settings.MessageTransformer == nil {
		subInstance.settings.MessageTransformer = transformReceivedMessage
	}
	if subInstance.settings.NackHandler == nil {
		subInstance.settings.NackHandler = handleNack
	}
	return subInstance, nil
}

func (si *subscriberInstance) onReassignment(before, after wire.PartitionSet) error {
	if si.settings.ReassignmentHandler != nil {
		return si.settings.ReassignmentHandler(before.SortedInts(), after.SortedInts())
	}
	return nil
}

func (si *subscriberInstance) transformMessage(in *wire.ReceivedMessage, out *pubsub.Message) error {
	if err := si.settings.MessageTransformer(in.Msg, out); err != nil {
		return err
	}
	if len(out.ID) > 0 {
		return errMessageIDSet
	}
	metadata := &MessageMetadata{Partition: in.Partition, Offset: in.Msg.GetCursor().GetOffset()}
	out.ID = metadata.String()
	return nil
}

func (si *subscriberInstance) onMessage(msg *wire.ReceivedMessage) {
	pslAckh := &pslAckHandler{
		ackh:        msg.Ack,
		nackh:       si.settings.NackHandler,
		partition:   msg.Partition,
		subInstance: si,
	}
	psMsg := ipubsub.NewMessage(pslAckh)
	pslAckh.msg = psMsg
	if err := si.transformMessage(msg, psMsg); err != nil {
		si.Terminate(err)
		return
	}

	si.activeReceivers.Add(1)
	si.receiver(si.recvCtx, psMsg)
	si.activeReceivers.Done()
}

// shutdown starts shutting down the subscriber client. The wire subscriber can
// optionally wait for all outstanding messages to be acked/nacked.
func (si *subscriberInstance) shutdown(waitForAcks bool, err error) {
	si.mu.Lock()
	defer si.mu.Unlock()

	// Don't clobber original error.
	if si.err == nil {
		si.err = err
	}

	// Cancel recvCtx to notify message receiver funcs of shutdown.
	si.recvCancel()

	// Either wait for acks, or terminate quickly upon fatal error.
	if waitForAcks {
		si.wireSub.Stop()
	} else {
		si.wireSub.Terminate()
	}
}

// Terminate shuts down the subscriber client without waiting for outstanding
// messages to be acked/nacked.
func (si *subscriberInstance) Terminate(err error) {
	si.shutdown(false, err)
}

// Wait for the subscriber to stop, or the context is done, whichever occurs
// first.
func (si *subscriberInstance) Wait(ctx context.Context) error {
	si.wireSub.Start()
	if err := si.wireSub.WaitStarted(); err != nil {
		return err
	}

	// Start a goroutine to monitor when the context is done.
	subscriberStopped := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			si.shutdown(true, nil)
		case <-subscriberStopped:
		}
	}()
	err := si.wireSub.WaitStopped()

	// End goroutine above if the wire subscriber terminated due to fatal error
	// and ctx is not done.
	close(subscriberStopped)
	// And also wait for all the receivers to finish.
	si.recvCancel()
	si.activeReceivers.Wait()

	si.mu.Lock()
	defer si.mu.Unlock()

	if si.err != nil {
		return si.err
	}
	return err
}

// SubscriberClient is a Pub/Sub Lite client to receive messages for a given
// subscription.
//
// See https://cloud.google.com/pubsub/lite/docs/subscribing for more
// information about receiving messages.
type SubscriberClient struct {
	clientCtx      context.Context
	settings       ReceiveSettings
	wireSubFactory wireSubscriberFactory

	// Fields below must be guarded with mu.
	mu            sync.Mutex
	receiveActive bool
}

// NewSubscriberClient creates a new Pub/Sub Lite client to receive messages for
// a given subscription, using DefaultReceiveSettings. A valid subscription path
// has the format:
// "projects/PROJECT_ID/locations/LOCATION/subscriptions/SUBSCRIPTION_ID".
func NewSubscriberClient(ctx context.Context, subscription string, opts ...option.ClientOption) (*SubscriberClient, error) {
	return NewSubscriberClientWithSettings(ctx, subscription, DefaultReceiveSettings, opts...)
}

// NewSubscriberClientWithSettings creates a new Pub/Sub Lite client to receive
// messages for a given subscription, using the specified ReceiveSettings. A
// valid subscription path has the format:
// "projects/PROJECT_ID/locations/LOCATION/subscriptions/SUBSCRIPTION_ID".
func NewSubscriberClientWithSettings(ctx context.Context, subscription string, settings ReceiveSettings, opts ...option.ClientOption) (*SubscriberClient, error) {
	subscriptionPath, err := wire.ParseSubscriptionPath(subscription)
	if err != nil {
		return nil, err
	}
	region, err := wire.LocationToRegion(subscriptionPath.Location)
	if err != nil {
		return nil, err
	}
	factory := &wireSubscriberFactoryImpl{
		settings:     settings.toWireSettings(),
		region:       region,
		subscription: subscriptionPath,
		options:      opts,
	}
	subClient := &SubscriberClient{
		clientCtx:      ctx,
		settings:       settings,
		wireSubFactory: factory,
	}
	return subClient, nil
}

// Receive calls f with the messages from the subscription. It blocks until ctx
// is done, or the service returns a non-retryable error.
//
// The standard way to terminate a Receive is to cancel its context:
//
//	cctx, cancel := context.WithCancel(ctx)
//	err := sub.Receive(cctx, callback)
//	// Call cancel from callback, or another goroutine.
//
// If there is a fatal service error, Receive returns that error after all of
// the outstanding calls to f have returned. If ctx is done, Receive returns nil
// after all of the outstanding calls to f have returned and all messages have
// been acknowledged. The context passed to f will be canceled when ctx is Done
// or there is a fatal service error.
//
// Receive calls f concurrently from multiple goroutines if the SubscriberClient
// is connected to multiple partitions. Only one call from any connected
// partition will be outstanding at a time, and blocking in the receiver
// callback f will block the delivery of subsequent messages for the partition.
//
// All messages received by f must be ACKed or NACKed. Failure to do so can
// prevent Receive from returning.
//
// Each SubscriberClient may have only one invocation of Receive active at a
// time.
func (s *SubscriberClient) Receive(ctx context.Context, f func(context.Context, *pubsub.Message)) error {
	if err := s.setReceiveActive(true); err != nil {
		return err
	}
	defer s.setReceiveActive(false)

	// Initialize a subscriber instance.
	subInstance, err := newSubscriberInstance(ctx, s.clientCtx, s.wireSubFactory, s.settings, f)
	if err != nil {
		return err
	}

	// Wait for the subscriber without mutex held. Overlapping Receive invocations
	// will return an error.
	return subInstance.Wait(ctx)
}

func (s *SubscriberClient) setReceiveActive(active bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if active && s.receiveActive {
		return errDuplicateReceive
	}
	s.receiveActive = active
	return nil
}
