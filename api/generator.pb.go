// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generator.proto

package modoki

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
func (OperateKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type KubernetesConfig struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *KubernetesConfig) Reset()                    { *m = KubernetesConfig{} }
func (m *KubernetesConfig) String() string            { return proto.CompactTextString(m) }
func (*KubernetesConfig) ProtoMessage()               {}
func (*KubernetesConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *KubernetesConfig) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type YAML struct {
	Config string `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
}

func (m *YAML) Reset()                    { *m = YAML{} }
func (m *YAML) String() string            { return proto.CompactTextString(m) }
func (*YAML) ProtoMessage()               {}
func (*YAML) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *YAML) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type OperateRequest struct {
	Id         string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Kind       OperateKind       `protobuf:"varint,2,opt,name=kind,enum=modoki.OperateKind" json:"kind,omitempty"`
	Performer  int32             `protobuf:"varint,3,opt,name=performer" json:"performer,omitempty"`
	Spec       *AppSpec          `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	DeleteYaml *YAML             `protobuf:"bytes,5,opt,name=delete_yaml,json=deleteYaml" json:"delete_yaml,omitempty"`
	ApplyYaml  *YAML             `protobuf:"bytes,6,opt,name=apply_yaml,json=applyYaml" json:"apply_yaml,omitempty"`
	K8SConfig  *KubernetesConfig `protobuf:"bytes,7,opt,name=k8s_config,json=k8sConfig" json:"k8s_config,omitempty"`
}

func (m *OperateRequest) Reset()                    { *m = OperateRequest{} }
func (m *OperateRequest) String() string            { return proto.CompactTextString(m) }
func (*OperateRequest) ProtoMessage()               {}
func (*OperateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *OperateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OperateRequest) GetKind() OperateKind {
	if m != nil {
		return m.Kind
	}
	return OperateKind_Apply
}

func (m *OperateRequest) GetPerformer() int32 {
	if m != nil {
		return m.Performer
	}
	return 0
}

func (m *OperateRequest) GetSpec() *AppSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *OperateRequest) GetDeleteYaml() *YAML {
	if m != nil {
		return m.DeleteYaml
	}
	return nil
}

func (m *OperateRequest) GetApplyYaml() *YAML {
	if m != nil {
		return m.ApplyYaml
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
	DeleteYaml *YAML `protobuf:"bytes,1,opt,name=delete_yaml,json=deleteYaml" json:"delete_yaml,omitempty"`
	ApplyYaml  *YAML `protobuf:"bytes,2,opt,name=apply_yaml,json=applyYaml" json:"apply_yaml,omitempty"`
}

func (m *OperateResponse) Reset()                    { *m = OperateResponse{} }
func (m *OperateResponse) String() string            { return proto.CompactTextString(m) }
func (*OperateResponse) ProtoMessage()               {}
func (*OperateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *OperateResponse) GetDeleteYaml() *YAML {
	if m != nil {
		return m.DeleteYaml
	}
	return nil
}

func (m *OperateResponse) GetApplyYaml() *YAML {
	if m != nil {
		return m.ApplyYaml
	}
	return nil
}

type MetricsRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *MetricsRequest) Reset()                    { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()               {}
func (*MetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *MetricsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MetricsResponse struct {
	Status *AppStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *MetricsResponse) Reset()                    { *m = MetricsResponse{} }
func (m *MetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*MetricsResponse) ProtoMessage()               {}
func (*MetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *MetricsResponse) GetStatus() *AppStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*KubernetesConfig)(nil), "modoki.KubernetesConfig")
	proto.RegisterType((*YAML)(nil), "modoki.YAML")
	proto.RegisterType((*OperateRequest)(nil), "modoki.OperateRequest")
	proto.RegisterType((*OperateResponse)(nil), "modoki.OperateResponse")
	proto.RegisterType((*MetricsRequest)(nil), "modoki.MetricsRequest")
	proto.RegisterType((*MetricsResponse)(nil), "modoki.MetricsResponse")
	proto.RegisterEnum("modoki.OperateKind", OperateKind_name, OperateKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Generator service

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
	err := grpc.Invoke(ctx, "/modoki.Generator/Operate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generatorClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := grpc.Invoke(ctx, "/modoki.Generator/Metrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Generator service

type GeneratorServer interface {
	Operate(context.Context, *OperateRequest) (*OperateResponse, error)
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
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

func init() { proto.RegisterFile("generator.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x8b, 0xd4, 0x30,
	0x14, 0xc0, 0x4d, 0x9d, 0xe9, 0xd8, 0x37, 0xd2, 0x8e, 0x11, 0xd6, 0xb2, 0x88, 0x94, 0x2a, 0x58,
	0x15, 0x07, 0x19, 0x0f, 0x2e, 0x8b, 0x97, 0x51, 0xc1, 0xc3, 0xba, 0x08, 0xf1, 0xb4, 0x5e, 0x96,
	0x6c, 0xfb, 0x76, 0x29, 0xfd, 0x93, 0x98, 0x64, 0x0e, 0x7b, 0xf6, 0x93, 0xf8, 0x4d, 0xa5, 0x69,
	0xba, 0x1d, 0xeb, 0x1f, 0xf0, 0xd6, 0xbe, 0xf7, 0x7b, 0xc9, 0xef, 0xbd, 0x24, 0x10, 0x5d, 0x61,
	0x8b, 0x8a, 0x1b, 0xa1, 0xd6, 0x52, 0x09, 0x23, 0xa8, 0xdf, 0x88, 0x42, 0x54, 0xe5, 0x61, 0xc0,
	0xa5, 0xec, 0x43, 0xe9, 0x2b, 0x58, 0x9d, 0xec, 0x2e, 0x50, 0xb5, 0x68, 0x50, 0xbf, 0x17, 0xed,
	0x65, 0x79, 0x45, 0x1f, 0x42, 0xd0, 0xf2, 0x06, 0xb5, 0xe4, 0x39, 0xc6, 0x24, 0x21, 0x59, 0xc0,
	0xc6, 0x40, 0xfa, 0x08, 0x66, 0x67, 0xdb, 0xd3, 0x4f, 0xf4, 0x00, 0xfc, 0xdc, 0xf2, 0x0e, 0x71,
	0x7f, 0xe9, 0x0f, 0x0f, 0xc2, 0xcf, 0xb2, 0xdb, 0x17, 0x19, 0x7e, 0xdb, 0xa1, 0x36, 0x34, 0x04,
	0xaf, 0x2c, 0x1c, 0xe6, 0x95, 0x05, 0x7d, 0x0a, 0xb3, 0xaa, 0x6c, 0x8b, 0xd8, 0x4b, 0x48, 0x16,
	0x6e, 0xee, 0xaf, 0x7b, 0xad, 0xb5, 0xab, 0x3a, 0x29, 0xdb, 0x82, 0x59, 0xa0, 0x33, 0x91, 0xa8,
	0x2e, 0x85, 0x6a, 0x50, 0xc5, 0xb7, 0x13, 0x92, 0xcd, 0xd9, 0x18, 0xa0, 0x8f, 0x61, 0xa6, 0x25,
	0xe6, 0xf1, 0x2c, 0x21, 0xd9, 0x72, 0x13, 0x0d, 0xcb, 0x6c, 0xa5, 0xfc, 0x22, 0x31, 0x67, 0x36,
	0x49, 0x5f, 0xc2, 0xb2, 0xc0, 0x1a, 0x0d, 0x9e, 0x5f, 0xf3, 0xa6, 0x8e, 0xe7, 0x96, 0xbd, 0x3b,
	0xb0, 0x5d, 0x27, 0x0c, 0x7a, 0xe0, 0x8c, 0x37, 0x35, 0x7d, 0x01, 0xc0, 0xa5, 0xac, 0xaf, 0x7b,
	0xda, 0xff, 0x03, 0x1d, 0xd8, 0xbc, 0x85, 0xdf, 0x00, 0x54, 0x47, 0xfa, 0xdc, 0x8d, 0x61, 0x61,
	0xe1, 0x78, 0x80, 0xa7, 0x63, 0x65, 0x41, 0x75, 0xe4, 0x3e, 0xd3, 0x06, 0xa2, 0x9b, 0x11, 0x69,
	0x29, 0x5a, 0x8d, 0x53, 0x4f, 0xf2, 0x5f, 0x9e, 0xde, 0x3f, 0x3d, 0xd3, 0x04, 0xc2, 0x53, 0x34,
	0xaa, 0xcc, 0xf5, 0x5f, 0x4e, 0x24, 0x7d, 0x0b, 0xd1, 0x0d, 0xe1, 0x84, 0x9e, 0x81, 0xaf, 0x0d,
	0x37, 0x3b, 0xed, 0x5c, 0xee, 0xed, 0xcf, 0xd7, 0x26, 0x98, 0x03, 0x9e, 0x3f, 0x81, 0xe5, 0xde,
	0xd9, 0xd1, 0x00, 0xe6, 0xdb, 0x6e, 0xef, 0xd5, 0x2d, 0x0a, 0xe0, 0x7f, 0xb0, 0xd2, 0x2b, 0xb2,
	0xf9, 0x4e, 0x20, 0xf8, 0x38, 0xdc, 0x48, 0x7a, 0x0c, 0x0b, 0x57, 0x43, 0x0f, 0x26, 0x17, 0xc0,
	0x49, 0x1e, 0x3e, 0xf8, 0x2d, 0xee, 0xd4, 0x8e, 0x61, 0xe1, 0x6c, 0xc7, 0xda, 0x5f, 0x1b, 0x1c,
	0x6b, 0x27, 0x6d, 0xbd, 0xbb, 0xf3, 0xd5, 0xbd, 0x82, 0x0b, 0xdf, 0xbe, 0x80, 0xd7, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0xb9, 0x3c, 0x54, 0x27, 0x03, 0x00, 0x00,
}
