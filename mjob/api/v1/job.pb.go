// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: api/v1/job.proto

package v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	v1 "github.com/jeffrom/job-manager/mjob/resource/job/v1"
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

type ListJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue        []string             `protobuf:"bytes,1,rep,name=queue,proto3" json:"queue,omitempty"`
	Status       []string             `protobuf:"bytes,2,rep,name=status,proto3" json:"status,omitempty"`
	Selector     []string             `protobuf:"bytes,3,rep,name=selector,proto3" json:"selector,omitempty"`
	Claim        []string             `protobuf:"bytes,4,rep,name=claim,proto3" json:"claim,omitempty"`
	NoUnclaimed  bool                 `protobuf:"varint,5,opt,name=no_unclaimed,json=noUnclaimed,proto3" json:"no_unclaimed,omitempty"`
	CreatedSince *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_since,json=createdSince,proto3" json:"created_since,omitempty"`
	CreatedUntil *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_until,json=createdUntil,proto3" json:"created_until,omitempty"`
	Page         *Pagination          `protobuf:"bytes,8,opt,name=page,proto3" json:"page,omitempty"`
	Include      []string             `protobuf:"bytes,9,rep,name=include,proto3" json:"include,omitempty"`
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_job_proto_rawDescGZIP(), []int{0}
}

func (x *ListJobsRequest) GetQueue() []string {
	if x != nil {
		return x.Queue
	}
	return nil
}

func (x *ListJobsRequest) GetStatus() []string {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ListJobsRequest) GetSelector() []string {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *ListJobsRequest) GetClaim() []string {
	if x != nil {
		return x.Claim
	}
	return nil
}

func (x *ListJobsRequest) GetNoUnclaimed() bool {
	if x != nil {
		return x.NoUnclaimed
	}
	return false
}

func (x *ListJobsRequest) GetCreatedSince() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedSince
	}
	return nil
}

func (x *ListJobsRequest) GetCreatedUntil() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedUntil
	}
	return nil
}

func (x *ListJobsRequest) GetPage() *Pagination {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListJobsRequest) GetInclude() []string {
	if x != nil {
		return x.Include
	}
	return nil
}

type ListJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items  []*v1.Job        `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Errors []*ErrorResponse `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_job_proto_rawDescGZIP(), []int{1}
}

func (x *ListJobsResponse) GetItems() []*v1.Job {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListJobsResponse) GetErrors() []*ErrorResponse {
	if x != nil {
		return x.Errors
	}
	return nil
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Include []string `protobuf:"bytes,2,rep,name=include,proto3" json:"include,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_job_proto_rawDescGZIP(), []int{2}
}

func (x *GetJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetJobRequest) GetInclude() []string {
	if x != nil {
		return x.Include
	}
	return nil
}

var File_api_v1_job_proto protoreflect.FileDescriptor

var file_api_v1_job_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6a, 0x6f, 0x62,
	0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x12,
	0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x75, 0x6e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e, 0x6f, 0x55, 0x6e, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75,
	0x6e, 0x74, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55,
	0x6e, 0x74, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x22, 0x64, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f,
	0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2d, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x39, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x66, 0x66, 0x72, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f,
	0x62, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x6a, 0x6f, 0x62, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_job_proto_rawDescOnce sync.Once
	file_api_v1_job_proto_rawDescData = file_api_v1_job_proto_rawDesc
)

func file_api_v1_job_proto_rawDescGZIP() []byte {
	file_api_v1_job_proto_rawDescOnce.Do(func() {
		file_api_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_job_proto_rawDescData)
	})
	return file_api_v1_job_proto_rawDescData
}

var file_api_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_job_proto_goTypes = []interface{}{
	(*ListJobsRequest)(nil),     // 0: api.v1.ListJobsRequest
	(*ListJobsResponse)(nil),    // 1: api.v1.ListJobsResponse
	(*GetJobRequest)(nil),       // 2: api.v1.GetJobRequest
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Pagination)(nil),          // 4: api.v1.Pagination
	(*v1.Job)(nil),              // 5: job.v1.Job
	(*ErrorResponse)(nil),       // 6: api.v1.ErrorResponse
}
var file_api_v1_job_proto_depIdxs = []int32{
	3, // 0: api.v1.ListJobsRequest.created_since:type_name -> google.protobuf.Timestamp
	3, // 1: api.v1.ListJobsRequest.created_until:type_name -> google.protobuf.Timestamp
	4, // 2: api.v1.ListJobsRequest.page:type_name -> api.v1.Pagination
	5, // 3: api.v1.ListJobsResponse.items:type_name -> job.v1.Job
	6, // 4: api.v1.ListJobsResponse.errors:type_name -> api.v1.ErrorResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_job_proto_init() }
func file_api_v1_job_proto_init() {
	if File_api_v1_job_proto != nil {
		return
	}
	file_api_v1_errors_proto_init()
	file_api_v1_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsRequest); i {
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
		file_api_v1_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsResponse); i {
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
		file_api_v1_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
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
			RawDescriptor: file_api_v1_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_job_proto_goTypes,
		DependencyIndexes: file_api_v1_job_proto_depIdxs,
		MessageInfos:      file_api_v1_job_proto_msgTypes,
	}.Build()
	File_api_v1_job_proto = out.File
	file_api_v1_job_proto_rawDesc = nil
	file_api_v1_job_proto_goTypes = nil
	file_api_v1_job_proto_depIdxs = nil
}
