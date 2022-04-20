// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb/rpcbff.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommandResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandCode int64            `protobuf:"varint,1,opt,name=command_code,json=commandCode,proto3" json:"command_code,omitempty"`
	CommandMsg  string           `protobuf:"bytes,2,opt,name=command_msg,json=commandMsg,proto3" json:"command_msg,omitempty"`
	CommandData *structpb.Struct `protobuf:"bytes,3,opt,name=command_data,json=commandData,proto3" json:"command_data,omitempty"`
}

func (x *CommandResp) Reset() {
	*x = CommandResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpcbff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResp) ProtoMessage() {}

func (x *CommandResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpcbff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResp.ProtoReflect.Descriptor instead.
func (*CommandResp) Descriptor() ([]byte, []int) {
	return file_pb_rpcbff_proto_rawDescGZIP(), []int{0}
}

func (x *CommandResp) GetCommandCode() int64 {
	if x != nil {
		return x.CommandCode
	}
	return 0
}

func (x *CommandResp) GetCommandMsg() string {
	if x != nil {
		return x.CommandMsg
	}
	return ""
}

func (x *CommandResp) GetCommandData() *structpb.Struct {
	if x != nil {
		return x.CommandData
	}
	return nil
}

type CommandReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CommandReq) Reset() {
	*x = CommandReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpcbff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandReq) ProtoMessage() {}

func (x *CommandReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpcbff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandReq.ProtoReflect.Descriptor instead.
func (*CommandReq) Descriptor() ([]byte, []int) {
	return file_pb_rpcbff_proto_rawDescGZIP(), []int{1}
}

func (x *CommandReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_pb_rpcbff_proto protoreflect.FileDescriptor

var file_pb_rpcbff_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x65, 0x0a, 0x06, 0x52, 0x70, 0x63, 0x62, 0x66,
	0x66, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x30, 0x01, 0x12,
	0x2d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_rpcbff_proto_rawDescOnce sync.Once
	file_pb_rpcbff_proto_rawDescData = file_pb_rpcbff_proto_rawDesc
)

func file_pb_rpcbff_proto_rawDescGZIP() []byte {
	file_pb_rpcbff_proto_rawDescOnce.Do(func() {
		file_pb_rpcbff_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_rpcbff_proto_rawDescData)
	})
	return file_pb_rpcbff_proto_rawDescData
}

var file_pb_rpcbff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_rpcbff_proto_goTypes = []interface{}{
	(*CommandResp)(nil),     // 0: pb.CommandResp
	(*CommandReq)(nil),      // 1: pb.CommandReq
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
}
var file_pb_rpcbff_proto_depIdxs = []int32{
	2, // 0: pb.CommandResp.command_data:type_name -> google.protobuf.Struct
	1, // 1: pb.Rpcbff.connect:input_type -> pb.CommandReq
	1, // 2: pb.Rpcbff.disconnect:input_type -> pb.CommandReq
	0, // 3: pb.Rpcbff.connect:output_type -> pb.CommandResp
	0, // 4: pb.Rpcbff.disconnect:output_type -> pb.CommandResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_rpcbff_proto_init() }
func file_pb_rpcbff_proto_init() {
	if File_pb_rpcbff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_rpcbff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResp); i {
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
		file_pb_rpcbff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandReq); i {
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
			RawDescriptor: file_pb_rpcbff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_rpcbff_proto_goTypes,
		DependencyIndexes: file_pb_rpcbff_proto_depIdxs,
		MessageInfos:      file_pb_rpcbff_proto_msgTypes,
	}.Build()
	File_pb_rpcbff_proto = out.File
	file_pb_rpcbff_proto_rawDesc = nil
	file_pb_rpcbff_proto_goTypes = nil
	file_pb_rpcbff_proto_depIdxs = nil
}
