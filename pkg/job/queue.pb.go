// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: job/v1/queue.proto

package job

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Concurrency int32                `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	MaxRetries  int32                `protobuf:"varint,3,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	Duration    *duration.Duration   `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Labels      map[string]string    `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_v1_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_job_v1_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_job_v1_queue_proto_rawDescGZIP(), []int{0}
}

func (x *Queue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Queue) GetConcurrency() int32 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

func (x *Queue) GetMaxRetries() int32 {
	if x != nil {
		return x.MaxRetries
	}
	return 0
}

func (x *Queue) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Queue) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Queue) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Queue) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Queue) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type Queues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queues []*Queue `protobuf:"bytes,1,rep,name=queues,proto3" json:"queues,omitempty"`
}

func (x *Queues) Reset() {
	*x = Queues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_v1_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queues) ProtoMessage() {}

func (x *Queues) ProtoReflect() protoreflect.Message {
	mi := &file_job_v1_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queues.ProtoReflect.Descriptor instead.
func (*Queues) Descriptor() ([]byte, []int) {
	return file_job_v1_queue_proto_rawDescGZIP(), []int{1}
}

func (x *Queues) GetQueues() []*Queue {
	if x != nil {
		return x.Queues
	}
	return nil
}

var File_job_v1_queue_proto protoreflect.FileDescriptor

var file_job_v1_queue_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x03,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x2f, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x52, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x66, 0x66, 0x72, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6f, 0x62, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_v1_queue_proto_rawDescOnce sync.Once
	file_job_v1_queue_proto_rawDescData = file_job_v1_queue_proto_rawDesc
)

func file_job_v1_queue_proto_rawDescGZIP() []byte {
	file_job_v1_queue_proto_rawDescOnce.Do(func() {
		file_job_v1_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_v1_queue_proto_rawDescData)
	})
	return file_job_v1_queue_proto_rawDescData
}

var file_job_v1_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_job_v1_queue_proto_goTypes = []interface{}{
	(*Queue)(nil),               // 0: job.v1.Queue
	(*Queues)(nil),              // 1: job.v1.Queues
	nil,                         // 2: job.v1.Queue.LabelsEntry
	(*duration.Duration)(nil),   // 3: google.protobuf.Duration
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_job_v1_queue_proto_depIdxs = []int32{
	3, // 0: job.v1.Queue.duration:type_name -> google.protobuf.Duration
	2, // 1: job.v1.Queue.labels:type_name -> job.v1.Queue.LabelsEntry
	4, // 2: job.v1.Queue.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: job.v1.Queue.updated_at:type_name -> google.protobuf.Timestamp
	4, // 4: job.v1.Queue.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 5: job.v1.Queues.queues:type_name -> job.v1.Queue
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_job_v1_queue_proto_init() }
func file_job_v1_queue_proto_init() {
	if File_job_v1_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_v1_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
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
		file_job_v1_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queues); i {
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
	file_job_v1_queue_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_v1_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_job_v1_queue_proto_goTypes,
		DependencyIndexes: file_job_v1_queue_proto_depIdxs,
		MessageInfos:      file_job_v1_queue_proto_msgTypes,
	}.Build()
	File_job_v1_queue_proto = out.File
	file_job_v1_queue_proto_rawDesc = nil
	file_job_v1_queue_proto_goTypes = nil
	file_job_v1_queue_proto_depIdxs = nil
}
