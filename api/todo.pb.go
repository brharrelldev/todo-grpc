// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/todo.proto

package todo_api

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

type Todo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskName             string   `protobuf:"bytes,2,opt,name=taskName,proto3" json:"taskName,omitempty"`
	Done                 bool     `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *Todo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type GetTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{1}
}

func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (m *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(m, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAllTodoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllTodoRequest) Reset()         { *m = GetAllTodoRequest{} }
func (m *GetAllTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllTodoRequest) ProtoMessage()    {}
func (*GetAllTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{2}
}

func (m *GetAllTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllTodoRequest.Unmarshal(m, b)
}
func (m *GetAllTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllTodoRequest.Marshal(b, m, deterministic)
}
func (m *GetAllTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllTodoRequest.Merge(m, src)
}
func (m *GetAllTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllTodoRequest.Size(m)
}
func (m *GetAllTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllTodoRequest proto.InternalMessageInfo

type GetAllTodoResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllTodoResponse) Reset()         { *m = GetAllTodoResponse{} }
func (m *GetAllTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllTodoResponse) ProtoMessage()    {}
func (*GetAllTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{3}
}

func (m *GetAllTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllTodoResponse.Unmarshal(m, b)
}
func (m *GetAllTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllTodoResponse.Marshal(b, m, deterministic)
}
func (m *GetAllTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllTodoResponse.Merge(m, src)
}
func (m *GetAllTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllTodoResponse.Size(m)
}
func (m *GetAllTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllTodoResponse proto.InternalMessageInfo

func (m *GetAllTodoResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type GetTodoResponse struct {
	Response             *Todo    `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoResponse) Reset()         { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()    {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{4}
}

func (m *GetTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoResponse.Unmarshal(m, b)
}
func (m *GetTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoResponse.Marshal(b, m, deterministic)
}
func (m *GetTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoResponse.Merge(m, src)
}
func (m *GetTodoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTodoResponse.Size(m)
}
func (m *GetTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoResponse proto.InternalMessageInfo

func (m *GetTodoResponse) GetResponse() *Todo {
	if m != nil {
		return m.Response
	}
	return nil
}

type CreateTodoRequest struct {
	TaskName             string   `protobuf:"bytes,1,opt,name=taskName,proto3" json:"taskName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{5}
}

func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(m, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

type CreateTodoResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{6}
}

func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(m, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateTodoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRequest) Reset()         { *m = DeleteTodoRequest{} }
func (m *DeleteTodoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRequest) ProtoMessage()    {}
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{7}
}

func (m *DeleteTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRequest.Unmarshal(m, b)
}
func (m *DeleteTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRequest.Merge(m, src)
}
func (m *DeleteTodoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRequest.Size(m)
}
func (m *DeleteTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRequest proto.InternalMessageInfo

func (m *DeleteTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoResponse) Reset()         { *m = DeleteTodoResponse{} }
func (m *DeleteTodoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoResponse) ProtoMessage()    {}
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca2a0928dca4b1b7, []int{8}
}

func (m *DeleteTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoResponse.Unmarshal(m, b)
}
func (m *DeleteTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoResponse.Merge(m, src)
}
func (m *DeleteTodoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoResponse.Size(m)
}
func (m *DeleteTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoResponse proto.InternalMessageInfo

func (m *DeleteTodoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.api.Todo")
	proto.RegisterType((*GetTodoRequest)(nil), "todo.api.GetTodoRequest")
	proto.RegisterType((*GetAllTodoRequest)(nil), "todo.api.GetAllTodoRequest")
	proto.RegisterType((*GetAllTodoResponse)(nil), "todo.api.GetAllTodoResponse")
	proto.RegisterType((*GetTodoResponse)(nil), "todo.api.GetTodoResponse")
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.api.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.api.CreateTodoResponse")
	proto.RegisterType((*DeleteTodoRequest)(nil), "todo.api.DeleteTodoRequest")
	proto.RegisterType((*DeleteTodoResponse)(nil), "todo.api.DeleteTodoResponse")
}

func init() { proto.RegisterFile("api/todo.proto", fileDescriptor_ca2a0928dca4b1b7) }

var fileDescriptor_ca2a0928dca4b1b7 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0xed, 0xb6, 0xfd, 0xbe, 0xae, 0x53, 0x58, 0xe9, 0x78, 0x59, 0xab, 0x87, 0x25, 0x7a, 0x28,
	0x1e, 0xb6, 0x50, 0x6f, 0x82, 0xa2, 0x28, 0x16, 0x11, 0x3c, 0x14, 0xd1, 0x73, 0x64, 0x07, 0x59,
	0x6c, 0x9b, 0xb5, 0x89, 0xff, 0x5f, 0x36, 0xa6, 0x9b, 0xd9, 0xa6, 0xb7, 0x99, 0xcc, 0xcc, 0x9b,
	0xf7, 0x5e, 0x12, 0x48, 0x64, 0x55, 0x4e, 0x8d, 0x2a, 0x54, 0x5e, 0x6d, 0x94, 0x51, 0x18, 0xdb,
	0x58, 0x56, 0xa5, 0x78, 0x84, 0xfe, 0xab, 0x2a, 0x14, 0x26, 0xd0, 0x2d, 0x8b, 0x34, 0xca, 0xa2,
	0xc9, 0xc1, 0xa2, 0x5b, 0x16, 0x38, 0x86, 0xd8, 0x48, 0xfd, 0xf5, 0x22, 0x57, 0x94, 0x76, 0xed,
	0x69, 0x93, 0x23, 0x42, 0xbf, 0x50, 0x6b, 0x4a, 0x7b, 0x59, 0x34, 0x89, 0x17, 0x36, 0x16, 0x19,
	0x24, 0x73, 0x32, 0x35, 0xd4, 0x82, 0xbe, 0x7f, 0x48, 0x9b, 0x5d, 0x44, 0x71, 0x04, 0xa3, 0x39,
	0x99, 0xbb, 0xe5, 0x92, 0x35, 0x89, 0x2b, 0x40, 0x7e, 0xa8, 0x2b, 0xb5, 0xd6, 0x84, 0xe7, 0xf0,
	0xaf, 0x26, 0xa8, 0xd3, 0x28, 0xeb, 0x4d, 0x86, 0xb3, 0x24, 0xdf, 0xd2, 0xcd, 0x6d, 0xdb, 0x5f,
	0x51, 0x5c, 0xc3, 0x61, 0xb3, 0xd2, 0x0d, 0x5e, 0x40, 0xbc, 0x71, 0xb1, 0xdd, 0x1c, 0xce, 0x36,
	0x75, 0x31, 0x85, 0xd1, 0xfd, 0x86, 0xa4, 0x21, 0x4e, 0x9a, 0xcb, 0x8e, 0xda, 0xb2, 0xc5, 0x0d,
	0x20, 0x1f, 0x70, 0x2b, 0x77, 0x8d, 0x4b, 0x61, 0xb0, 0x22, 0xad, 0xe5, 0xe7, 0xd6, 0xb7, 0x6d,
	0x2a, 0xce, 0x60, 0xf4, 0x40, 0x4b, 0x6a, 0x2f, 0xdc, 0x75, 0x29, 0x07, 0xe4, 0x4d, 0x6e, 0x09,
	0x03, 0x8d, 0x5a, 0xa0, 0xb3, 0x77, 0x00, 0x4f, 0x0a, 0x9f, 0x5a, 0xd9, 0x89, 0xd7, 0x1e, 0x28,
	0x1d, 0x9f, 0xee, 0x2f, 0x3a, 0x73, 0x3a, 0xb3, 0x67, 0x18, 0x38, 0x77, 0xf1, 0xd6, 0x87, 0xa9,
	0x9f, 0x6a, 0x5f, 0xf7, 0xf8, 0x78, 0x4f, 0xa5, 0x01, 0x7b, 0x83, 0xa1, 0xbf, 0x66, 0x8d, 0x73,
	0x88, 0x5d, 0x8f, 0xe6, 0x24, 0x83, 0xe7, 0xc1, 0x49, 0x86, 0xcf, 0x44, 0x74, 0x6a, 0xf5, 0xde,
	0xad, 0x5a, 0x3d, 0xcb, 0x18, 0x70, 0x60, 0x3b, 0x07, 0x0e, 0xed, 0x16, 0x9d, 0x8f, 0xff, 0xf6,
	0x9f, 0x5c, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x99, 0xf8, 0xc4, 0x39, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CreateTodoClient is the client API for CreateTodo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateTodoClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
}

type createTodoClient struct {
	cc *grpc.ClientConn
}

func NewCreateTodoClient(cc *grpc.ClientConn) CreateTodoClient {
	return &createTodoClient{cc}
}

func (c *createTodoClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.api.CreateTodo/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateTodoServer is the server API for CreateTodo service.
type CreateTodoServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
}

// UnimplementedCreateTodoServer can be embedded to have forward compatible implementations.
type UnimplementedCreateTodoServer struct {
}

func (*UnimplementedCreateTodoServer) CreateTodo(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}

func RegisterCreateTodoServer(s *grpc.Server, srv CreateTodoServer) {
	s.RegisterService(&_CreateTodo_serviceDesc, srv)
}

func _CreateTodo_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateTodoServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.api.CreateTodo/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateTodoServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateTodo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.api.CreateTodo",
	HandlerType: (*CreateTodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _CreateTodo_CreateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todo.proto",
}

// GetTodoClient is the client API for GetTodo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetTodoClient interface {
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error)
}

type getTodoClient struct {
	cc *grpc.ClientConn
}

func NewGetTodoClient(cc *grpc.ClientConn) GetTodoClient {
	return &getTodoClient{cc}
}

func (c *getTodoClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error) {
	out := new(GetTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.api.GetTodo/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetTodoServer is the server API for GetTodo service.
type GetTodoServer interface {
	GetTodo(context.Context, *GetTodoRequest) (*GetTodoResponse, error)
}

// UnimplementedGetTodoServer can be embedded to have forward compatible implementations.
type UnimplementedGetTodoServer struct {
}

func (*UnimplementedGetTodoServer) GetTodo(ctx context.Context, req *GetTodoRequest) (*GetTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}

func RegisterGetTodoServer(s *grpc.Server, srv GetTodoServer) {
	s.RegisterService(&_GetTodo_serviceDesc, srv)
}

func _GetTodo_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTodoServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.api.GetTodo/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTodoServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetTodo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.api.GetTodo",
	HandlerType: (*GetTodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodo",
			Handler:    _GetTodo_GetTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todo.proto",
}

// GetAllTodosClient is the client API for GetAllTodos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetAllTodosClient interface {
	GetTodos(ctx context.Context, in *GetAllTodoRequest, opts ...grpc.CallOption) (*GetAllTodoResponse, error)
}

type getAllTodosClient struct {
	cc *grpc.ClientConn
}

func NewGetAllTodosClient(cc *grpc.ClientConn) GetAllTodosClient {
	return &getAllTodosClient{cc}
}

func (c *getAllTodosClient) GetTodos(ctx context.Context, in *GetAllTodoRequest, opts ...grpc.CallOption) (*GetAllTodoResponse, error) {
	out := new(GetAllTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.api.GetAllTodos/GetTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetAllTodosServer is the server API for GetAllTodos service.
type GetAllTodosServer interface {
	GetTodos(context.Context, *GetAllTodoRequest) (*GetAllTodoResponse, error)
}

// UnimplementedGetAllTodosServer can be embedded to have forward compatible implementations.
type UnimplementedGetAllTodosServer struct {
}

func (*UnimplementedGetAllTodosServer) GetTodos(ctx context.Context, req *GetAllTodoRequest) (*GetAllTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodos not implemented")
}

func RegisterGetAllTodosServer(s *grpc.Server, srv GetAllTodosServer) {
	s.RegisterService(&_GetAllTodos_serviceDesc, srv)
}

func _GetAllTodos_GetTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetAllTodosServer).GetTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.api.GetAllTodos/GetTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetAllTodosServer).GetTodos(ctx, req.(*GetAllTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetAllTodos_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.api.GetAllTodos",
	HandlerType: (*GetAllTodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodos",
			Handler:    _GetAllTodos_GetTodos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todo.proto",
}

// DeleteTodoClient is the client API for DeleteTodo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeleteTodoClient interface {
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
}

type deleteTodoClient struct {
	cc *grpc.ClientConn
}

func NewDeleteTodoClient(cc *grpc.ClientConn) DeleteTodoClient {
	return &deleteTodoClient{cc}
}

func (c *deleteTodoClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.api.DeleteTodo/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteTodoServer is the server API for DeleteTodo service.
type DeleteTodoServer interface {
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
}

// UnimplementedDeleteTodoServer can be embedded to have forward compatible implementations.
type UnimplementedDeleteTodoServer struct {
}

func (*UnimplementedDeleteTodoServer) DeleteTodo(ctx context.Context, req *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}

func RegisterDeleteTodoServer(s *grpc.Server, srv DeleteTodoServer) {
	s.RegisterService(&_DeleteTodo_serviceDesc, srv)
}

func _DeleteTodo_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteTodoServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.api.DeleteTodo/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteTodoServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeleteTodo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.api.DeleteTodo",
	HandlerType: (*DeleteTodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteTodo",
			Handler:    _DeleteTodo_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todo.proto",
}
