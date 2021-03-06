// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb/transfer.proto

package pb

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

type TransferReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action int32  `protobuf:"varint,1,opt,name=action,proto3" json:"action,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransferReq) Reset() {
	*x = TransferReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferReq) ProtoMessage() {}

func (x *TransferReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferReq.ProtoReflect.Descriptor instead.
func (*TransferReq) Descriptor() ([]byte, []int) {
	return file_pb_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferReq) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *TransferReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type TransferResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransferResp) Reset() {
	*x = TransferResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResp) ProtoMessage() {}

func (x *TransferResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResp.ProtoReflect.Descriptor instead.
func (*TransferResp) Descriptor() ([]byte, []int) {
	return file_pb_transfer_proto_rawDescGZIP(), []int{1}
}

var File_pb_transfer_proto protoreflect.FileDescriptor

var file_pb_transfer_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x39, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x0e, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x32, 0x37, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x06, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_transfer_proto_rawDescOnce sync.Once
	file_pb_transfer_proto_rawDescData = file_pb_transfer_proto_rawDesc
)

func file_pb_transfer_proto_rawDescGZIP() []byte {
	file_pb_transfer_proto_rawDescOnce.Do(func() {
		file_pb_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_transfer_proto_rawDescData)
	})
	return file_pb_transfer_proto_rawDescData
}

var file_pb_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_transfer_proto_goTypes = []interface{}{
	(*TransferReq)(nil),  // 0: pb.TransferReq
	(*TransferResp)(nil), // 1: pb.TransferResp
}
var file_pb_transfer_proto_depIdxs = []int32{
	0, // 0: pb.Transfer.invoke:input_type -> pb.TransferReq
	1, // 1: pb.Transfer.invoke:output_type -> pb.TransferResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_transfer_proto_init() }
func file_pb_transfer_proto_init() {
	if File_pb_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferReq); i {
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
		file_pb_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResp); i {
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
			RawDescriptor: file_pb_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_transfer_proto_goTypes,
		DependencyIndexes: file_pb_transfer_proto_depIdxs,
		MessageInfos:      file_pb_transfer_proto_msgTypes,
	}.Build()
	File_pb_transfer_proto = out.File
	file_pb_transfer_proto_rawDesc = nil
	file_pb_transfer_proto_goTypes = nil
	file_pb_transfer_proto_depIdxs = nil
}
