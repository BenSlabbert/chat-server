// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/chat.proto

package chat

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

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{3}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateRoomRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomRequest) Reset()         { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()    {}
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{4}
}

func (m *CreateRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomRequest.Unmarshal(m, b)
}
func (m *CreateRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomRequest.Merge(m, src)
}
func (m *CreateRoomRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoomRequest.Size(m)
}
func (m *CreateRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomRequest proto.InternalMessageInfo

func (m *CreateRoomRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRoomResponse struct {
	Room                 *Room    `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomResponse) Reset()         { *m = CreateRoomResponse{} }
func (m *CreateRoomResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoomResponse) ProtoMessage()    {}
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{5}
}

func (m *CreateRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomResponse.Unmarshal(m, b)
}
func (m *CreateRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomResponse.Marshal(b, m, deterministic)
}
func (m *CreateRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomResponse.Merge(m, src)
}
func (m *CreateRoomResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRoomResponse.Size(m)
}
func (m *CreateRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomResponse proto.InternalMessageInfo

func (m *CreateRoomResponse) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

type GetRoomsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoomsRequest) Reset()         { *m = GetRoomsRequest{} }
func (m *GetRoomsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRoomsRequest) ProtoMessage()    {}
func (*GetRoomsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{6}
}

func (m *GetRoomsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomsRequest.Unmarshal(m, b)
}
func (m *GetRoomsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomsRequest.Marshal(b, m, deterministic)
}
func (m *GetRoomsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomsRequest.Merge(m, src)
}
func (m *GetRoomsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRoomsRequest.Size(m)
}
func (m *GetRoomsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomsRequest proto.InternalMessageInfo

type GetRoomsResponse struct {
	Rooms                map[string]*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetRoomsResponse) Reset()         { *m = GetRoomsResponse{} }
func (m *GetRoomsResponse) String() string { return proto.CompactTextString(m) }
func (*GetRoomsResponse) ProtoMessage()    {}
func (*GetRoomsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{7}
}

func (m *GetRoomsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomsResponse.Unmarshal(m, b)
}
func (m *GetRoomsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomsResponse.Marshal(b, m, deterministic)
}
func (m *GetRoomsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomsResponse.Merge(m, src)
}
func (m *GetRoomsResponse) XXX_Size() int {
	return xxx_messageInfo_GetRoomsResponse.Size(m)
}
func (m *GetRoomsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomsResponse proto.InternalMessageInfo

func (m *GetRoomsResponse) GetRooms() map[string]*Room {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type StreamRoomRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRoomRequest) Reset()         { *m = StreamRoomRequest{} }
func (m *StreamRoomRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRoomRequest) ProtoMessage()    {}
func (*StreamRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{8}
}

func (m *StreamRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRoomRequest.Unmarshal(m, b)
}
func (m *StreamRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRoomRequest.Marshal(b, m, deterministic)
}
func (m *StreamRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRoomRequest.Merge(m, src)
}
func (m *StreamRoomRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRoomRequest.Size(m)
}
func (m *StreamRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRoomRequest proto.InternalMessageInfo

func (m *StreamRoomRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamRoomResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRoomResponse) Reset()         { *m = StreamRoomResponse{} }
func (m *StreamRoomResponse) String() string { return proto.CompactTextString(m) }
func (*StreamRoomResponse) ProtoMessage()    {}
func (*StreamRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{9}
}

func (m *StreamRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRoomResponse.Unmarshal(m, b)
}
func (m *StreamRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRoomResponse.Marshal(b, m, deterministic)
}
func (m *StreamRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRoomResponse.Merge(m, src)
}
func (m *StreamRoomResponse) XXX_Size() int {
	return xxx_messageInfo_StreamRoomResponse.Size(m)
}
func (m *StreamRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRoomResponse proto.InternalMessageInfo

func (m *StreamRoomResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Room struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Users                map[string]*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{10}
}

func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Room) GetUsers() map[string]*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_767a49518b351716, []int{11}
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

func init() {
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*PingResponse)(nil), "PingResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "CreateUserResponse")
	proto.RegisterType((*CreateRoomRequest)(nil), "CreateRoomRequest")
	proto.RegisterType((*CreateRoomResponse)(nil), "CreateRoomResponse")
	proto.RegisterType((*GetRoomsRequest)(nil), "GetRoomsRequest")
	proto.RegisterType((*GetRoomsResponse)(nil), "GetRoomsResponse")
	proto.RegisterMapType((map[string]*Room)(nil), "GetRoomsResponse.RoomsEntry")
	proto.RegisterType((*StreamRoomRequest)(nil), "StreamRoomRequest")
	proto.RegisterType((*StreamRoomResponse)(nil), "StreamRoomResponse")
	proto.RegisterType((*Room)(nil), "Room")
	proto.RegisterMapType((map[string]*User)(nil), "Room.UsersEntry")
	proto.RegisterType((*User)(nil), "User")
}

func init() {
	proto.RegisterFile("chat/chat.proto", fileDescriptor_767a49518b351716)
}

var fileDescriptor_767a49518b351716 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x4f, 0xf2, 0x40,
	0x10, 0x4d, 0x4b, 0xcb, 0xf7, 0x7d, 0x03, 0x7c, 0x94, 0xe1, 0x52, 0xab, 0x07, 0x52, 0x13, 0x25,
	0x1e, 0x16, 0x53, 0x63, 0x34, 0x5e, 0x4c, 0x24, 0xc6, 0xab, 0xa9, 0xf1, 0xe2, 0xad, 0xca, 0x06,
	0x88, 0xb6, 0x8b, 0xdd, 0xc5, 0x84, 0x33, 0x07, 0xff, 0x6d, 0xb3, 0xdb, 0x1f, 0x5b, 0x28, 0x89,
	0x5c, 0x9a, 0xed, 0xdb, 0x37, 0x2f, 0xf3, 0x66, 0xde, 0x42, 0xf7, 0x6d, 0x16, 0x89, 0x91, 0xfc,
	0x90, 0x45, 0xca, 0x04, 0xf3, 0x3b, 0xd0, 0x7a, 0x9c, 0x27, 0xd3, 0x90, 0x7e, 0x2e, 0x29, 0x17,
	0xbe, 0x0f, 0xed, 0xec, 0x97, 0x2f, 0x58, 0xc2, 0x29, 0x22, 0x58, 0x0b, 0x96, 0x4c, 0x5d, 0x63,
	0x60, 0x0c, 0xff, 0x85, 0xea, 0xec, 0x9f, 0x42, 0x6f, 0x9c, 0xd2, 0x48, 0xd0, 0x67, 0x4e, 0xd3,
	0xbc, 0x50, 0x12, 0x93, 0x28, 0xa6, 0x05, 0x51, 0x9e, 0xfd, 0x11, 0x60, 0x95, 0x98, 0x4b, 0x1e,
	0x80, 0xb5, 0xe4, 0x34, 0x55, 0xcc, 0x56, 0x60, 0x13, 0x75, 0xa9, 0x20, 0xad, 0x1c, 0x32, 0x16,
	0xef, 0xa5, 0x9c, 0x11, 0xb5, 0x72, 0xca, 0x58, 0x5c, 0x2a, 0xab, 0x4b, 0x05, 0xf9, 0x3d, 0xe8,
	0x3e, 0x50, 0x21, 0x01, 0x5e, 0x58, 0xfd, 0x36, 0xc0, 0xd1, 0x58, 0x2e, 0x11, 0x80, 0x2d, 0xf9,
	0xdc, 0x35, 0x06, 0x8d, 0x61, 0x2b, 0x38, 0x22, 0xdb, 0x0c, 0x25, 0xca, 0xef, 0x13, 0x91, 0xae,
	0xc2, 0x8c, 0xea, 0xdd, 0x02, 0x68, 0x10, 0x1d, 0x68, 0xbc, 0xd3, 0x55, 0xde, 0xad, 0x3c, 0xe2,
	0x21, 0xd8, 0x5f, 0xd1, 0xc7, 0x92, 0xba, 0x66, 0xb5, 0xaf, 0x0c, 0xbb, 0x31, 0xaf, 0x0d, 0x69,
	0xfb, 0x49, 0xa4, 0x34, 0x8a, 0x7f, 0xb3, 0x4d, 0x00, 0xab, 0xc4, 0xbc, 0x67, 0x17, 0xfe, 0xc4,
	0x94, 0xf3, 0x68, 0x5a, 0x90, 0x8b, 0x5f, 0x7f, 0x6d, 0x80, 0x25, 0xa9, 0xbb, 0xc4, 0xf0, 0x04,
	0x6c, 0x39, 0x74, 0xee, 0x9a, 0xca, 0xaa, 0xa3, 0xda, 0x52, 0xdb, 0x28, 0xec, 0xa9, 0x6b, 0x69,
	0x4f, 0x83, 0xfb, 0xd8, 0x53, 0x0b, 0xad, 0xd8, 0x3b, 0x03, 0x4b, 0x42, 0xf8, 0x1f, 0xcc, 0xf9,
	0x24, 0xaf, 0x34, 0xe7, 0x93, 0xb2, 0x29, 0x53, 0x37, 0x15, 0xac, 0x4d, 0xb0, 0xc6, 0xb3, 0x48,
	0xe0, 0x31, 0x58, 0x32, 0x88, 0xd8, 0x26, 0x95, 0x78, 0x7a, 0x1d, 0xb2, 0x91, 0xce, 0x4b, 0x00,
	0x1d, 0x30, 0x44, 0x52, 0x8b, 0xa5, 0xd7, 0x27, 0x3b, 0x12, 0x58, 0x96, 0x65, 0xb3, 0x21, 0xb5,
	0xcc, 0x95, 0x65, 0x1b, 0x73, 0x1e, 0xc1, 0xdf, 0x22, 0x0d, 0xe8, 0x90, 0xad, 0x38, 0x79, 0xbd,
	0x5a, 0x54, 0xf0, 0x0a, 0x40, 0xaf, 0x0b, 0x91, 0xd4, 0x96, 0xec, 0xf5, 0x49, 0x7d, 0x9f, 0xe7,
	0xc6, 0x5d, 0xf3, 0xc5, 0x92, 0x4f, 0xf4, 0xb5, 0xa9, 0xde, 0xe8, 0xc5, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdf, 0x77, 0xcc, 0x6b, 0xb6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error)
	StreamRoom(ctx context.Context, in *StreamRoomRequest, opts ...grpc.CallOption) (Chat_StreamRoomClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/Chat/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/Chat/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, "/Chat/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error) {
	out := new(GetRoomsResponse)
	err := c.cc.Invoke(ctx, "/Chat/GetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) StreamRoom(ctx context.Context, in *StreamRoomRequest, opts ...grpc.CallOption) (Chat_StreamRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/Chat/StreamRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_StreamRoomClient interface {
	Recv() (*StreamRoomResponse, error)
	grpc.ClientStream
}

type chatStreamRoomClient struct {
	grpc.ClientStream
}

func (x *chatStreamRoomClient) Recv() (*StreamRoomResponse, error) {
	m := new(StreamRoomResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error)
	StreamRoom(*StreamRoomRequest, Chat_StreamRoomServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedChatServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedChatServer) CreateRoom(ctx context.Context, req *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedChatServer) GetRooms(ctx context.Context, req *GetRoomsRequest) (*GetRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (*UnimplementedChatServer) StreamRoom(req *StreamRoomRequest, srv Chat_StreamRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRoom not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat/GetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_StreamRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).StreamRoom(m, &chatStreamRoomServer{stream})
}

type Chat_StreamRoomServer interface {
	Send(*StreamRoomResponse) error
	grpc.ServerStream
}

type chatStreamRoomServer struct {
	grpc.ServerStream
}

func (x *chatStreamRoomServer) Send(m *StreamRoomResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Chat_Ping_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Chat_CreateUser_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _Chat_CreateRoom_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _Chat_GetRooms_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoom",
			Handler:       _Chat_StreamRoom_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat/chat.proto",
}
