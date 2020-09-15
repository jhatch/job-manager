// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: job/v1/queue.proto

package v1

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

	Id              string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	V               int32                `protobuf:"varint,2,opt,name=v,proto3" json:"v,omitempty"`
	Concurrency     int32                `protobuf:"varint,3,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Retries         int32                `protobuf:"varint,4,opt,name=retries,proto3" json:"retries,omitempty"`
	Duration        *duration.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	ClaimDuration   *duration.Duration   `protobuf:"bytes,6,opt,name=claim_duration,json=claimDuration,proto3" json:"claim_duration,omitempty"`
	CheckinDuration *duration.Duration   `protobuf:"bytes,7,opt,name=checkin_duration,json=checkinDuration,proto3" json:"checkin_duration,omitempty"`
	Labels          map[string]string    `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Schema          []byte               `protobuf:"bytes,9,opt,name=schema,proto3" json:"schema,omitempty"`
	Unique          bool                 `protobuf:"varint,10,opt,name=unique,proto3" json:"unique,omitempty"`
	CreatedAt       *timestamp.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamp.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       *timestamp.Timestamp `protobuf:"bytes,13,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
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

func (x *Queue) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Queue) GetV() int32 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *Queue) GetConcurrency() int32 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

func (x *Queue) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *Queue) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Queue) GetClaimDuration() *duration.Duration {
	if x != nil {
		return x.ClaimDuration
	}
	return nil
}

func (x *Queue) GetCheckinDuration() *duration.Duration {
	if x != nil {
		return x.CheckinDuration
	}
	return nil
}

func (x *Queue) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Queue) GetSchema() []byte {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *Queue) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
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

type QueueListParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names     []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	Selectors []string `protobuf:"bytes,2,rep,name=selectors,proto3" json:"selectors,omitempty"`
}

func (x *QueueListParams) Reset() {
	*x = QueueListParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_v1_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueListParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueListParams) ProtoMessage() {}

func (x *QueueListParams) ProtoReflect() protoreflect.Message {
	mi := &file_job_v1_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueListParams.ProtoReflect.Descriptor instead.
func (*QueueListParams) Descriptor() ([]byte, []int) {
	return file_job_v1_queue_proto_rawDescGZIP(), []int{2}
}

func (x *QueueListParams) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *QueueListParams) GetSelectors() []string {
	if x != nil {
		return x.Selectors
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
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6a,
	0x6f, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xef, 0x04, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x76, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x10,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x2f, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x06, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x73, 0x22, 0x45, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x66, 0x66, 0x72, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6f, 0x62, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_job_v1_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_job_v1_queue_proto_goTypes = []interface{}{
	(*Queue)(nil),               // 0: job.v1.Queue
	(*Queues)(nil),              // 1: job.v1.Queues
	(*QueueListParams)(nil),     // 2: job.v1.QueueListParams
	nil,                         // 3: job.v1.Queue.LabelsEntry
	(*duration.Duration)(nil),   // 4: google.protobuf.Duration
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_job_v1_queue_proto_depIdxs = []int32{
	4, // 0: job.v1.Queue.duration:type_name -> google.protobuf.Duration
	4, // 1: job.v1.Queue.claim_duration:type_name -> google.protobuf.Duration
	4, // 2: job.v1.Queue.checkin_duration:type_name -> google.protobuf.Duration
	3, // 3: job.v1.Queue.labels:type_name -> job.v1.Queue.LabelsEntry
	5, // 4: job.v1.Queue.created_at:type_name -> google.protobuf.Timestamp
	5, // 5: job.v1.Queue.updated_at:type_name -> google.protobuf.Timestamp
	5, // 6: job.v1.Queue.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 7: job.v1.Queues.queues:type_name -> job.v1.Queue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_job_v1_queue_proto_init() }
func file_job_v1_queue_proto_init() {
	if File_job_v1_queue_proto != nil {
		return
	}
	file_job_v1_job_proto_init()
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
		file_job_v1_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueListParams); i {
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
			RawDescriptor: file_job_v1_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
