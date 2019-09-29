// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cyclone.proto

package cyclone_test

import (
	_ "cyclone/healthy"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Response_ResponseResult int32

const (
	Response_Success Response_ResponseResult = 0
	Response_Fail    Response_ResponseResult = -1
)

var Response_ResponseResult_name = map[int32]string{
	0:  "Success",
	-1: "Fail",
}

var Response_ResponseResult_value = map[string]int32{
	"Success": 0,
	"Fail":    -1,
}

func (x Response_ResponseResult) String() string {
	return proto.EnumName(Response_ResponseResult_name, int32(x))
}

func (Response_ResponseResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_985224f23f3b1254, []int{2, 0}
}

type ErrorStatus struct {
	Message              string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Details              []*any.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ErrorStatus) Reset()         { *m = ErrorStatus{} }
func (m *ErrorStatus) String() string { return proto.CompactTextString(m) }
func (*ErrorStatus) ProtoMessage()    {}
func (*ErrorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_985224f23f3b1254, []int{0}
}

func (m *ErrorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorStatus.Unmarshal(m, b)
}
func (m *ErrorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorStatus.Marshal(b, m, deterministic)
}
func (m *ErrorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorStatus.Merge(m, src)
}
func (m *ErrorStatus) XXX_Size() int {
	return xxx_messageInfo_ErrorStatus.Size(m)
}
func (m *ErrorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorStatus proto.InternalMessageInfo

func (m *ErrorStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorStatus) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_985224f23f3b1254, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	Code                 Response_ResponseResult `protobuf:"varint,1,opt,name=code,proto3,enum=cyclone.test.Response_ResponseResult" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_985224f23f3b1254, []int{2}
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

func (m *Response) GetCode() Response_ResponseResult {
	if m != nil {
		return m.Code
	}
	return Response_Success
}

func init() {
	proto.RegisterEnum("cyclone.test.Response_ResponseResult", Response_ResponseResult_name, Response_ResponseResult_value)
	proto.RegisterType((*ErrorStatus)(nil), "cyclone.test.ErrorStatus")
	proto.RegisterType((*Request)(nil), "cyclone.test.Request")
	proto.RegisterType((*Response)(nil), "cyclone.test.Response")
}

func init() { proto.RegisterFile("cyclone.proto", fileDescriptor_985224f23f3b1254) }

var fileDescriptor_985224f23f3b1254 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x37, 0x1d, 0xc6, 0xbd, 0xba, 0x31, 0x83, 0x4a, 0xed, 0xc5, 0x52, 0x10, 0x76, 0xca,
	0xa4, 0x9e, 0xf4, 0x26, 0x32, 0xf1, 0xe4, 0x21, 0x3b, 0x88, 0xc7, 0xac, 0x7b, 0x76, 0x83, 0xd8,
	0xcc, 0xbe, 0x04, 0xe9, 0x5f, 0xaf, 0xd0, 0x2c, 0x9b, 0x8a, 0xf6, 0xf4, 0xbe, 0xbe, 0x2f, 0xbf,
	0xf7, 0xf1, 0xc1, 0xa0, 0x68, 0x0a, 0x6d, 0x2a, 0x14, 0xeb, 0xda, 0x58, 0xc3, 0x8f, 0x82, 0xb4,
	0x48, 0x36, 0x19, 0x2c, 0x51, 0x69, 0xbb, 0x6c, 0xfc, 0x32, 0x39, 0x2f, 0x8d, 0x29, 0x35, 0x4e,
	0x5a, 0x35, 0x77, 0xaf, 0x13, 0x55, 0x6d, 0x56, 0xd9, 0x33, 0x44, 0xd3, 0xba, 0x36, 0xf5, 0xcc,
	0x2a, 0xeb, 0x88, 0xc7, 0xc0, 0xde, 0x90, 0x48, 0x95, 0x18, 0x77, 0xd3, 0xee, 0xb8, 0x2f, 0x83,
	0xe4, 0x02, 0xd8, 0x02, 0xad, 0x5a, 0x69, 0x8a, 0xf7, 0xd2, 0xfd, 0x71, 0x94, 0x9f, 0x08, 0x4f,
	0x15, 0x81, 0x2a, 0xee, 0xaa, 0x46, 0x06, 0x53, 0xd6, 0x07, 0x26, 0xf1, 0xdd, 0x21, 0xd9, 0xec,
	0x03, 0x0e, 0x25, 0xd2, 0xda, 0x54, 0x84, 0xfc, 0x06, 0x7a, 0x85, 0x59, 0x78, 0xfa, 0x30, 0xbf,
	0x14, 0xdf, 0x63, 0x8b, 0xe0, 0xda, 0x0e, 0x12, 0xc9, 0x69, 0x2b, 0xdb, 0x27, 0xd9, 0x15, 0x0c,
	0x7f, 0xfe, 0xe7, 0x11, 0xb0, 0x99, 0x2b, 0x0a, 0x24, 0x1a, 0x75, 0xf8, 0x31, 0xf4, 0x1e, 0xd4,
	0x4a, 0x8f, 0x3e, 0xc3, 0xd7, 0xcd, 0xa7, 0xc0, 0xee, 0x3d, 0x9f, 0xdf, 0xee, 0xc6, 0xd3, 0xdf,
	0x47, 0xdb, 0x94, 0xc9, 0xd9, 0xdf, 0x59, 0xb2, 0x4e, 0xfe, 0x02, 0xec, 0xd1, 0xf7, 0xc9, 0x9f,
	0x76, 0xe3, 0xc5, 0xd6, 0x1f, 0xca, 0xde, 0x1c, 0x08, 0xc0, 0xf4, 0x7f, 0x43, 0x40, 0xcf, 0x0f,
	0xda, 0xf2, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42, 0x3a, 0xdc, 0xe1, 0xce, 0x01, 0x00,
	0x00,
}
