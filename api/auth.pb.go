// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

/*
Package modoki is a generated protocol buffer package.

It is generated from these files:
	auth.proto
	generator.proto
	service.proto
	user_org.proto

It has these top-level messages:
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
	AppSpec
	AppStatus
	AppCreateRequest
	AppCreateResponse
	User
	Organization
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

type SignInRequest struct {
}

func (m *SignInRequest) Reset()                    { *m = SignInRequest{} }
func (m *SignInRequest) String() string            { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()               {}
func (*SignInRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SignInResponse struct {
	RedirectUri string `protobuf:"bytes,1,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
}

func (m *SignInResponse) Reset()                    { *m = SignInResponse{} }
func (m *SignInResponse) String() string            { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()               {}
func (*SignInResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignInResponse) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

type SignOutRequest struct {
}

func (m *SignOutRequest) Reset()                    { *m = SignOutRequest{} }
func (m *SignOutRequest) String() string            { return proto.CompactTextString(m) }
func (*SignOutRequest) ProtoMessage()               {}
func (*SignOutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SignOutResponse struct {
}

func (m *SignOutResponse) Reset()                    { *m = SignOutResponse{} }
func (m *SignOutResponse) String() string            { return proto.CompactTextString(m) }
func (*SignOutResponse) ProtoMessage()               {}
func (*SignOutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CallbackRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=AccessToken" json:"AccessToken,omitempty"`
	IDToken     string `protobuf:"bytes,2,opt,name=IDToken" json:"IDToken,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=State" json:"State,omitempty"`
}

func (m *CallbackRequest) Reset()                    { *m = CallbackRequest{} }
func (m *CallbackRequest) String() string            { return proto.CompactTextString(m) }
func (*CallbackRequest) ProtoMessage()               {}
func (*CallbackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CallbackRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CallbackRequest) GetIDToken() string {
	if m != nil {
		return m.IDToken
	}
	return ""
}

func (m *CallbackRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type CallbackResponse struct {
}

func (m *CallbackResponse) Reset()                    { *m = CallbackResponse{} }
func (m *CallbackResponse) String() string            { return proto.CompactTextString(m) }
func (*CallbackResponse) ProtoMessage()               {}
func (*CallbackResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type IssueTokenRequest struct {
	Owner int32  `protobuf:"varint,1,opt,name=owner" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *IssueTokenRequest) Reset()                    { *m = IssueTokenRequest{} }
func (m *IssueTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*IssueTokenRequest) ProtoMessage()               {}
func (*IssueTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IssueTokenRequest) GetOwner() int32 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *IssueTokenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IssueTokenResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *IssueTokenResponse) Reset()                    { *m = IssueTokenResponse{} }
func (m *IssueTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*IssueTokenResponse) ProtoMessage()               {}
func (*IssueTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IssueTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*SignInRequest)(nil), "modoki.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "modoki.SignInResponse")
	proto.RegisterType((*SignOutRequest)(nil), "modoki.SignOutRequest")
	proto.RegisterType((*SignOutResponse)(nil), "modoki.SignOutResponse")
	proto.RegisterType((*CallbackRequest)(nil), "modoki.CallbackRequest")
	proto.RegisterType((*CallbackResponse)(nil), "modoki.CallbackResponse")
	proto.RegisterType((*IssueTokenRequest)(nil), "modoki.IssueTokenRequest")
	proto.RegisterType((*IssueTokenResponse)(nil), "modoki.IssueTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authority service

type AuthorityClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
	IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error)
}

type authorityClient struct {
	cc *grpc.ClientConn
}

func NewAuthorityClient(cc *grpc.ClientConn) AuthorityClient {
	return &authorityClient{cc}
}

func (c *authorityClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := grpc.Invoke(ctx, "/modoki.Authority/SignIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := grpc.Invoke(ctx, "/modoki.Authority/SignOut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := grpc.Invoke(ctx, "/modoki.Authority/Callback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityClient) IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error) {
	out := new(IssueTokenResponse)
	err := grpc.Invoke(ctx, "/modoki.Authority/IssueToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authority service

type AuthorityServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	Callback(context.Context, *CallbackRequest) (*CallbackResponse, error)
	IssueToken(context.Context, *IssueTokenRequest) (*IssueTokenResponse, error)
}

func RegisterAuthorityServer(s *grpc.Server, srv AuthorityServer) {
	s.RegisterService(&_Authority_serviceDesc, srv)
}

func _Authority_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Authority/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authority_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Authority/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authority_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Authority/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authority_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modoki.Authority/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServer).IssueToken(ctx, req.(*IssueTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "modoki.Authority",
	HandlerType: (*AuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _Authority_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _Authority_SignOut_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _Authority_Callback_Handler,
		},
		{
			MethodName: "IssueToken",
			Handler:    _Authority_IssueToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbb, 0x4f, 0x32, 0x41,
	0x14, 0xc5, 0x03, 0xdf, 0xb7, 0x3c, 0x2e, 0x2a, 0x70, 0x83, 0xba, 0x6e, 0x85, 0x53, 0x19, 0x0b,
	0x0a, 0x29, 0x4c, 0x4c, 0x28, 0x10, 0x1b, 0x2a, 0x13, 0xd0, 0xc6, 0xc6, 0x2c, 0xcb, 0x44, 0x26,
	0xc0, 0x0c, 0xce, 0x23, 0xc6, 0xde, 0x3f, 0xdc, 0x64, 0x1e, 0xec, 0xc2, 0xda, 0x71, 0xcf, 0xbd,
	0xe7, 0x1c, 0xf6, 0x97, 0x01, 0x48, 0x8d, 0x5e, 0x0d, 0x76, 0x52, 0x68, 0x81, 0xb5, 0xad, 0x58,
	0x8a, 0x35, 0x23, 0x6d, 0x38, 0x9d, 0xb3, 0x0f, 0x3e, 0xe5, 0x33, 0xfa, 0x69, 0xa8, 0xd2, 0x64,
	0x08, 0x67, 0x41, 0x50, 0x3b, 0xc1, 0x15, 0xc5, 0x6b, 0x38, 0x91, 0x74, 0xc9, 0x24, 0xcd, 0xf4,
	0xbb, 0x91, 0x2c, 0xae, 0xf4, 0x2b, 0x37, 0xcd, 0x59, 0x2b, 0x68, 0xaf, 0x92, 0x91, 0x8e, 0x33,
	0x3d, 0x1b, 0x1d, 0x62, 0xba, 0xd0, 0xde, 0x2b, 0x2e, 0x87, 0x64, 0xd0, 0x9e, 0xa4, 0x9b, 0xcd,
	0x22, 0xcd, 0xd6, 0xfe, 0x0a, 0xfb, 0xd0, 0x1a, 0x67, 0x19, 0x55, 0xea, 0x45, 0xac, 0x29, 0x0f,
	0xc9, 0x05, 0x09, 0x63, 0xa8, 0x4f, 0x9f, 0xdc, 0xb6, 0x6a, 0xb7, 0x61, 0xc4, 0x1e, 0x44, 0x73,
	0x9d, 0x6a, 0x1a, 0xff, 0xb3, 0xba, 0x1b, 0x08, 0x42, 0x27, 0x2f, 0xf1, 0xc5, 0x23, 0xe8, 0x4e,
	0x95, 0x32, 0xd4, 0xfa, 0x42, 0x75, 0x0f, 0x22, 0xf1, 0xc5, 0xa9, 0xb4, 0xa5, 0xd1, 0xcc, 0x0d,
	0x88, 0xf0, 0x9f, 0xa7, 0x5b, 0xea, 0xbb, 0xec, 0x6f, 0x72, 0x0b, 0x58, 0xb4, 0x7b, 0x2a, 0x3d,
	0x88, 0x74, 0xe1, 0x4f, 0xbb, 0xe1, 0xee, 0xa7, 0x0a, 0xcd, 0xb1, 0xd1, 0x2b, 0x21, 0x99, 0xfe,
	0xc6, 0x7b, 0xa8, 0x39, 0x96, 0x78, 0x3e, 0x70, 0xbc, 0x07, 0x07, 0xb0, 0x93, 0x8b, 0x63, 0xd9,
	0x87, 0x3f, 0x40, 0xdd, 0xd3, 0xc3, 0x83, 0x93, 0x1c, 0x70, 0x72, 0x59, 0xd2, 0xbd, 0x77, 0x04,
	0x8d, 0x40, 0x00, 0xf7, 0x47, 0x47, 0xe0, 0x93, 0xb8, 0xbc, 0xf0, 0xf6, 0x09, 0x40, 0xfe, 0xb5,
	0x78, 0x15, 0xee, 0x4a, 0x00, 0x93, 0xe4, 0xaf, 0x95, 0x0b, 0x79, 0x6c, 0xbc, 0xf9, 0xf7, 0xb5,
	0xa8, 0xd9, 0xe7, 0x36, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x3f, 0x21, 0xd9, 0x7c, 0x02,
	0x00, 0x00,
}