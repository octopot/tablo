// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/service.proto

package v1

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

type Tablo struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Columns              []*Column `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Tablo) Reset()         { *m = Tablo{} }
func (m *Tablo) String() string { return proto.CompactTextString(m) }
func (*Tablo) ProtoMessage()    {}
func (*Tablo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e34d69331f2f1a, []int{0}
}

func (m *Tablo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tablo.Unmarshal(m, b)
}
func (m *Tablo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tablo.Marshal(b, m, deterministic)
}
func (m *Tablo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tablo.Merge(m, src)
}
func (m *Tablo) XXX_Size() int {
	return xxx_messageInfo_Tablo.Size(m)
}
func (m *Tablo) XXX_DiscardUnknown() {
	xxx_messageInfo_Tablo.DiscardUnknown(m)
}

var xxx_messageInfo_Tablo proto.InternalMessageInfo

func (m *Tablo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tablo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Tablo) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

type Column struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Cards                []*Card  `protobuf:"bytes,3,rep,name=cards,proto3" json:"cards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e34d69331f2f1a, []int{1}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Column) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Column) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

type Card struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Labels               []string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e34d69331f2f1a, []int{2}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Card) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Card) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Card) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Criteria struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Criteria) Reset()         { *m = Criteria{} }
func (m *Criteria) String() string { return proto.CompactTextString(m) }
func (*Criteria) ProtoMessage()    {}
func (*Criteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e34d69331f2f1a, []int{3}
}

func (m *Criteria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Criteria.Unmarshal(m, b)
}
func (m *Criteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Criteria.Marshal(b, m, deterministic)
}
func (m *Criteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Criteria.Merge(m, src)
}
func (m *Criteria) XXX_Size() int {
	return xxx_messageInfo_Criteria.Size(m)
}
func (m *Criteria) XXX_DiscardUnknown() {
	xxx_messageInfo_Criteria.DiscardUnknown(m)
}

var xxx_messageInfo_Criteria proto.InternalMessageInfo

type List struct {
	List                 []*Tablo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3e34d69331f2f1a, []int{4}
}

func (m *List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List.Unmarshal(m, b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List.Marshal(b, m, deterministic)
}
func (m *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(m, src)
}
func (m *List) XXX_Size() int {
	return xxx_messageInfo_List.Size(m)
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

func (m *List) GetList() []*Tablo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Tablo)(nil), "octolab.api.tablo.v1.Tablo")
	proto.RegisterType((*Column)(nil), "octolab.api.tablo.v1.Column")
	proto.RegisterType((*Card)(nil), "octolab.api.tablo.v1.Card")
	proto.RegisterType((*Criteria)(nil), "octolab.api.tablo.v1.Criteria")
	proto.RegisterType((*List)(nil), "octolab.api.tablo.v1.List")
}

func init() { proto.RegisterFile("v1/service.proto", fileDescriptor_d3e34d69331f2f1a) }

var fileDescriptor_d3e34d69331f2f1a = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x4f, 0x83, 0x30,
	0x18, 0x0d, 0x14, 0x98, 0xfb, 0x34, 0xba, 0x34, 0x8b, 0xa9, 0x68, 0x0c, 0xe1, 0xc4, 0xa9, 0xc8,
	0x4c, 0xf4, 0xe4, 0xc5, 0x69, 0x16, 0x13, 0x4f, 0xf8, 0xe3, 0xe0, 0xc9, 0x02, 0x3d, 0x34, 0x29,
	0xeb, 0x52, 0x3a, 0xfe, 0x12, 0xff, 0x60, 0x43, 0xd9, 0x6e, 0xb0, 0xec, 0xd6, 0xf7, 0xfa, 0xbe,
	0xef, 0xbd, 0xbc, 0x16, 0x66, 0x6d, 0x96, 0x36, 0x5c, 0xb7, 0xa2, 0xe4, 0x74, 0xa3, 0x95, 0x51,
	0x78, 0xae, 0x4a, 0xa3, 0x24, 0x2b, 0x28, 0xdb, 0x08, 0x6a, 0x58, 0x21, 0x15, 0x6d, 0xb3, 0xf0,
	0xa2, 0xcd, 0xd2, 0x52, 0xd5, 0xb5, 0x5a, 0xf7, 0xb2, 0x98, 0x83, 0xff, 0xd9, 0x5d, 0xe2, 0x73,
	0x70, 0x45, 0x45, 0x9c, 0xc8, 0x49, 0xa6, 0xb9, 0x2b, 0x2a, 0x3c, 0x07, 0xdf, 0x08, 0x23, 0x39,
	0x71, 0x2d, 0xd5, 0x03, 0xfc, 0x00, 0x93, 0x52, 0xc9, 0x6d, 0xbd, 0x6e, 0x08, 0x8a, 0x50, 0x72,
	0xba, 0xb8, 0xa1, 0x43, 0x3e, 0x74, 0x69, 0x45, 0xf9, 0x5e, 0x1c, 0xff, 0x42, 0xd0, 0x53, 0x47,
	0xfa, 0xdc, 0x81, 0x5f, 0x32, 0x5d, 0xed, 0x5d, 0xc2, 0x11, 0x17, 0xa6, 0xab, 0xbc, 0x17, 0xc6,
	0xdf, 0xe0, 0x75, 0xf0, 0xc8, 0xfd, 0x33, 0x40, 0x5b, 0x2d, 0x09, 0xb2, 0x5c, 0x77, 0xc4, 0x97,
	0x10, 0x48, 0x56, 0x70, 0xd9, 0x10, 0x2f, 0x42, 0xc9, 0x34, 0xdf, 0xa1, 0x18, 0xe0, 0x64, 0xa9,
	0x85, 0xe1, 0x5a, 0xb0, 0xf8, 0x11, 0xbc, 0x77, 0xd1, 0x18, 0x9c, 0x82, 0x27, 0x45, 0x63, 0x88,
	0x6b, 0xc3, 0x5d, 0x0f, 0x87, 0xb3, 0xb5, 0xe6, 0x56, 0xb8, 0xf8, 0x73, 0xe0, 0xcc, 0xe2, 0x8f,
	0xfe, 0x8d, 0xf0, 0x13, 0xa0, 0x15, 0x37, 0xf8, 0x6a, 0x78, 0xf4, 0xeb, 0xed, 0x25, 0x3c, 0xb4,
	0x15, 0xbf, 0xc2, 0x64, 0xc5, 0x8d, 0xcd, 0x72, 0x3b, 0x52, 0xcd, 0x2e, 0x73, 0x38, 0x52, 0x5d,
	0x37, 0xfb, 0xec, 0xfd, 0xb8, 0x6d, 0x56, 0x04, 0xf6, 0x27, 0xdc, 0xff, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x97, 0x45, 0x50, 0x99, 0x44, 0x02, 0x00, 0x00,
}
