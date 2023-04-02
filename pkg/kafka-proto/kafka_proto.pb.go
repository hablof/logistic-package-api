// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: hablof/logistic_package_api/v1/kafka_proto.proto

package kafka_proto

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

type EventType int32

const (
	EventType_Created EventType = 0
	EventType_Updated EventType = 1
	EventType_Removed EventType = 2
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "Created",
		1: "Updated",
		2: "Removed",
	}
	EventType_value = map[string]int32{
		"Created": 0,
		"Updated": 1,
		"Removed": 2,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_hablof_logistic_package_api_v1_kafka_proto_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_hablof_logistic_package_api_v1_kafka_proto_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescGZIP(), []int{0}
}

type PackageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PackageID uint64                 `protobuf:"varint,2,opt,name=PackageID,proto3" json:"PackageID,omitempty"`
	Type      EventType              `protobuf:"varint,3,opt,name=Type,proto3,enum=hablof.proto.v1.EventType" json:"Type,omitempty"`
	Created   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=Created,proto3" json:"Created,omitempty"`
	Payload   []byte                 `protobuf:"bytes,5,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *PackageEvent) Reset() {
	*x = PackageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hablof_logistic_package_api_v1_kafka_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageEvent) ProtoMessage() {}

func (x *PackageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_hablof_logistic_package_api_v1_kafka_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageEvent.ProtoReflect.Descriptor instead.
func (*PackageEvent) Descriptor() ([]byte, []int) {
	return file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescGZIP(), []int{0}
}

func (x *PackageEvent) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PackageEvent) GetPackageID() uint64 {
	if x != nil {
		return x.PackageID
	}
	return 0
}

func (x *PackageEvent) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_Created
}

func (x *PackageEvent) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *PackageEvent) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_hablof_logistic_package_api_v1_kafka_proto_proto protoreflect.FileDescriptor

var file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDesc = []byte{
	0x0a, 0x30, 0x68, 0x61, 0x62, 0x6c, 0x6f, 0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x68, 0x61, 0x62, 0x6c, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x68, 0x61, 0x62, 0x6c, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2a, 0x32, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x02, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x62, 0x6c, 0x6f, 0x66, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescOnce sync.Once
	file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescData = file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDesc
)

func file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescGZIP() []byte {
	file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescOnce.Do(func() {
		file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescData)
	})
	return file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDescData
}

var file_hablof_logistic_package_api_v1_kafka_proto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hablof_logistic_package_api_v1_kafka_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hablof_logistic_package_api_v1_kafka_proto_proto_goTypes = []interface{}{
	(EventType)(0),                // 0: hablof.proto.v1.EventType
	(*PackageEvent)(nil),          // 1: hablof.proto.v1.PackageEvent
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_hablof_logistic_package_api_v1_kafka_proto_proto_depIdxs = []int32{
	0, // 0: hablof.proto.v1.PackageEvent.Type:type_name -> hablof.proto.v1.EventType
	2, // 1: hablof.proto.v1.PackageEvent.Created:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hablof_logistic_package_api_v1_kafka_proto_proto_init() }
func file_hablof_logistic_package_api_v1_kafka_proto_proto_init() {
	if File_hablof_logistic_package_api_v1_kafka_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hablof_logistic_package_api_v1_kafka_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageEvent); i {
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
			RawDescriptor: file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hablof_logistic_package_api_v1_kafka_proto_proto_goTypes,
		DependencyIndexes: file_hablof_logistic_package_api_v1_kafka_proto_proto_depIdxs,
		EnumInfos:         file_hablof_logistic_package_api_v1_kafka_proto_proto_enumTypes,
		MessageInfos:      file_hablof_logistic_package_api_v1_kafka_proto_proto_msgTypes,
	}.Build()
	File_hablof_logistic_package_api_v1_kafka_proto_proto = out.File
	file_hablof_logistic_package_api_v1_kafka_proto_proto_rawDesc = nil
	file_hablof_logistic_package_api_v1_kafka_proto_proto_goTypes = nil
	file_hablof_logistic_package_api_v1_kafka_proto_proto_depIdxs = nil
}
