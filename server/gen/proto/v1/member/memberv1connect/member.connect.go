// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/v1/member/member.proto

package memberv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	member "github.com/gilwong00/go-discord/gen/proto/v1/member"
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
	// MemberServiceName is the fully-qualified name of the MemberService service.
	MemberServiceName = "proto.member.v1.MemberService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MemberServiceCreateServerMemberProcedure is the fully-qualified name of the MemberService's
	// CreateServerMember RPC.
	MemberServiceCreateServerMemberProcedure = "/proto.member.v1.MemberService/CreateServerMember"
)

// MemberServiceClient is a client for the proto.member.v1.MemberService service.
type MemberServiceClient interface {
	CreateServerMember(context.Context, *connect.Request[member.CreateServerMemberRequest]) (*connect.Response[member.CreateServerMemberResponse], error)
}

// NewMemberServiceClient constructs a client for the proto.member.v1.MemberService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMemberServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MemberServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &memberServiceClient{
		createServerMember: connect.NewClient[member.CreateServerMemberRequest, member.CreateServerMemberResponse](
			httpClient,
			baseURL+MemberServiceCreateServerMemberProcedure,
			opts...,
		),
	}
}

// memberServiceClient implements MemberServiceClient.
type memberServiceClient struct {
	createServerMember *connect.Client[member.CreateServerMemberRequest, member.CreateServerMemberResponse]
}

// CreateServerMember calls proto.member.v1.MemberService.CreateServerMember.
func (c *memberServiceClient) CreateServerMember(ctx context.Context, req *connect.Request[member.CreateServerMemberRequest]) (*connect.Response[member.CreateServerMemberResponse], error) {
	return c.createServerMember.CallUnary(ctx, req)
}

// MemberServiceHandler is an implementation of the proto.member.v1.MemberService service.
type MemberServiceHandler interface {
	CreateServerMember(context.Context, *connect.Request[member.CreateServerMemberRequest]) (*connect.Response[member.CreateServerMemberResponse], error)
}

// NewMemberServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMemberServiceHandler(svc MemberServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	memberServiceCreateServerMemberHandler := connect.NewUnaryHandler(
		MemberServiceCreateServerMemberProcedure,
		svc.CreateServerMember,
		opts...,
	)
	return "/proto.member.v1.MemberService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MemberServiceCreateServerMemberProcedure:
			memberServiceCreateServerMemberHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMemberServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMemberServiceHandler struct{}

func (UnimplementedMemberServiceHandler) CreateServerMember(context.Context, *connect.Request[member.CreateServerMemberRequest]) (*connect.Response[member.CreateServerMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.member.v1.MemberService.CreateServerMember is not implemented"))
}
