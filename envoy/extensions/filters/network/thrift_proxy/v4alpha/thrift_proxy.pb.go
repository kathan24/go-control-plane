// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v4alpha/thrift_proxy.proto

package envoy_extensions_filters_network_thrift_proxy_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type TransportType int32

const (
	TransportType_AUTO_TRANSPORT TransportType = 0
	TransportType_FRAMED         TransportType = 1
	TransportType_UNFRAMED       TransportType = 2
	TransportType_HEADER         TransportType = 3
)

var TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x TransportType) String() string {
	return proto.EnumName(TransportType_name, int32(x))
}

func (TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39167cae9dfa2168, []int{0}
}

type ProtocolType int32

const (
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	ProtocolType_BINARY        ProtocolType = 1
	ProtocolType_LAX_BINARY    ProtocolType = 2
	ProtocolType_COMPACT       ProtocolType = 3
	ProtocolType_TWITTER       ProtocolType = 4
)

var ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
	4: "TWITTER",
}

var ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
	"TWITTER":       4,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39167cae9dfa2168, []int{1}
}

type ThriftProxy struct {
	Transport            TransportType       `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v4alpha.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType        `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v4alpha.ProtocolType" json:"protocol,omitempty"`
	StatPrefix           string              `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	RouteConfig          *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	ThriftFilters        []*ThriftFilter     `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_39167cae9dfa2168, []int{0}
}

func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (m *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(m, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return xxx_messageInfo_ThriftProxy.Size(m)
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if m != nil {
		return m.ThriftFilters
	}
	return nil
}

type ThriftFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ThriftFilter_TypedConfig
	ConfigType           isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_39167cae9dfa2168, []int{1}
}

func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftFilter.Unmarshal(m, b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
}
func (m *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(m, src)
}
func (m *ThriftFilter) XXX_Size() int {
	return xxx_messageInfo_ThriftFilter.Size(m)
}
func (m *ThriftFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftFilter proto.InternalMessageInfo

func (m *ThriftFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ThriftFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ThriftFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ThriftFilter_TypedConfig)(nil),
	}
}

type ThriftProtocolOptions struct {
	Transport            TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v4alpha.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType  `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v4alpha.ProtocolType" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ThriftProtocolOptions) Reset()         { *m = ThriftProtocolOptions{} }
func (m *ThriftProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*ThriftProtocolOptions) ProtoMessage()    {}
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_39167cae9dfa2168, []int{2}
}

func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProtocolOptions.Unmarshal(m, b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
}
func (m *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(m, src)
}
func (m *ThriftProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_ThriftProtocolOptions.Size(m)
}
func (m *ThriftProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProtocolOptions proto.InternalMessageInfo

func (m *ThriftProtocolOptions) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.network.thrift_proxy.v4alpha.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.extensions.filters.network.thrift_proxy.v4alpha.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*ThriftProxy)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.ThriftProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/v4alpha/thrift_proxy.proto", fileDescriptor_39167cae9dfa2168)
}

var fileDescriptor_39167cae9dfa2168 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0xed, 0x34, 0x4d, 0xaf, 0x93, 0xc8, 0xdf, 0xe8, 0x43, 0x84, 0xf2, 0xa3, 0xd0, 0x55,
	0xd4, 0x85, 0x2d, 0xb5, 0xb0, 0xa0, 0xe2, 0x47, 0xb6, 0x93, 0x2a, 0x41, 0x6d, 0x6d, 0x0d, 0xae,
	0x0a, 0xab, 0xc8, 0x6d, 0x27, 0xa9, 0xc1, 0x78, 0xac, 0xf1, 0x24, 0x34, 0x3b, 0xc4, 0x8a, 0x67,
	0xe0, 0x25, 0xd8, 0xb3, 0x61, 0x85, 0x04, 0x4b, 0x5e, 0xa7, 0x2b, 0xe4, 0xb1, 0xd3, 0x3a, 0x61,
	0x85, 0x25, 0xd8, 0xf9, 0xde, 0x33, 0x3e, 0xe7, 0xcc, 0x99, 0x7b, 0xa1, 0x4f, 0xa2, 0x29, 0x9d,
	0x19, 0xe4, 0x82, 0x93, 0x28, 0x09, 0x68, 0x94, 0x18, 0xa3, 0x20, 0xe4, 0x84, 0x25, 0x46, 0x44,
	0xf8, 0x3b, 0xca, 0xde, 0x18, 0xfc, 0x9c, 0x05, 0x23, 0x3e, 0x8c, 0x19, 0xbd, 0x98, 0x19, 0xd3,
	0x07, 0x7e, 0x18, 0x9f, 0xfb, 0x0b, 0x4d, 0x3d, 0x66, 0x94, 0x53, 0xf4, 0x50, 0x30, 0xe9, 0xd7,
	0x4c, 0x7a, 0xce, 0xa4, 0xe7, 0x4c, 0xfa, 0xc2, 0x4f, 0x39, 0xd3, 0x86, 0x59, 0xce, 0x00, 0xa3,
	0x13, 0x4e, 0x32, 0xe5, 0x8d, 0x5b, 0x63, 0x4a, 0xc7, 0x21, 0x31, 0x44, 0x75, 0x32, 0x19, 0x19,
	0x7e, 0x94, 0x9b, 0xda, 0xb8, 0xb3, 0x0c, 0x25, 0x9c, 0x4d, 0x4e, 0x79, 0x8e, 0xde, 0x9d, 0x9c,
	0xc5, 0xbe, 0xe1, 0x47, 0x11, 0xe5, 0x3e, 0x17, 0xda, 0x09, 0xf7, 0xf9, 0x24, 0xc9, 0xe1, 0xfb,
	0xbf, 0xc1, 0x53, 0xc2, 0x52, 0x8f, 0x41, 0x34, 0xce, 0x8f, 0xdc, 0x9c, 0xfa, 0x61, 0x70, 0xe6,
	0x73, 0x62, 0xcc, 0x3f, 0x32, 0x60, 0xf3, 0x73, 0x05, 0x54, 0x4f, 0x18, 0x77, 0x53, 0xdf, 0x28,
	0x84, 0x75, 0xce, 0xfc, 0x28, 0x89, 0x29, 0xe3, 0x2d, 0xb9, 0x2d, 0x75, 0x9a, 0xdb, 0x5d, 0xbd,
	0x54, 0x62, 0xba, 0x37, 0xe7, 0xf1, 0x66, 0x31, 0xb1, 0x6a, 0x97, 0xd6, 0xea, 0x07, 0x49, 0xd6,
	0x24, 0x7c, 0x2d, 0x80, 0x02, 0xa8, 0x09, 0x1b, 0xa7, 0x34, 0x6c, 0x29, 0x42, 0xcc, 0x2e, 0x29,
	0xe6, 0xe6, 0x34, 0x4b, 0x5a, 0x57, 0xf4, 0xa8, 0x03, 0x6a, 0x1a, 0xda, 0x30, 0x66, 0x64, 0x14,
	0x5c, 0xb4, 0xa4, 0xb6, 0xd4, 0x59, 0xb7, 0xd6, 0x2e, 0xad, 0x0a, 0x93, 0xdb, 0x12, 0x86, 0x14,
	0x73, 0x05, 0x84, 0x42, 0xa8, 0x8b, 0x57, 0x1b, 0x9e, 0xd2, 0x68, 0x14, 0x8c, 0x5b, 0x95, 0xb6,
	0xd4, 0x51, 0xb7, 0x07, 0x25, 0x8d, 0xe1, 0x94, 0xca, 0x16, 0x4c, 0x13, 0x26, 0x9e, 0x07, 0xab,
	0xec, 0xba, 0x87, 0x5e, 0x43, 0x33, 0xff, 0x2f, 0xa7, 0x6b, 0xad, 0xb6, 0x95, 0x8e, 0x5a, 0x3a,
	0x88, 0xec, 0x31, 0xf7, 0xc4, 0x51, 0xdc, 0xe0, 0x85, 0x2a, 0xd9, 0xb5, 0x3f, 0x7d, 0xfb, 0x78,
	0xef, 0x29, 0x3c, 0xfe, 0x43, 0xe6, 0x1d, 0xbd, 0x30, 0x21, 0x9b, 0x3f, 0x24, 0xa8, 0x17, 0x45,
	0xd0, 0x6d, 0xa8, 0x44, 0xfe, 0x5b, 0xb2, 0x1c, 0xa9, 0x68, 0xa2, 0x47, 0x50, 0xe7, 0xb3, 0x98,
	0x9c, 0xcd, 0xc3, 0x54, 0x44, 0x98, 0xff, 0xeb, 0xd9, 0xbc, 0xeb, 0xf3, 0x79, 0xd7, 0xcd, 0x68,
	0xd6, 0x5f, 0xc1, 0xaa, 0x38, 0x9b, 0x25, 0xb3, 0xdb, 0x4d, 0xdd, 0x3e, 0x83, 0x27, 0x25, 0xdd,
	0x66, 0xee, 0xac, 0x06, 0xa8, 0x99, 0xf4, 0x30, 0xe5, 0x7e, 0x5e, 0xa9, 0xc9, 0x9a, 0x82, 0xab,
	0x59, 0x6b, 0xf3, 0xab, 0x0c, 0x37, 0xae, 0xee, 0x26, 0xe6, 0xc4, 0x89, 0xc5, 0x06, 0x2d, 0xee,
	0x81, 0xf4, 0x2f, 0xf7, 0x40, 0xfe, 0xab, 0x7b, 0xb0, 0x7b, 0x90, 0xa6, 0xda, 0x87, 0xbd, 0xf2,
	0x33, 0x50, 0xcc, 0x69, 0x6b, 0x00, 0x8d, 0x85, 0xfb, 0x21, 0x04, 0x4d, 0xf3, 0xc8, 0x73, 0x86,
	0x1e, 0x36, 0x0f, 0x5f, 0xb8, 0x0e, 0xf6, 0xb4, 0x15, 0x04, 0x50, 0xdd, 0xc3, 0xe6, 0x41, 0xaf,
	0xab, 0x49, 0xa8, 0x0e, 0xb5, 0xa3, 0xc3, 0xbc, 0x92, 0x53, 0xa4, 0xdf, 0x33, 0xbb, 0x3d, 0xac,
	0x29, 0x5b, 0xc7, 0x50, 0x2f, 0xba, 0x47, 0xff, 0x41, 0x43, 0x30, 0xb9, 0xd8, 0xf1, 0x1c, 0xdb,
	0xd9, 0xcf, 0x88, 0xac, 0xc1, 0xa1, 0x89, 0x5f, 0x69, 0x12, 0x6a, 0x02, 0xec, 0x9b, 0x2f, 0x87,
	0x79, 0x2d, 0x23, 0x15, 0xd6, 0x6c, 0xe7, 0xc0, 0x35, 0x6d, 0x4f, 0x53, 0xd2, 0xc2, 0x3b, 0x1e,
	0x78, 0x5e, 0x0f, 0x6b, 0x15, 0xcb, 0xff, 0xf2, 0xfe, 0xfb, 0xcf, 0xaa, 0xac, 0x29, 0x60, 0x07,
	0x34, 0xcb, 0x35, 0xbb, 0x55, 0xa9, 0x88, 0x2d, 0xad, 0xb0, 0x0d, 0xc2, 0xb0, 0x2b, 0x9d, 0x54,
	0x45, 0xbe, 0x3b, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xc7, 0xa4, 0xe5, 0xa7, 0x06, 0x00,
	0x00,
}
