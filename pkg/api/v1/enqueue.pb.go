// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: api/v1/enqueue.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	v1 "github.com/jeffrom/job-manager/pkg/resource/job/v1"
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

type EnqueueJobsRequestArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job  string           `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Args []*_struct.Value `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Data *v1.Data         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EnqueueJobsRequestArgs) Reset() {
	*x = EnqueueJobsRequestArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_enqueue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueJobsRequestArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueJobsRequestArgs) ProtoMessage() {}

func (x *EnqueueJobsRequestArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_enqueue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueJobsRequestArgs.ProtoReflect.Descriptor instead.
func (*EnqueueJobsRequestArgs) Descriptor() ([]byte, []int) {
	return file_api_v1_enqueue_proto_rawDescGZIP(), []int{0}
}

func (x *EnqueueJobsRequestArgs) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *EnqueueJobsRequestArgs) GetArgs() []*_struct.Value {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *EnqueueJobsRequestArgs) GetData() *v1.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type EnqueueJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*EnqueueJobsRequestArgs `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *EnqueueJobsRequest) Reset() {
	*x = EnqueueJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_enqueue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueJobsRequest) ProtoMessage() {}

func (x *EnqueueJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_enqueue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueJobsRequest.ProtoReflect.Descriptor instead.
func (*EnqueueJobsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_enqueue_proto_rawDescGZIP(), []int{1}
}

func (x *EnqueueJobsRequest) GetJobs() []*EnqueueJobsRequestArgs {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type EnqueueJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs   []string         `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Errors []*ErrorResponse `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *EnqueueJobsResponse) Reset() {
	*x = EnqueueJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_enqueue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueJobsResponse) ProtoMessage() {}

func (x *EnqueueJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_enqueue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueJobsResponse.ProtoReflect.Descriptor instead.
func (*EnqueueJobsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_enqueue_proto_rawDescGZIP(), []int{2}
}

func (x *EnqueueJobsResponse) GetJobs() []string {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *EnqueueJobsResponse) GetErrors() []*ErrorResponse {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_api_v1_enqueue_proto protoreflect.FileDescriptor

var file_api_v1_enqueue_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6a, 0x6f,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x16, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4a,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62,
	0x12, 0x2a, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x48,
	0x0a, 0x12, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x72,
	0x67, 0x73, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x58, 0x0a, 0x13, 0x45, 0x6e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6a,
	0x6f, 0x62, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x65, 0x66, 0x66, 0x72, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_enqueue_proto_rawDescOnce sync.Once
	file_api_v1_enqueue_proto_rawDescData = file_api_v1_enqueue_proto_rawDesc
)

func file_api_v1_enqueue_proto_rawDescGZIP() []byte {
	file_api_v1_enqueue_proto_rawDescOnce.Do(func() {
		file_api_v1_enqueue_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_enqueue_proto_rawDescData)
	})
	return file_api_v1_enqueue_proto_rawDescData
}

var file_api_v1_enqueue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_enqueue_proto_goTypes = []interface{}{
	(*EnqueueJobsRequestArgs)(nil), // 0: api.v1.EnqueueJobsRequestArgs
	(*EnqueueJobsRequest)(nil),     // 1: api.v1.EnqueueJobsRequest
	(*EnqueueJobsResponse)(nil),    // 2: api.v1.EnqueueJobsResponse
	(*_struct.Value)(nil),          // 3: google.protobuf.Value
	(*v1.Data)(nil),                // 4: job.v1.Data
	(*ErrorResponse)(nil),          // 5: api.v1.ErrorResponse
}
var file_api_v1_enqueue_proto_depIdxs = []int32{
	3, // 0: api.v1.EnqueueJobsRequestArgs.args:type_name -> google.protobuf.Value
	4, // 1: api.v1.EnqueueJobsRequestArgs.data:type_name -> job.v1.Data
	0, // 2: api.v1.EnqueueJobsRequest.jobs:type_name -> api.v1.EnqueueJobsRequestArgs
	5, // 3: api.v1.EnqueueJobsResponse.errors:type_name -> api.v1.ErrorResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_enqueue_proto_init() }
func file_api_v1_enqueue_proto_init() {
	if File_api_v1_enqueue_proto != nil {
		return
	}
	file_api_v1_errors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_enqueue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueJobsRequestArgs); i {
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
		file_api_v1_enqueue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueJobsRequest); i {
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
		file_api_v1_enqueue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueJobsResponse); i {
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
			RawDescriptor: file_api_v1_enqueue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_enqueue_proto_goTypes,
		DependencyIndexes: file_api_v1_enqueue_proto_depIdxs,
		MessageInfos:      file_api_v1_enqueue_proto_msgTypes,
	}.Build()
	File_api_v1_enqueue_proto = out.File
	file_api_v1_enqueue_proto_rawDesc = nil
	file_api_v1_enqueue_proto_goTypes = nil
	file_api_v1_enqueue_proto_depIdxs = nil
}
