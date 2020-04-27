// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=maxWeight,proto3" json:"maxWeight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=maxWeight,proto3" json:"maxWeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
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

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() {
	proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716)
}

var fileDescriptor_04ef66862bb50716 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x65, 0xd3, 0x36, 0x4d, 0x47, 0x52, 0x64, 0x40, 0x58, 0x8b, 0x87, 0x90, 0x83, 0xe4, 0x54,
	0xa1, 0x5e, 0xbc, 0x7a, 0x11, 0x72, 0x4d, 0x41, 0x2f, 0x5e, 0xb6, 0xc9, 0xa8, 0x03, 0xe9, 0x6e,
	0xc8, 0x86, 0xa8, 0x3f, 0xe3, 0xb7, 0x8a, 0x9b, 0x6c, 0x4b, 0x45, 0x3c, 0x65, 0xde, 0x7b, 0x79,
	0xc3, 0x9b, 0xb7, 0x70, 0xd9, 0xb4, 0xa6, 0x33, 0x37, 0x3d, 0x59, 0x4b, 0xf5, 0xf8, 0x59, 0x3b,
	0x0e, 0xc3, 0x01, 0xa5, 0x5f, 0x02, 0xc2, 0x47, 0x37, 0xe2, 0x12, 0x02, 0xae, 0xa4, 0x48, 0x44,
	0xb6, 0x28, 0x02, 0xae, 0x70, 0x05, 0x51, 0xa9, 0x1a, 0x55, 0x72, 0xf7, 0x29, 0x83, 0x44, 0x64,
	0xb3, 0xe2, 0x80, 0xf1, 0x0a, 0x16, 0x7b, 0xf5, 0xf1, 0x44, 0xfc, 0xfa, 0xd6, 0xc9, 0x89, 0x13,
	0x8f, 0x04, 0x22, 0x4c, 0xb5, 0xda, 0x93, 0x9c, 0xba, 0x5d, 0x6e, 0xfe, 0x71, 0xa8, 0x5e, 0x71,
	0xad, 0x76, 0x35, 0xc9, 0x59, 0x22, 0xb2, 0xa8, 0x38, 0x12, 0x28, 0x61, 0x6e, 0xde, 0x35, 0xb5,
	0x79, 0x25, 0x43, 0x67, 0xf2, 0x30, 0xcd, 0x21, 0xde, 0x36, 0x54, 0xf2, 0x0b, 0x97, 0xaa, 0x63,
	0xa3, 0x4f, 0x62, 0x89, 0xff, 0x62, 0x05, 0xbf, 0x62, 0xa5, 0xcf, 0x10, 0x15, 0x64, 0x1b, 0xa3,
	0x2d, 0xe1, 0x35, 0x8c, 0x0d, 0xb8, 0x1d, 0x67, 0x9b, 0xe5, 0x7a, 0xac, 0x67, 0x28, 0xa3, 0x18,
	0x55, 0xcc, 0x60, 0x3e, 0x4c, 0x56, 0x06, 0xc9, 0xe4, 0x8f, 0x1f, 0xbd, 0xbc, 0xc9, 0x21, 0x1e,
	0xa8, 0x2d, 0xb5, 0x3d, 0x97, 0x84, 0x77, 0x10, 0x3f, 0xb0, 0xae, 0xee, 0x0f, 0x47, 0x5e, 0x78,
	0xeb, 0xc9, 0x41, 0xab, 0x73, 0x4f, 0xfb, 0x70, 0xbb, 0xd0, 0xbd, 0xd1, 0xed, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0x73, 0xec, 0xef, 0xc0, 0x01, 0x00, 0x00,
}
