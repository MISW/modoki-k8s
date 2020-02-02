// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generator.proto

package modoki

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OperateKind int32

const (
	OperateKind_Apply  OperateKind = 0
	OperateKind_Delete OperateKind = 1
)

var OperateKind_name = map[int32]string{
	0: "Apply",
	1: "Delete",
}

var OperateKind_value = map[string]int32{
	"Apply":  0,
	"Delete": 1,
}

func (x OperateKind) String() string {
	return proto.EnumName(OperateKind_name, int32(x))
}

func (OperateKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25df606994424d60, []int{0}
}

type KubernetesConfig struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesConfig) Reset()         { *m = KubernetesConfig{} }
func (m *KubernetesConfig) String() string { return proto.CompactTextString(m) }
func (*KubernetesConfig) ProtoMessage()    {}
func (*KubernetesConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_25df606994424d60, []int{0}
}

func (m *KubernetesConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesConfig.Unmarshal(m, b)
}
func (m *KubernetesConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesConfig.Marshal(b, m, deterministic)
}
func (m *KubernetesConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesConfig.Merge(m, src)
}
func (m *KubernetesConfig) XXX_Size() int {
	return xxx_messageInfo_KubernetesConfig.Size(m)
}
func (m *KubernetesConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesConfig.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesConfig proto.InternalMessageInfo

func (m *KubernetesConfig) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type YAML struct {
	Config               string   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YAML) Reset()         { *m = YAML{} }
func (m *YAML) String() string { return proto.CompactTextString(m) }
func (*YAML) ProtoMessage()    {}
func (*YAML) Descriptor() ([]byte, []int) {
	return fileDescriptor_25df606994424d60, []int{1}
}

func (m *YAML) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YAML.Unmarshal(m, b)
}
func (m *YAML) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YAML.Marshal(b, m, deterministic)
}
func (m *YAML) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YAML.Merge(m, src)
}
func (m *YAML) XXX_Size() int {
	return xxx_messageInfo_YAML.Size(m)
}
func (m *YAML) XXX_DiscardUnknown() {
	xxx_messageInfo_YAML.DiscardUnknown(m)
}

var xxx_messageInfo_YAML proto.InternalMessageInfo

func (m *YAML) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type OperateRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain               string            `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Kind                 OperateKind       `protobuf:"varint,3,opt,name=kind,proto3,enum=modoki.OperateKind" json:"kind,omitempty"`
	Spec                 *AppSpec          `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	Yaml                 *YAML             `protobuf:"bytes,5,opt,name=yaml,proto3" json:"yaml,omitempty"`
	K8SConfig            *KubernetesConfig `protobuf:"bytes,7,opt,name=k8s_config,json=k8sConfig,proto3" json:"k8s_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OperateRequest) Reset()         { *m = OperateRequest{} }
func (m *OperateRequest) String() string { return proto.CompactTextString(m) }
func (*OperateRequest) ProtoMessage()    {}
func (*OperateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25df606994424d60, []int{2}
}

func (m *OperateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperateRequest.Unmarshal(m, b)
}
func (m *OperateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperateRequest.Marshal(b, m, deterministic)
}
func (m *OperateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperateRequest.Merge(m, src)
}
func (m *OperateRequest) XXX_Size() int {
	return xxx_messageInfo_OperateRequest.Size(m)
}
func (m *OperateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperateRequest proto.InternalMessageInfo

func (m *OperateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OperateRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *OperateRequest) GetKind() OperateKind {
	if m != nil {
		return m.Kind
	}
	return OperateKind_Apply
}

func (m *OperateRequest) GetSpec() *AppSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *OperateRequest) GetYaml() *YAML {
	if m != nil {
		return m.Yaml
	}
	return nil
}

func (m *OperateRequest) GetK8SConfig() *KubernetesConfig {
	if m != nil {
		return m.K8SConfig
	}
	return nil
}

type OperateResponse struct {
	Yaml                 *YAML    `protobuf:"bytes,1,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperateResponse) Reset()         { *m = OperateResponse{} }
func (m *OperateResponse) String() string { return proto.CompactTextString(m) }
func (*OperateResponse) ProtoMessage()    {}
func (*OperateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25df606994424d60, []int{3}
}

func (m *OperateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperateResponse.Unmarshal(m, b)
}
func (m *OperateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperateResponse.Marshal(b, m, deterministic)
}
func (m *OperateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperateResponse.Merge(m, src)
}
func (m *OperateResponse) XXX_Size() int {
	return xxx_messageInfo_OperateResponse.Size(m)
}
func (m *OperateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OperateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OperateResponse proto.InternalMessageInfo

func (m *OperateResponse) GetYaml() *YAML {
	if m != nil {
		return m.Yaml
	}
	return nil
}

type MetricsRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsRequest) Reset()         { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()    {}
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25df606994424d60, []int{4}
}

func (m *MetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsRequest.Unmarshal(m, b)
}
func (m *MetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsRequest.Marshal(b, m, deterministic)
}
func (m *MetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsRequest.Merge(m, src)
}
func (m *MetricsRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsRequest.Size(m)
}
func (m *MetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsRequest proto.InternalMessageInfo

func (m *MetricsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MetricsResponse struct {
	Status               *AppStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MetricsResponse) Reset()         { *m = MetricsResponse{} }
func (m *MetricsResponse) String() string { return proto.CompactTextString(m) }
func (*MetricsResponse) ProtoMessage()    {}
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25df606994424d60, []int{5}
}

func (m *MetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsResponse.Unmarshal(m, b)
}
func (m *MetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsResponse.Marshal(b, m, deterministic)
}
func (m *MetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsResponse.Merge(m, src)
}
func (m *MetricsResponse) XXX_Size() int {
	return xxx_messageInfo_MetricsResponse.Size(m)
}
func (m *MetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsResponse proto.InternalMessageInfo

func (m *MetricsResponse) GetStatus() *AppStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("modoki.OperateKind", OperateKind_name, OperateKind_value)
	proto.RegisterType((*KubernetesConfig)(nil), "modoki.KubernetesConfig")
	proto.RegisterType((*YAML)(nil), "modoki.YAML")
	proto.RegisterType((*OperateRequest)(nil), "modoki.OperateRequest")
	proto.RegisterType((*OperateResponse)(nil), "modoki.OperateResponse")
	proto.RegisterType((*MetricsRequest)(nil), "modoki.MetricsRequest")
	proto.RegisterType((*MetricsResponse)(nil), "modoki.MetricsResponse")
}

func init() { proto.RegisterFile("generator.proto", fileDescriptor_25df606994424d60) }

var fileDescriptor_25df606994424d60 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x86, 0x6b, 0x1a, 0x42, 0x33, 0x54, 0x09, 0x75, 0x25, 0x1a, 0xa1, 0xaa, 0x8a, 0xd2, 0x4a,
	0xa5, 0x3d, 0xa0, 0x0a, 0x0e, 0x45, 0xa8, 0x17, 0xda, 0x4a, 0x3d, 0x50, 0x54, 0x29, 0x3d, 0xed,
	0x5e, 0x56, 0x21, 0x99, 0x45, 0x56, 0x88, 0xed, 0x8d, 0xcd, 0x81, 0xf3, 0x3e, 0xe9, 0xbe, 0xc9,
	0x0a, 0xc7, 0x01, 0x96, 0x15, 0x37, 0x7b, 0xe6, 0x9b, 0x99, 0xff, 0x1f, 0x0d, 0x04, 0x6b, 0xe4,
	0x58, 0xa5, 0x5a, 0x54, 0x23, 0x59, 0x09, 0x2d, 0xa8, 0x5b, 0x8a, 0x5c, 0x14, 0x6c, 0xe0, 0xa5,
	0x52, 0xd6, 0xa1, 0xf8, 0x1b, 0xf4, 0x16, 0xdb, 0x15, 0x56, 0x1c, 0x35, 0xaa, 0x5f, 0x82, 0xdf,
	0xb2, 0x35, 0x7d, 0x0f, 0x1e, 0x4f, 0x4b, 0x54, 0x32, 0xcd, 0x30, 0x24, 0x11, 0x19, 0x7a, 0xc9,
	0x31, 0x10, 0x7f, 0x00, 0xe7, 0x6a, 0xbe, 0xfc, 0x4b, 0xfb, 0xe0, 0x66, 0x86, 0xb7, 0x88, 0xfd,
	0xc5, 0x0f, 0x04, 0xfc, 0x7f, 0x72, 0x3f, 0x17, 0x13, 0xbc, 0xdb, 0xa2, 0xd2, 0xd4, 0x87, 0x16,
	0xcb, 0x2d, 0xd6, 0x62, 0xf9, 0xbe, 0x34, 0x17, 0x65, 0xca, 0x78, 0xd8, 0xaa, 0x4b, 0xeb, 0x1f,
	0xfd, 0x0c, 0x4e, 0xc1, 0x78, 0x1e, 0xbe, 0x8c, 0xc8, 0xd0, 0x1f, 0xbf, 0x1d, 0xd5, 0x72, 0x47,
	0xb6, 0xdb, 0x82, 0xf1, 0x3c, 0x31, 0x00, 0xfd, 0x08, 0x8e, 0x92, 0x98, 0x85, 0x4e, 0x44, 0x86,
	0xdd, 0x71, 0xd0, 0x80, 0x73, 0x29, 0xff, 0x4b, 0xcc, 0x12, 0x93, 0xa4, 0x11, 0x38, 0xbb, 0xb4,
	0xdc, 0x84, 0x6d, 0x03, 0xbd, 0x6e, 0xa0, 0xbd, 0xf8, 0xc4, 0x64, 0xe8, 0x77, 0x80, 0x62, 0xaa,
	0x6e, 0xac, 0x8d, 0x8e, 0xe1, 0xc2, 0x86, 0x3b, 0x5f, 0x4b, 0xe2, 0x15, 0x53, 0xfb, 0x8c, 0x27,
	0x10, 0x1c, 0x2c, 0x2a, 0x29, 0xb8, 0xc2, 0xc3, 0x34, 0x72, 0x69, 0x5a, 0x1c, 0x81, 0xbf, 0x44,
	0x5d, 0xb1, 0x4c, 0x5d, 0xd8, 0x4b, 0xfc, 0x03, 0x82, 0x03, 0x61, 0xdb, 0x7e, 0x01, 0x57, 0xe9,
	0x54, 0x6f, 0x95, 0x6d, 0xfc, 0xe6, 0xd4, 0xab, 0x49, 0x24, 0x16, 0xf8, 0xfa, 0x09, 0xba, 0x27,
	0x9b, 0xa2, 0x1e, 0xb4, 0xe7, 0x52, 0x6e, 0x76, 0xbd, 0x17, 0x14, 0xc0, 0xfd, 0x8d, 0x1b, 0xd4,
	0xd8, 0x23, 0xe3, 0x7b, 0x02, 0xde, 0x9f, 0xe6, 0x2e, 0xe8, 0x0c, 0x3a, 0xb6, 0x86, 0xf6, 0xcf,
	0xd6, 0x6d, 0x45, 0x0e, 0xde, 0x3d, 0x8b, 0x5b, 0x69, 0x33, 0xe8, 0x58, 0xb5, 0xc7, 0xda, 0xa7,
	0x06, 0x8f, 0xb5, 0x67, 0xb6, 0x7e, 0xbe, 0xba, 0xb6, 0xb7, 0xb8, 0x72, 0xcd, 0x1d, 0x4e, 0x1e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x95, 0xde, 0x94, 0x45, 0xad, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeneratorClient is the client API for Generator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeneratorClient interface {
	Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error)
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type generatorClient struct {
	cc *grpc.ClientConn
}

func NewGeneratorClient(cc *grpc.ClientConn) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error) {
	out := new(OperateResponse)
	err := c.cc.Invoke(ctx, "/modoki.Generator/Operate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/modoki.Generator/Metrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneratorServer is the server API for Generator service.
type GeneratorServer interface {
	Operate(context.Context, *OperateRequest) (*OperateResponse, error)
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
}

// UnimplementedGeneratorServer can be embedded to have forward compatible implementations.
type UnimplementedGeneratorServer struct {
}

func (*UnimplementedGeneratorServer) Operate(ctx context.Context, req *OperateRequest) (*OperateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operate not implemented")
}
func (*UnimplementedGeneratorServer) Metrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}

func RegisterGeneratorServer(s *grpc.Server, srv GeneratorServer) {
	s.RegisterService(&_Generator_serviceDesc, srv)
}

func _Generator_Operate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).Operate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Generator/Operate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).Operate(ctx, req.(*OperateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generator_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Generator/Metrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).Metrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modoki.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Operate",
			Handler:    _Generator_Operate_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _Generator_Metrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generator.proto",
}
