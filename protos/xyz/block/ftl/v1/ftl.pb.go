// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: xyz/block/ftl/v1/ftl.proto

package ftlv1

import (
	schema "github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1/schema"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{1}
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Metadata_Pair `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetValues() []*Metadata_Pair {
	if x != nil {
		return x.Values
	}
	return nil
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Verb     *schema.VerbRef `protobuf:"bytes,2,opt,name=verb,proto3" json:"verb,omitempty"`
	Body     []byte          `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{3}
}

func (x *CallRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CallRequest) GetVerb() *schema.VerbRef {
	if x != nil {
		return x.Verb
	}
	return nil
}

func (x *CallRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*CallResponse_Body
	//	*CallResponse_Error_
	Response isCallResponse_Response `protobuf_oneof:"response"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{4}
}

func (m *CallResponse) GetResponse() isCallResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *CallResponse) GetBody() []byte {
	if x, ok := x.GetResponse().(*CallResponse_Body); ok {
		return x.Body
	}
	return nil
}

func (x *CallResponse) GetError() *CallResponse_Error {
	if x, ok := x.GetResponse().(*CallResponse_Error_); ok {
		return x.Error
	}
	return nil
}

type isCallResponse_Response interface {
	isCallResponse_Response()
}

type CallResponse_Body struct {
	Body []byte `protobuf:"bytes,1,opt,name=body,proto3,oneof"`
}

type CallResponse_Error_ struct {
	Error *CallResponse_Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*CallResponse_Body) isCallResponse_Response() {}

func (*CallResponse_Error_) isCallResponse_Response() {}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{5}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verbs []*schema.VerbRef `protobuf:"bytes,1,rep,name=verbs,proto3" json:"verbs,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{6}
}

func (x *ListResponse) GetVerbs() []*schema.VerbRef {
	if x != nil {
		return x.Verbs
	}
	return nil
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Verb     *schema.VerbRef `protobuf:"bytes,2,opt,name=verb,proto3" json:"verb,omitempty"`
	Body     []byte          `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{7}
}

func (x *SendRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SendRequest) GetVerb() *schema.VerbRef {
	if x != nil {
		return x.Verb
	}
	return nil
}

func (x *SendRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{8}
}

type SyncSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema *schema.Module `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *SyncSchemaRequest) Reset() {
	*x = SyncSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncSchemaRequest) ProtoMessage() {}

func (x *SyncSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncSchemaRequest.ProtoReflect.Descriptor instead.
func (*SyncSchemaRequest) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{9}
}

func (x *SyncSchemaRequest) GetSchema() *schema.Module {
	if x != nil {
		return x.Schema
	}
	return nil
}

type SyncSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema *schema.Module `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	// If true, there are more schema changes immediately following this one.
	// If false, there still may be more schema changes in the future.
	More bool `protobuf:"varint,2,opt,name=more,proto3" json:"more,omitempty"`
}

func (x *SyncSchemaResponse) Reset() {
	*x = SyncSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncSchemaResponse) ProtoMessage() {}

func (x *SyncSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncSchemaResponse.ProtoReflect.Descriptor instead.
func (*SyncSchemaResponse) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{10}
}

func (x *SyncSchemaResponse) GetSchema() *schema.Module {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *SyncSchemaResponse) GetMore() bool {
	if x != nil {
		return x.More
	}
	return false
}

type Metadata_Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Metadata_Pair) Reset() {
	*x = Metadata_Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata_Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata_Pair) ProtoMessage() {}

func (x *Metadata_Pair) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata_Pair.ProtoReflect.Descriptor instead.
func (*Metadata_Pair) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Metadata_Pair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Metadata_Pair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CallResponse_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // TODO: Richer error type.
}

func (x *CallResponse_Error) Reset() {
	*x = CallResponse_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse_Error) ProtoMessage() {}

func (x *CallResponse_Error) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_ftl_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse_Error.ProtoReflect.Descriptor instead.
func (*CallResponse_Error) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CallResponse_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_xyz_block_ftl_v1_ftl_proto protoreflect.FileDescriptor

var file_xyz_block_ftl_v1_ftl_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x78, 0x79, 0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x66, 0x74, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x78, 0x79,
	0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x24,
	0x78, 0x79, 0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x66, 0x74, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x73, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x37, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x2e, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x79, 0x7a,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x34, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52, 0x65, 0x66,
	0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a,
	0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x05, 0x76, 0x65, 0x72, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78,
	0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52, 0x65, 0x66, 0x52, 0x05,
	0x76, 0x65, 0x72, 0x62, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a,
	0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x79,
	0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52, 0x65, 0x66, 0x52, 0x04, 0x76,
	0x65, 0x72, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x78,
	0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x61, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x79,
	0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x6d, 0x6f, 0x72, 0x65, 0x32, 0xa9, 0x02, 0x0a, 0x0b, 0x56, 0x65, 0x72,
	0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x1d, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1d,
	0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb2, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e,
	0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x78,
	0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0a,
	0x53, 0x79, 0x6e, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x23, 0x2e, 0x78, 0x79, 0x7a,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x42, 0x44, 0x35, 0x34, 0x35,
	0x36, 0x36, 0x39, 0x37, 0x35, 0x2f, 0x66, 0x74, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x78, 0x79, 0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x66, 0x74, 0x6c, 0x2f, 0x76,
	0x31, 0x3b, 0x66, 0x74, 0x6c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xyz_block_ftl_v1_ftl_proto_rawDescOnce sync.Once
	file_xyz_block_ftl_v1_ftl_proto_rawDescData = file_xyz_block_ftl_v1_ftl_proto_rawDesc
)

func file_xyz_block_ftl_v1_ftl_proto_rawDescGZIP() []byte {
	file_xyz_block_ftl_v1_ftl_proto_rawDescOnce.Do(func() {
		file_xyz_block_ftl_v1_ftl_proto_rawDescData = protoimpl.X.CompressGZIP(file_xyz_block_ftl_v1_ftl_proto_rawDescData)
	})
	return file_xyz_block_ftl_v1_ftl_proto_rawDescData
}

var file_xyz_block_ftl_v1_ftl_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_xyz_block_ftl_v1_ftl_proto_goTypes = []interface{}{
	(*PingRequest)(nil),        // 0: xyz.block.ftl.v1.PingRequest
	(*PingResponse)(nil),       // 1: xyz.block.ftl.v1.PingResponse
	(*Metadata)(nil),           // 2: xyz.block.ftl.v1.Metadata
	(*CallRequest)(nil),        // 3: xyz.block.ftl.v1.CallRequest
	(*CallResponse)(nil),       // 4: xyz.block.ftl.v1.CallResponse
	(*ListRequest)(nil),        // 5: xyz.block.ftl.v1.ListRequest
	(*ListResponse)(nil),       // 6: xyz.block.ftl.v1.ListResponse
	(*SendRequest)(nil),        // 7: xyz.block.ftl.v1.SendRequest
	(*SendResponse)(nil),       // 8: xyz.block.ftl.v1.SendResponse
	(*SyncSchemaRequest)(nil),  // 9: xyz.block.ftl.v1.SyncSchemaRequest
	(*SyncSchemaResponse)(nil), // 10: xyz.block.ftl.v1.SyncSchemaResponse
	(*Metadata_Pair)(nil),      // 11: xyz.block.ftl.v1.Metadata.Pair
	(*CallResponse_Error)(nil), // 12: xyz.block.ftl.v1.CallResponse.Error
	(*schema.VerbRef)(nil),     // 13: xyz.block.ftl.v1.schema.VerbRef
	(*schema.Module)(nil),      // 14: xyz.block.ftl.v1.schema.Module
}
var file_xyz_block_ftl_v1_ftl_proto_depIdxs = []int32{
	11, // 0: xyz.block.ftl.v1.Metadata.values:type_name -> xyz.block.ftl.v1.Metadata.Pair
	2,  // 1: xyz.block.ftl.v1.CallRequest.metadata:type_name -> xyz.block.ftl.v1.Metadata
	13, // 2: xyz.block.ftl.v1.CallRequest.verb:type_name -> xyz.block.ftl.v1.schema.VerbRef
	12, // 3: xyz.block.ftl.v1.CallResponse.error:type_name -> xyz.block.ftl.v1.CallResponse.Error
	13, // 4: xyz.block.ftl.v1.ListResponse.verbs:type_name -> xyz.block.ftl.v1.schema.VerbRef
	2,  // 5: xyz.block.ftl.v1.SendRequest.metadata:type_name -> xyz.block.ftl.v1.Metadata
	13, // 6: xyz.block.ftl.v1.SendRequest.verb:type_name -> xyz.block.ftl.v1.schema.VerbRef
	14, // 7: xyz.block.ftl.v1.SyncSchemaRequest.schema:type_name -> xyz.block.ftl.v1.schema.Module
	14, // 8: xyz.block.ftl.v1.SyncSchemaResponse.schema:type_name -> xyz.block.ftl.v1.schema.Module
	0,  // 9: xyz.block.ftl.v1.VerbService.Ping:input_type -> xyz.block.ftl.v1.PingRequest
	3,  // 10: xyz.block.ftl.v1.VerbService.Call:input_type -> xyz.block.ftl.v1.CallRequest
	7,  // 11: xyz.block.ftl.v1.VerbService.Send:input_type -> xyz.block.ftl.v1.SendRequest
	5,  // 12: xyz.block.ftl.v1.VerbService.List:input_type -> xyz.block.ftl.v1.ListRequest
	0,  // 13: xyz.block.ftl.v1.DevelService.Ping:input_type -> xyz.block.ftl.v1.PingRequest
	9,  // 14: xyz.block.ftl.v1.DevelService.SyncSchema:input_type -> xyz.block.ftl.v1.SyncSchemaRequest
	1,  // 15: xyz.block.ftl.v1.VerbService.Ping:output_type -> xyz.block.ftl.v1.PingResponse
	4,  // 16: xyz.block.ftl.v1.VerbService.Call:output_type -> xyz.block.ftl.v1.CallResponse
	8,  // 17: xyz.block.ftl.v1.VerbService.Send:output_type -> xyz.block.ftl.v1.SendResponse
	6,  // 18: xyz.block.ftl.v1.VerbService.List:output_type -> xyz.block.ftl.v1.ListResponse
	1,  // 19: xyz.block.ftl.v1.DevelService.Ping:output_type -> xyz.block.ftl.v1.PingResponse
	10, // 20: xyz.block.ftl.v1.DevelService.SyncSchema:output_type -> xyz.block.ftl.v1.SyncSchemaResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_xyz_block_ftl_v1_ftl_proto_init() }
func file_xyz_block_ftl_v1_ftl_proto_init() {
	if File_xyz_block_ftl_v1_ftl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncSchemaRequest); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncSchemaResponse); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata_Pair); i {
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
		file_xyz_block_ftl_v1_ftl_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse_Error); i {
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
	file_xyz_block_ftl_v1_ftl_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*CallResponse_Body)(nil),
		(*CallResponse_Error_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xyz_block_ftl_v1_ftl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_xyz_block_ftl_v1_ftl_proto_goTypes,
		DependencyIndexes: file_xyz_block_ftl_v1_ftl_proto_depIdxs,
		MessageInfos:      file_xyz_block_ftl_v1_ftl_proto_msgTypes,
	}.Build()
	File_xyz_block_ftl_v1_ftl_proto = out.File
	file_xyz_block_ftl_v1_ftl_proto_rawDesc = nil
	file_xyz_block_ftl_v1_ftl_proto_goTypes = nil
	file_xyz_block_ftl_v1_ftl_proto_depIdxs = nil
}
