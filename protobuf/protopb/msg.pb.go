// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package protopb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Msg struct {
	Cmd                  uint32   `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	ReqUid               int64    `protobuf:"varint,2,opt,name=req_uid,json=reqUid,proto3" json:"req_uid,omitempty"`
	Tms                  int64    `protobuf:"varint,3,opt,name=tms,proto3" json:"tms,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *Msg) GetReqUid() int64 {
	if m != nil {
		return m.ReqUid
	}
	return 0
}

func (m *Msg) GetTms() int64 {
	if m != nil {
		return m.Tms
	}
	return 0
}

func (m *Msg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Msg)(nil), "protopb.Msg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x53, 0x05, 0x49, 0x4a, 0x61, 0x5c, 0xcc, 0xbe,
	0xc5, 0xe9, 0x42, 0x02, 0x5c, 0xcc, 0xc9, 0xb9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41,
	0x20, 0xa6, 0x90, 0x38, 0x17, 0x7b, 0x51, 0x6a, 0x61, 0x7c, 0x69, 0x66, 0x8a, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x73, 0x10, 0x5b, 0x51, 0x6a, 0x61, 0x68, 0x66, 0x0a, 0x48, 0x69, 0x49, 0x6e, 0xb1,
	0x04, 0x33, 0x58, 0x10, 0xc4, 0x14, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x51,
	0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x9d, 0x44, 0xa2, 0x84, 0xf4, 0xf4, 0xf4, 0xa1, 0xb6, 0x58,
	0x43, 0xe9, 0x24, 0x36, 0x30, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x82, 0x00, 0xeb, 0xd1,
	0x8a, 0x00, 0x00, 0x00,
}
