// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: solo-kit.proto

package core

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Resource struct {
	// becomes the kubernetes short name for the generated crd
	ShortName string `protobuf:"bytes,1,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	// becomes the kubernetes plural name for the generated crd
	PluralName string `protobuf:"bytes,2,opt,name=plural_name,json=pluralName,proto3" json:"plural_name,omitempty"`
	// the resource lives at the cluster level, namespace is ignored by the server
	ClusterScoped bool `protobuf:"varint,3,opt,name=cluster_scoped,json=clusterScoped,proto3" json:"cluster_scoped,omitempty"`
	// indicates whether documentation generation has to be skipped for the given resource, defaults to false
	SkipDocsGen bool `protobuf:"varint,4,opt,name=skip_docs_gen,json=skipDocsGen,proto3" json:"skip_docs_gen,omitempty"`
	// indicates whether annotations should be excluded from the resource's generated hash function.
	// if set to true, changes in annotations will not cause a new snapshot to be emitted
	SkipHashingAnnotations bool     `protobuf:"varint,5,opt,name=skip_hashing_annotations,json=skipHashingAnnotations,proto3" json:"skip_hashing_annotations,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2ca748c58e61805, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func (m *Resource) GetPluralName() string {
	if m != nil {
		return m.PluralName
	}
	return ""
}

func (m *Resource) GetClusterScoped() bool {
	if m != nil {
		return m.ClusterScoped
	}
	return false
}

func (m *Resource) GetSkipDocsGen() bool {
	if m != nil {
		return m.SkipDocsGen
	}
	return false
}

func (m *Resource) GetSkipHashingAnnotations() bool {
	if m != nil {
		return m.SkipHashingAnnotations
	}
	return false
}

var E_Resource = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*Resource)(nil),
	Field:         10000,
	Name:          "core.solo.io.resource",
	Tag:           "bytes,10000,opt,name=resource",
	Filename:      "solo-kit.proto",
}

var E_SkipHashing = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         10000,
	Name:          "core.solo.io.skip_hashing",
	Tag:           "varint,10000,opt,name=skip_hashing",
	Filename:      "solo-kit.proto",
}

func init() {
	proto.RegisterType((*Resource)(nil), "core.solo.io.Resource")
	proto.RegisterExtension(E_Resource)
	proto.RegisterExtension(E_SkipHashing)
}

func init() { proto.RegisterFile("solo-kit.proto", fileDescriptor_d2ca748c58e61805) }

var fileDescriptor_d2ca748c58e61805 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xc9, 0xff, 0x8d, 0x76, 0xfb, 0x72, 0xc8, 0xa1, 0x84, 0x3f, 0x94, 0x86, 0x82, 0xd0,
	0x8b, 0xbb, 0xa8, 0x88, 0xd2, 0x5b, 0x45, 0xd4, 0x8b, 0x15, 0xa2, 0x27, 0x2f, 0x61, 0xbb, 0x19,
	0x93, 0xa5, 0x69, 0x66, 0xd9, 0xd9, 0xf8, 0x39, 0xfc, 0x66, 0x7e, 0x25, 0xc9, 0xa6, 0xd1, 0x8a,
	0xd7, 0xe7, 0xf7, 0xf0, 0xcc, 0xcc, 0x33, 0x6c, 0x4c, 0x58, 0xe2, 0xf1, 0x56, 0x3b, 0x6e, 0x2c,
	0x3a, 0x0c, 0x87, 0x0a, 0x2d, 0xf0, 0x46, 0xe4, 0x1a, 0xff, 0xc7, 0x39, 0x62, 0x5e, 0x82, 0xf0,
	0x6c, 0x53, 0xbf, 0x88, 0x0c, 0x48, 0x59, 0x6d, 0x1c, 0xda, 0xd6, 0x3f, 0x7f, 0x0f, 0x58, 0x2f,
	0x01, 0xc2, 0xda, 0x2a, 0x08, 0xa7, 0x8c, 0x51, 0x81, 0xd6, 0xa5, 0x95, 0xdc, 0x41, 0x14, 0xc4,
	0xc1, 0xa2, 0x9f, 0xf4, 0xbd, 0xb2, 0x96, 0x3b, 0x08, 0x67, 0x6c, 0x60, 0xca, 0xda, 0xca, 0xb2,
	0xe5, 0xbf, 0x3c, 0x67, 0xad, 0xe4, 0x0d, 0x47, 0x6c, 0xac, 0xca, 0x9a, 0x1c, 0xd8, 0x94, 0x14,
	0x1a, 0xc8, 0xa2, 0xdf, 0x71, 0xb0, 0xe8, 0x25, 0xa3, 0xbd, 0xfa, 0xe8, 0xc5, 0x70, 0xce, 0x46,
	0xb4, 0xd5, 0x26, 0xcd, 0x50, 0x51, 0x9a, 0x43, 0x15, 0xfd, 0xf1, 0xae, 0x41, 0x23, 0x5e, 0xa3,
	0xa2, 0x5b, 0xa8, 0xc2, 0x4b, 0x16, 0x79, 0x4f, 0x21, 0xa9, 0xd0, 0x55, 0x9e, 0xca, 0xaa, 0x42,
	0x27, 0x9d, 0xc6, 0x8a, 0xa2, 0xbf, 0xde, 0x3e, 0x69, 0xf8, 0x5d, 0x8b, 0x57, 0x5f, 0x74, 0xf9,
	0xc4, 0x7a, 0xb6, 0x3b, 0x68, 0xc6, 0xdb, 0x02, 0x78, 0x57, 0x00, 0xbf, 0x07, 0x22, 0x99, 0xc3,
	0x83, 0x69, 0xa3, 0xde, 0xd6, 0x71, 0xb0, 0x18, 0x9c, 0x4e, 0xf8, 0x61, 0x6d, 0xbc, 0x2b, 0x24,
	0xf9, 0x4c, 0x5a, 0xae, 0xd8, 0xf0, 0x70, 0x9f, 0x70, 0xfa, 0x23, 0xf9, 0x46, 0x43, 0x99, 0x7d,
	0xcb, 0xdd, 0x9f, 0xb4, 0xdf, 0xf1, 0xea, 0xe2, 0xf9, 0x3c, 0xd7, 0xae, 0xa8, 0x37, 0x5c, 0xe1,
	0x4e, 0xf8, 0xbf, 0x69, 0x14, 0xdd, 0xff, 0x84, 0xd9, 0xe6, 0x42, 0x1a, 0x2d, 0x5e, 0x4f, 0x44,
	0x37, 0x95, 0x44, 0xb3, 0xd6, 0xe6, 0x9f, 0x9f, 0x71, 0xf6, 0x11, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x9a, 0x46, 0x9b, 0xec, 0x01, 0x00, 0x00,
}