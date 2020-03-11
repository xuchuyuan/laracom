// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package laracom_user_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	StripeId             string   `protobuf:"bytes,6,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	CardBrand            string   `protobuf:"bytes,7,opt,name=card_brand,json=cardBrand,proto3" json:"card_brand,omitempty"`
	CardLastFour         string   `protobuf:"bytes,8,opt,name=card_last_four,json=cardLastFour,proto3" json:"card_last_four,omitempty"`
	TrialEndsAt          string   `protobuf:"bytes,9,opt,name=trial_ends_at,json=trialEndsAt,proto3" json:"trial_ends_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	RememberToken        string   `protobuf:"bytes,11,opt,name=remember_token,json=rememberToken,proto3" json:"remember_token,omitempty"`
	CreatedAt            string   `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *User) GetStripeId() string {
	if m != nil {
		return m.StripeId
	}
	return ""
}

func (m *User) GetCardBrand() string {
	if m != nil {
		return m.CardBrand
	}
	return ""
}

func (m *User) GetCardLastFour() string {
	if m != nil {
		return m.CardLastFour
	}
	return ""
}

func (m *User) GetTrialEndsAt() string {
	if m != nil {
		return m.TrialEndsAt
	}
	return ""
}

func (m *User) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func (m *User) GetRememberToken() string {
	if m != nil {
		return m.RememberToken
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "laracom.user.service.User")
	proto.RegisterType((*Response)(nil), "laracom.user.service.Response")
	proto.RegisterType((*Error)(nil), "laracom.user.service.Error")
}

func init() {
	proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7)
}

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8b, 0xd4, 0x4c,
	0x10, 0x86, 0xbf, 0xc9, 0x4c, 0xf2, 0xcd, 0x54, 0x76, 0xe6, 0x50, 0xac, 0xd2, 0xcc, 0xa2, 0x0c,
	0x41, 0xc1, 0x53, 0x94, 0xdd, 0xb3, 0x87, 0x51, 0x76, 0x41, 0xf0, 0x14, 0xdd, 0x73, 0xe8, 0x49,
	0x97, 0x10, 0x4c, 0xd2, 0xa1, 0xba, 0xa3, 0xbf, 0xc7, 0x3f, 0x2a, 0xd2, 0xd5, 0x19, 0xf1, 0x30,
	0xe8, 0x25, 0xe9, 0x7a, 0xde, 0xb7, 0xba, 0x8a, 0xae, 0x82, 0x27, 0x23, 0x5b, 0x6f, 0x5f, 0x4f,
	0x8e, 0x58, 0x3e, 0xa5, 0xc4, 0x78, 0xdd, 0x69, 0xd6, 0x8d, 0xed, 0x4b, 0x61, 0x8e, 0xf8, 0x5b,
	0xdb, 0x50, 0xf1, 0x33, 0x81, 0xd5, 0xa3, 0x23, 0xc6, 0x1d, 0x24, 0xad, 0x51, 0x8b, 0xc3, 0xe2,
	0xd5, 0xa6, 0x4a, 0x5a, 0x83, 0x08, 0xab, 0x41, 0xf7, 0xa4, 0x12, 0x21, 0x72, 0xc6, 0x6b, 0x48,
	0xa9, 0xd7, 0x6d, 0xa7, 0x96, 0x02, 0x63, 0x80, 0x7b, 0x58, 0x8f, 0xda, 0xb9, 0xef, 0x96, 0x8d,
	0x5a, 0x89, 0xf0, 0x3b, 0xc6, 0xa7, 0x90, 0x39, 0xaf, 0xfd, 0xe4, 0x54, 0x2a, 0xca, 0x1c, 0xe1,
	0x0d, 0x6c, 0x9c, 0xe7, 0x76, 0xa4, 0xba, 0x35, 0x2a, 0x8b, 0x49, 0x11, 0x7c, 0x30, 0xf8, 0x0c,
	0xa0, 0xd1, 0x6c, 0xea, 0x13, 0xeb, 0xc1, 0xa8, 0xff, 0x45, 0xdd, 0x04, 0xf2, 0x2e, 0x00, 0x7c,
	0x01, 0x3b, 0x91, 0x3b, 0xed, 0x7c, 0xfd, 0xc5, 0x4e, 0xac, 0xd6, 0x62, 0xb9, 0x0a, 0xf4, 0xa3,
	0x76, 0xfe, 0xc1, 0x4e, 0x8c, 0x05, 0x6c, 0x3d, 0xb7, 0xba, 0xab, 0x69, 0x30, 0xae, 0xd6, 0x5e,
	0x6d, 0xc4, 0x94, 0x0b, 0xbc, 0x1f, 0x8c, 0x3b, 0xfa, 0x50, 0xc8, 0x50, 0x47, 0x9e, 0x4c, 0x30,
	0x40, 0x2c, 0x34, 0x93, 0xa3, 0xc7, 0x97, 0xb0, 0x63, 0xea, 0xa9, 0x3f, 0x11, 0xd7, 0xde, 0x7e,
	0xa5, 0x41, 0xe5, 0x62, 0xd9, 0x9e, 0xe9, 0xe7, 0x00, 0xa5, 0x5d, 0x26, 0x3d, 0xdf, 0x72, 0x35,
	0xb7, 0x1b, 0x49, 0x2c, 0x32, 0x8d, 0xe6, 0x2c, 0x6f, 0xa3, 0x3c, 0x93, 0xa3, 0x2f, 0x7e, 0x2c,
	0x60, 0x5d, 0x91, 0x1b, 0xed, 0xe0, 0x08, 0x4b, 0x58, 0x85, 0xe9, 0xc8, 0x18, 0xf2, 0xdb, 0x7d,
	0x79, 0x69, 0x64, 0x65, 0x18, 0x57, 0x25, 0x3e, 0x7c, 0x03, 0x69, 0xf8, 0x3b, 0x95, 0x1c, 0x96,
	0xff, 0x48, 0x88, 0x46, 0xbc, 0x83, 0x8c, 0x98, 0x2d, 0x3b, 0xb5, 0x94, 0x94, 0x9b, 0xcb, 0x29,
	0xf7, 0xc1, 0x53, 0xcd, 0xd6, 0xe2, 0x2d, 0xa4, 0x02, 0xc2, 0x52, 0x34, 0xd6, 0x90, 0xf4, 0x97,
	0x56, 0x72, 0xc6, 0x03, 0xe4, 0x86, 0x5c, 0xc3, 0xed, 0xe8, 0x5b, 0x3b, 0xcc, 0xfb, 0xf2, 0x27,
	0xba, 0x7d, 0x84, 0x3c, 0x5c, 0xfe, 0x29, 0xde, 0x8d, 0x0f, 0x90, 0xbd, 0x97, 0xd7, 0xc1, 0xbf,
	0xf4, 0xbb, 0x7f, 0x7e, 0x59, 0x3b, 0x3f, 0x55, 0xf1, 0xdf, 0x29, 0x93, 0xbd, 0xbe, 0xfb, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xa8, 0xfb, 0x3e, 0x3e, 0xf0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "laracom.user.service"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}