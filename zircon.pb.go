// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: zircon.proto

package zircon_proto

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

type BlockResponse_BlockStatus int32

const (
	BlockResponse_INVALID   BlockResponse_BlockStatus = 0
	BlockResponse_LOW_SCORE BlockResponse_BlockStatus = 1
	BlockResponse_VALID     BlockResponse_BlockStatus = 2
	BlockResponse_WINNER    BlockResponse_BlockStatus = 3
)

// Enum value maps for BlockResponse_BlockStatus.
var (
	BlockResponse_BlockStatus_name = map[int32]string{
		0: "INVALID",
		1: "LOW_SCORE",
		2: "VALID",
		3: "WINNER",
	}
	BlockResponse_BlockStatus_value = map[string]int32{
		"INVALID":   0,
		"LOW_SCORE": 1,
		"VALID":     2,
		"WINNER":    3,
	}
)

func (x BlockResponse_BlockStatus) Enum() *BlockResponse_BlockStatus {
	p := new(BlockResponse_BlockStatus)
	*p = x
	return p
}

func (x BlockResponse_BlockStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockResponse_BlockStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_zircon_proto_enumTypes[0].Descriptor()
}

func (BlockResponse_BlockStatus) Type() protoreflect.EnumType {
	return &file_zircon_proto_enumTypes[0]
}

func (x BlockResponse_BlockStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockResponse_BlockStatus.Descriptor instead.
func (BlockResponse_BlockStatus) EnumDescriptor() ([]byte, []int) {
	return file_zircon_proto_rawDescGZIP(), []int{1, 0}
}

// A constructed block (hex_blob+nonces+seed_hash) should equal hex_result
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexBlob    string `protobuf:"bytes,1,opt,name=hex_blob,json=hexBlob,proto3" json:"hex_blob,omitempty"`
	HexNonce   string `protobuf:"bytes,2,opt,name=hex_nonce,json=hexNonce,proto3" json:"hex_nonce,omitempty"`
	ExtraNonce uint32 `protobuf:"varint,3,opt,name=extra_nonce,json=extraNonce,proto3" json:"extra_nonce,omitempty"`
	SeedHash   string `protobuf:"bytes,4,opt,name=seed_hash,json=seedHash,proto3" json:"seed_hash,omitempty"`
	// Actual result from miner, need to verify via rebuilding block
	HexResult string `protobuf:"bytes,5,opt,name=hex_result,json=hexResult,proto3" json:"hex_result,omitempty"`
	// Really a 256bit int
	LocalDiff  string `protobuf:"bytes,6,opt,name=local_diff,json=localDiff,proto3" json:"local_diff,omitempty"`
	GlobalDiff string `protobuf:"bytes,7,opt,name=global_diff,json=globalDiff,proto3" json:"global_diff,omitempty"`
	MinerId    string `protobuf:"bytes,8,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zircon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_zircon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_zircon_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetHexBlob() string {
	if x != nil {
		return x.HexBlob
	}
	return ""
}

func (x *Block) GetHexNonce() string {
	if x != nil {
		return x.HexNonce
	}
	return ""
}

func (x *Block) GetExtraNonce() uint32 {
	if x != nil {
		return x.ExtraNonce
	}
	return 0
}

func (x *Block) GetSeedHash() string {
	if x != nil {
		return x.SeedHash
	}
	return ""
}

func (x *Block) GetHexResult() string {
	if x != nil {
		return x.HexResult
	}
	return ""
}

func (x *Block) GetLocalDiff() string {
	if x != nil {
		return x.LocalDiff
	}
	return ""
}

func (x *Block) GetGlobalDiff() string {
	if x != nil {
		return x.GlobalDiff
	}
	return ""
}

func (x *Block) GetMinerId() string {
	if x != nil {
		return x.MinerId
	}
	return ""
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockStatus BlockResponse_BlockStatus `protobuf:"varint,1,opt,name=block_status,json=blockStatus,proto3,enum=zircon.BlockResponse_BlockStatus" json:"block_status,omitempty"`
	ShouldBan   bool                      `protobuf:"varint,2,opt,name=should_ban,json=shouldBan,proto3" json:"should_ban,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zircon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zircon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_zircon_proto_rawDescGZIP(), []int{1}
}

func (x *BlockResponse) GetBlockStatus() BlockResponse_BlockStatus {
	if x != nil {
		return x.BlockStatus
	}
	return BlockResponse_INVALID
}

func (x *BlockResponse) GetShouldBan() bool {
	if x != nil {
		return x.ShouldBan
	}
	return false
}

var File_zircon_proto protoreflect.FileDescriptor

var file_zircon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x7a, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x7a, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x22, 0xf7, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x65, 0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x78, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x68,
	0x65, 0x78, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x65, 0x78, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x65,
	0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x78, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65, 0x78, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x66, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x44, 0x69, 0x66, 0x66, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x66, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x44, 0x69, 0x66, 0x66, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xb6, 0x01, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x7a, 0x69, 0x72, 0x63, 0x6f,
	0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x5f, 0x62, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x68,
	0x6f, 0x75, 0x6c, 0x64, 0x42, 0x61, 0x6e, 0x22, 0x40, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x57, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x32, 0x41, 0x0a, 0x06, 0x5a, 0x69, 0x72,
	0x63, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x7a, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x1a, 0x15, 0x2e, 0x7a, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x72, 0x69, 0x61,
	0x64, 0x65, 0x69, 0x6e, 0x63, 0x2f, 0x7a, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zircon_proto_rawDescOnce sync.Once
	file_zircon_proto_rawDescData = file_zircon_proto_rawDesc
)

func file_zircon_proto_rawDescGZIP() []byte {
	file_zircon_proto_rawDescOnce.Do(func() {
		file_zircon_proto_rawDescData = protoimpl.X.CompressGZIP(file_zircon_proto_rawDescData)
	})
	return file_zircon_proto_rawDescData
}

var file_zircon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zircon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_zircon_proto_goTypes = []interface{}{
	(BlockResponse_BlockStatus)(0), // 0: zircon.BlockResponse.BlockStatus
	(*Block)(nil),                  // 1: zircon.Block
	(*BlockResponse)(nil),          // 2: zircon.BlockResponse
}
var file_zircon_proto_depIdxs = []int32{
	0, // 0: zircon.BlockResponse.block_status:type_name -> zircon.BlockResponse.BlockStatus
	1, // 1: zircon.Zircon.validateBlock:input_type -> zircon.Block
	2, // 2: zircon.Zircon.validateBlock:output_type -> zircon.BlockResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_zircon_proto_init() }
func file_zircon_proto_init() {
	if File_zircon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zircon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_zircon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
			RawDescriptor: file_zircon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zircon_proto_goTypes,
		DependencyIndexes: file_zircon_proto_depIdxs,
		EnumInfos:         file_zircon_proto_enumTypes,
		MessageInfos:      file_zircon_proto_msgTypes,
	}.Build()
	File_zircon_proto = out.File
	file_zircon_proto_rawDesc = nil
	file_zircon_proto_goTypes = nil
	file_zircon_proto_depIdxs = nil
}
