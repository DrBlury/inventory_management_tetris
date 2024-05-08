// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: Inventory.proto

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

type Inventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InventoryMeta *InventoryMeta   `protobuf:"bytes,1,opt,name=inventoryMeta,proto3" json:"inventoryMeta,omitempty"`
	Items         []*InventoryItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Inventory) Reset() {
	*x = Inventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inventory) ProtoMessage() {}

func (x *Inventory) ProtoReflect() protoreflect.Message {
	mi := &file_Inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inventory.ProtoReflect.Descriptor instead.
func (*Inventory) Descriptor() ([]byte, []int) {
	return file_Inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Inventory) GetInventoryMeta() *InventoryMeta {
	if x != nil {
		return x.InventoryMeta
	}
	return nil
}

func (x *Inventory) GetItems() []*InventoryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_Inventory_proto protoreflect.FileDescriptor

var file_Inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x13, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x3b, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0d,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2b, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Inventory_proto_rawDescOnce sync.Once
	file_Inventory_proto_rawDescData = file_Inventory_proto_rawDesc
)

func file_Inventory_proto_rawDescGZIP() []byte {
	file_Inventory_proto_rawDescOnce.Do(func() {
		file_Inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_Inventory_proto_rawDescData)
	})
	return file_Inventory_proto_rawDescData
}

var file_Inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Inventory_proto_goTypes = []interface{}{
	(*Inventory)(nil),     // 0: domain.Inventory
	(*InventoryMeta)(nil), // 1: domain.InventoryMeta
	(*InventoryItem)(nil), // 2: domain.InventoryItem
}
var file_Inventory_proto_depIdxs = []int32{
	1, // 0: domain.Inventory.inventoryMeta:type_name -> domain.InventoryMeta
	2, // 1: domain.Inventory.items:type_name -> domain.InventoryItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Inventory_proto_init() }
func file_Inventory_proto_init() {
	if File_Inventory_proto != nil {
		return
	}
	file_InventoryMeta_proto_init()
	file_InventoryItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inventory); i {
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
			RawDescriptor: file_Inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Inventory_proto_goTypes,
		DependencyIndexes: file_Inventory_proto_depIdxs,
		MessageInfos:      file_Inventory_proto_msgTypes,
	}.Build()
	File_Inventory_proto = out.File
	file_Inventory_proto_rawDesc = nil
	file_Inventory_proto_goTypes = nil
	file_Inventory_proto_depIdxs = nil
}
