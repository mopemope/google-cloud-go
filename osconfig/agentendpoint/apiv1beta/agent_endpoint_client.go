// Copyright 2022 Google LLC
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
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package agentendpoint

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	agentendpointpb "cloud.google.com/go/osconfig/agentendpoint/apiv1beta/agentendpointpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ReceiveTaskNotification    []gax.CallOption
	StartNextTask              []gax.CallOption
	ReportTaskProgress         []gax.CallOption
	ReportTaskComplete         []gax.CallOption
	LookupEffectiveGuestPolicy []gax.CallOption
	RegisterAgent              []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("osconfig.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("osconfig.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://osconfig.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ReceiveTaskNotification: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Canceled,
					codes.Aborted,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		StartNextTask: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ReportTaskProgress: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ReportTaskComplete: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		LookupEffectiveGuestPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RegisterAgent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		ReceiveTaskNotification: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					499,
					http.StatusConflict,
					http.StatusInternalServerError,
					http.StatusServiceUnavailable)
			}),
		},
		StartNextTask: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ReportTaskProgress: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ReportTaskComplete: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		LookupEffectiveGuestPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		RegisterAgent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from OS Config API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ReceiveTaskNotification(context.Context, *agentendpointpb.ReceiveTaskNotificationRequest, ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error)
	StartNextTask(context.Context, *agentendpointpb.StartNextTaskRequest, ...gax.CallOption) (*agentendpointpb.StartNextTaskResponse, error)
	ReportTaskProgress(context.Context, *agentendpointpb.ReportTaskProgressRequest, ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error)
	ReportTaskComplete(context.Context, *agentendpointpb.ReportTaskCompleteRequest, ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error)
	LookupEffectiveGuestPolicy(context.Context, *agentendpointpb.LookupEffectiveGuestPolicyRequest, ...gax.CallOption) (*agentendpointpb.EffectiveGuestPolicy, error)
	RegisterAgent(context.Context, *agentendpointpb.RegisterAgentRequest, ...gax.CallOption) (*agentendpointpb.RegisterAgentResponse, error)
}

// Client is a client for interacting with OS Config API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// OS Config agent endpoint API.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ReceiveTaskNotification stream established by client to receive Task notifications.
func (c *Client) ReceiveTaskNotification(ctx context.Context, req *agentendpointpb.ReceiveTaskNotificationRequest, opts ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error) {
	return c.internalClient.ReceiveTaskNotification(ctx, req, opts...)
}

// StartNextTask signals the start of a task execution and returns the task info.
func (c *Client) StartNextTask(ctx context.Context, req *agentendpointpb.StartNextTaskRequest, opts ...gax.CallOption) (*agentendpointpb.StartNextTaskResponse, error) {
	return c.internalClient.StartNextTask(ctx, req, opts...)
}

// ReportTaskProgress signals an intermediary progress checkpoint in task execution.
func (c *Client) ReportTaskProgress(ctx context.Context, req *agentendpointpb.ReportTaskProgressRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error) {
	return c.internalClient.ReportTaskProgress(ctx, req, opts...)
}

// ReportTaskComplete signals that the task execution is complete and optionally returns the next
// task.
func (c *Client) ReportTaskComplete(ctx context.Context, req *agentendpointpb.ReportTaskCompleteRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error) {
	return c.internalClient.ReportTaskComplete(ctx, req, opts...)
}

// LookupEffectiveGuestPolicy lookup the effective guest policy that applies to a VM instance. This
// lookup merges all policies that are assigned to the instance ancestry.
func (c *Client) LookupEffectiveGuestPolicy(ctx context.Context, req *agentendpointpb.LookupEffectiveGuestPolicyRequest, opts ...gax.CallOption) (*agentendpointpb.EffectiveGuestPolicy, error) {
	return c.internalClient.LookupEffectiveGuestPolicy(ctx, req, opts...)
}

// RegisterAgent registers the agent running on the VM.
func (c *Client) RegisterAgent(ctx context.Context, req *agentendpointpb.RegisterAgentRequest, opts ...gax.CallOption) (*agentendpointpb.RegisterAgentResponse, error) {
	return c.internalClient.RegisterAgent(ctx, req, opts...)
}

// gRPCClient is a client for interacting with OS Config API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client agentendpointpb.AgentEndpointServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new agent endpoint service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// OS Config agent endpoint API.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           agentendpointpb.NewAgentEndpointServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions
}

// NewRESTClient creates a new agent endpoint service rest client.
//
// OS Config agent endpoint API.
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://osconfig.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://osconfig.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://osconfig.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) ReceiveTaskNotification(ctx context.Context, req *agentendpointpb.ReceiveTaskNotificationRequest, opts ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReceiveTaskNotification(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) StartNextTask(ctx context.Context, req *agentendpointpb.StartNextTaskRequest, opts ...gax.CallOption) (*agentendpointpb.StartNextTaskResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).StartNextTask[0:len((*c.CallOptions).StartNextTask):len((*c.CallOptions).StartNextTask)], opts...)
	var resp *agentendpointpb.StartNextTaskResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.StartNextTask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ReportTaskProgress(ctx context.Context, req *agentendpointpb.ReportTaskProgressRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ReportTaskProgress[0:len((*c.CallOptions).ReportTaskProgress):len((*c.CallOptions).ReportTaskProgress)], opts...)
	var resp *agentendpointpb.ReportTaskProgressResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportTaskProgress(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ReportTaskComplete(ctx context.Context, req *agentendpointpb.ReportTaskCompleteRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ReportTaskComplete[0:len((*c.CallOptions).ReportTaskComplete):len((*c.CallOptions).ReportTaskComplete)], opts...)
	var resp *agentendpointpb.ReportTaskCompleteResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportTaskComplete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) LookupEffectiveGuestPolicy(ctx context.Context, req *agentendpointpb.LookupEffectiveGuestPolicyRequest, opts ...gax.CallOption) (*agentendpointpb.EffectiveGuestPolicy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).LookupEffectiveGuestPolicy[0:len((*c.CallOptions).LookupEffectiveGuestPolicy):len((*c.CallOptions).LookupEffectiveGuestPolicy)], opts...)
	var resp *agentendpointpb.EffectiveGuestPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.LookupEffectiveGuestPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) RegisterAgent(ctx context.Context, req *agentendpointpb.RegisterAgentRequest, opts ...gax.CallOption) (*agentendpointpb.RegisterAgentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).RegisterAgent[0:len((*c.CallOptions).RegisterAgent):len((*c.CallOptions).RegisterAgent)], opts...)
	var resp *agentendpointpb.RegisterAgentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.RegisterAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ReceiveTaskNotification stream established by client to receive Task notifications.
func (c *restClient) ReceiveTaskNotification(ctx context.Context, req *agentendpointpb.ReceiveTaskNotificationRequest, opts ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	params.Add("agentVersion", fmt.Sprintf("%v", req.GetAgentVersion()))
	params.Add("instanceIdToken", fmt.Sprintf("%v", req.GetInstanceIdToken()))

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	var streamClient *receiveTaskNotificationRESTClient
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		streamClient = &receiveTaskNotificationRESTClient{
			ctx:    ctx,
			md:     metadata.MD(httpRsp.Header),
			stream: gax.NewProtoJSONStreamReader(httpRsp.Body, (&agentendpointpb.ReceiveTaskNotificationResponse{}).ProtoReflect().Type()),
		}
		return nil
	}, opts...)

	return streamClient, e
}

// receiveTaskNotificationRESTClient is the stream client used to consume the server stream created by
// the REST implementation of ReceiveTaskNotification.
type receiveTaskNotificationRESTClient struct {
	ctx    context.Context
	md     metadata.MD
	stream *gax.ProtoJSONStream
}

func (c *receiveTaskNotificationRESTClient) Recv() (*agentendpointpb.ReceiveTaskNotificationResponse, error) {
	if err := c.ctx.Err(); err != nil {
		defer c.stream.Close()
		return nil, err
	}
	msg, err := c.stream.Recv()
	if err != nil {
		defer c.stream.Close()
		return nil, err
	}
	res := msg.(*agentendpointpb.ReceiveTaskNotificationResponse)
	return res, nil
}

func (c *receiveTaskNotificationRESTClient) Header() (metadata.MD, error) {
	return c.md, nil
}

func (c *receiveTaskNotificationRESTClient) Trailer() metadata.MD {
	return c.md
}

func (c *receiveTaskNotificationRESTClient) CloseSend() error {
	// This is a no-op to fulfill the interface.
	return fmt.Errorf("this method is not implemented for a server-stream")
}

func (c *receiveTaskNotificationRESTClient) Context() context.Context {
	return c.ctx
}

func (c *receiveTaskNotificationRESTClient) SendMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return fmt.Errorf("this method is not implemented for a server-stream")
}

func (c *receiveTaskNotificationRESTClient) RecvMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return fmt.Errorf("this method is not implemented, use Recv")
}

// StartNextTask signals the start of a task execution and returns the task info.
func (c *restClient) StartNextTask(ctx context.Context, req *agentendpointpb.StartNextTaskRequest, opts ...gax.CallOption) (*agentendpointpb.StartNextTaskResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	params.Add("instanceIdToken", fmt.Sprintf("%v", req.GetInstanceIdToken()))

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).StartNextTask[0:len((*c.CallOptions).StartNextTask):len((*c.CallOptions).StartNextTask)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &agentendpointpb.StartNextTaskResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ReportTaskProgress signals an intermediary progress checkpoint in task execution.
func (c *restClient) ReportTaskProgress(ctx context.Context, req *agentendpointpb.ReportTaskProgressRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	params.Add("applyPatchesTaskProgress.state", fmt.Sprintf("%v", req.GetApplyPatchesTaskProgress().GetState()))
	params.Add("execStepTaskProgress.state", fmt.Sprintf("%v", req.GetExecStepTaskProgress().GetState()))
	params.Add("instanceIdToken", fmt.Sprintf("%v", req.GetInstanceIdToken()))
	params.Add("taskId", fmt.Sprintf("%v", req.GetTaskId()))
	params.Add("taskType", fmt.Sprintf("%v", req.GetTaskType()))

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).ReportTaskProgress[0:len((*c.CallOptions).ReportTaskProgress):len((*c.CallOptions).ReportTaskProgress)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &agentendpointpb.ReportTaskProgressResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ReportTaskComplete signals that the task execution is complete and optionally returns the next
// task.
func (c *restClient) ReportTaskComplete(ctx context.Context, req *agentendpointpb.ReportTaskCompleteRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	params.Add("applyPatchesTaskOutput.state", fmt.Sprintf("%v", req.GetApplyPatchesTaskOutput().GetState()))
	if req.GetErrorMessage() != "" {
		params.Add("errorMessage", fmt.Sprintf("%v", req.GetErrorMessage()))
	}
	params.Add("execStepTaskOutput.exitCode", fmt.Sprintf("%v", req.GetExecStepTaskOutput().GetExitCode()))
	params.Add("execStepTaskOutput.state", fmt.Sprintf("%v", req.GetExecStepTaskOutput().GetState()))
	params.Add("instanceIdToken", fmt.Sprintf("%v", req.GetInstanceIdToken()))
	params.Add("taskId", fmt.Sprintf("%v", req.GetTaskId()))
	params.Add("taskType", fmt.Sprintf("%v", req.GetTaskType()))

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).ReportTaskComplete[0:len((*c.CallOptions).ReportTaskComplete):len((*c.CallOptions).ReportTaskComplete)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &agentendpointpb.ReportTaskCompleteResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// LookupEffectiveGuestPolicy lookup the effective guest policy that applies to a VM instance. This
// lookup merges all policies that are assigned to the instance ancestry.
func (c *restClient) LookupEffectiveGuestPolicy(ctx context.Context, req *agentendpointpb.LookupEffectiveGuestPolicyRequest, opts ...gax.CallOption) (*agentendpointpb.EffectiveGuestPolicy, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	params.Add("instanceIdToken", fmt.Sprintf("%v", req.GetInstanceIdToken()))
	if req.GetOsArchitecture() != "" {
		params.Add("osArchitecture", fmt.Sprintf("%v", req.GetOsArchitecture()))
	}
	if req.GetOsShortName() != "" {
		params.Add("osShortName", fmt.Sprintf("%v", req.GetOsShortName()))
	}
	if req.GetOsVersion() != "" {
		params.Add("osVersion", fmt.Sprintf("%v", req.GetOsVersion()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).LookupEffectiveGuestPolicy[0:len((*c.CallOptions).LookupEffectiveGuestPolicy):len((*c.CallOptions).LookupEffectiveGuestPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &agentendpointpb.EffectiveGuestPolicy{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// RegisterAgent registers the agent running on the VM.
func (c *restClient) RegisterAgent(ctx context.Context, req *agentendpointpb.RegisterAgentRequest, opts ...gax.CallOption) (*agentendpointpb.RegisterAgentResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	params.Add("agentVersion", fmt.Sprintf("%v", req.GetAgentVersion()))
	params.Add("instanceIdToken", fmt.Sprintf("%v", req.GetInstanceIdToken()))
	if req.GetOsArchitecture() != "" {
		params.Add("osArchitecture", fmt.Sprintf("%v", req.GetOsArchitecture()))
	}
	if req.GetOsLongName() != "" {
		params.Add("osLongName", fmt.Sprintf("%v", req.GetOsLongName()))
	}
	if req.GetOsShortName() != "" {
		params.Add("osShortName", fmt.Sprintf("%v", req.GetOsShortName()))
	}
	if req.GetOsVersion() != "" {
		params.Add("osVersion", fmt.Sprintf("%v", req.GetOsVersion()))
	}
	if req.GetSupportedCapabilities() != nil {
		params.Add("supportedCapabilities", fmt.Sprintf("%v", req.GetSupportedCapabilities()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).RegisterAgent[0:len((*c.CallOptions).RegisterAgent):len((*c.CallOptions).RegisterAgent)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &agentendpointpb.RegisterAgentResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
