// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: job/v1/status.proto

package v1

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

type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_STATUS_QUEUED      Status = 1
	Status_STATUS_RUNNING     Status = 2
	Status_STATUS_COMPLETE    Status = 3
	Status_STATUS_DEAD        Status = 4
	Status_STATUS_CANCELLED   Status = 5
	Status_STATUS_INVALID     Status = 6
	Status_STATUS_FAILED      Status = 7
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_QUEUED",
		2: "STATUS_RUNNING",
		3: "STATUS_COMPLETE",
		4: "STATUS_DEAD",
		5: "STATUS_CANCELLED",
		6: "STATUS_INVALID",
		7: "STATUS_FAILED",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_QUEUED":      1,
		"STATUS_RUNNING":     2,
		"STATUS_COMPLETE":    3,
		"STATUS_DEAD":        4,
		"STATUS_CANCELLED":   5,
		"STATUS_INVALID":     6,
		"STATUS_FAILED":      7,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_job_v1_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_job_v1_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_job_v1_status_proto_rawDescGZIP(), []int{0}
}

var File_job_v1_status_proto protoreflect.FileDescriptor

var file_job_v1_status_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2a, 0xaa, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x10, 0x04, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x66, 0x66, 0x72, 0x6f, 0x6d,
	0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_v1_status_proto_rawDescOnce sync.Once
	file_job_v1_status_proto_rawDescData = file_job_v1_status_proto_rawDesc
)

func file_job_v1_status_proto_rawDescGZIP() []byte {
	file_job_v1_status_proto_rawDescOnce.Do(func() {
		file_job_v1_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_v1_status_proto_rawDescData)
	})
	return file_job_v1_status_proto_rawDescData
}

var file_job_v1_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_job_v1_status_proto_goTypes = []interface{}{
	(Status)(0), // 0: job.v1.Status
}
var file_job_v1_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_job_v1_status_proto_init() }
func file_job_v1_status_proto_init() {
	if File_job_v1_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_v1_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_job_v1_status_proto_goTypes,
		DependencyIndexes: file_job_v1_status_proto_depIdxs,
		EnumInfos:         file_job_v1_status_proto_enumTypes,
	}.Build()
	File_job_v1_status_proto = out.File
	file_job_v1_status_proto_rawDesc = nil
	file_job_v1_status_proto_goTypes = nil
	file_job_v1_status_proto_depIdxs = nil
}
