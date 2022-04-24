// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: langya_platform_app.proto

package app

import (
	common "gitee.com/langya_platform/langya/platform/common"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender common.GenderType `protobuf:"varint,3,opt,name=gender,proto3,enum=common.GenderType" json:"gender,omitempty"`
	Number string            `protobuf:"bytes,4,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_langya_platform_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_langya_platform_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_langya_platform_app_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetGender() common.GenderType {
	if x != nil {
		return x.Gender
	}
	return common.GenderType(0)
}

func (x *Person) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type ContractBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
}

func (x *ContractBook) Reset() {
	*x = ContractBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_langya_platform_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractBook) ProtoMessage() {}

func (x *ContractBook) ProtoReflect() protoreflect.Message {
	mi := &file_langya_platform_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractBook.ProtoReflect.Descriptor instead.
func (*ContractBook) Descriptor() ([]byte, []int) {
	return file_langya_platform_app_proto_rawDescGZIP(), []int{1}
}

func (x *ContractBook) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

type ContractBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContractBookReq) Reset() {
	*x = ContractBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_langya_platform_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractBookReq) ProtoMessage() {}

func (x *ContractBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_langya_platform_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractBookReq.ProtoReflect.Descriptor instead.
func (*ContractBookReq) Descriptor() ([]byte, []int) {
	return file_langya_platform_app_proto_rawDescGZIP(), []int{2}
}

var File_langya_platform_app_proto protoreflect.FileDescriptor

var file_langya_platform_app_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x61, 0x6e, 0x67, 0x79, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x70,
	0x1a, 0x1c, 0x6c, 0x61, 0x6e, 0x67, 0x79, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x06, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x25, 0x0a, 0x07,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x32, 0xc7, 0x01, 0x0a, 0x0e, 0x4c, 0x61, 0x6e, 0x67, 0x59,
	0x61, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x56, 0x0a, 0x13, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x12, 0x5d, 0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x61, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a,
	0x42, 0x8f, 0x02, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x79, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x79, 0x61, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61,
	0x70, 0x70, 0x92, 0x41, 0xdc, 0x01, 0x12, 0xa6, 0x01, 0x0a, 0x1c, 0x4c, 0x61, 0x6e, 0x67, 0x20,
	0x59, 0x61, 0x20, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x61, 0x70, 0x70, 0x20,
	0x41, 0x70, 0x69, 0x20, 0x44, 0x6f, 0x63, 0x12, 0x23, 0xe7, 0x90, 0x85, 0xe7, 0x90, 0x8a, 0xe5,
	0x86, 0x85, 0xe5, 0xae, 0xb9, 0xe5, 0xb9, 0xb3, 0xe5, 0x8f, 0xb0, 0x61, 0x70, 0x70, 0xe7, 0xab,
	0xaf, 0x20, 0x41, 0x70, 0x69, 0x20, 0xe6, 0x96, 0x87, 0xe6, 0xa1, 0xa3, 0x22, 0x5c, 0x0a, 0x0c,
	0x6d, 0x69, 0x6e, 0x67, 0x68, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x65, 0x6e, 0x12, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x69, 0x6e, 0x67, 0x31, 0x30, 0x32, 0x38, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x79, 0x61,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x1a, 0x1b, 0x6d,
	0x69, 0x6e, 0x67, 0x68, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x65, 0x6e, 0x40, 0x68, 0x75, 0x61, 0x73,
	0x68, 0x65, 0x6e, 0x67, 0x66, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x22,
	0x09, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x2a, 0x02, 0x02, 0x01, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_langya_platform_app_proto_rawDescOnce sync.Once
	file_langya_platform_app_proto_rawDescData = file_langya_platform_app_proto_rawDesc
)

func file_langya_platform_app_proto_rawDescGZIP() []byte {
	file_langya_platform_app_proto_rawDescOnce.Do(func() {
		file_langya_platform_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_langya_platform_app_proto_rawDescData)
	})
	return file_langya_platform_app_proto_rawDescData
}

var file_langya_platform_app_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_langya_platform_app_proto_goTypes = []interface{}{
	(*Person)(nil),          // 0: app.Person
	(*ContractBook)(nil),    // 1: app.ContractBook
	(*ContractBookReq)(nil), // 2: app.ContractBookReq
	(common.GenderType)(0),  // 3: common.GenderType
}
var file_langya_platform_app_proto_depIdxs = []int32{
	3, // 0: app.Person.gender:type_name -> common.GenderType
	0, // 1: app.ContractBook.persons:type_name -> app.Person
	2, // 2: app.LangYaPlatform.ServiceContractBook:input_type -> app.ContractBookReq
	2, // 3: app.LangYaPlatform.ServiceContractBookSave:input_type -> app.ContractBookReq
	1, // 4: app.LangYaPlatform.ServiceContractBook:output_type -> app.ContractBook
	1, // 5: app.LangYaPlatform.ServiceContractBookSave:output_type -> app.ContractBook
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_langya_platform_app_proto_init() }
func file_langya_platform_app_proto_init() {
	if File_langya_platform_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_langya_platform_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_langya_platform_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractBook); i {
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
		file_langya_platform_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractBookReq); i {
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
			RawDescriptor: file_langya_platform_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_langya_platform_app_proto_goTypes,
		DependencyIndexes: file_langya_platform_app_proto_depIdxs,
		MessageInfos:      file_langya_platform_app_proto_msgTypes,
	}.Build()
	File_langya_platform_app_proto = out.File
	file_langya_platform_app_proto_rawDesc = nil
	file_langya_platform_app_proto_goTypes = nil
	file_langya_platform_app_proto_depIdxs = nil
}
