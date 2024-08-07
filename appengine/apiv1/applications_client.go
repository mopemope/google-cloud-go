// Copyright 2024 Google LLC
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

package appengine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newApplicationsClientHook clientHook

// ApplicationsCallOptions contains the retry settings for each method of ApplicationsClient.
type ApplicationsCallOptions struct {
	GetApplication    []gax.CallOption
	CreateApplication []gax.CallOption
	UpdateApplication []gax.CallOption
	RepairApplication []gax.CallOption
}

func defaultApplicationsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("appengine.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultApplicationsCallOptions() *ApplicationsCallOptions {
	return &ApplicationsCallOptions{
		GetApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		RepairApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultApplicationsRESTCallOptions() *ApplicationsCallOptions {
	return &ApplicationsCallOptions{
		GetApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		RepairApplication: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalApplicationsClient is an interface that defines the methods available from App Engine Admin API.
type internalApplicationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetApplication(context.Context, *appenginepb.GetApplicationRequest, ...gax.CallOption) (*appenginepb.Application, error)
	CreateApplication(context.Context, *appenginepb.CreateApplicationRequest, ...gax.CallOption) (*CreateApplicationOperation, error)
	CreateApplicationOperation(name string) *CreateApplicationOperation
	UpdateApplication(context.Context, *appenginepb.UpdateApplicationRequest, ...gax.CallOption) (*UpdateApplicationOperation, error)
	UpdateApplicationOperation(name string) *UpdateApplicationOperation
	RepairApplication(context.Context, *appenginepb.RepairApplicationRequest, ...gax.CallOption) (*RepairApplicationOperation, error)
	RepairApplicationOperation(name string) *RepairApplicationOperation
}

// ApplicationsClient is a client for interacting with App Engine Admin API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages App Engine applications.
type ApplicationsClient struct {
	// The internal transport-dependent client.
	internalClient internalApplicationsClient

	// The call options for this service.
	CallOptions *ApplicationsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ApplicationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ApplicationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ApplicationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetApplication gets information about an application.
func (c *ApplicationsClient) GetApplication(ctx context.Context, req *appenginepb.GetApplicationRequest, opts ...gax.CallOption) (*appenginepb.Application, error) {
	return c.internalClient.GetApplication(ctx, req, opts...)
}

// CreateApplication creates an App Engine application for a Google Cloud Platform project.
// Required fields:
//
//	id - The ID of the target Cloud Platform project.
//
//	location - The region (at https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.
//
// For more information about App Engine applications, see Managing Projects, Applications, and Billing (at https://cloud.google.com/appengine/docs/standard/python/console/).
func (c *ApplicationsClient) CreateApplication(ctx context.Context, req *appenginepb.CreateApplicationRequest, opts ...gax.CallOption) (*CreateApplicationOperation, error) {
	return c.internalClient.CreateApplication(ctx, req, opts...)
}

// CreateApplicationOperation returns a new CreateApplicationOperation from a given name.
// The name must be that of a previously created CreateApplicationOperation, possibly from a different process.
func (c *ApplicationsClient) CreateApplicationOperation(name string) *CreateApplicationOperation {
	return c.internalClient.CreateApplicationOperation(name)
}

// UpdateApplication updates the specified Application resource.
// You can update the following fields:
//
//	auth_domain - Google authentication domain for controlling user access to the application.
//
//	default_cookie_expiration - Cookie expiration policy for the application.
//
//	iap - Identity-Aware Proxy properties for the application.
func (c *ApplicationsClient) UpdateApplication(ctx context.Context, req *appenginepb.UpdateApplicationRequest, opts ...gax.CallOption) (*UpdateApplicationOperation, error) {
	return c.internalClient.UpdateApplication(ctx, req, opts...)
}

// UpdateApplicationOperation returns a new UpdateApplicationOperation from a given name.
// The name must be that of a previously created UpdateApplicationOperation, possibly from a different process.
func (c *ApplicationsClient) UpdateApplicationOperation(name string) *UpdateApplicationOperation {
	return c.internalClient.UpdateApplicationOperation(name)
}

// RepairApplication recreates the required App Engine features for the specified App Engine
// application, for example a Cloud Storage bucket or App Engine service
// account.
// Use this method if you receive an error message about a missing feature,
// for example, Error retrieving the App Engine service account.
// If you have deleted your App Engine service account, this will
// not be able to recreate it. Instead, you should attempt to use the
// IAM undelete API if possible at https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts/undelete?apix_params={ (at https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts/undelete?apix_params=%7B)“name”%3A"projects%2F-%2FserviceAccounts%2Funique_id"%2C"resource"%3A%7B%7D%7D .
// If the deletion was recent, the numeric ID can be found in the Cloud
// Console Activity Log.
func (c *ApplicationsClient) RepairApplication(ctx context.Context, req *appenginepb.RepairApplicationRequest, opts ...gax.CallOption) (*RepairApplicationOperation, error) {
	return c.internalClient.RepairApplication(ctx, req, opts...)
}

// RepairApplicationOperation returns a new RepairApplicationOperation from a given name.
// The name must be that of a previously created RepairApplicationOperation, possibly from a different process.
func (c *ApplicationsClient) RepairApplicationOperation(name string) *RepairApplicationOperation {
	return c.internalClient.RepairApplicationOperation(name)
}

// applicationsGRPCClient is a client for interacting with App Engine Admin API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type applicationsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ApplicationsClient
	CallOptions **ApplicationsCallOptions

	// The gRPC API client.
	applicationsClient appenginepb.ApplicationsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewApplicationsClient creates a new applications client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages App Engine applications.
func NewApplicationsClient(ctx context.Context, opts ...option.ClientOption) (*ApplicationsClient, error) {
	clientOpts := defaultApplicationsGRPCClientOptions()
	if newApplicationsClientHook != nil {
		hookOpts, err := newApplicationsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ApplicationsClient{CallOptions: defaultApplicationsCallOptions()}

	c := &applicationsGRPCClient{
		connPool:           connPool,
		applicationsClient: appenginepb.NewApplicationsClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *applicationsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *applicationsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *applicationsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type applicationsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing ApplicationsClient
	CallOptions **ApplicationsCallOptions
}

// NewApplicationsRESTClient creates a new applications rest client.
//
// Manages App Engine applications.
func NewApplicationsRESTClient(ctx context.Context, opts ...option.ClientOption) (*ApplicationsClient, error) {
	clientOpts := append(defaultApplicationsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultApplicationsRESTCallOptions()
	c := &applicationsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &ApplicationsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultApplicationsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://appengine.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://appengine.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://appengine.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *applicationsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *applicationsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *applicationsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *applicationsGRPCClient) GetApplication(ctx context.Context, req *appenginepb.GetApplicationRequest, opts ...gax.CallOption) (*appenginepb.Application, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetApplication[0:len((*c.CallOptions).GetApplication):len((*c.CallOptions).GetApplication)], opts...)
	var resp *appenginepb.Application
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.GetApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *applicationsGRPCClient) CreateApplication(ctx context.Context, req *appenginepb.CreateApplicationRequest, opts ...gax.CallOption) (*CreateApplicationOperation, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateApplication[0:len((*c.CallOptions).CreateApplication):len((*c.CallOptions).CreateApplication)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.CreateApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *applicationsGRPCClient) UpdateApplication(ctx context.Context, req *appenginepb.UpdateApplicationRequest, opts ...gax.CallOption) (*UpdateApplicationOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateApplication[0:len((*c.CallOptions).UpdateApplication):len((*c.CallOptions).UpdateApplication)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.UpdateApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *applicationsGRPCClient) RepairApplication(ctx context.Context, req *appenginepb.RepairApplicationRequest, opts ...gax.CallOption) (*RepairApplicationOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RepairApplication[0:len((*c.CallOptions).RepairApplication):len((*c.CallOptions).RepairApplication)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.applicationsClient.RepairApplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RepairApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// GetApplication gets information about an application.
func (c *applicationsRESTClient) GetApplication(ctx context.Context, req *appenginepb.GetApplicationRequest, opts ...gax.CallOption) (*appenginepb.Application, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetApplication[0:len((*c.CallOptions).GetApplication):len((*c.CallOptions).GetApplication)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &appenginepb.Application{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
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

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// CreateApplication creates an App Engine application for a Google Cloud Platform project.
// Required fields:
//
//	id - The ID of the target Cloud Platform project.
//
//	location - The region (at https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.
//
// For more information about App Engine applications, see Managing Projects, Applications, and Billing (at https://cloud.google.com/appengine/docs/standard/python/console/).
func (c *applicationsRESTClient) CreateApplication(ctx context.Context, req *appenginepb.CreateApplicationRequest, opts ...gax.CallOption) (*CreateApplicationOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetApplication()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/apps")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &CreateApplicationOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// UpdateApplication updates the specified Application resource.
// You can update the following fields:
//
//	auth_domain - Google authentication domain for controlling user access to the application.
//
//	default_cookie_expiration - Cookie expiration policy for the application.
//
//	iap - Identity-Aware Proxy properties for the application.
func (c *applicationsRESTClient) UpdateApplication(ctx context.Context, req *appenginepb.UpdateApplicationRequest, opts ...gax.CallOption) (*UpdateApplicationOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetApplication()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask[1:len(updateMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
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

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &UpdateApplicationOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// RepairApplication recreates the required App Engine features for the specified App Engine
// application, for example a Cloud Storage bucket or App Engine service
// account.
// Use this method if you receive an error message about a missing feature,
// for example, Error retrieving the App Engine service account.
// If you have deleted your App Engine service account, this will
// not be able to recreate it. Instead, you should attempt to use the
// IAM undelete API if possible at https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts/undelete?apix_params={ (at https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts/undelete?apix_params=%7B)“name”%3A"projects%2F-%2FserviceAccounts%2Funique_id"%2C"resource"%3A%7B%7D%7D .
// If the deletion was recent, the numeric ID can be found in the Cloud
// Console Activity Log.
func (c *applicationsRESTClient) RepairApplication(ctx context.Context, req *appenginepb.RepairApplicationRequest, opts ...gax.CallOption) (*RepairApplicationOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:repair", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &RepairApplicationOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// CreateApplicationOperation returns a new CreateApplicationOperation from a given name.
// The name must be that of a previously created CreateApplicationOperation, possibly from a different process.
func (c *applicationsGRPCClient) CreateApplicationOperation(name string) *CreateApplicationOperation {
	return &CreateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// CreateApplicationOperation returns a new CreateApplicationOperation from a given name.
// The name must be that of a previously created CreateApplicationOperation, possibly from a different process.
func (c *applicationsRESTClient) CreateApplicationOperation(name string) *CreateApplicationOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &CreateApplicationOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// RepairApplicationOperation returns a new RepairApplicationOperation from a given name.
// The name must be that of a previously created RepairApplicationOperation, possibly from a different process.
func (c *applicationsGRPCClient) RepairApplicationOperation(name string) *RepairApplicationOperation {
	return &RepairApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// RepairApplicationOperation returns a new RepairApplicationOperation from a given name.
// The name must be that of a previously created RepairApplicationOperation, possibly from a different process.
func (c *applicationsRESTClient) RepairApplicationOperation(name string) *RepairApplicationOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &RepairApplicationOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// UpdateApplicationOperation returns a new UpdateApplicationOperation from a given name.
// The name must be that of a previously created UpdateApplicationOperation, possibly from a different process.
func (c *applicationsGRPCClient) UpdateApplicationOperation(name string) *UpdateApplicationOperation {
	return &UpdateApplicationOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UpdateApplicationOperation returns a new UpdateApplicationOperation from a given name.
// The name must be that of a previously created UpdateApplicationOperation, possibly from a different process.
func (c *applicationsRESTClient) UpdateApplicationOperation(name string) *UpdateApplicationOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &UpdateApplicationOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
