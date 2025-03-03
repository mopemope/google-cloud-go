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

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
	monitoredrespb "google.golang.org/genproto/googleapis/api/monitoredres"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newMetricClientHook clientHook

// MetricCallOptions contains the retry settings for each method of MetricClient.
type MetricCallOptions struct {
	ListMonitoredResourceDescriptors []gax.CallOption
	GetMonitoredResourceDescriptor   []gax.CallOption
	ListMetricDescriptors            []gax.CallOption
	GetMetricDescriptor              []gax.CallOption
	CreateMetricDescriptor           []gax.CallOption
	DeleteMetricDescriptor           []gax.CallOption
	ListTimeSeries                   []gax.CallOption
	CreateTimeSeries                 []gax.CallOption
	CreateServiceTimeSeries          []gax.CallOption
}

func defaultMetricGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMetricCallOptions() *MetricCallOptions {
	return &MetricCallOptions{
		ListMonitoredResourceDescriptors: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetMonitoredResourceDescriptor: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListMetricDescriptors: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetMetricDescriptor: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateMetricDescriptor: []gax.CallOption{},
		DeleteMetricDescriptor: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListTimeSeries: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateTimeSeries:        []gax.CallOption{},
		CreateServiceTimeSeries: []gax.CallOption{},
	}
}

// internalMetricClient is an interface that defines the methods available from Cloud Monitoring API.
type internalMetricClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListMonitoredResourceDescriptors(context.Context, *monitoringpb.ListMonitoredResourceDescriptorsRequest, ...gax.CallOption) *MonitoredResourceDescriptorIterator
	GetMonitoredResourceDescriptor(context.Context, *monitoringpb.GetMonitoredResourceDescriptorRequest, ...gax.CallOption) (*monitoredrespb.MonitoredResourceDescriptor, error)
	ListMetricDescriptors(context.Context, *monitoringpb.ListMetricDescriptorsRequest, ...gax.CallOption) *MetricDescriptorIterator
	GetMetricDescriptor(context.Context, *monitoringpb.GetMetricDescriptorRequest, ...gax.CallOption) (*metricpb.MetricDescriptor, error)
	CreateMetricDescriptor(context.Context, *monitoringpb.CreateMetricDescriptorRequest, ...gax.CallOption) (*metricpb.MetricDescriptor, error)
	DeleteMetricDescriptor(context.Context, *monitoringpb.DeleteMetricDescriptorRequest, ...gax.CallOption) error
	ListTimeSeries(context.Context, *monitoringpb.ListTimeSeriesRequest, ...gax.CallOption) *TimeSeriesIterator
	CreateTimeSeries(context.Context, *monitoringpb.CreateTimeSeriesRequest, ...gax.CallOption) error
	CreateServiceTimeSeries(context.Context, *monitoringpb.CreateTimeSeriesRequest, ...gax.CallOption) error
}

// MetricClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages metric descriptors, monitored resource descriptors, and
// time series data.
type MetricClient struct {
	// The internal transport-dependent client.
	internalClient internalMetricClient

	// The call options for this service.
	CallOptions *MetricCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MetricClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MetricClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *MetricClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListMonitoredResourceDescriptors lists monitored resource descriptors that match a filter. This method does not require a Workspace.
func (c *MetricClient) ListMonitoredResourceDescriptors(ctx context.Context, req *monitoringpb.ListMonitoredResourceDescriptorsRequest, opts ...gax.CallOption) *MonitoredResourceDescriptorIterator {
	return c.internalClient.ListMonitoredResourceDescriptors(ctx, req, opts...)
}

// GetMonitoredResourceDescriptor gets a single monitored resource descriptor. This method does not require a Workspace.
func (c *MetricClient) GetMonitoredResourceDescriptor(ctx context.Context, req *monitoringpb.GetMonitoredResourceDescriptorRequest, opts ...gax.CallOption) (*monitoredrespb.MonitoredResourceDescriptor, error) {
	return c.internalClient.GetMonitoredResourceDescriptor(ctx, req, opts...)
}

// ListMetricDescriptors lists metric descriptors that match a filter. This method does not require a Workspace.
func (c *MetricClient) ListMetricDescriptors(ctx context.Context, req *monitoringpb.ListMetricDescriptorsRequest, opts ...gax.CallOption) *MetricDescriptorIterator {
	return c.internalClient.ListMetricDescriptors(ctx, req, opts...)
}

// GetMetricDescriptor gets a single metric descriptor. This method does not require a Workspace.
func (c *MetricClient) GetMetricDescriptor(ctx context.Context, req *monitoringpb.GetMetricDescriptorRequest, opts ...gax.CallOption) (*metricpb.MetricDescriptor, error) {
	return c.internalClient.GetMetricDescriptor(ctx, req, opts...)
}

// CreateMetricDescriptor creates a new metric descriptor.
// The creation is executed asynchronously and callers may check the returned
// operation to track its progress.
// User-created metric descriptors define
// custom metrics (at https://cloud.google.com/monitoring/custom-metrics).
func (c *MetricClient) CreateMetricDescriptor(ctx context.Context, req *monitoringpb.CreateMetricDescriptorRequest, opts ...gax.CallOption) (*metricpb.MetricDescriptor, error) {
	return c.internalClient.CreateMetricDescriptor(ctx, req, opts...)
}

// DeleteMetricDescriptor deletes a metric descriptor. Only user-created
// custom metrics (at https://cloud.google.com/monitoring/custom-metrics) can be
// deleted.
func (c *MetricClient) DeleteMetricDescriptor(ctx context.Context, req *monitoringpb.DeleteMetricDescriptorRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteMetricDescriptor(ctx, req, opts...)
}

// ListTimeSeries lists time series that match a filter. This method does not require a Workspace.
func (c *MetricClient) ListTimeSeries(ctx context.Context, req *monitoringpb.ListTimeSeriesRequest, opts ...gax.CallOption) *TimeSeriesIterator {
	return c.internalClient.ListTimeSeries(ctx, req, opts...)
}

// CreateTimeSeries creates or adds data to one or more time series.
// The response is empty if all time series in the request were written.
// If any time series could not be written, a corresponding failure message is
// included in the error response.
func (c *MetricClient) CreateTimeSeries(ctx context.Context, req *monitoringpb.CreateTimeSeriesRequest, opts ...gax.CallOption) error {
	return c.internalClient.CreateTimeSeries(ctx, req, opts...)
}

// CreateServiceTimeSeries creates or adds data to one or more service time series. A service time
// series is a time series for a metric from a Google Cloud service. The
// response is empty if all time series in the request were written. If any
// time series could not be written, a corresponding failure message is
// included in the error response. This endpoint rejects writes to
// user-defined metrics.
// This method is only for use by Google Cloud services. Use
// projects.timeSeries.create
// instead.
func (c *MetricClient) CreateServiceTimeSeries(ctx context.Context, req *monitoringpb.CreateTimeSeriesRequest, opts ...gax.CallOption) error {
	return c.internalClient.CreateServiceTimeSeries(ctx, req, opts...)
}

// metricGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type metricGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing MetricClient
	CallOptions **MetricCallOptions

	// The gRPC API client.
	metricClient monitoringpb.MetricServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMetricClient creates a new metric service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages metric descriptors, monitored resource descriptors, and
// time series data.
func NewMetricClient(ctx context.Context, opts ...option.ClientOption) (*MetricClient, error) {
	clientOpts := defaultMetricGRPCClientOptions()
	if newMetricClientHook != nil {
		hookOpts, err := newMetricClientHook(ctx, clientHookParams{})
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
	client := MetricClient{CallOptions: defaultMetricCallOptions()}

	c := &metricGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		metricClient:     monitoringpb.NewMetricServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *metricGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *metricGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *metricGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *metricGRPCClient) ListMonitoredResourceDescriptors(ctx context.Context, req *monitoringpb.ListMonitoredResourceDescriptorsRequest, opts ...gax.CallOption) *MonitoredResourceDescriptorIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListMonitoredResourceDescriptors[0:len((*c.CallOptions).ListMonitoredResourceDescriptors):len((*c.CallOptions).ListMonitoredResourceDescriptors)], opts...)
	it := &MonitoredResourceDescriptorIterator{}
	req = proto.Clone(req).(*monitoringpb.ListMonitoredResourceDescriptorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoredrespb.MonitoredResourceDescriptor, string, error) {
		resp := &monitoringpb.ListMonitoredResourceDescriptorsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.metricClient.ListMonitoredResourceDescriptors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResourceDescriptors(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *metricGRPCClient) GetMonitoredResourceDescriptor(ctx context.Context, req *monitoringpb.GetMonitoredResourceDescriptorRequest, opts ...gax.CallOption) (*monitoredrespb.MonitoredResourceDescriptor, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMonitoredResourceDescriptor[0:len((*c.CallOptions).GetMonitoredResourceDescriptor):len((*c.CallOptions).GetMonitoredResourceDescriptor)], opts...)
	var resp *monitoredrespb.MonitoredResourceDescriptor
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricClient.GetMonitoredResourceDescriptor(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricGRPCClient) ListMetricDescriptors(ctx context.Context, req *monitoringpb.ListMetricDescriptorsRequest, opts ...gax.CallOption) *MetricDescriptorIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListMetricDescriptors[0:len((*c.CallOptions).ListMetricDescriptors):len((*c.CallOptions).ListMetricDescriptors)], opts...)
	it := &MetricDescriptorIterator{}
	req = proto.Clone(req).(*monitoringpb.ListMetricDescriptorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*metricpb.MetricDescriptor, string, error) {
		resp := &monitoringpb.ListMetricDescriptorsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.metricClient.ListMetricDescriptors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMetricDescriptors(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *metricGRPCClient) GetMetricDescriptor(ctx context.Context, req *monitoringpb.GetMetricDescriptorRequest, opts ...gax.CallOption) (*metricpb.MetricDescriptor, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMetricDescriptor[0:len((*c.CallOptions).GetMetricDescriptor):len((*c.CallOptions).GetMetricDescriptor)], opts...)
	var resp *metricpb.MetricDescriptor
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricClient.GetMetricDescriptor(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricGRPCClient) CreateMetricDescriptor(ctx context.Context, req *monitoringpb.CreateMetricDescriptorRequest, opts ...gax.CallOption) (*metricpb.MetricDescriptor, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 12000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateMetricDescriptor[0:len((*c.CallOptions).CreateMetricDescriptor):len((*c.CallOptions).CreateMetricDescriptor)], opts...)
	var resp *metricpb.MetricDescriptor
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.metricClient.CreateMetricDescriptor(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *metricGRPCClient) DeleteMetricDescriptor(ctx context.Context, req *monitoringpb.DeleteMetricDescriptorRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteMetricDescriptor[0:len((*c.CallOptions).DeleteMetricDescriptor):len((*c.CallOptions).DeleteMetricDescriptor)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.metricClient.DeleteMetricDescriptor(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metricGRPCClient) ListTimeSeries(ctx context.Context, req *monitoringpb.ListTimeSeriesRequest, opts ...gax.CallOption) *TimeSeriesIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTimeSeries[0:len((*c.CallOptions).ListTimeSeries):len((*c.CallOptions).ListTimeSeries)], opts...)
	it := &TimeSeriesIterator{}
	req = proto.Clone(req).(*monitoringpb.ListTimeSeriesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.TimeSeries, string, error) {
		resp := &monitoringpb.ListTimeSeriesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.metricClient.ListTimeSeries(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTimeSeries(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *metricGRPCClient) CreateTimeSeries(ctx context.Context, req *monitoringpb.CreateTimeSeriesRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 12000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateTimeSeries[0:len((*c.CallOptions).CreateTimeSeries):len((*c.CallOptions).CreateTimeSeries)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.metricClient.CreateTimeSeries(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *metricGRPCClient) CreateServiceTimeSeries(ctx context.Context, req *monitoringpb.CreateTimeSeriesRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateServiceTimeSeries[0:len((*c.CallOptions).CreateServiceTimeSeries):len((*c.CallOptions).CreateServiceTimeSeries)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.metricClient.CreateServiceTimeSeries(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// MetricDescriptorIterator manages a stream of *metricpb.MetricDescriptor.
type MetricDescriptorIterator struct {
	items    []*metricpb.MetricDescriptor
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*metricpb.MetricDescriptor, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *MetricDescriptorIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *MetricDescriptorIterator) Next() (*metricpb.MetricDescriptor, error) {
	var item *metricpb.MetricDescriptor
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *MetricDescriptorIterator) bufLen() int {
	return len(it.items)
}

func (it *MetricDescriptorIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// MonitoredResourceDescriptorIterator manages a stream of *monitoredrespb.MonitoredResourceDescriptor.
type MonitoredResourceDescriptorIterator struct {
	items    []*monitoredrespb.MonitoredResourceDescriptor
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoredrespb.MonitoredResourceDescriptor, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *MonitoredResourceDescriptorIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *MonitoredResourceDescriptorIterator) Next() (*monitoredrespb.MonitoredResourceDescriptor, error) {
	var item *monitoredrespb.MonitoredResourceDescriptor
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *MonitoredResourceDescriptorIterator) bufLen() int {
	return len(it.items)
}

func (it *MonitoredResourceDescriptorIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TimeSeriesIterator manages a stream of *monitoringpb.TimeSeries.
type TimeSeriesIterator struct {
	items    []*monitoringpb.TimeSeries
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.TimeSeries, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TimeSeriesIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TimeSeriesIterator) Next() (*monitoringpb.TimeSeries, error) {
	var item *monitoringpb.TimeSeries
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TimeSeriesIterator) bufLen() int {
	return len(it.items)
}

func (it *TimeSeriesIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
