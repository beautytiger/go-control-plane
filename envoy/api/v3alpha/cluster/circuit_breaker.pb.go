// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/cluster/circuit_breaker.proto

package envoy_api_v3alpha_cluster

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type CircuitBreakers struct {
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f46fa043d540e79, []int{0}
}

func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers.Unmarshal(m, b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers.Size(m)
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

type CircuitBreakers_Thresholds struct {
	Priority             core.RoutingPriority  `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.api.v3alpha.core.RoutingPriority" json:"priority,omitempty"`
	MaxConnections       *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests   *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests          *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries           *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	TrackRemaining       bool                  `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	MaxConnectionPools   *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f46fa043d540e79, []int{0, 0}
}

func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Size(m)
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() core.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return core.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.api.v3alpha.cluster.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.api.v3alpha.cluster.CircuitBreakers.Thresholds")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/cluster/circuit_breaker.proto", fileDescriptor_9f46fa043d540e79)
}

var fileDescriptor_9f46fa043d540e79 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x15, 0xba, 0x2c, 0x2b, 0x17, 0xb5, 0x92, 0xe1, 0x10, 0x2a, 0x84, 0x0a, 0x97, 0xf6,
	0x64, 0x4b, 0xad, 0x38, 0x22, 0x50, 0x2b, 0x0e, 0x5c, 0x50, 0x14, 0xb1, 0x5c, 0x23, 0x37, 0x3b,
	0xa4, 0xd6, 0x26, 0x1e, 0x33, 0x76, 0x96, 0xf4, 0x1d, 0x79, 0x09, 0xde, 0x04, 0x35, 0x6e, 0x13,
	0x58, 0xed, 0x4a, 0x39, 0xc6, 0xf6, 0xf7, 0xfd, 0xf9, 0xed, 0x61, 0x12, 0xcc, 0x1d, 0x1e, 0xa4,
	0xb2, 0x5a, 0xde, 0xad, 0x55, 0x69, 0xf7, 0x4a, 0xe6, 0x65, 0xed, 0x3c, 0x90, 0xcc, 0x35, 0xe5,
	0xb5, 0xf6, 0xd9, 0x8e, 0x40, 0xdd, 0x02, 0x09, 0x4b, 0xe8, 0x91, 0xbf, 0x6a, 0x01, 0xa1, 0xac,
	0x16, 0x27, 0x40, 0x9c, 0x80, 0xd9, 0xdb, 0x07, 0x5c, 0x48, 0x20, 0x77, 0xca, 0x41, 0xa0, 0x67,
	0x6f, 0x0a, 0xc4, 0xa2, 0x04, 0xd9, 0x7e, 0xed, 0xea, 0x1f, 0xf2, 0x17, 0x29, 0x6b, 0x81, 0x5c,
	0xd8, 0x7f, 0xf7, 0xfb, 0x82, 0x4d, 0xb7, 0x21, 0x77, 0x13, 0x62, 0x1d, 0xbf, 0x66, 0xcc, 0xef,
	0x09, 0xdc, 0x1e, 0xcb, 0x1b, 0x17, 0x47, 0xf3, 0xd1, 0x72, 0xbc, 0x7a, 0x2f, 0x1e, 0xfd, 0x0d,
	0x71, 0x8f, 0x17, 0xdf, 0x3a, 0x38, 0xfd, 0x47, 0x34, 0xfb, 0x33, 0x62, 0xac, 0xdf, 0xe2, 0x5b,
	0x76, 0x65, 0x49, 0x23, 0x69, 0x7f, 0x88, 0xa3, 0x79, 0xb4, 0x9c, 0xac, 0x16, 0x0f, 0x65, 0x20,
	0x81, 0x48, 0xb1, 0xf6, 0xda, 0x14, 0xc9, 0xe9, 0x78, 0xda, 0x81, 0xfc, 0x33, 0x9b, 0x56, 0xaa,
	0xc9, 0x72, 0x34, 0x06, 0x72, 0xaf, 0xd1, 0xb8, 0xf8, 0xc9, 0x3c, 0x5a, 0x8e, 0x57, 0xaf, 0x45,
	0x28, 0x2e, 0xce, 0xc5, 0xc5, 0xf5, 0x17, 0xe3, 0xd7, 0xab, 0xef, 0xaa, 0xac, 0x21, 0x9d, 0x54,
	0xaa, 0xd9, 0xf6, 0x0c, 0xff, 0xca, 0x5e, 0x1e, 0x35, 0x16, 0xcc, 0x8d, 0x36, 0x45, 0x46, 0xf0,
	0xb3, 0x06, 0xe7, 0x5d, 0x3c, 0x1a, 0xe0, 0xe2, 0x95, 0x6a, 0x92, 0x00, 0xa6, 0x27, 0x8e, 0x7f,
	0x64, 0xcf, 0x8f, 0xbe, 0xce, 0x73, 0x31, 0xc0, 0x33, 0xae, 0x54, 0xd3, 0x09, 0x3e, 0xb0, 0x71,
	0x10, 0x78, 0xd2, 0xe0, 0xe2, 0xa7, 0x03, 0x78, 0xd6, 0xf2, 0xed, 0x79, 0xbe, 0x60, 0x53, 0x4f,
	0x2a, 0xbf, 0xcd, 0x08, 0x2a, 0xa5, 0x8d, 0x36, 0x45, 0x7c, 0x39, 0x8f, 0x96, 0x57, 0xe9, 0xa4,
	0x5d, 0x4e, 0xcf, 0xab, 0xe7, 0xe2, 0xfd, 0xfd, 0x65, 0x16, 0xb1, 0x74, 0xf1, 0xb3, 0x81, 0xc5,
	0xfb, 0x4b, 0x4c, 0x8e, 0xdc, 0xe6, 0x13, 0x5b, 0x68, 0x0c, 0xcf, 0x68, 0x09, 0x9b, 0xc3, 0xe3,
	0x53, 0xb3, 0x79, 0xf1, 0xff, 0xd8, 0x24, 0xc7, 0x88, 0x24, 0xda, 0x5d, 0xb6, 0x59, 0xeb, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x24, 0x65, 0x90, 0x14, 0x28, 0x03, 0x00, 0x00,
}
