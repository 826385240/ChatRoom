// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.15.8
// source: login.proto

package login

import (
	proto "github.com/golang/protobuf/proto"
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

type MSG_StartLogin_CS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MSG_StartLogin_CS) Reset() {
	*x = MSG_StartLogin_CS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSG_StartLogin_CS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSG_StartLogin_CS) ProtoMessage() {}

func (x *MSG_StartLogin_CS) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSG_StartLogin_CS.ProtoReflect.Descriptor instead.
func (*MSG_StartLogin_CS) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *MSG_StartLogin_CS) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MSG_StartLogin_SC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MSG_StartLogin_SC) Reset() {
	*x = MSG_StartLogin_SC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSG_StartLogin_SC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSG_StartLogin_SC) ProtoMessage() {}

func (x *MSG_StartLogin_SC) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSG_StartLogin_SC.ProtoReflect.Descriptor instead.
func (*MSG_StartLogin_SC) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

type MSG_CreateRole_CS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MSG_CreateRole_CS) Reset() {
	*x = MSG_CreateRole_CS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSG_CreateRole_CS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSG_CreateRole_CS) ProtoMessage() {}

func (x *MSG_CreateRole_CS) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSG_CreateRole_CS.ProtoReflect.Descriptor instead.
func (*MSG_CreateRole_CS) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *MSG_CreateRole_CS) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MSG_CreateRole_SC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode bool `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *MSG_CreateRole_SC) Reset() {
	*x = MSG_CreateRole_SC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MSG_CreateRole_SC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MSG_CreateRole_SC) ProtoMessage() {}

func (x *MSG_CreateRole_SC) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MSG_CreateRole_SC.ProtoReflect.Descriptor instead.
func (*MSG_CreateRole_SC) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

func (x *MSG_CreateRole_SC) GetRetcode() bool {
	if x != nil {
		return x.Retcode
	}
	return false
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x22, 0x27, 0x0a, 0x11, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x43, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x53, 0x43, 0x22, 0x27, 0x0a, 0x11, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x5f, 0x43, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x4d,
	0x53, 0x47, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x5f, 0x53, 0x43,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_login_proto_goTypes = []interface{}{
	(*MSG_StartLogin_CS)(nil), // 0: login.MSG_StartLogin_CS
	(*MSG_StartLogin_SC)(nil), // 1: login.MSG_StartLogin_SC
	(*MSG_CreateRole_CS)(nil), // 2: login.MSG_CreateRole_CS
	(*MSG_CreateRole_SC)(nil), // 3: login.MSG_CreateRole_SC
}
var file_login_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSG_StartLogin_CS); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSG_StartLogin_SC); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSG_CreateRole_CS); i {
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
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MSG_CreateRole_SC); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
