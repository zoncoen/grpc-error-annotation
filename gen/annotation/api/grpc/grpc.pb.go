// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/api/grpc/grpc.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Error struct {
	Cancelled            *ErrorDescription `protobuf:"bytes,1,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	Unknown              *ErrorDescription `protobuf:"bytes,2,opt,name=unknown,proto3" json:"unknown,omitempty"`
	InvalidArgument      *ErrorDescription `protobuf:"bytes,3,opt,name=invalid_argument,json=invalidArgument,proto3" json:"invalid_argument,omitempty"`
	DeadlineExceeded     *ErrorDescription `protobuf:"bytes,4,opt,name=deadline_exceeded,json=deadlineExceeded,proto3" json:"deadline_exceeded,omitempty"`
	NotFound             *ErrorDescription `protobuf:"bytes,5,opt,name=not_found,json=notFound,proto3" json:"not_found,omitempty"`
	AlreadyExists        *ErrorDescription `protobuf:"bytes,6,opt,name=already_exists,json=alreadyExists,proto3" json:"already_exists,omitempty"`
	PermissionDenied     *ErrorDescription `protobuf:"bytes,7,opt,name=permission_denied,json=permissionDenied,proto3" json:"permission_denied,omitempty"`
	Unauthenticated      *ErrorDescription `protobuf:"bytes,16,opt,name=unauthenticated,proto3" json:"unauthenticated,omitempty"`
	ResourceExhausted    *ErrorDescription `protobuf:"bytes,8,opt,name=resource_exhausted,json=resourceExhausted,proto3" json:"resource_exhausted,omitempty"`
	FailedPrecondition   *ErrorDescription `protobuf:"bytes,9,opt,name=failed_precondition,json=failedPrecondition,proto3" json:"failed_precondition,omitempty"`
	Aborted              *ErrorDescription `protobuf:"bytes,10,opt,name=aborted,proto3" json:"aborted,omitempty"`
	OutOfRange           *ErrorDescription `protobuf:"bytes,11,opt,name=out_of_range,json=outOfRange,proto3" json:"out_of_range,omitempty"`
	Unimplemented        *ErrorDescription `protobuf:"bytes,12,opt,name=unimplemented,proto3" json:"unimplemented,omitempty"`
	Internal             *ErrorDescription `protobuf:"bytes,13,opt,name=internal,proto3" json:"internal,omitempty"`
	Unavailable          *ErrorDescription `protobuf:"bytes,14,opt,name=unavailable,proto3" json:"unavailable,omitempty"`
	DataLoss             *ErrorDescription `protobuf:"bytes,15,opt,name=data_loss,json=dataLoss,proto3" json:"data_loss,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbbfdc47eb9f8cc, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCancelled() *ErrorDescription {
	if m != nil {
		return m.Cancelled
	}
	return nil
}

func (m *Error) GetUnknown() *ErrorDescription {
	if m != nil {
		return m.Unknown
	}
	return nil
}

func (m *Error) GetInvalidArgument() *ErrorDescription {
	if m != nil {
		return m.InvalidArgument
	}
	return nil
}

func (m *Error) GetDeadlineExceeded() *ErrorDescription {
	if m != nil {
		return m.DeadlineExceeded
	}
	return nil
}

func (m *Error) GetNotFound() *ErrorDescription {
	if m != nil {
		return m.NotFound
	}
	return nil
}

func (m *Error) GetAlreadyExists() *ErrorDescription {
	if m != nil {
		return m.AlreadyExists
	}
	return nil
}

func (m *Error) GetPermissionDenied() *ErrorDescription {
	if m != nil {
		return m.PermissionDenied
	}
	return nil
}

func (m *Error) GetUnauthenticated() *ErrorDescription {
	if m != nil {
		return m.Unauthenticated
	}
	return nil
}

func (m *Error) GetResourceExhausted() *ErrorDescription {
	if m != nil {
		return m.ResourceExhausted
	}
	return nil
}

func (m *Error) GetFailedPrecondition() *ErrorDescription {
	if m != nil {
		return m.FailedPrecondition
	}
	return nil
}

func (m *Error) GetAborted() *ErrorDescription {
	if m != nil {
		return m.Aborted
	}
	return nil
}

func (m *Error) GetOutOfRange() *ErrorDescription {
	if m != nil {
		return m.OutOfRange
	}
	return nil
}

func (m *Error) GetUnimplemented() *ErrorDescription {
	if m != nil {
		return m.Unimplemented
	}
	return nil
}

func (m *Error) GetInternal() *ErrorDescription {
	if m != nil {
		return m.Internal
	}
	return nil
}

func (m *Error) GetUnavailable() *ErrorDescription {
	if m != nil {
		return m.Unavailable
	}
	return nil
}

func (m *Error) GetDataLoss() *ErrorDescription {
	if m != nil {
		return m.DataLoss
	}
	return nil
}

type ErrorDescription struct {
	Causes               []string `protobuf:"bytes,1,rep,name=causes,proto3" json:"causes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDescription) Reset()         { *m = ErrorDescription{} }
func (m *ErrorDescription) String() string { return proto.CompactTextString(m) }
func (*ErrorDescription) ProtoMessage()    {}
func (*ErrorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbbfdc47eb9f8cc, []int{1}
}
func (m *ErrorDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDescription.Unmarshal(m, b)
}
func (m *ErrorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDescription.Marshal(b, m, deterministic)
}
func (m *ErrorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDescription.Merge(m, src)
}
func (m *ErrorDescription) XXX_Size() int {
	return xxx_messageInfo_ErrorDescription.Size(m)
}
func (m *ErrorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDescription proto.InternalMessageInfo

func (m *ErrorDescription) GetCauses() []string {
	if m != nil {
		return m.Causes
	}
	return nil
}

var E_Error = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Error)(nil),
	Field:         56789,
	Name:          "zoncoen.api.grpc.error",
	Tag:           "bytes,56789,opt,name=error",
	Filename:      "proto/api/grpc/grpc.proto",
}

func init() {
	proto.RegisterType((*Error)(nil), "zoncoen.api.grpc.Error")
	proto.RegisterType((*ErrorDescription)(nil), "zoncoen.api.grpc.ErrorDescription")
	proto.RegisterExtension(E_Error)
}

func init() { proto.RegisterFile("proto/api/grpc/grpc.proto", fileDescriptor_fdbbfdc47eb9f8cc) }

var fileDescriptor_fdbbfdc47eb9f8cc = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5d, 0x8b, 0x13, 0x3d,
	0x14, 0xc7, 0xd9, 0x67, 0x9f, 0x7d, 0x69, 0xf6, 0xa5, 0xdd, 0x08, 0x1a, 0xbd, 0x90, 0xd2, 0xab,
	0x45, 0xd8, 0x19, 0xd0, 0x3b, 0x15, 0xdf, 0x68, 0x7d, 0x81, 0x5d, 0xaa, 0xf5, 0xce, 0x9b, 0x21,
	0x9d, 0x9c, 0x4e, 0x83, 0xe9, 0x39, 0x43, 0x5e, 0xd6, 0xd5, 0x2f, 0xe1, 0x07, 0xf2, 0x73, 0xf8,
	0x7d, 0x24, 0xd3, 0x19, 0xb6, 0x16, 0x84, 0xdc, 0x0c, 0x24, 0x39, 0xbf, 0x1f, 0x27, 0xe1, 0xcc,
	0x9f, 0xdd, 0xaf, 0x2d, 0x79, 0xca, 0x65, 0xad, 0xf3, 0xca, 0xd6, 0x65, 0xf3, 0xc9, 0x9a, 0x3d,
	0x3e, 0xf8, 0x41, 0x58, 0x12, 0x60, 0x26, 0x6b, 0x9d, 0xc5, 0xfd, 0x07, 0xc3, 0x8a, 0xa8, 0x32,
	0x90, 0x37, 0xe7, 0xf3, 0xb0, 0xc8, 0x15, 0xb8, 0xd2, 0xea, 0xda, 0x93, 0x5d, 0x33, 0xa3, 0x9f,
	0x3d, 0xb6, 0x37, 0xb1, 0x96, 0x2c, 0x7f, 0xc5, 0x7a, 0xa5, 0xc4, 0x12, 0x8c, 0x01, 0x25, 0x76,
	0x86, 0x3b, 0xe7, 0x47, 0x8f, 0x47, 0xd9, 0xb6, 0x31, 0x6b, 0x6a, 0xc7, 0xad, 0x45, 0x13, 0xce,
	0x6e, 0x21, 0xfe, 0x9c, 0x1d, 0x04, 0xfc, 0x8a, 0xf4, 0x0d, 0xc5, 0x7f, 0xc9, 0x7c, 0x87, 0xf0,
	0x2b, 0x36, 0xd0, 0x78, 0x2d, 0x8d, 0x56, 0x85, 0xb4, 0x55, 0x58, 0x01, 0x7a, 0xb1, 0x9b, 0xac,
	0xe9, 0xb7, 0xec, 0xeb, 0x16, 0xe5, 0x53, 0x76, 0xa6, 0x40, 0x2a, 0xa3, 0x11, 0x0a, 0xb8, 0x29,
	0x01, 0x14, 0x28, 0xf1, 0x7f, 0xb2, 0x6f, 0xd0, 0xc1, 0x93, 0x96, 0xe5, 0x2f, 0x59, 0x0f, 0xc9,
	0x17, 0x0b, 0x0a, 0xa8, 0xc4, 0x5e, 0xb2, 0xe8, 0x10, 0xc9, 0xbf, 0x8d, 0x0c, 0xff, 0xc0, 0x4e,
	0xa5, 0xb1, 0x20, 0xd5, 0xf7, 0x02, 0x6e, 0xb4, 0xf3, 0x4e, 0xec, 0x27, 0x5b, 0x4e, 0x5a, 0x72,
	0xd2, 0x80, 0xf1, 0x72, 0x35, 0xd8, 0x95, 0x76, 0x4e, 0x13, 0x16, 0x0a, 0x50, 0x83, 0x12, 0x07,
	0xe9, 0x97, 0xbb, 0x85, 0xc7, 0x0d, 0xcb, 0x2f, 0x59, 0x3f, 0xa0, 0x0c, 0x7e, 0x09, 0xe8, 0x75,
	0x29, 0x3d, 0x28, 0x31, 0x48, 0x7f, 0xfb, 0x2d, 0x94, 0x7f, 0x62, 0xdc, 0x82, 0xa3, 0x60, 0xcb,
	0xf8, 0xf6, 0x4b, 0x19, 0x5c, 0x14, 0x1e, 0x26, 0x0b, 0xcf, 0x3a, 0x7a, 0xd2, 0xc1, 0xfc, 0x33,
	0xbb, 0xb3, 0x90, 0xda, 0x80, 0x2a, 0x6a, 0x0b, 0x25, 0xa1, 0xd2, 0xb1, 0x52, 0xf4, 0x92, 0x9d,
	0x7c, 0x8d, 0x7f, 0xdc, 0xa0, 0xe3, 0xc0, 0xca, 0x39, 0xd9, 0xd8, 0x1c, 0x4b, 0x1f, 0xd8, 0x16,
	0xe1, 0x63, 0x76, 0x4c, 0xc1, 0x17, 0xb4, 0x28, 0xac, 0xc4, 0x0a, 0xc4, 0x51, 0xb2, 0x82, 0x51,
	0xf0, 0xd3, 0xc5, 0x2c, 0x52, 0xfc, 0x3d, 0x3b, 0x09, 0xa8, 0x57, 0xb5, 0x81, 0x38, 0xb6, 0xa0,
	0xc4, 0x71, 0xfa, 0x50, 0xfc, 0x05, 0xf2, 0x17, 0xec, 0x50, 0xa3, 0x07, 0x8b, 0xd2, 0x88, 0x93,
	0xf4, 0xf9, 0xec, 0x18, 0x3e, 0x66, 0x47, 0x01, 0xe5, 0xb5, 0xd4, 0x46, 0xce, 0x0d, 0x88, 0xd3,
	0x64, 0xc5, 0x26, 0x16, 0x7f, 0x13, 0x25, 0xbd, 0x2c, 0x0c, 0x39, 0x27, 0xfa, 0xe9, 0x6d, 0x44,
	0xe8, 0x92, 0x9c, 0x1b, 0x3d, 0x62, 0x83, 0xed, 0x53, 0x7e, 0x97, 0xed, 0x97, 0x32, 0x38, 0x70,
	0x62, 0x67, 0xb8, 0x7b, 0xde, 0x9b, 0xb5, 0xab, 0xa7, 0x53, 0xb6, 0x07, 0x4d, 0x78, 0x3d, 0xcc,
	0xd6, 0x49, 0x97, 0x75, 0x49, 0x97, 0x5d, 0x81, 0x5f, 0x92, 0x9a, 0x36, 0xbc, 0x13, 0xbf, 0x7f,
	0xad, 0xa3, 0xe4, 0xde, 0x3f, 0x5a, 0x99, 0xad, 0x3d, 0x6f, 0xde, 0x7d, 0x99, 0x54, 0xda, 0x2f,
	0xc3, 0x3c, 0x2b, 0x69, 0x95, 0xb7, 0xb5, 0x4d, 0xc6, 0x5e, 0x34, 0x05, 0x17, 0x12, 0x91, 0xbc,
	0x8c, 0xd2, 0xbc, 0x02, 0xcc, 0x37, 0x96, 0x5d, 0x22, 0x3f, 0x8b, 0x9f, 0xf9, 0x7e, 0xd3, 0xc8,
	0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x14, 0x5f, 0xc6, 0xaf, 0x05, 0x00, 0x00,
}