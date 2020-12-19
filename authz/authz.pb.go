// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: authz/authz.proto

package authz

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Roles int32

const (
	Roles_UNKNOWN    Roles = 0
	Roles_REFLECTION Roles = 1
	Roles_VIEWER     Roles = 2
	Roles_EDITOR     Roles = 3
	Roles_OWNER      Roles = 4
)

// Enum value maps for Roles.
var (
	Roles_name = map[int32]string{
		0: "UNKNOWN",
		1: "REFLECTION",
		2: "VIEWER",
		3: "EDITOR",
		4: "OWNER",
	}
	Roles_value = map[string]int32{
		"UNKNOWN":    0,
		"REFLECTION": 1,
		"VIEWER":     2,
		"EDITOR":     3,
		"OWNER":      4,
	}
)

func (x Roles) Enum() *Roles {
	p := new(Roles)
	*p = x
	return p
}

func (x Roles) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Roles) Descriptor() protoreflect.EnumDescriptor {
	return file_authz_authz_proto_enumTypes[0].Descriptor()
}

func (Roles) Type() protoreflect.EnumType {
	return &file_authz_authz_proto_enumTypes[0]
}

func (x Roles) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Roles.Descriptor instead.
func (Roles) EnumDescriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{0}
}

type MethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role        Roles    `protobuf:"varint,1,opt,name=role,proto3,enum=authz.Roles" json:"role,omitempty"`
	CustomRoles []string `protobuf:"bytes,2,rep,name=custom_roles,json=customRoles,proto3" json:"custom_roles,omitempty"`
}

func (x *MethodOptions) Reset() {
	*x = MethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptions) ProtoMessage() {}

func (x *MethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptions.ProtoReflect.Descriptor instead.
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{0}
}

func (x *MethodOptions) GetRole() Roles {
	if x != nil {
		return x.Role
	}
	return Roles_UNKNOWN
}

func (x *MethodOptions) GetCustomRoles() []string {
	if x != nil {
		return x.CustomRoles
	}
	return nil
}

type ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *ServiceOptions) Reset() {
	*x = ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOptions) ProtoMessage() {}

func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_authz_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return file_authz_authz_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceOptions) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

var file_authz_authz_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*MethodOptions)(nil),
		Field:         5999,
		Name:          "authz.method",
		Tag:           "bytes,5999,opt,name=method",
		Filename:      "authz/authz.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*ServiceOptions)(nil),
		Field:         5999,
		Name:          "authz.service",
		Tag:           "bytes,5999,opt,name=service",
		Filename:      "authz/authz.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// optional authz.MethodOptions method = 5999;
	E_Method = &file_authz_authz_proto_extTypes[0]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// optional authz.ServiceOptions service = 5999;
	E_Service = &file_authz_authz_proto_extTypes[1]
)

var File_authz_authz_proto protoreflect.FileDescriptor

var file_authz_authz_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x22, 0x2c, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x2a, 0x47, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x46, 0x4c, 0x45, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x04, 0x3a, 0x4d, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xef, 0x2e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x51, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xef, 0x2e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x71, 0x75, 0x65, 0x72, 0x6e,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authz_authz_proto_rawDescOnce sync.Once
	file_authz_authz_proto_rawDescData = file_authz_authz_proto_rawDesc
)

func file_authz_authz_proto_rawDescGZIP() []byte {
	file_authz_authz_proto_rawDescOnce.Do(func() {
		file_authz_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_authz_authz_proto_rawDescData)
	})
	return file_authz_authz_proto_rawDescData
}

var file_authz_authz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authz_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_authz_authz_proto_goTypes = []interface{}{
	(Roles)(0),                        // 0: authz.Roles
	(*MethodOptions)(nil),             // 1: authz.MethodOptions
	(*ServiceOptions)(nil),            // 2: authz.ServiceOptions
	(*descriptor.MethodOptions)(nil),  // 3: google.protobuf.MethodOptions
	(*descriptor.ServiceOptions)(nil), // 4: google.protobuf.ServiceOptions
}
var file_authz_authz_proto_depIdxs = []int32{
	0, // 0: authz.MethodOptions.role:type_name -> authz.Roles
	3, // 1: authz.method:extendee -> google.protobuf.MethodOptions
	4, // 2: authz.service:extendee -> google.protobuf.ServiceOptions
	1, // 3: authz.method:type_name -> authz.MethodOptions
	2, // 4: authz.service:type_name -> authz.ServiceOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_authz_authz_proto_init() }
func file_authz_authz_proto_init() {
	if File_authz_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authz_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_authz_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authz_authz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_authz_authz_proto_goTypes,
		DependencyIndexes: file_authz_authz_proto_depIdxs,
		EnumInfos:         file_authz_authz_proto_enumTypes,
		MessageInfos:      file_authz_authz_proto_msgTypes,
		ExtensionInfos:    file_authz_authz_proto_extTypes,
	}.Build()
	File_authz_authz_proto = out.File
	file_authz_authz_proto_rawDesc = nil
	file_authz_authz_proto_goTypes = nil
	file_authz_authz_proto_depIdxs = nil
}
