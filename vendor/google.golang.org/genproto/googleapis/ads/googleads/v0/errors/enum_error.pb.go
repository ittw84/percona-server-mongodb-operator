// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/enum_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible enum errors.
type EnumErrorEnum_EnumError int32

const (
	// Enum unspecified.
	EnumErrorEnum_UNSPECIFIED EnumErrorEnum_EnumError = 0
	// The received error code is not known in this version.
	EnumErrorEnum_UNKNOWN EnumErrorEnum_EnumError = 1
	// The enum value is not permitted.
	EnumErrorEnum_ENUM_VALUE_NOT_PERMITTED EnumErrorEnum_EnumError = 3
)

var EnumErrorEnum_EnumError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "ENUM_VALUE_NOT_PERMITTED",
}
var EnumErrorEnum_EnumError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"ENUM_VALUE_NOT_PERMITTED": 3,
}

func (x EnumErrorEnum_EnumError) String() string {
	return proto.EnumName(EnumErrorEnum_EnumError_name, int32(x))
}
func (EnumErrorEnum_EnumError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_error_80a12d5eeb48484c, []int{0, 0}
}

// Container for enum describing possible enum errors.
type EnumErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumErrorEnum) Reset()         { *m = EnumErrorEnum{} }
func (m *EnumErrorEnum) String() string { return proto.CompactTextString(m) }
func (*EnumErrorEnum) ProtoMessage()    {}
func (*EnumErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_enum_error_80a12d5eeb48484c, []int{0}
}
func (m *EnumErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumErrorEnum.Unmarshal(m, b)
}
func (m *EnumErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumErrorEnum.Marshal(b, m, deterministic)
}
func (dst *EnumErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumErrorEnum.Merge(dst, src)
}
func (m *EnumErrorEnum) XXX_Size() int {
	return xxx_messageInfo_EnumErrorEnum.Size(m)
}
func (m *EnumErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_EnumErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EnumErrorEnum)(nil), "google.ads.googleads.v0.errors.EnumErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.EnumErrorEnum_EnumError", EnumErrorEnum_EnumError_name, EnumErrorEnum_EnumError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/enum_error.proto", fileDescriptor_enum_error_80a12d5eeb48484c)
}

var fileDescriptor_enum_error_80a12d5eeb48484c = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x6d, 0x07, 0x8a, 0x19, 0x6a, 0xe9, 0xc9, 0x83, 0xec, 0xd0, 0x07, 0x48, 0x03, 0xde,
	0xe2, 0x29, 0xb3, 0xb1, 0x14, 0x5d, 0x56, 0xb4, 0xad, 0x22, 0x85, 0x52, 0x4d, 0x09, 0xc2, 0xda,
	0x8c, 0xc4, 0xed, 0x81, 0x3c, 0xfa, 0x28, 0xbe, 0x89, 0x3e, 0x85, 0xa4, 0xd9, 0x72, 0xd3, 0x53,
	0x7e, 0xf9, 0xf8, 0x7d, 0xf9, 0xfe, 0x5f, 0x40, 0x2c, 0xa4, 0x14, 0xab, 0x2e, 0x6e, 0xb9, 0xde,
	0xa1, 0xa1, 0x2d, 0x8a, 0x3b, 0xa5, 0xa4, 0xd2, 0x71, 0x37, 0x6c, 0xfa, 0x66, 0x64, 0xb8, 0x56,
	0xf2, 0x5d, 0x86, 0x33, 0x6b, 0xc1, 0x96, 0x6b, 0xe8, 0x1a, 0xe0, 0x16, 0x41, 0xdb, 0x10, 0x3d,
	0x81, 0x13, 0x3a, 0x6c, 0x7a, 0x6a, 0x6e, 0x06, 0xa2, 0x14, 0x1c, 0xbb, 0x42, 0x78, 0x06, 0xa6,
	0x25, 0x7b, 0xc8, 0xe9, 0x75, 0x76, 0x93, 0xd1, 0x24, 0x38, 0x08, 0xa7, 0xe0, 0xa8, 0x64, 0xb7,
	0x6c, 0xf9, 0xc8, 0x02, 0x2f, 0xbc, 0x00, 0xe7, 0x94, 0x95, 0x8b, 0xa6, 0x22, 0x77, 0x25, 0x6d,
	0xd8, 0xb2, 0x68, 0x72, 0x7a, 0xbf, 0xc8, 0x8a, 0x82, 0x26, 0xc1, 0x64, 0xfe, 0xed, 0x81, 0xe8,
	0x55, 0xf6, 0xf0, 0xff, 0x00, 0xf3, 0x53, 0x37, 0x2d, 0x37, 0x81, 0x73, 0xef, 0x39, 0xd9, 0x75,
	0x08, 0xb9, 0x6a, 0x07, 0x01, 0xa5, 0x12, 0xb1, 0xe8, 0x86, 0x71, 0x9d, 0xfd, 0xce, 0xeb, 0x37,
	0xfd, 0xd7, 0x17, 0x5c, 0xd9, 0xe3, 0xc3, 0x9f, 0xa4, 0x84, 0x7c, 0xfa, 0xb3, 0xd4, 0x3e, 0x46,
	0xb8, 0x86, 0x16, 0x0d, 0x55, 0x08, 0x8e, 0x23, 0xf5, 0xd7, 0x5e, 0xa8, 0x09, 0xd7, 0xb5, 0x13,
	0xea, 0x0a, 0xd5, 0x56, 0xf8, 0xf1, 0x23, 0x5b, 0xc5, 0x98, 0x70, 0x8d, 0xb1, 0x53, 0x30, 0xae,
	0x10, 0xc6, 0x56, 0x7a, 0x39, 0x1c, 0xd3, 0x5d, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x02,
	0x98, 0x9c, 0x9f, 0x01, 0x00, 0x00,
}