// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

/*
Package modoki is a generated protocol buffer package.

It is generated from these files:
	app.proto
	auth.proto
	generator.proto
	user_org.proto

It has these top-level messages:
	AppSpec
	AppStatus
	AppCreateRequest
	AppCreateResponse
	AppApplyRequest
	AppApplyResponse
	AppDeleteRequest
	AppDeleteResponse
	AppStatusRequest
	AppStatusResponse
	SignInRequest
	SignInResponse
	SignOutRequest
	SignOutResponse
	CallbackRequest
	CallbackResponse
	IssueTokenRequest
	IssueTokenResponse
	KubernetesConfig
	YAML
	OperateRequest
	OperateResponse
	MetricsRequest
	MetricsResponse
	User
	Organization
	UserAddRequest
	UserAddResponse
	UserDeleteRequest
	UserDeleteResponse
	UserFindByIDRequest
	UserFindByIDResponse
	OrganizationAddRequest
	OrganizationAddResponse
	OrganizationDeleteRequest
	OrganizationDeleteResponse
	OrganizationInviteUserRequest
	OrganizationInviteUserResponse
	OrganizationListUserRequest
	OrganizationListUserResponse
*/
package modoki

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AppSpec struct {
	Owner       int32             `protobuf:"varint,1,opt,name=owner" json:"owner,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Domain      string            `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	Image       string            `protobuf:"bytes,4,opt,name=image" json:"image,omitempty"`
	Command     []string          `protobuf:"bytes,5,rep,name=command" json:"command,omitempty"`
	Args        []string          `protobuf:"bytes,6,rep,name=args" json:"args,omitempty"`
	Env         map[string]string `protobuf:"bytes,7,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Annotations map[string]string `protobuf:"bytes,10,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Labels      map[string]string `protobuf:"bytes,11,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AppSpec) Reset()                    { *m = AppSpec{} }
func (m *AppSpec) String() string            { return proto.CompactTextString(m) }
func (*AppSpec) ProtoMessage()               {}
func (*AppSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AppSpec) GetOwner() int32 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *AppSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppSpec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *AppSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *AppSpec) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *AppSpec) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *AppSpec) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *AppSpec) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *AppSpec) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type AppStatus struct {
	Id         int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AppId      string                     `protobuf:"bytes,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	State      string                     `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	StartedAt  *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	ExitCode   int32                      `protobuf:"varint,5,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	CreatedAt  *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt  *google_protobuf.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Attributes map[string]string          `protobuf:"bytes,12,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AppStatus) Reset()                    { *m = AppStatus{} }
func (m *AppStatus) String() string            { return proto.CompactTextString(m) }
func (*AppStatus) ProtoMessage()               {}
func (*AppStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppStatus) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppStatus) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *AppStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *AppStatus) GetStartedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *AppStatus) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *AppStatus) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *AppStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *AppStatus) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type AppCreateRequest struct {
	Spec *AppSpec `protobuf:"bytes,1,opt,name=spec" json:"spec,omitempty"`
}

func (m *AppCreateRequest) Reset()                    { *m = AppCreateRequest{} }
func (m *AppCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AppCreateRequest) ProtoMessage()               {}
func (*AppCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppCreateRequest) GetSpec() *AppSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type AppCreateResponse struct {
	Id   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec *AppSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
}

func (m *AppCreateResponse) Reset()                    { *m = AppCreateResponse{} }
func (m *AppCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*AppCreateResponse) ProtoMessage()               {}
func (*AppCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AppCreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AppCreateResponse) GetSpec() *AppSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type AppApplyRequest struct {
	Spec *AppSpec `protobuf:"bytes,1,opt,name=spec" json:"spec,omitempty"`
}

func (m *AppApplyRequest) Reset()                    { *m = AppApplyRequest{} }
func (m *AppApplyRequest) String() string            { return proto.CompactTextString(m) }
func (*AppApplyRequest) ProtoMessage()               {}
func (*AppApplyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AppApplyRequest) GetSpec() *AppSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type AppApplyResponse struct {
}

func (m *AppApplyResponse) Reset()                    { *m = AppApplyResponse{} }
func (m *AppApplyResponse) String() string            { return proto.CompactTextString(m) }
func (*AppApplyResponse) ProtoMessage()               {}
func (*AppApplyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AppDeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *AppDeleteRequest) Reset()                    { *m = AppDeleteRequest{} }
func (m *AppDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AppDeleteRequest) ProtoMessage()               {}
func (*AppDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AppDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AppDeleteResponse struct {
}

func (m *AppDeleteResponse) Reset()                    { *m = AppDeleteResponse{} }
func (m *AppDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*AppDeleteResponse) ProtoMessage()               {}
func (*AppDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AppStatusRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *AppStatusRequest) Reset()                    { *m = AppStatusRequest{} }
func (m *AppStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*AppStatusRequest) ProtoMessage()               {}
func (*AppStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AppStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AppStatusResponse struct {
	Status *AppStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *AppStatusResponse) Reset()                    { *m = AppStatusResponse{} }
func (m *AppStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*AppStatusResponse) ProtoMessage()               {}
func (*AppStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AppStatusResponse) GetStatus() *AppStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AppSpec)(nil), "modoki.AppSpec")
	proto.RegisterType((*AppStatus)(nil), "modoki.AppStatus")
	proto.RegisterType((*AppCreateRequest)(nil), "modoki.AppCreateRequest")
	proto.RegisterType((*AppCreateResponse)(nil), "modoki.AppCreateResponse")
	proto.RegisterType((*AppApplyRequest)(nil), "modoki.AppApplyRequest")
	proto.RegisterType((*AppApplyResponse)(nil), "modoki.AppApplyResponse")
	proto.RegisterType((*AppDeleteRequest)(nil), "modoki.AppDeleteRequest")
	proto.RegisterType((*AppDeleteResponse)(nil), "modoki.AppDeleteResponse")
	proto.RegisterType((*AppStatusRequest)(nil), "modoki.AppStatusRequest")
	proto.RegisterType((*AppStatusResponse)(nil), "modoki.AppStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for App service

type AppClient interface {
	Create(ctx context.Context, in *AppCreateRequest, opts ...grpc.CallOption) (*AppCreateResponse, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Create(ctx context.Context, in *AppCreateRequest, opts ...grpc.CallOption) (*AppCreateResponse, error) {
	out := new(AppCreateResponse)
	err := grpc.Invoke(ctx, "/modoki.App/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppServer interface {
	Create(context.Context, *AppCreateRequest) (*AppCreateResponse, error)
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.App/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Create(ctx, req.(*AppCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modoki.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _App_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0x95, 0xb8, 0x71, 0xeb, 0xf1, 0xa3, 0xa7, 0xe9, 0xf2, 0xa2, 0xc5, 0x3d, 0x10, 0xcc,
	0x25, 0x70, 0x70, 0xa5, 0x54, 0x2a, 0x14, 0xa9, 0x95, 0xdc, 0x17, 0x09, 0x24, 0x4e, 0x86, 0x13,
	0x97, 0x68, 0x13, 0x2f, 0x91, 0x55, 0xdb, 0xbb, 0x78, 0xd7, 0x81, 0x7c, 0x13, 0xbe, 0x21, 0x5f,
	0x03, 0x79, 0x76, 0x4d, 0x5d, 0xf3, 0x52, 0xe5, 0xb6, 0x33, 0xfb, 0xff, 0x8d, 0xff, 0x3b, 0x33,
	0x32, 0x78, 0x4c, 0xca, 0x48, 0x56, 0x42, 0x0b, 0xe2, 0x16, 0x22, 0x15, 0x37, 0x59, 0xf0, 0x74,
	0x25, 0xc4, 0x2a, 0xe7, 0x47, 0x98, 0x5d, 0xd4, 0x9f, 0x8f, 0x74, 0x56, 0x70, 0xa5, 0x59, 0x61,
	0x85, 0xe1, 0x0f, 0x07, 0x76, 0x63, 0x29, 0x3f, 0x48, 0xbe, 0x24, 0x0f, 0x61, 0x24, 0xbe, 0x96,
	0xbc, 0xa2, 0x83, 0xc9, 0x60, 0x3a, 0x4a, 0x4c, 0x40, 0x08, 0xec, 0x94, 0xac, 0xe0, 0x74, 0x38,
	0x19, 0x4c, 0xbd, 0x04, 0xcf, 0xe4, 0x31, 0xb8, 0xa9, 0x28, 0x58, 0x56, 0x52, 0x07, 0xb3, 0x36,
	0x6a, 0x2a, 0x64, 0x05, 0x5b, 0x71, 0xba, 0x83, 0x69, 0x13, 0x10, 0x0a, 0xbb, 0x4b, 0x51, 0x14,
	0xac, 0x4c, 0xe9, 0x68, 0xe2, 0x4c, 0xbd, 0xa4, 0x0d, 0x9b, 0xda, 0xac, 0x5a, 0x29, 0xea, 0x62,
	0x1a, 0xcf, 0xe4, 0x25, 0x38, 0xbc, 0x5c, 0xd3, 0xdd, 0x89, 0x33, 0xf5, 0x67, 0x34, 0x32, 0x0f,
	0x89, 0xac, 0xc7, 0xe8, 0xba, 0x5c, 0x5f, 0x97, 0xba, 0xda, 0x24, 0x8d, 0x88, 0x5c, 0x80, 0xcf,
	0xca, 0x52, 0x68, 0xa6, 0x33, 0x51, 0x2a, 0x0a, 0xc8, 0x4c, 0xfa, 0x4c, 0x7c, 0x2b, 0x31, 0x6c,
	0x17, 0x22, 0xc7, 0xe0, 0xe6, 0x6c, 0xc1, 0x73, 0x45, 0x7d, 0xc4, 0x0f, 0xfb, 0xf8, 0x7b, 0xbc,
	0x35, 0xa4, 0x95, 0x06, 0x27, 0xb0, 0xd7, 0x3a, 0x21, 0x63, 0x70, 0x6e, 0xf8, 0x06, 0x9b, 0xe6,
	0x25, 0xcd, 0xb1, 0x69, 0xc3, 0x9a, 0xe5, 0x75, 0xdb, 0x33, 0x13, 0xbc, 0x19, 0xbe, 0x1e, 0x04,
	0xe7, 0x30, 0xee, 0xbb, 0xd9, 0x8a, 0x3f, 0x05, 0xbf, 0x63, 0x67, 0x1b, 0x34, 0xfc, 0xee, 0x80,
	0xd7, 0x3c, 0x49, 0x33, 0x5d, 0x2b, 0xf2, 0x3f, 0x0c, 0xb3, 0xd4, 0x0e, 0x7a, 0x98, 0xa5, 0xe4,
	0x11, 0xb8, 0x4c, 0xca, 0x79, 0x96, 0xb6, 0x20, 0x93, 0xf2, 0x5d, 0xda, 0x94, 0x53, 0x9a, 0x69,
	0x6e, 0xe7, 0x6c, 0x02, 0x72, 0x0a, 0xa0, 0x34, 0xab, 0x34, 0x4f, 0xe7, 0x4c, 0xe3, 0xac, 0xfd,
	0x59, 0x10, 0x99, 0x55, 0x8b, 0xda, 0x55, 0x8b, 0x3e, 0xb6, 0xab, 0x96, 0x78, 0x56, 0x1d, 0x6b,
	0x72, 0x08, 0x1e, 0xff, 0x96, 0xe9, 0xf9, 0x52, 0xa4, 0x9c, 0x8e, 0xf0, 0xf3, 0x7b, 0x4d, 0xe2,
	0x52, 0xa4, 0x58, 0x77, 0x59, 0x71, 0x66, 0xeb, 0xc2, 0xfd, 0x75, 0xad, 0x3a, 0xd6, 0x0d, 0x5a,
	0xcb, 0xb4, 0x45, 0xfd, 0xfb, 0x51, 0xab, 0x8e, 0x35, 0x89, 0x01, 0x98, 0xd6, 0x55, 0xb6, 0xa8,
	0x35, 0x57, 0xf4, 0x3f, 0x5c, 0x82, 0x67, 0xdd, 0x25, 0xc0, 0x8e, 0x45, 0xf1, 0x2f, 0x8d, 0x59,
	0x85, 0x0e, 0x14, 0x9c, 0xc1, 0x7e, 0xef, 0x7a, 0xab, 0xd1, 0xbc, 0x82, 0x71, 0x2c, 0xe5, 0x25,
	0x3e, 0x26, 0xe1, 0x5f, 0x6a, 0xae, 0x34, 0x79, 0x0e, 0x3b, 0x4a, 0xf2, 0x25, 0x16, 0xf0, 0x67,
	0xfb, 0xbd, 0xa5, 0x4c, 0xf0, 0x32, 0x7c, 0x0b, 0x07, 0x1d, 0x50, 0x49, 0x51, 0x2a, 0xde, 0x19,
	0xad, 0x87, 0xa3, 0x6d, 0x2b, 0x0d, 0xff, 0x55, 0xe9, 0x04, 0xf6, 0x63, 0x29, 0x63, 0x29, 0xf3,
	0xcd, 0x56, 0x0e, 0x08, 0x5a, 0xb7, 0x9c, 0x31, 0x10, 0x86, 0x98, 0xbb, 0xe2, 0x39, 0xbf, 0x7d,
	0x4e, 0xcf, 0x54, 0xf8, 0x00, 0x9d, 0xb7, 0x9a, 0x3b, 0xa0, 0xe9, 0xf7, 0xdf, 0xc0, 0x73, 0x04,
	0x5b, 0x8d, 0x7d, 0xf2, 0x0b, 0x70, 0x15, 0x66, 0xac, 0xd9, 0x83, 0xdf, 0xc6, 0x97, 0x58, 0xc1,
	0xec, 0x0a, 0x9c, 0x58, 0x4a, 0x72, 0x06, 0xae, 0x69, 0x1b, 0xe9, 0xfe, 0x62, 0xee, 0x8c, 0x20,
	0x78, 0xf2, 0x87, 0x1b, 0xf3, 0xc1, 0x8b, 0xbd, 0x4f, 0xf6, 0x0f, 0xbb, 0x70, 0x71, 0xb9, 0x8e,
	0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x22, 0xca, 0xa8, 0x14, 0x7d, 0x05, 0x00, 0x00,
}