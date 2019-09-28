// Code generated by protoc-gen-go. DO NOT EDIT.
// source: healthy.proto

package cyclone_healthy

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

type Response_ResponseResult int32

const (
	Response_Success Response_ResponseResult = 0
	Response_Fail    Response_ResponseResult = -1
	Response_Healthy Response_ResponseResult = 1
)

var Response_ResponseResult_name = map[int32]string{
	0:  "Success",
	-1: "Fail",
	1:  "Healthy",
}

var Response_ResponseResult_value = map[string]int32{
	"Success": 0,
	"Fail":    -1,
	"Healthy": 1,
}

func (x Response_ResponseResult) String() string {
	return proto.EnumName(Response_ResponseResult_name, int32(x))
}

func (Response_ResponseResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_741f22f95cb14d6b, []int{1, 0}
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
	return fileDescriptor_741f22f95cb14d6b, []int{0}
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
	Code                 Response_ResponseResult `protobuf:"varint,1,opt,name=code,proto3,enum=cyclone.healthy.Response_ResponseResult" json:"code,omitempty"`
	Msg                  string                  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_741f22f95cb14d6b, []int{1}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("cyclone.healthy.Response_ResponseResult", Response_ResponseResult_name, Response_ResponseResult_value)
	proto.RegisterType((*Request)(nil), "cyclone.healthy.Request")
	proto.RegisterType((*Response)(nil), "cyclone.healthy.Response")
}

func init() { proto.RegisterFile("healthy.proto", fileDescriptor_741f22f95cb14d6b) }

var fileDescriptor_741f22f95cb14d6b = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0xae, 0x4c, 0xce, 0xc9,
	0xcf, 0x4b, 0xd5, 0x83, 0x0a, 0x2b, 0x71, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x28, 0xcd, 0x64, 0xe4, 0xe2, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2, 0xe1,
	0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xd0, 0x43, 0xd3,
	0xa7, 0x07, 0x53, 0x08, 0x67, 0x04, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x04, 0x81, 0x75, 0x09, 0x09,
	0x70, 0x31, 0xe7, 0x16, 0xa7, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0xb6,
	0x5c, 0x7c, 0xa8, 0x2a, 0x85, 0xb8, 0xb9, 0xd8, 0x83, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x05,
	0x18, 0x84, 0x04, 0xb9, 0x58, 0xdc, 0x12, 0x33, 0x73, 0x04, 0xfe, 0xc3, 0x00, 0x23, 0x48, 0xde,
	0x03, 0x62, 0x99, 0x00, 0xa3, 0x91, 0x37, 0x9c, 0x23, 0xe4, 0x80, 0x60, 0x4a, 0x60, 0x71, 0x16,
	0xd8, 0x2f, 0x52, 0x92, 0x38, 0x1d, 0xac, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x0b, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x02, 0xd6, 0x7b, 0xcd, 0x1c, 0x01, 0x00, 0x00,
}