// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v2alpha/external_auth.proto

package envoy_service_auth_v2alpha

import (
	context "context"
	fmt "fmt"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("envoy/service/auth/v2alpha/external_auth.proto", fileDescriptor_878c0ddb0c43de8d)
}

var fileDescriptor_878c0ddb0c43de8d = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xce, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x06, 0xe0, 0x30, 0xe8, 0xd0, 0xc4, 0x85, 0x91, 0x55, 0x07, 0x5d, 0xae, 0x09, 0x8e, 0x4e,
	0xc2, 0x0b, 0x10, 0x5e, 0xc0, 0x54, 0x72, 0x49, 0x1b, 0x49, 0xaf, 0xb6, 0x07, 0x01, 0x9f, 0xc0,
	0xc7, 0x36, 0x54, 0xc6, 0xea, 0x7a, 0xff, 0x97, 0xff, 0x7e, 0x01, 0x68, 0x47, 0x9a, 0x65, 0x40,
	0x3f, 0x9a, 0x0e, 0xa5, 0x1a, 0x58, 0xcb, 0xb1, 0x54, 0xbd, 0xd3, 0x4a, 0xe2, 0xc4, 0xe8, 0xad,
	0xea, 0x6f, 0xcb, 0x15, 0x9c, 0x27, 0xa6, 0xbc, 0x88, 0x1e, 0x56, 0x0f, 0x31, 0x59, 0x7d, 0x71,
	0x4a, 0x76, 0xa5, 0x6a, 0xca, 0x4e, 0xec, 0xae, 0x03, 0x6b, 0xf2, 0xe6, 0xa5, 0xd8, 0x90, 0xcd,
	0x5b, 0xb1, 0xa9, 0x35, 0x76, 0x8f, 0x7c, 0x0f, 0xc9, 0x0f, 0x10, 0xd3, 0x16, 0x9f, 0x03, 0x06,
	0x2e, 0x0e, 0xff, 0x51, 0x70, 0x64, 0x03, 0x56, 0x17, 0x71, 0x34, 0xf4, 0x95, 0xce, 0xd3, 0x34,
	0xc3, 0xef, 0xed, 0x95, 0xa8, 0xd1, 0x73, 0x68, 0x96, 0x71, 0x4d, 0xf6, 0xce, 0xb2, 0xfb, 0x36,
	0x0e, 0x3d, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xed, 0xe3, 0x41, 0xb2, 0x21, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	Check(ctx context.Context, in *v2.CheckRequest, opts ...grpc.CallOption) (*v2.CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *v2.CheckRequest, opts ...grpc.CallOption) (*v2.CheckResponse, error) {
	out := new(v2.CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v2alpha.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	Check(context.Context, *v2.CheckRequest) (*v2.CheckResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v2alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*v2.CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v2alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v2alpha/external_auth.proto",
}
