// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messaging/v1/message_list.proto

package messagingv1

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

type CreateMessageRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageRequest) Reset()         { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()    {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_505ca9b1b3a6ca17, []int{0}
}

func (m *CreateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRequest.Unmarshal(m, b)
}
func (m *CreateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRequest.Marshal(b, m, deterministic)
}
func (m *CreateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRequest.Merge(m, src)
}
func (m *CreateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRequest.Size(m)
}
func (m *CreateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRequest proto.InternalMessageInfo

func (m *CreateMessageRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type CreateMessageResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessageResponse) Reset()         { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()    {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_505ca9b1b3a6ca17, []int{1}
}

func (m *CreateMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageResponse.Unmarshal(m, b)
}
func (m *CreateMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageResponse.Marshal(b, m, deterministic)
}
func (m *CreateMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageResponse.Merge(m, src)
}
func (m *CreateMessageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMessageResponse.Size(m)
}
func (m *CreateMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageResponse proto.InternalMessageInfo

func (m *CreateMessageResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListMessagesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMessagesRequest) Reset()         { *m = ListMessagesRequest{} }
func (m *ListMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMessagesRequest) ProtoMessage()    {}
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_505ca9b1b3a6ca17, []int{2}
}

func (m *ListMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesRequest.Unmarshal(m, b)
}
func (m *ListMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesRequest.Marshal(b, m, deterministic)
}
func (m *ListMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesRequest.Merge(m, src)
}
func (m *ListMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMessagesRequest.Size(m)
}
func (m *ListMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesRequest proto.InternalMessageInfo

type ListMessagesResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMessagesResponse) Reset()         { *m = ListMessagesResponse{} }
func (m *ListMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMessagesResponse) ProtoMessage()    {}
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_505ca9b1b3a6ca17, []int{3}
}

func (m *ListMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMessagesResponse.Unmarshal(m, b)
}
func (m *ListMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMessagesResponse.Marshal(b, m, deterministic)
}
func (m *ListMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMessagesResponse.Merge(m, src)
}
func (m *ListMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMessagesResponse.Size(m)
}
func (m *ListMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMessagesResponse proto.InternalMessageInfo

func (m *ListMessagesResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type GetMessageRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageRequest) Reset()         { *m = GetMessageRequest{} }
func (m *GetMessageRequest) String() string { return proto.CompactTextString(m) }
func (*GetMessageRequest) ProtoMessage()    {}
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_505ca9b1b3a6ca17, []int{4}
}

func (m *GetMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageRequest.Unmarshal(m, b)
}
func (m *GetMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageRequest.Marshal(b, m, deterministic)
}
func (m *GetMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageRequest.Merge(m, src)
}
func (m *GetMessageRequest) XXX_Size() int {
	return xxx_messageInfo_GetMessageRequest.Size(m)
}
func (m *GetMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageRequest proto.InternalMessageInfo

func (m *GetMessageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetMessageResponse struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageResponse) Reset()         { *m = GetMessageResponse{} }
func (m *GetMessageResponse) String() string { return proto.CompactTextString(m) }
func (*GetMessageResponse) ProtoMessage()    {}
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_505ca9b1b3a6ca17, []int{5}
}

func (m *GetMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageResponse.Unmarshal(m, b)
}
func (m *GetMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageResponse.Marshal(b, m, deterministic)
}
func (m *GetMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageResponse.Merge(m, src)
}
func (m *GetMessageResponse) XXX_Size() int {
	return xxx_messageInfo_GetMessageResponse.Size(m)
}
func (m *GetMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageResponse proto.InternalMessageInfo

func (m *GetMessageResponse) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateMessageRequest)(nil), "messaging.v1.CreateMessageRequest")
	proto.RegisterType((*CreateMessageResponse)(nil), "messaging.v1.CreateMessageResponse")
	proto.RegisterType((*ListMessagesRequest)(nil), "messaging.v1.ListMessagesRequest")
	proto.RegisterType((*ListMessagesResponse)(nil), "messaging.v1.ListMessagesResponse")
	proto.RegisterType((*GetMessageRequest)(nil), "messaging.v1.GetMessageRequest")
	proto.RegisterType((*GetMessageResponse)(nil), "messaging.v1.GetMessageResponse")
}

func init() {
	proto.RegisterFile("messaging/v1/message_list.proto", fileDescriptor_505ca9b1b3a6ca17)
}

var fileDescriptor_505ca9b1b3a6ca17 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0xcf, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0xd4, 0x87, 0x70, 0x52, 0xe3, 0x73, 0x32, 0x8b,
	0x4b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0xe0, 0x0a, 0xf4, 0xca, 0x0c, 0xa5, 0xa4,
	0xb0, 0x29, 0x87, 0xa8, 0x54, 0xd2, 0xe2, 0x12, 0x71, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0xf5, 0x85,
	0x08, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xea, 0x5c, 0xa2, 0x68, 0x6a, 0x8b,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0xa0, 0x4a, 0x99, 0x32, 0x53,
	0x94, 0x44, 0xb9, 0x84, 0x7d, 0x32, 0x8b, 0x4b, 0xa0, 0xca, 0x8a, 0xa1, 0x66, 0x2a, 0x79, 0x72,
	0x89, 0xa0, 0x0a, 0x43, 0xb5, 0x1b, 0x72, 0x71, 0x40, 0x1d, 0x55, 0x2c, 0xc1, 0xa8, 0xc0, 0xac,
	0xc1, 0x6d, 0x24, 0xaa, 0x87, 0xec, 0x01, 0x3d, 0x98, 0x7d, 0x70, 0x65, 0x4a, 0xca, 0x5c, 0x82,
	0xee, 0xa9, 0x25, 0x68, 0x6e, 0x46, 0x77, 0x86, 0x2b, 0x97, 0x10, 0xb2, 0x22, 0xa8, 0x6d, 0xfa,
	0x5c, 0xec, 0x50, 0x63, 0xc0, 0x4a, 0x71, 0x5a, 0x06, 0x53, 0x65, 0xd4, 0xcf, 0xc4, 0xc5, 0x0d,
	0x15, 0x04, 0x39, 0x5f, 0x28, 0x82, 0x8b, 0x17, 0x25, 0x18, 0x84, 0x94, 0x50, 0x0d, 0xc0, 0x16,
	0x9e, 0x52, 0xca, 0x78, 0xd5, 0x40, 0x9d, 0x16, 0xca, 0xc5, 0x83, 0x1c, 0x40, 0x42, 0x8a, 0xa8,
	0x9a, 0xb0, 0x84, 0xa9, 0x94, 0x12, 0x3e, 0x25, 0x50, 0x63, 0xfd, 0xb9, 0xb8, 0x10, 0xe1, 0x20,
	0x24, 0x8f, 0xaa, 0x03, 0x23, 0x18, 0xa5, 0x14, 0x70, 0x2b, 0x80, 0x18, 0xe8, 0x14, 0xce, 0x25,
	0x90, 0x9c, 0x9f, 0x8b, 0xa2, 0xcc, 0x49, 0x00, 0x29, 0x88, 0x02, 0x40, 0x49, 0x2b, 0x80, 0x31,
	0x8a, 0x1b, 0xae, 0xa2, 0xcc, 0x70, 0x11, 0x13, 0x73, 0x48, 0x44, 0xc4, 0x2a, 0x26, 0x1e, 0x5f,
	0xb8, 0xae, 0x30, 0xc3, 0x53, 0x48, 0xdc, 0x98, 0x30, 0xc3, 0x24, 0x36, 0x70, 0xa2, 0x34, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x10, 0xc4, 0x4e, 0x7c, 0xe1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageListClient is the client API for MessageList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageListClient interface {
	// CreateMessage adds a new message to the message list.
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
	// ListMessages returns the list of messages.
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error)
	// GetMessage returns message by id
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
}

type messageListClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageListClient(cc grpc.ClientConnInterface) MessageListClient {
	return &messageListClient{cc}
}

func (c *messageListClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, "/messaging.v1.MessageList/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageListClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (*ListMessagesResponse, error) {
	out := new(ListMessagesResponse)
	err := c.cc.Invoke(ctx, "/messaging.v1.MessageList/ListMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageListClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	out := new(GetMessageResponse)
	err := c.cc.Invoke(ctx, "/messaging.v1.MessageList/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageListServer is the server API for MessageList service.
type MessageListServer interface {
	// CreateMessage adds a new message to the message list.
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
	// ListMessages returns the list of messages.
	ListMessages(context.Context, *ListMessagesRequest) (*ListMessagesResponse, error)
	// GetMessage returns message by id
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
}

// UnimplementedMessageListServer can be embedded to have forward compatible implementations.
type UnimplementedMessageListServer struct {
}

func (*UnimplementedMessageListServer) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (*UnimplementedMessageListServer) ListMessages(ctx context.Context, req *ListMessagesRequest) (*ListMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (*UnimplementedMessageListServer) GetMessage(ctx context.Context, req *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}

func RegisterMessageListServer(s *grpc.Server, srv MessageListServer) {
	s.RegisterService(&_MessageList_serviceDesc, srv)
}

func _MessageList_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageListServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.v1.MessageList/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageListServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageList_ListMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageListServer).ListMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.v1.MessageList/ListMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageListServer).ListMessages(ctx, req.(*ListMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageList_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageListServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.v1.MessageList/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageListServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.v1.MessageList",
	HandlerType: (*MessageListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _MessageList_CreateMessage_Handler,
		},
		{
			MethodName: "ListMessages",
			Handler:    _MessageList_ListMessages_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _MessageList_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messaging/v1/message_list.proto",
}
