// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/client_ssl_auth/v3/client_ssl_auth.proto

package envoy_extensions_filters_network_client_ssl_auth_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ClientSSLAuth struct {
	AuthApiCluster       string             `protobuf:"bytes,1,opt,name=auth_api_cluster,json=authApiCluster,proto3" json:"auth_api_cluster,omitempty"`
	StatPrefix           string             `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	RefreshDelay         *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	IpWhiteList          []*v3.CidrRange    `protobuf:"bytes,4,rep,name=ip_white_list,json=ipWhiteList,proto3" json:"ip_white_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClientSSLAuth) Reset()         { *m = ClientSSLAuth{} }
func (m *ClientSSLAuth) String() string { return proto.CompactTextString(m) }
func (*ClientSSLAuth) ProtoMessage()    {}
func (*ClientSSLAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ad74cf4e5256d1, []int{0}
}

func (m *ClientSSLAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSSLAuth.Unmarshal(m, b)
}
func (m *ClientSSLAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSSLAuth.Marshal(b, m, deterministic)
}
func (m *ClientSSLAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSSLAuth.Merge(m, src)
}
func (m *ClientSSLAuth) XXX_Size() int {
	return xxx_messageInfo_ClientSSLAuth.Size(m)
}
func (m *ClientSSLAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSSLAuth.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSSLAuth proto.InternalMessageInfo

func (m *ClientSSLAuth) GetAuthApiCluster() string {
	if m != nil {
		return m.AuthApiCluster
	}
	return ""
}

func (m *ClientSSLAuth) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ClientSSLAuth) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ClientSSLAuth) GetIpWhiteList() []*v3.CidrRange {
	if m != nil {
		return m.IpWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientSSLAuth)(nil), "envoy.extensions.filters.network.client_ssl_auth.v3.ClientSSLAuth")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/client_ssl_auth/v3/client_ssl_auth.proto", fileDescriptor_30ad74cf4e5256d1)
}

var fileDescriptor_30ad74cf4e5256d1 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0xdf, 0xe9, 0x10, 0x1b, 0x82, 0x4e, 0x6e, 0x08, 0x27, 0x71, 0x84, 0xab, 0x52,
	0xed, 0x8a, 0xb8, 0x43, 0xe8, 0xa4, 0xc4, 0xd7, 0x20, 0x5d, 0x11, 0xe5, 0x0a, 0x4a, 0x6b, 0x2f,
	0x1e, 0xdb, 0x23, 0x56, 0xbb, 0xd6, 0xee, 0xd8, 0x97, 0x74, 0x94, 0x3c, 0x03, 0x0f, 0x42, 0x41,
	0x8f, 0x44, 0xcb, 0xeb, 0x50, 0xa1, 0xf5, 0x3a, 0x42, 0x09, 0xa2, 0xb9, 0xce, 0x9e, 0xff, 0x9f,
	0x7f, 0x46, 0xdf, 0x2c, 0xfb, 0x00, 0xba, 0x33, 0x3b, 0x01, 0x5b, 0x02, 0xed, 0xd0, 0x68, 0x27,
	0x4a, 0x54, 0x04, 0xd6, 0x09, 0x0d, 0xf4, 0x60, 0xec, 0x27, 0xb1, 0x51, 0x08, 0x9a, 0x72, 0xe7,
	0x54, 0x2e, 0x5b, 0xaa, 0x45, 0x97, 0x1e, 0x97, 0x78, 0x63, 0x0d, 0x99, 0x24, 0xed, 0xa3, 0xf8,
	0xdf, 0x28, 0x3e, 0x44, 0xf1, 0x21, 0x8a, 0x1f, 0xf7, 0x75, 0xe9, 0xc5, 0x55, 0x98, 0xbf, 0x31,
	0xba, 0xc4, 0x4a, 0x6c, 0x8c, 0x05, 0x3f, 0x40, 0x16, 0x85, 0x05, 0xe7, 0x42, 0xf0, 0xc5, 0x65,
	0x65, 0x4c, 0xa5, 0x40, 0xf4, 0x7f, 0xf7, 0x6d, 0x29, 0x8a, 0xd6, 0x4a, 0x42, 0xa3, 0x07, 0xfd,
	0x55, 0x5b, 0x34, 0x52, 0x48, 0xad, 0x0d, 0xf5, 0x65, 0x27, 0x1c, 0x49, 0x6a, 0xf7, 0xed, 0x6f,
	0xfe, 0x91, 0x3b, 0xb0, 0x7e, 0x41, 0xd4, 0xd5, 0x60, 0x79, 0xd1, 0x49, 0x85, 0x85, 0x24, 0x10,
	0xfb, 0x8f, 0x20, 0x5c, 0x7d, 0x8b, 0xd9, 0x38, 0xeb, 0xb7, 0xbe, 0xbb, 0xbb, 0x5d, 0xb4, 0x54,
	0x27, 0x6f, 0xd9, 0xb9, 0xdf, 0x3d, 0x97, 0x0d, 0xe6, 0x1b, 0xd5, 0x3a, 0x02, 0x3b, 0x89, 0xa6,
	0xd1, 0xec, 0xe9, 0xf2, 0xc9, 0xef, 0xe5, 0xa9, 0x8d, 0xa7, 0xd1, 0xfa, 0xb9, 0x37, 0x2c, 0x1a,
	0xcc, 0x82, 0x9c, 0xcc, 0xd8, 0xc8, 0x2f, 0x94, 0x37, 0x16, 0x4a, 0xdc, 0x4e, 0xe2, 0x43, 0x37,
	0xf3, 0xda, 0xaa, 0x97, 0x92, 0x6b, 0x36, 0xb6, 0x50, 0x5a, 0x70, 0x75, 0x5e, 0x80, 0x92, 0xbb,
	0xc9, 0xc9, 0x34, 0x9a, 0x8d, 0xe6, 0x2f, 0x79, 0x20, 0xc0, 0xf7, 0x04, 0xf8, 0xcd, 0x40, 0x60,
	0xfd, 0x6c, 0xf0, 0xdf, 0x78, 0x7b, 0x92, 0xb1, 0x31, 0x36, 0xf9, 0x43, 0x8d, 0x04, 0xb9, 0x42,
	0x47, 0x93, 0xd3, 0xe9, 0xc9, 0x6c, 0x34, 0x7f, 0xcd, 0xc3, 0x69, 0x02, 0x65, 0xee, 0x29, 0xf3,
	0x2e, 0xe5, 0x19, 0x16, 0x76, 0x2d, 0x75, 0x05, 0xeb, 0x11, 0x36, 0x1f, 0x7d, 0xd3, 0x2d, 0x3a,
	0x7a, 0x97, 0x7d, 0xfd, 0xf1, 0xe5, 0xf2, 0x9a, 0xbd, 0x3f, 0xe8, 0x09, 0xa7, 0xfc, 0xff, 0x25,
	0xe7, 0xfc, 0x00, 0xd3, 0x52, 0x7e, 0xff, 0xfc, 0xf3, 0xd7, 0x59, 0x7c, 0x1e, 0xb3, 0x05, 0x9a,
	0x30, 0xbe, 0xb1, 0x66, 0xbb, 0xe3, 0x8f, 0x78, 0x24, 0xcb, 0x64, 0xc8, 0x76, 0xca, 0x67, 0xaf,
	0x3c, 0x84, 0x55, 0x74, 0x7f, 0xd6, 0xd3, 0x48, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x8a,
	0xd3, 0x15, 0xc3, 0x02, 0x00, 0x00,
}
