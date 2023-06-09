// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: service.rpc.assets/proto/asset.proto

package assetsproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AssetType int32

const (
	AssetType_ASSET_TYPE_UNKNOWN      AssetType = 0
	AssetType_ASSET_TYPE_IE00B52L4369 AssetType = 1
	AssetType_ASSET_TYPE_GB00BQ1YHQ70 AssetType = 2
	AssetType_ASSET_TYPE_GB00B3X7QG63 AssetType = 3
	AssetType_ASSET_TYPE_GB00BG0QP828 AssetType = 4
	AssetType_ASSET_TYPE_GB00BPN5P238 AssetType = 5
	AssetType_ASSET_TYPE_IE00B1S74Q32 AssetType = 6
)

// Enum value maps for AssetType.
var (
	AssetType_name = map[int32]string{
		0: "ASSET_TYPE_UNKNOWN",
		1: "ASSET_TYPE_IE00B52L4369",
		2: "ASSET_TYPE_GB00BQ1YHQ70",
		3: "ASSET_TYPE_GB00B3X7QG63",
		4: "ASSET_TYPE_GB00BG0QP828",
		5: "ASSET_TYPE_GB00BPN5P238",
		6: "ASSET_TYPE_IE00B1S74Q32",
	}
	AssetType_value = map[string]int32{
		"ASSET_TYPE_UNKNOWN":      0,
		"ASSET_TYPE_IE00B52L4369": 1,
		"ASSET_TYPE_GB00BQ1YHQ70": 2,
		"ASSET_TYPE_GB00B3X7QG63": 3,
		"ASSET_TYPE_GB00BG0QP828": 4,
		"ASSET_TYPE_GB00BPN5P238": 5,
		"ASSET_TYPE_IE00B1S74Q32": 6,
	}
)

func (x AssetType) Enum() *AssetType {
	p := new(AssetType)
	*p = x
	return p
}

func (x AssetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_rpc_assets_proto_asset_proto_enumTypes[0].Descriptor()
}

func (AssetType) Type() protoreflect.EnumType {
	return &file_service_rpc_assets_proto_asset_proto_enumTypes[0]
}

func (x AssetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetType.Descriptor instead.
func (AssetType) EnumDescriptor() ([]byte, []int) {
	return file_service_rpc_assets_proto_asset_proto_rawDescGZIP(), []int{0}
}

type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UserId    string                 `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Type      AssetType              `protobuf:"varint,5,opt,name=type,proto3,enum=emre.AssetType" json:"type,omitempty"`
	Units     int32                  `protobuf:"varint,6,opt,name=units,proto3" json:"units,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rpc_assets_proto_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_service_rpc_assets_proto_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_service_rpc_assets_proto_asset_proto_rawDescGZIP(), []int{0}
}

func (x *Asset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Asset) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Asset) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Asset) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Asset) GetType() AssetType {
	if x != nil {
		return x.Type
	}
	return AssetType_ASSET_TYPE_UNKNOWN
}

func (x *Asset) GetUnits() int32 {
	if x != nil {
		return x.Units
	}
	return 0
}

var File_service_rpc_assets_proto_asset_proto protoreflect.FileDescriptor

var file_service_rpc_assets_proto_asset_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x6d, 0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01,
	0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x6d, 0x72, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x2a, 0xd1, 0x01, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x45, 0x30, 0x30, 0x42, 0x35, 0x32, 0x4c, 0x34, 0x33,
	0x36, 0x39, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x47, 0x42, 0x30, 0x30, 0x42, 0x51, 0x31, 0x59, 0x48, 0x51, 0x37, 0x30, 0x10,
	0x02, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x47, 0x42, 0x30, 0x30, 0x42, 0x33, 0x58, 0x37, 0x51, 0x47, 0x36, 0x33, 0x10, 0x03, 0x12, 0x1b,
	0x0a, 0x17, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x42, 0x30,
	0x30, 0x42, 0x47, 0x30, 0x51, 0x50, 0x38, 0x32, 0x38, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x41,
	0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x42, 0x30, 0x30, 0x42, 0x50,
	0x4e, 0x35, 0x50, 0x32, 0x33, 0x38, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x53, 0x53, 0x45,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x45, 0x30, 0x30, 0x42, 0x31, 0x53, 0x37, 0x34,
	0x51, 0x33, 0x32, 0x10, 0x06, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x73, 0x6b, 0x69, 0x73, 0x61, 0x6e,
	0x2f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_rpc_assets_proto_asset_proto_rawDescOnce sync.Once
	file_service_rpc_assets_proto_asset_proto_rawDescData = file_service_rpc_assets_proto_asset_proto_rawDesc
)

func file_service_rpc_assets_proto_asset_proto_rawDescGZIP() []byte {
	file_service_rpc_assets_proto_asset_proto_rawDescOnce.Do(func() {
		file_service_rpc_assets_proto_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_rpc_assets_proto_asset_proto_rawDescData)
	})
	return file_service_rpc_assets_proto_asset_proto_rawDescData
}

var file_service_rpc_assets_proto_asset_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_rpc_assets_proto_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_rpc_assets_proto_asset_proto_goTypes = []interface{}{
	(AssetType)(0),                // 0: emre.AssetType
	(*Asset)(nil),                 // 1: emre.Asset
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_service_rpc_assets_proto_asset_proto_depIdxs = []int32{
	2, // 0: emre.Asset.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: emre.Asset.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: emre.Asset.type:type_name -> emre.AssetType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_rpc_assets_proto_asset_proto_init() }
func file_service_rpc_assets_proto_asset_proto_init() {
	if File_service_rpc_assets_proto_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_rpc_assets_proto_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
			RawDescriptor: file_service_rpc_assets_proto_asset_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_rpc_assets_proto_asset_proto_goTypes,
		DependencyIndexes: file_service_rpc_assets_proto_asset_proto_depIdxs,
		EnumInfos:         file_service_rpc_assets_proto_asset_proto_enumTypes,
		MessageInfos:      file_service_rpc_assets_proto_asset_proto_msgTypes,
	}.Build()
	File_service_rpc_assets_proto_asset_proto = out.File
	file_service_rpc_assets_proto_asset_proto_rawDesc = nil
	file_service_rpc_assets_proto_asset_proto_goTypes = nil
	file_service_rpc_assets_proto_asset_proto_depIdxs = nil
}
