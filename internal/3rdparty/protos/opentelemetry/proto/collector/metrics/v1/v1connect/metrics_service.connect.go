// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: opentelemetry/proto/collector/metrics/v1/metrics_service.proto

package v1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect_go "github.com/bufbuild/connect-go"
	v1 "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// MetricsServiceName is the fully-qualified name of the MetricsService service.
	MetricsServiceName = "opentelemetry.proto.collector.metrics.v1.MetricsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MetricsServiceExportProcedure is the fully-qualified name of the MetricsService's Export RPC.
	MetricsServiceExportProcedure = "/opentelemetry.proto.collector.metrics.v1.MetricsService/Export"
)

// MetricsServiceClient is a client for the opentelemetry.proto.collector.metrics.v1.MetricsService
// service.
type MetricsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *connect_go.Request[v1.ExportMetricsServiceRequest]) (*connect_go.Response[v1.ExportMetricsServiceResponse], error)
}

// NewMetricsServiceClient constructs a client for the
// opentelemetry.proto.collector.metrics.v1.MetricsService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMetricsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MetricsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &metricsServiceClient{
		export: connect_go.NewClient[v1.ExportMetricsServiceRequest, v1.ExportMetricsServiceResponse](
			httpClient,
			baseURL+MetricsServiceExportProcedure,
			opts...,
		),
	}
}

// metricsServiceClient implements MetricsServiceClient.
type metricsServiceClient struct {
	export *connect_go.Client[v1.ExportMetricsServiceRequest, v1.ExportMetricsServiceResponse]
}

// Export calls opentelemetry.proto.collector.metrics.v1.MetricsService.Export.
func (c *metricsServiceClient) Export(ctx context.Context, req *connect_go.Request[v1.ExportMetricsServiceRequest]) (*connect_go.Response[v1.ExportMetricsServiceResponse], error) {
	return c.export.CallUnary(ctx, req)
}

// MetricsServiceHandler is an implementation of the
// opentelemetry.proto.collector.metrics.v1.MetricsService service.
type MetricsServiceHandler interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *connect_go.Request[v1.ExportMetricsServiceRequest]) (*connect_go.Response[v1.ExportMetricsServiceResponse], error)
}

// NewMetricsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMetricsServiceHandler(svc MetricsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(MetricsServiceExportProcedure, connect_go.NewUnaryHandler(
		MetricsServiceExportProcedure,
		svc.Export,
		opts...,
	))
	return "/opentelemetry.proto.collector.metrics.v1.MetricsService/", mux
}

// UnimplementedMetricsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMetricsServiceHandler struct{}

func (UnimplementedMetricsServiceHandler) Export(context.Context, *connect_go.Request[v1.ExportMetricsServiceRequest]) (*connect_go.Response[v1.ExportMetricsServiceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("opentelemetry.proto.collector.metrics.v1.MetricsService.Export is not implemented"))
}
