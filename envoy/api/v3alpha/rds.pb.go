// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/rds.proto

package envoy_api_v3alpha

import (
	context "context"
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RouteConfiguration struct {
	Name                    string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VirtualHosts            []*route.VirtualHost      `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	Vhds                    *Vhds                     `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	InternalOnlyHeaders     []string                  `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	ResponseHeadersToAdd    []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	ResponseHeadersToRemove []string                  `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	RequestHeadersToAdd     []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove  []string                  `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	ValidateClusters        *wrappers.BoolValue       `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaba07addf704458, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*route.VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetVhds() *Vhds {
	if m != nil {
		return m.Vhds
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

type Vhds struct {
	ConfigSource         *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Vhds) Reset()         { *m = Vhds{} }
func (m *Vhds) String() string { return proto.CompactTextString(m) }
func (*Vhds) ProtoMessage()    {}
func (*Vhds) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaba07addf704458, []int{1}
}

func (m *Vhds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vhds.Unmarshal(m, b)
}
func (m *Vhds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vhds.Marshal(b, m, deterministic)
}
func (m *Vhds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vhds.Merge(m, src)
}
func (m *Vhds) XXX_Size() int {
	return xxx_messageInfo_Vhds.Size(m)
}
func (m *Vhds) XXX_DiscardUnknown() {
	xxx_messageInfo_Vhds.DiscardUnknown(m)
}

var xxx_messageInfo_Vhds proto.InternalMessageInfo

func (m *Vhds) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.api.v3alpha.RouteConfiguration")
	proto.RegisterType((*Vhds)(nil), "envoy.api.v3alpha.Vhds")
}

func init() { proto.RegisterFile("envoy/api/v3alpha/rds.proto", fileDescriptor_eaba07addf704458) }

var fileDescriptor_eaba07addf704458 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x26, 0x49, 0x7f, 0x32, 0x69, 0xa5, 0x76, 0xfa, 0xb5, 0x49, 0x53, 0x54, 0x42, 0xda,
	0x45, 0x5a, 0x24, 0xbb, 0x4a, 0x57, 0x94, 0x15, 0x6e, 0x05, 0x65, 0xd5, 0xca, 0x41, 0xd9, 0xb0,
	0xb0, 0xa6, 0xf6, 0x6d, 0x6d, 0xe1, 0x7a, 0xcc, 0xcc, 0xd8, 0x90, 0x0d, 0x0b, 0x56, 0x08, 0x89,
	0x0d, 0x3c, 0x1a, 0x6f, 0x80, 0x58, 0xb0, 0xe1, 0x05, 0xba, 0x42, 0x9e, 0xb1, 0xa1, 0xa9, 0x0d,
	0x42, 0x02, 0x36, 0x91, 0x9d, 0x73, 0xce, 0xbd, 0x67, 0xee, 0x9c, 0x6b, 0xbc, 0x01, 0x51, 0xca,
	0x26, 0x26, 0x8d, 0x03, 0x33, 0xdd, 0xa7, 0x61, 0xec, 0x53, 0x93, 0x7b, 0xc2, 0x88, 0x39, 0x93,
	0x8c, 0x2c, 0x2b, 0xd0, 0xa0, 0x71, 0x60, 0xe4, 0x60, 0xf7, 0x4e, 0x99, 0xef, 0x32, 0x0e, 0xe6,
	0x19, 0x15, 0xa0, 0x55, 0xdd, 0xdd, 0x9f, 0x50, 0x5c, 0x16, 0x9d, 0x07, 0x17, 0x8e, 0x60, 0x09,
	0x77, 0x0b, 0x6e, 0x45, 0x39, 0x2f, 0x10, 0x2e, 0x4b, 0x81, 0x4f, 0x72, 0xca, 0x56, 0x85, 0x43,
	0x96, 0x48, 0xd0, 0xbf, 0x39, 0xe9, 0xd6, 0x05, 0x63, 0x17, 0x21, 0x28, 0x16, 0x8d, 0x22, 0x26,
	0xa9, 0x0c, 0x58, 0x94, 0x9f, 0xa3, 0xbb, 0x99, 0xa3, 0xea, 0xed, 0x2c, 0x39, 0x37, 0x5f, 0x70,
	0x1a, 0xc7, 0xc0, 0x0b, 0xbc, 0x9d, 0xd2, 0x30, 0xf0, 0xa8, 0x04, 0xb3, 0x78, 0xd0, 0x40, 0xff,
	0x6b, 0x03, 0x13, 0x3b, 0x6b, 0x73, 0xa8, 0xbc, 0x27, 0x5c, 0x95, 0x25, 0x04, 0x37, 0x22, 0x7a,
	0x09, 0x1d, 0xd4, 0x43, 0x83, 0xa6, 0xad, 0x9e, 0xc9, 0x63, 0xbc, 0x98, 0x06, 0x5c, 0x26, 0x34,
	0x74, 0x7c, 0x26, 0xa4, 0xe8, 0xd4, 0x7a, 0xf5, 0x41, 0x6b, 0xb8, 0x6d, 0x94, 0x66, 0x68, 0x68,
	0xe3, 0x63, 0xcd, 0x3e, 0x66, 0x42, 0xda, 0x0b, 0xe9, 0x8f, 0x17, 0x41, 0xee, 0xe2, 0x46, 0xea,
	0x7b, 0xa2, 0xd3, 0xec, 0xa1, 0x41, 0x6b, 0xd8, 0xae, 0xa8, 0x30, 0xf6, 0x3d, 0x61, 0x2b, 0x12,
	0x19, 0xe2, 0xd5, 0x20, 0x92, 0xc0, 0x23, 0x1a, 0x3a, 0x2c, 0x0a, 0x27, 0x8e, 0x0f, 0xd4, 0x03,
	0x2e, 0x3a, 0xf5, 0x5e, 0x7d, 0xd0, 0xb4, 0x57, 0x0a, 0xf0, 0x24, 0x0a, 0x27, 0xc7, 0x1a, 0x22,
	0xcf, 0x70, 0x9b, 0x83, 0x88, 0x59, 0x24, 0xa0, 0xa0, 0x3b, 0x92, 0x39, 0xd4, 0xf3, 0x3a, 0x0d,
	0xe5, 0x7a, 0xa7, 0xa2, 0x67, 0x76, 0x87, 0x86, 0xae, 0x30, 0xa6, 0x61, 0x02, 0x27, 0x71, 0x36,
	0x0b, 0xab, 0x79, 0x65, 0xcd, 0xbe, 0x47, 0xf5, 0xa5, 0x2f, 0x73, 0xf6, 0xff, 0x45, 0xd1, 0xbc,
	0xcf, 0x13, 0xf6, 0xc0, 0xf3, 0xc8, 0x7d, 0xdc, 0xad, 0x6a, 0xc6, 0xe1, 0x92, 0xa5, 0xd0, 0x99,
	0x51, 0x2e, 0xdb, 0x25, 0xa5, 0xad, 0x60, 0x12, 0xe0, 0x35, 0x0e, 0xcf, 0x13, 0x10, 0xf2, 0xa6,
	0xd1, 0xd9, 0x3f, 0x30, 0xba, 0x92, 0xd7, 0x9c, 0xf2, 0x79, 0x0f, 0xaf, 0x57, 0xb4, 0xca, 0x6d,
	0xce, 0x2b, 0x9b, 0x6b, 0x37, 0x75, 0xb9, 0xcb, 0x47, 0x78, 0xb9, 0x08, 0x8e, 0xe3, 0x86, 0x89,
	0x90, 0xd9, 0xfc, 0xe7, 0xd4, 0xed, 0x75, 0x0d, 0x9d, 0x3d, 0xa3, 0xc8, 0x9e, 0x61, 0x31, 0x16,
	0x2a, 0x5f, 0xf6, 0x52, 0x21, 0x3a, 0xcc, 0x35, 0xfd, 0xa7, 0xb8, 0x91, 0x5d, 0x2d, 0x19, 0xe1,
	0xc5, 0xa9, 0x6d, 0x51, 0x49, 0xab, 0x0e, 0x93, 0x3a, 0xad, 0x8e, 0xe7, 0x48, 0x71, 0xad, 0xf9,
	0x2b, 0x6b, 0xe6, 0x2d, 0xaa, 0x2d, 0x21, 0x7b, 0xc1, 0xbd, 0xf6, 0xff, 0xf0, 0x53, 0x0d, 0xaf,
	0xaa, 0x30, 0x1f, 0x15, 0x1b, 0x36, 0x02, 0x9e, 0x06, 0x2e, 0x10, 0x07, 0x2f, 0x8c, 0x24, 0x07,
	0x7a, 0xa9, 0x60, 0x41, 0xb6, 0x2a, 0xfa, 0x7c, 0x17, 0xd9, 0x7a, 0x06, 0xdd, 0xed, 0x5f, 0x93,
	0xf4, 0x7d, 0xf6, 0xff, 0x1b, 0xa0, 0x3d, 0x44, 0x7c, 0xdc, 0x3a, 0x82, 0x50, 0xd2, 0xbc, 0xfe,
	0xa0, 0x4a, 0x9a, 0xe1, 0xa5, 0x26, 0x3b, 0xbf, 0xc1, 0x9c, 0xea, 0xf4, 0x0a, 0xb7, 0x1e, 0x82,
	0x74, 0xfd, 0xbf, 0x7f, 0x92, 0xed, 0xd7, 0x1f, 0x3f, 0x7f, 0xa8, 0x6d, 0xf6, 0xd7, 0xcb, 0x5f,
	0xaa, 0x03, 0xb5, 0xd1, 0xe2, 0x00, 0xed, 0x0e, 0xdf, 0x21, 0xbc, 0x71, 0x6d, 0xb3, 0x4b, 0xa3,
	0x8e, 0xf0, 0xb2, 0xf2, 0x3f, 0xbe, 0xbe, 0xf0, 0xff, 0x6e, 0x1e, 0xd6, 0x1e, 0xbe, 0x1d, 0x30,
	0x2d, 0x8a, 0x39, 0x7b, 0x39, 0x29, 0xeb, 0xad, 0x79, 0xdb, 0x13, 0xa7, 0x59, 0x3a, 0x4f, 0xd1,
	0x1b, 0x84, 0xce, 0x66, 0x55, 0x52, 0xf7, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0x80, 0xe8, 0x67,
	0xed, 0x1a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteDiscoveryServiceClient is the client API for RouteDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteDiscoveryServiceClient interface {
	StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error)
	DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error)
	FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v3alpha.RouteDiscoveryService/StreamRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceStreamRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_StreamRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceStreamRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceStreamRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v3alpha.RouteDiscoveryService/DeltaRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceDeltaRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_DeltaRoutesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceDeltaRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v3alpha.RouteDiscoveryService/FetchRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteDiscoveryServiceServer is the server API for RouteDiscoveryService service.
type RouteDiscoveryServiceServer interface {
	StreamRoutes(RouteDiscoveryService_StreamRoutesServer) error
	DeltaRoutes(RouteDiscoveryService_DeltaRoutesServer) error
	FetchRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_StreamRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).StreamRoutes(&routeDiscoveryServiceStreamRoutesServer{stream})
}

type RouteDiscoveryService_StreamRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceStreamRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceStreamRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_DeltaRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).DeltaRoutes(&routeDiscoveryServiceDeltaRoutesServer{stream})
}

type RouteDiscoveryService_DeltaRoutesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceDeltaRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_FetchRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v3alpha.RouteDiscoveryService/FetchRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v3alpha.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRoutes",
			Handler:    _RouteDiscoveryService_FetchRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoutes",
			Handler:       _RouteDiscoveryService_StreamRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRoutes",
			Handler:       _RouteDiscoveryService_DeltaRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v3alpha/rds.proto",
}

// VirtualHostDiscoveryServiceClient is the client API for VirtualHostDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHostDiscoveryServiceClient interface {
	DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error)
}

type virtualHostDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHostDiscoveryServiceClient(cc *grpc.ClientConn) VirtualHostDiscoveryServiceClient {
	return &virtualHostDiscoveryServiceClient{cc}
}

func (c *virtualHostDiscoveryServiceClient) DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VirtualHostDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v3alpha.VirtualHostDiscoveryService/DeltaVirtualHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &virtualHostDiscoveryServiceDeltaVirtualHostsClient{stream}
	return x, nil
}

type VirtualHostDiscoveryService_DeltaVirtualHostsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsClient struct {
	grpc.ClientStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VirtualHostDiscoveryServiceServer is the server API for VirtualHostDiscoveryService service.
type VirtualHostDiscoveryServiceServer interface {
	DeltaVirtualHosts(VirtualHostDiscoveryService_DeltaVirtualHostsServer) error
}

func RegisterVirtualHostDiscoveryServiceServer(s *grpc.Server, srv VirtualHostDiscoveryServiceServer) {
	s.RegisterService(&_VirtualHostDiscoveryService_serviceDesc, srv)
}

func _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VirtualHostDiscoveryServiceServer).DeltaVirtualHosts(&virtualHostDiscoveryServiceDeltaVirtualHostsServer{stream})
}

type VirtualHostDiscoveryService_DeltaVirtualHostsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsServer struct {
	grpc.ServerStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VirtualHostDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v3alpha.VirtualHostDiscoveryService",
	HandlerType: (*VirtualHostDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaVirtualHosts",
			Handler:       _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v3alpha/rds.proto",
}
