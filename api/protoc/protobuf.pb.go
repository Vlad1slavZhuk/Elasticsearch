// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf.proto

package api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Account struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Ad struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Brand                string   `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Color                string   `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Price                int32    `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{1}
}

func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (m *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(m, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Ad) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Ad) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Ad) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Ad) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

//Request
//Add in grpc... todo
type RequestUpdateAcc struct {
	Acc                  *Account `protobuf:"bytes,1,opt,name=acc,proto3" json:"acc,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUpdateAcc) Reset()         { *m = RequestUpdateAcc{} }
func (m *RequestUpdateAcc) String() string { return proto.CompactTextString(m) }
func (*RequestUpdateAcc) ProtoMessage()    {}
func (*RequestUpdateAcc) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{2}
}

func (m *RequestUpdateAcc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdateAcc.Unmarshal(m, b)
}
func (m *RequestUpdateAcc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdateAcc.Marshal(b, m, deterministic)
}
func (m *RequestUpdateAcc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdateAcc.Merge(m, src)
}
func (m *RequestUpdateAcc) XXX_Size() int {
	return xxx_messageInfo_RequestUpdateAcc.Size(m)
}
func (m *RequestUpdateAcc) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdateAcc.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdateAcc proto.InternalMessageInfo

func (m *RequestUpdateAcc) GetAcc() *Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *RequestUpdateAcc) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RequestID struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestID) Reset()         { *m = RequestID{} }
func (m *RequestID) String() string { return proto.CompactTextString(m) }
func (*RequestID) ProtoMessage()    {}
func (*RequestID) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{3}
}

func (m *RequestID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestID.Unmarshal(m, b)
}
func (m *RequestID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestID.Marshal(b, m, deterministic)
}
func (m *RequestID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestID.Merge(m, src)
}
func (m *RequestID) XXX_Size() int {
	return xxx_messageInfo_RequestID.Size(m)
}
func (m *RequestID) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestID.DiscardUnknown(m)
}

var xxx_messageInfo_RequestID proto.InternalMessageInfo

func (m *RequestID) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RequestUpdateAd struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ad                   *Ad      `protobuf:"bytes,2,opt,name=ad,proto3" json:"ad,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUpdateAd) Reset()         { *m = RequestUpdateAd{} }
func (m *RequestUpdateAd) String() string { return proto.CompactTextString(m) }
func (*RequestUpdateAd) ProtoMessage()    {}
func (*RequestUpdateAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{4}
}

func (m *RequestUpdateAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdateAd.Unmarshal(m, b)
}
func (m *RequestUpdateAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdateAd.Marshal(b, m, deterministic)
}
func (m *RequestUpdateAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdateAd.Merge(m, src)
}
func (m *RequestUpdateAd) XXX_Size() int {
	return xxx_messageInfo_RequestUpdateAd.Size(m)
}
func (m *RequestUpdateAd) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdateAd.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdateAd proto.InternalMessageInfo

func (m *RequestUpdateAd) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RequestUpdateAd) GetAd() *Ad {
	if m != nil {
		return m.Ad
	}
	return nil
}

type RequestAdd struct {
	Ad                   *Ad      `protobuf:"bytes,1,opt,name=ad,proto3" json:"ad,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAdd) Reset()         { *m = RequestAdd{} }
func (m *RequestAdd) String() string { return proto.CompactTextString(m) }
func (*RequestAdd) ProtoMessage()    {}
func (*RequestAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{5}
}

func (m *RequestAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAdd.Unmarshal(m, b)
}
func (m *RequestAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAdd.Marshal(b, m, deterministic)
}
func (m *RequestAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAdd.Merge(m, src)
}
func (m *RequestAdd) XXX_Size() int {
	return xxx_messageInfo_RequestAdd.Size(m)
}
func (m *RequestAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAdd.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAdd proto.InternalMessageInfo

func (m *RequestAdd) GetAd() *Ad {
	if m != nil {
		return m.Ad
	}
	return nil
}

//Response
type ResponseGetAds struct {
	Ads                  []*Ad    `protobuf:"bytes,1,rep,name=ads,proto3" json:"ads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseGetAds) Reset()         { *m = ResponseGetAds{} }
func (m *ResponseGetAds) String() string { return proto.CompactTextString(m) }
func (*ResponseGetAds) ProtoMessage()    {}
func (*ResponseGetAds) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{6}
}

func (m *ResponseGetAds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseGetAds.Unmarshal(m, b)
}
func (m *ResponseGetAds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseGetAds.Marshal(b, m, deterministic)
}
func (m *ResponseGetAds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGetAds.Merge(m, src)
}
func (m *ResponseGetAds) XXX_Size() int {
	return xxx_messageInfo_ResponseGetAds.Size(m)
}
func (m *ResponseGetAds) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGetAds.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGetAds proto.InternalMessageInfo

func (m *ResponseGetAds) GetAds() []*Ad {
	if m != nil {
		return m.Ads
	}
	return nil
}

type ResponseStatus struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseStatus) Reset()         { *m = ResponseStatus{} }
func (m *ResponseStatus) String() string { return proto.CompactTextString(m) }
func (*ResponseStatus) ProtoMessage()    {}
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{7}
}

func (m *ResponseStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseStatus.Unmarshal(m, b)
}
func (m *ResponseStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseStatus.Marshal(b, m, deterministic)
}
func (m *ResponseStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseStatus.Merge(m, src)
}
func (m *ResponseStatus) XXX_Size() int {
	return xxx_messageInfo_ResponseStatus.Size(m)
}
func (m *ResponseStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseStatus proto.InternalMessageInfo

func (m *ResponseStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ResponseSize struct {
	Size                 uint32   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseSize) Reset()         { *m = ResponseSize{} }
func (m *ResponseSize) String() string { return proto.CompactTextString(m) }
func (*ResponseSize) ProtoMessage()    {}
func (*ResponseSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{8}
}

func (m *ResponseSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseSize.Unmarshal(m, b)
}
func (m *ResponseSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseSize.Marshal(b, m, deterministic)
}
func (m *ResponseSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSize.Merge(m, src)
}
func (m *ResponseSize) XXX_Size() int {
	return xxx_messageInfo_ResponseSize.Size(m)
}
func (m *ResponseSize) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSize.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSize proto.InternalMessageInfo

func (m *ResponseSize) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ResponseAccounts struct {
	Acc                  []*Account `protobuf:"bytes,1,rep,name=acc,proto3" json:"acc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ResponseAccounts) Reset()         { *m = ResponseAccounts{} }
func (m *ResponseAccounts) String() string { return proto.CompactTextString(m) }
func (*ResponseAccounts) ProtoMessage()    {}
func (*ResponseAccounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{9}
}

func (m *ResponseAccounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAccounts.Unmarshal(m, b)
}
func (m *ResponseAccounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAccounts.Marshal(b, m, deterministic)
}
func (m *ResponseAccounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAccounts.Merge(m, src)
}
func (m *ResponseAccounts) XXX_Size() int {
	return xxx_messageInfo_ResponseAccounts.Size(m)
}
func (m *ResponseAccounts) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAccounts.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAccounts proto.InternalMessageInfo

func (m *ResponseAccounts) GetAcc() []*Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "serviceProtobuf.Account")
	proto.RegisterType((*Ad)(nil), "serviceProtobuf.Ad")
	proto.RegisterType((*RequestUpdateAcc)(nil), "serviceProtobuf.RequestUpdateAcc")
	proto.RegisterType((*RequestID)(nil), "serviceProtobuf.RequestID")
	proto.RegisterType((*RequestUpdateAd)(nil), "serviceProtobuf.RequestUpdateAd")
	proto.RegisterType((*RequestAdd)(nil), "serviceProtobuf.RequestAdd")
	proto.RegisterType((*ResponseGetAds)(nil), "serviceProtobuf.ResponseGetAds")
	proto.RegisterType((*ResponseStatus)(nil), "serviceProtobuf.ResponseStatus")
	proto.RegisterType((*ResponseSize)(nil), "serviceProtobuf.ResponseSize")
	proto.RegisterType((*ResponseAccounts)(nil), "serviceProtobuf.ResponseAccounts")
}

func init() { proto.RegisterFile("protobuf.proto", fileDescriptor_c77a803fcbc0c059) }

var fileDescriptor_c77a803fcbc0c059 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8b, 0xd3, 0x4c,
	0x14, 0xa6, 0x49, 0xdb, 0xf7, 0xdd, 0xd3, 0xb5, 0x95, 0x51, 0x97, 0xd0, 0x22, 0x76, 0x23, 0x42,
	0xf1, 0x22, 0x8b, 0xf5, 0xc2, 0x0b, 0x51, 0xc8, 0x7e, 0x58, 0xc5, 0x1b, 0xc9, 0xae, 0x08, 0x7a,
	0x35, 0x9d, 0x39, 0x2e, 0xc1, 0x34, 0x13, 0x33, 0x13, 0xc5, 0xfd, 0x35, 0xfe, 0x54, 0x99, 0x8f,
	0xb4, 0x76, 0x6b, 0x36, 0xe2, 0xdd, 0xf9, 0x7c, 0xe6, 0x9c, 0xe7, 0x99, 0x19, 0x18, 0x16, 0xa5,
	0x50, 0x62, 0x59, 0x7d, 0x8e, 0x8c, 0x41, 0x46, 0x12, 0xcb, 0x6f, 0x29, 0xc3, 0x77, 0x2e, 0x3c,
	0x9e, 0x5c, 0x0a, 0x71, 0x99, 0xe1, 0x51, 0x5d, 0x77, 0x84, 0xab, 0x42, 0xfd, 0xb0, 0xd5, 0xe1,
	0x07, 0xf8, 0x2f, 0x66, 0x4c, 0x54, 0xb9, 0x22, 0x63, 0xf8, 0xbf, 0x92, 0x58, 0xe6, 0x74, 0x85,
	0x41, 0x67, 0xda, 0x99, 0xed, 0x25, 0x6b, 0x5f, 0xe7, 0x0a, 0x2a, 0xe5, 0x77, 0x51, 0xf2, 0xc0,
	0xb3, 0xb9, 0xda, 0x27, 0x77, 0xa1, 0xa7, 0xc4, 0x17, 0xcc, 0x03, 0xdf, 0x24, 0xac, 0x13, 0x66,
	0xe0, 0xc5, 0x9c, 0x0c, 0xc1, 0x4b, 0xb9, 0x41, 0xbb, 0x95, 0x78, 0xa9, 0xa9, 0x5d, 0x09, 0x8e,
	0x99, 0x03, 0xb1, 0x8e, 0x8e, 0x2e, 0x4b, 0x9a, 0xf3, 0x1a, 0xc1, 0x38, 0x3a, 0xca, 0x44, 0x26,
	0xca, 0xa0, 0x6b, 0xa3, 0xc6, 0xd1, 0xd1, 0xa2, 0x4c, 0x19, 0x06, 0xbd, 0x69, 0x67, 0xd6, 0x4b,
	0xac, 0x13, 0x5e, 0xc0, 0xed, 0x04, 0xbf, 0x56, 0x28, 0xd5, 0xfb, 0x82, 0x53, 0x85, 0x31, 0x63,
	0xe4, 0x31, 0xf8, 0x94, 0x31, 0x73, 0xf8, 0x60, 0x1e, 0x44, 0xd7, 0x68, 0x89, 0xdc, 0xda, 0x89,
	0x2e, 0xda, 0xec, 0xe0, 0xfd, 0xbe, 0xc3, 0x04, 0xf6, 0x1c, 0xea, 0x9b, 0xd3, 0xeb, 0xab, 0x84,
	0xaf, 0x60, 0xb4, 0x7d, 0xe4, 0xee, 0xb6, 0x0f, 0xc1, 0xa3, 0x96, 0xaf, 0xc1, 0xfc, 0xce, 0xee,
	0x00, 0x3c, 0xf1, 0x28, 0x0f, 0x9f, 0x00, 0x38, 0x9c, 0x98, 0xd7, 0x2d, 0x9d, 0x9b, 0x5b, 0x9e,
	0xc1, 0x30, 0x41, 0x59, 0x88, 0x5c, 0xe2, 0x02, 0x55, 0xcc, 0x25, 0x79, 0x04, 0x3e, 0xe5, 0x32,
	0xe8, 0x4c, 0xfd, 0xa6, 0x3e, 0x9d, 0x0f, 0x67, 0x9b, 0xc6, 0x73, 0x45, 0x55, 0x25, 0xc9, 0x01,
	0xf4, 0xa5, 0xb1, 0x9c, 0xe4, 0xce, 0x0b, 0x43, 0xd8, 0x5f, 0x57, 0xa6, 0x57, 0x48, 0x08, 0x74,
	0x65, 0x7a, 0x85, 0x6e, 0x39, 0x63, 0x87, 0x2f, 0x35, 0xe9, 0xb6, 0xc6, 0x91, 0x29, 0x37, 0xa4,
	0xfb, 0xad, 0xa4, 0xcf, 0x7f, 0xf6, 0x60, 0x74, 0xbe, 0x5d, 0x40, 0x8e, 0xc1, 0x5f, 0xa0, 0x22,
	0xe3, 0x9d, 0xce, 0xb5, 0x10, 0xe3, 0x07, 0x7f, 0xc8, 0x6d, 0x91, 0x11, 0x43, 0x5f, 0x5b, 0x59,
	0x46, 0x0e, 0x22, 0x7b, 0xf7, 0xa3, 0xf5, 0x1b, 0x39, 0xd3, 0x77, 0xbf, 0x1d, 0xe2, 0x2d, 0xf4,
	0xad, 0xaa, 0x64, 0xda, 0x34, 0x49, 0xad, 0xfa, 0x0d, 0x60, 0x8e, 0xe3, 0x13, 0xf0, 0xb5, 0xb4,
	0x93, 0x26, 0xa4, 0x98, 0xff, 0x05, 0xc8, 0x19, 0xf4, 0x4f, 0x31, 0x43, 0x85, 0xff, 0xc8, 0x8d,
	0x83, 0x79, 0x01, 0x5d, 0xa3, 0x67, 0x13, 0x33, 0xf7, 0x9b, 0x01, 0x74, 0xdb, 0x02, 0x20, 0xe6,
	0xbc, 0xfe, 0x31, 0x1a, 0xf5, 0x6d, 0x9f, 0xe3, 0x35, 0x0c, 0x34, 0xd5, 0xf5, 0xb5, 0x69, 0x1a,
	0xe7, 0xb0, 0x11, 0x67, 0xdd, 0xfa, 0x09, 0xee, 0x59, 0x29, 0x2e, 0xf4, 0x9b, 0x3d, 0xa9, 0xca,
	0x12, 0x73, 0x0d, 0x4c, 0x0e, 0x5b, 0x94, 0x63, 0xac, 0x75, 0xcc, 0xe3, 0xe1, 0xc7, 0xfd, 0xc8,
	0x7e, 0x9c, 0xec, 0x39, 0x2d, 0xd2, 0x65, 0xdf, 0xd8, 0x4f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xca, 0xdc, 0x9d, 0xf5, 0x75, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceProtobufClient is the client API for ServiceProtobuf service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceProtobufClient interface {
	Get(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*ResponseGetAds, error)
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseGetAds, error)
	Update(ctx context.Context, in *RequestUpdateAd, opts ...grpc.CallOption) (*ResponseStatus, error)
	Add(ctx context.Context, in *RequestAdd, opts ...grpc.CallOption) (*ResponseStatus, error)
	Delete(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*ResponseStatus, error)
	Size(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseSize, error)
	AddAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*ResponseStatus, error)
	GetAccounts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseAccounts, error)
	UpdateTokenCurrentAcc(ctx context.Context, in *RequestUpdateAcc, opts ...grpc.CallOption) (*ResponseStatus, error)
}

type serviceProtobufClient struct {
	cc *grpc.ClientConn
}

func NewServiceProtobufClient(cc *grpc.ClientConn) ServiceProtobufClient {
	return &serviceProtobufClient{cc}
}

func (c *serviceProtobufClient) Get(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*ResponseGetAds, error) {
	out := new(ResponseGetAds)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseGetAds, error) {
	out := new(ResponseGetAds)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) Update(ctx context.Context, in *RequestUpdateAd, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) Add(ctx context.Context, in *RequestAdd, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) Delete(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) Size(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseSize, error) {
	out := new(ResponseSize)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/Size", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) AddAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/AddAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) GetAccounts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ResponseAccounts, error) {
	out := new(ResponseAccounts)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProtobufClient) UpdateTokenCurrentAcc(ctx context.Context, in *RequestUpdateAcc, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, "/serviceProtobuf.ServiceProtobuf/UpdateTokenCurrentAcc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceProtobufServer is the server API for ServiceProtobuf service.
type ServiceProtobufServer interface {
	Get(context.Context, *RequestID) (*ResponseGetAds, error)
	GetAll(context.Context, *empty.Empty) (*ResponseGetAds, error)
	Update(context.Context, *RequestUpdateAd) (*ResponseStatus, error)
	Add(context.Context, *RequestAdd) (*ResponseStatus, error)
	Delete(context.Context, *RequestID) (*ResponseStatus, error)
	Size(context.Context, *empty.Empty) (*ResponseSize, error)
	AddAccount(context.Context, *Account) (*ResponseStatus, error)
	GetAccounts(context.Context, *empty.Empty) (*ResponseAccounts, error)
	UpdateTokenCurrentAcc(context.Context, *RequestUpdateAcc) (*ResponseStatus, error)
}

// UnimplementedServiceProtobufServer can be embedded to have forward compatible implementations.
type UnimplementedServiceProtobufServer struct {
}

func (*UnimplementedServiceProtobufServer) Get(ctx context.Context, req *RequestID) (*ResponseGetAds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedServiceProtobufServer) GetAll(ctx context.Context, req *empty.Empty) (*ResponseGetAds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedServiceProtobufServer) Update(ctx context.Context, req *RequestUpdateAd) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedServiceProtobufServer) Add(ctx context.Context, req *RequestAdd) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedServiceProtobufServer) Delete(ctx context.Context, req *RequestID) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedServiceProtobufServer) Size(ctx context.Context, req *empty.Empty) (*ResponseSize, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Size not implemented")
}
func (*UnimplementedServiceProtobufServer) AddAccount(ctx context.Context, req *Account) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount not implemented")
}
func (*UnimplementedServiceProtobufServer) GetAccounts(ctx context.Context, req *empty.Empty) (*ResponseAccounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (*UnimplementedServiceProtobufServer) UpdateTokenCurrentAcc(ctx context.Context, req *RequestUpdateAcc) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTokenCurrentAcc not implemented")
}

func RegisterServiceProtobufServer(s *grpc.Server, srv ServiceProtobufServer) {
	s.RegisterService(&_ServiceProtobuf_serviceDesc, srv)
}

func _ServiceProtobuf_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).Get(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).GetAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateAd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).Update(ctx, req.(*RequestUpdateAd))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAdd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).Add(ctx, req.(*RequestAdd))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).Delete(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_Size_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).Size(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/Size",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).Size(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_AddAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).AddAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/AddAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).AddAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).GetAccounts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProtobuf_UpdateTokenCurrentAcc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateAcc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProtobufServer).UpdateTokenCurrentAcc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceProtobuf.ServiceProtobuf/UpdateTokenCurrentAcc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProtobufServer).UpdateTokenCurrentAcc(ctx, req.(*RequestUpdateAcc))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceProtobuf_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serviceProtobuf.ServiceProtobuf",
	HandlerType: (*ServiceProtobufServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ServiceProtobuf_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ServiceProtobuf_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServiceProtobuf_Update_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _ServiceProtobuf_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ServiceProtobuf_Delete_Handler,
		},
		{
			MethodName: "Size",
			Handler:    _ServiceProtobuf_Size_Handler,
		},
		{
			MethodName: "AddAccount",
			Handler:    _ServiceProtobuf_AddAccount_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _ServiceProtobuf_GetAccounts_Handler,
		},
		{
			MethodName: "UpdateTokenCurrentAcc",
			Handler:    _ServiceProtobuf_UpdateTokenCurrentAcc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}