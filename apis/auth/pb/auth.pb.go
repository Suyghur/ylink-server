// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb/auth.proto

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

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *structpb.Struct `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type PlayerAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *PlayerAuthReq) Reset() {
	*x = PlayerAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerAuthReq) ProtoMessage() {}

func (x *PlayerAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerAuthReq.ProtoReflect.Descriptor instead.
func (*PlayerAuthReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerAuthReq) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerAuthReq) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type CsAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uname    string `protobuf:"bytes,1,opt,name=uname,proto3" json:"uname,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CsAuthReq) Reset() {
	*x = CsAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsAuthReq) ProtoMessage() {}

func (x *CsAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsAuthReq.ProtoReflect.Descriptor instead.
func (*CsAuthReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CsAuthReq) GetUname() string {
	if x != nil {
		return x.Uname
	}
	return ""
}

func (x *CsAuthReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CheckAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckAuthReq) Reset() {
	*x = CheckAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAuthReq) ProtoMessage() {}

func (x *CheckAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAuthReq.ProtoReflect.Descriptor instead.
func (*CheckAuthReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CheckAuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *structpb.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthResp) Reset() {
	*x = AuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResp) ProtoMessage() {}

func (x *AuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResp.ProtoReflect.Descriptor instead.
func (*AuthResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{4}
}

func (x *AuthResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AuthResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *AuthResp) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pb_auth_proto protoreflect.FileDescriptor

var file_pb_auth_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x36, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x0d, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x22, 0x3d, 0x0a, 0x09, 0x43, 0x73, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x24, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x89, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x2d, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x06,
	0x63, 0x73, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x73, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_auth_proto_rawDescOnce sync.Once
	file_pb_auth_proto_rawDescData = file_pb_auth_proto_rawDesc
)

func file_pb_auth_proto_rawDescGZIP() []byte {
	file_pb_auth_proto_rawDescOnce.Do(func() {
		file_pb_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_auth_proto_rawDescData)
	})
	return file_pb_auth_proto_rawDescData
}

var file_pb_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_auth_proto_goTypes = []interface{}{
	(*AuthReq)(nil),         // 0: pb.AuthReq
	(*PlayerAuthReq)(nil),   // 1: pb.PlayerAuthReq
	(*CsAuthReq)(nil),       // 2: pb.CsAuthReq
	(*CheckAuthReq)(nil),    // 3: pb.CheckAuthReq
	(*AuthResp)(nil),        // 4: pb.AuthResp
	(*structpb.Struct)(nil), // 5: google.protobuf.Struct
}
var file_pb_auth_proto_depIdxs = []int32{
	5, // 0: pb.AuthReq.data:type_name -> google.protobuf.Struct
	5, // 1: pb.AuthResp.data:type_name -> google.protobuf.Struct
	1, // 2: pb.Auth.playerAuth:input_type -> pb.PlayerAuthReq
	2, // 3: pb.Auth.csAuth:input_type -> pb.CsAuthReq
	3, // 4: pb.Auth.checkAuth:input_type -> pb.CheckAuthReq
	4, // 5: pb.Auth.playerAuth:output_type -> pb.AuthResp
	4, // 6: pb.Auth.csAuth:output_type -> pb.AuthResp
	4, // 7: pb.Auth.checkAuth:output_type -> pb.AuthResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_auth_proto_init() }
func file_pb_auth_proto_init() {
	if File_pb_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_pb_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerAuthReq); i {
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
		file_pb_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsAuthReq); i {
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
		file_pb_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAuthReq); i {
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
		file_pb_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResp); i {
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
			RawDescriptor: file_pb_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_auth_proto_goTypes,
		DependencyIndexes: file_pb_auth_proto_depIdxs,
		MessageInfos:      file_pb_auth_proto_msgTypes,
	}.Build()
	File_pb_auth_proto = out.File
	file_pb_auth_proto_rawDesc = nil
	file_pb_auth_proto_goTypes = nil
	file_pb_auth_proto_depIdxs = nil
}