// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: hablof/logistic_package_api/v1/logistic_package_api.proto

package logistic_package_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Foo     uint64                 `protobuf:"varint,2,opt,name=foo,proto3" json:"foo,omitempty"`
	Created *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Package) GetFoo() uint64 {
	if x != nil {
		return x.Foo
	}
	return 0
}

func (x *Package) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type DescribePackageV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId uint64 `protobuf:"varint,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
}

func (x *DescribePackageV1Request) Reset() {
	*x = DescribePackageV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePackageV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePackageV1Request) ProtoMessage() {}

func (x *DescribePackageV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePackageV1Request.ProtoReflect.Descriptor instead.
func (*DescribePackageV1Request) Descriptor() ([]byte, []int) {
	return file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescGZIP(), []int{1}
}

func (x *DescribePackageV1Request) GetPackageId() uint64 {
	if x != nil {
		return x.PackageId
	}
	return 0
}

type DescribePackageV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Package `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DescribePackageV1Response) Reset() {
	*x = DescribePackageV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePackageV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePackageV1Response) ProtoMessage() {}

func (x *DescribePackageV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePackageV1Response.ProtoReflect.Descriptor instead.
func (*DescribePackageV1Response) Descriptor() ([]byte, []int) {
	return file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribePackageV1Response) GetValue() *Package {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_hablof_logistic_package_api_v1_logistic_package_api_proto protoreflect.FileDescriptor

var file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDesc = []byte{
	0x0a, 0x39, 0x68, 0x61, 0x62, 0x6c, 0x6f, 0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x6d, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x6f, 0x6f,
	0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x42, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x19, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xc9, 0x01, 0x0a, 0x19, 0x4c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xab, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x31, 0x12, 0x38, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x6d, 0x70, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x6d, 0x70, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x62, 0x6c, 0x6f, 0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescOnce sync.Once
	file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescData = file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDesc
)

func file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescGZIP() []byte {
	file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescOnce.Do(func() {
		file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescData)
	})
	return file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDescData
}

var file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hablof_logistic_package_api_v1_logistic_package_api_proto_goTypes = []interface{}{
	(*Package)(nil),                   // 0: ozonmp.logistic_package_api.v1.Package
	(*DescribePackageV1Request)(nil),  // 1: ozonmp.logistic_package_api.v1.DescribePackageV1Request
	(*DescribePackageV1Response)(nil), // 2: ozonmp.logistic_package_api.v1.DescribePackageV1Response
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
}
var file_hablof_logistic_package_api_v1_logistic_package_api_proto_depIdxs = []int32{
	3, // 0: ozonmp.logistic_package_api.v1.Package.created:type_name -> google.protobuf.Timestamp
	0, // 1: ozonmp.logistic_package_api.v1.DescribePackageV1Response.value:type_name -> ozonmp.logistic_package_api.v1.Package
	1, // 2: ozonmp.logistic_package_api.v1.LogisticPackageApiService.DescribePackageV1:input_type -> ozonmp.logistic_package_api.v1.DescribePackageV1Request
	2, // 3: ozonmp.logistic_package_api.v1.LogisticPackageApiService.DescribePackageV1:output_type -> ozonmp.logistic_package_api.v1.DescribePackageV1Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hablof_logistic_package_api_v1_logistic_package_api_proto_init() }
func file_hablof_logistic_package_api_v1_logistic_package_api_proto_init() {
	if File_hablof_logistic_package_api_v1_logistic_package_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePackageV1Request); i {
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
		file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePackageV1Response); i {
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
			RawDescriptor: file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hablof_logistic_package_api_v1_logistic_package_api_proto_goTypes,
		DependencyIndexes: file_hablof_logistic_package_api_v1_logistic_package_api_proto_depIdxs,
		MessageInfos:      file_hablof_logistic_package_api_v1_logistic_package_api_proto_msgTypes,
	}.Build()
	File_hablof_logistic_package_api_v1_logistic_package_api_proto = out.File
	file_hablof_logistic_package_api_v1_logistic_package_api_proto_rawDesc = nil
	file_hablof_logistic_package_api_v1_logistic_package_api_proto_goTypes = nil
	file_hablof_logistic_package_api_v1_logistic_package_api_proto_depIdxs = nil
}
