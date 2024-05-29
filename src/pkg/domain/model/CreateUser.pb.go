// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: CreateUser.proto

package domain

import (
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

type CreateUserParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateUserParams) Reset() {
	*x = CreateUserParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CreateUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserParams) ProtoMessage() {}

func (x *CreateUserParams) ProtoReflect() protoreflect.Message {
	mi := &file_CreateUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserParams.ProtoReflect.Descriptor instead.
func (*CreateUserParams) Descriptor() ([]byte, []int) {
	return file_CreateUser_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserParams) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_CreateUser_proto protoreflect.FileDescriptor

var file_CreateUser_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x44, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CreateUser_proto_rawDescOnce sync.Once
	file_CreateUser_proto_rawDescData = file_CreateUser_proto_rawDesc
)

func file_CreateUser_proto_rawDescGZIP() []byte {
	file_CreateUser_proto_rawDescOnce.Do(func() {
		file_CreateUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_CreateUser_proto_rawDescData)
	})
	return file_CreateUser_proto_rawDescData
}

var file_CreateUser_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CreateUser_proto_goTypes = []interface{}{
	(*CreateUserParams)(nil), // 0: domain.CreateUserParams
}
var file_CreateUser_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CreateUser_proto_init() }
func file_CreateUser_proto_init() {
	if File_CreateUser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CreateUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserParams); i {
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
			RawDescriptor: file_CreateUser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CreateUser_proto_goTypes,
		DependencyIndexes: file_CreateUser_proto_depIdxs,
		MessageInfos:      file_CreateUser_proto_msgTypes,
	}.Build()
	File_CreateUser_proto = out.File
	file_CreateUser_proto_rawDesc = nil
	file_CreateUser_proto_goTypes = nil
	file_CreateUser_proto_depIdxs = nil
}
