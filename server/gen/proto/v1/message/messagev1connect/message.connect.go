// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/v1/message/message.proto

package messagev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	message "github.com/gilwong00/go-discord/gen/proto/v1/message"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ServerServiceName is the fully-qualified name of the ServerService service.
	ServerServiceName = "proto.message.v1.ServerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ServerServiceSendDirectMessageProcedure is the fully-qualified name of the ServerService's
	// SendDirectMessage RPC.
	ServerServiceSendDirectMessageProcedure = "/proto.message.v1.ServerService/SendDirectMessage"
)

// ServerServiceClient is a client for the proto.message.v1.ServerService service.
type ServerServiceClient interface {
	SendDirectMessage(context.Context, *connect.Request[message.SendDirectMessageRequest]) (*connect.ServerStreamForClient[message.SendDirectMessageResponse], error)
}

// NewServerServiceClient constructs a client for the proto.message.v1.ServerService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewServerServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ServerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &serverServiceClient{
		sendDirectMessage: connect.NewClient[message.SendDirectMessageRequest, message.SendDirectMessageResponse](
			httpClient,
			baseURL+ServerServiceSendDirectMessageProcedure,
			opts...,
		),
	}
}

// serverServiceClient implements ServerServiceClient.
type serverServiceClient struct {
	sendDirectMessage *connect.Client[message.SendDirectMessageRequest, message.SendDirectMessageResponse]
}

// SendDirectMessage calls proto.message.v1.ServerService.SendDirectMessage.
func (c *serverServiceClient) SendDirectMessage(ctx context.Context, req *connect.Request[message.SendDirectMessageRequest]) (*connect.ServerStreamForClient[message.SendDirectMessageResponse], error) {
	return c.sendDirectMessage.CallServerStream(ctx, req)
}

// ServerServiceHandler is an implementation of the proto.message.v1.ServerService service.
type ServerServiceHandler interface {
	SendDirectMessage(context.Context, *connect.Request[message.SendDirectMessageRequest], *connect.ServerStream[message.SendDirectMessageResponse]) error
}

// NewServerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewServerServiceHandler(svc ServerServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	serverServiceSendDirectMessageHandler := connect.NewServerStreamHandler(
		ServerServiceSendDirectMessageProcedure,
		svc.SendDirectMessage,
		opts...,
	)
	return "/proto.message.v1.ServerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ServerServiceSendDirectMessageProcedure:
			serverServiceSendDirectMessageHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedServerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedServerServiceHandler struct{}

func (UnimplementedServerServiceHandler) SendDirectMessage(context.Context, *connect.Request[message.SendDirectMessageRequest], *connect.ServerStream[message.SendDirectMessageResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("proto.message.v1.ServerService.SendDirectMessage is not implemented"))
}
