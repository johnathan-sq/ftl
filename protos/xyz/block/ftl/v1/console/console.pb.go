// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: xyz/block/ftl/v1/console/console.proto

package pbconsole

import (
	v1 "github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1"
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

type Call struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RunnerKey     string `protobuf:"bytes,2,opt,name=runner_key,json=runnerKey,proto3" json:"runner_key,omitempty"`
	RequestId     int64  `protobuf:"varint,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ControllerKey string `protobuf:"bytes,4,opt,name=controller_key,json=controllerKey,proto3" json:"controller_key,omitempty"`
	TimeStamp     int64  `protobuf:"varint,5,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	SourceModule  string `protobuf:"bytes,6,opt,name=source_module,json=sourceModule,proto3" json:"source_module,omitempty"`
	SourceVerb    string `protobuf:"bytes,7,opt,name=source_verb,json=sourceVerb,proto3" json:"source_verb,omitempty"`
	DestModule    string `protobuf:"bytes,8,opt,name=dest_module,json=destModule,proto3" json:"dest_module,omitempty"`
	DestVerb      string `protobuf:"bytes,9,opt,name=dest_verb,json=destVerb,proto3" json:"dest_verb,omitempty"`
	DurationMs    int64  `protobuf:"varint,10,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	Request       []byte `protobuf:"bytes,11,opt,name=request,proto3" json:"request,omitempty"`
	Response      []byte `protobuf:"bytes,12,opt,name=response,proto3" json:"response,omitempty"`
	Error         string `protobuf:"bytes,13,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Call) Reset() {
	*x = Call{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Call) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Call) ProtoMessage() {}

func (x *Call) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Call.ProtoReflect.Descriptor instead.
func (*Call) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_console_console_proto_rawDescGZIP(), []int{0}
}

func (x *Call) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Call) GetRunnerKey() string {
	if x != nil {
		return x.RunnerKey
	}
	return ""
}

func (x *Call) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Call) GetControllerKey() string {
	if x != nil {
		return x.ControllerKey
	}
	return ""
}

func (x *Call) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *Call) GetSourceModule() string {
	if x != nil {
		return x.SourceModule
	}
	return ""
}

func (x *Call) GetSourceVerb() string {
	if x != nil {
		return x.SourceVerb
	}
	return ""
}

func (x *Call) GetDestModule() string {
	if x != nil {
		return x.DestModule
	}
	return ""
}

func (x *Call) GetDestVerb() string {
	if x != nil {
		return x.DestVerb
	}
	return ""
}

func (x *Call) GetDurationMs() int64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *Call) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Call) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Call) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Verb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verb  *schema.Verb `protobuf:"bytes,1,opt,name=verb,proto3" json:"verb,omitempty"`
	Calls []*Call      `protobuf:"bytes,2,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *Verb) Reset() {
	*x = Verb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verb) ProtoMessage() {}

func (x *Verb) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verb.ProtoReflect.Descriptor instead.
func (*Verb) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_console_console_proto_rawDescGZIP(), []int{1}
}

func (x *Verb) GetVerb() *schema.Verb {
	if x != nil {
		return x.Verb
	}
	return nil
}

func (x *Verb) GetCalls() []*Call {
	if x != nil {
		return x.Calls
	}
	return nil
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Language string         `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Verbs    []*Verb        `protobuf:"bytes,3,rep,name=verbs,proto3" json:"verbs,omitempty"`
	Data     []*schema.Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_console_console_proto_rawDescGZIP(), []int{2}
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Module) GetVerbs() []*Verb {
	if x != nil {
		return x.Verbs
	}
	return nil
}

func (x *Module) GetData() []*schema.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetModulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetModulesRequest) Reset() {
	*x = GetModulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModulesRequest) ProtoMessage() {}

func (x *GetModulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModulesRequest.ProtoReflect.Descriptor instead.
func (*GetModulesRequest) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_console_console_proto_rawDescGZIP(), []int{3}
}

type GetModulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*Module `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *GetModulesResponse) Reset() {
	*x = GetModulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModulesResponse) ProtoMessage() {}

func (x *GetModulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xyz_block_ftl_v1_console_console_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModulesResponse.ProtoReflect.Descriptor instead.
func (*GetModulesResponse) Descriptor() ([]byte, []int) {
	return file_xyz_block_ftl_v1_console_console_proto_rawDescGZIP(), []int{4}
}

func (x *GetModulesResponse) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

var File_xyz_block_ftl_v1_console_console_proto protoreflect.FileDescriptor

var file_xyz_block_ftl_v1_console_console_proto_rawDesc = []byte{
	0x0a, 0x26, 0x78, 0x79, 0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x66, 0x74, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x1a, 0x1a, 0x78, 0x79, 0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x66, 0x74,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x78, 0x79, 0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x66, 0x74, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x5f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x73, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73,
	0x74, 0x56, 0x65, 0x72, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x6f, 0x0a, 0x04, 0x56, 0x65, 0x72, 0x62, 0x12, 0x31, 0x0a, 0x04, 0x76, 0x65,
	0x72, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x34, 0x0a,
	0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x78,
	0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x63, 0x61,
	0x6c, 0x6c, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x34,
	0x0a, 0x05, 0x76, 0x65, 0x72, 0x62, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52, 0x05, 0x76,
	0x65, 0x72, 0x62, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66,
	0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x32, 0xc5,
	0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x78, 0x79, 0x7a, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x67, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x78, 0x79,
	0x7a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x78, 0x79, 0x7a, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x66, 0x74, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x48, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x42, 0x44, 0x35, 0x34, 0x35, 0x36, 0x36, 0x39,
	0x37, 0x35, 0x2f, 0x66, 0x74, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x78, 0x79,
	0x7a, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x66, 0x74, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x3b, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xyz_block_ftl_v1_console_console_proto_rawDescOnce sync.Once
	file_xyz_block_ftl_v1_console_console_proto_rawDescData = file_xyz_block_ftl_v1_console_console_proto_rawDesc
)

func file_xyz_block_ftl_v1_console_console_proto_rawDescGZIP() []byte {
	file_xyz_block_ftl_v1_console_console_proto_rawDescOnce.Do(func() {
		file_xyz_block_ftl_v1_console_console_proto_rawDescData = protoimpl.X.CompressGZIP(file_xyz_block_ftl_v1_console_console_proto_rawDescData)
	})
	return file_xyz_block_ftl_v1_console_console_proto_rawDescData
}

var file_xyz_block_ftl_v1_console_console_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_xyz_block_ftl_v1_console_console_proto_goTypes = []interface{}{
	(*Call)(nil),               // 0: xyz.block.ftl.v1.console.Call
	(*Verb)(nil),               // 1: xyz.block.ftl.v1.console.Verb
	(*Module)(nil),             // 2: xyz.block.ftl.v1.console.Module
	(*GetModulesRequest)(nil),  // 3: xyz.block.ftl.v1.console.GetModulesRequest
	(*GetModulesResponse)(nil), // 4: xyz.block.ftl.v1.console.GetModulesResponse
	(*schema.Verb)(nil),        // 5: xyz.block.ftl.v1.schema.Verb
	(*schema.Data)(nil),        // 6: xyz.block.ftl.v1.schema.Data
	(*v1.PingRequest)(nil),     // 7: xyz.block.ftl.v1.PingRequest
	(*v1.PingResponse)(nil),    // 8: xyz.block.ftl.v1.PingResponse
}
var file_xyz_block_ftl_v1_console_console_proto_depIdxs = []int32{
	5, // 0: xyz.block.ftl.v1.console.Verb.verb:type_name -> xyz.block.ftl.v1.schema.Verb
	0, // 1: xyz.block.ftl.v1.console.Verb.calls:type_name -> xyz.block.ftl.v1.console.Call
	1, // 2: xyz.block.ftl.v1.console.Module.verbs:type_name -> xyz.block.ftl.v1.console.Verb
	6, // 3: xyz.block.ftl.v1.console.Module.data:type_name -> xyz.block.ftl.v1.schema.Data
	2, // 4: xyz.block.ftl.v1.console.GetModulesResponse.modules:type_name -> xyz.block.ftl.v1.console.Module
	7, // 5: xyz.block.ftl.v1.console.ConsoleService.Ping:input_type -> xyz.block.ftl.v1.PingRequest
	3, // 6: xyz.block.ftl.v1.console.ConsoleService.GetModules:input_type -> xyz.block.ftl.v1.console.GetModulesRequest
	8, // 7: xyz.block.ftl.v1.console.ConsoleService.Ping:output_type -> xyz.block.ftl.v1.PingResponse
	4, // 8: xyz.block.ftl.v1.console.ConsoleService.GetModules:output_type -> xyz.block.ftl.v1.console.GetModulesResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_xyz_block_ftl_v1_console_console_proto_init() }
func file_xyz_block_ftl_v1_console_console_proto_init() {
	if File_xyz_block_ftl_v1_console_console_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xyz_block_ftl_v1_console_console_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Call); i {
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
		file_xyz_block_ftl_v1_console_console_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verb); i {
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
		file_xyz_block_ftl_v1_console_console_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_xyz_block_ftl_v1_console_console_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModulesRequest); i {
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
		file_xyz_block_ftl_v1_console_console_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModulesResponse); i {
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
			RawDescriptor: file_xyz_block_ftl_v1_console_console_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xyz_block_ftl_v1_console_console_proto_goTypes,
		DependencyIndexes: file_xyz_block_ftl_v1_console_console_proto_depIdxs,
		MessageInfos:      file_xyz_block_ftl_v1_console_console_proto_msgTypes,
	}.Build()
	File_xyz_block_ftl_v1_console_console_proto = out.File
	file_xyz_block_ftl_v1_console_console_proto_rawDesc = nil
	file_xyz_block_ftl_v1_console_console_proto_goTypes = nil
	file_xyz_block_ftl_v1_console_console_proto_depIdxs = nil
}
