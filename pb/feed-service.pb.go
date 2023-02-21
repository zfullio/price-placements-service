// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/grpc/feed-service.proto

package pb

import (
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

type Placement int32

const (
	Placement_PLACEMENT_YANDEX_REALTY Placement = 0
	Placement_PLACEMENT_CIAN          Placement = 1
	Placement_PLACEMENT_AVITO         Placement = 2
	Placement_PLACEMENT_DOMCLICK      Placement = 3
)

// Enum value maps for Placement.
var (
	Placement_name = map[int32]string{
		0: "PLACEMENT_YANDEX_REALTY",
		1: "PLACEMENT_CIAN",
		2: "PLACEMENT_AVITO",
		3: "PLACEMENT_DOMCLICK",
	}
	Placement_value = map[string]int32{
		"PLACEMENT_YANDEX_REALTY": 0,
		"PLACEMENT_CIAN":          1,
		"PLACEMENT_AVITO":         2,
		"PLACEMENT_DOMCLICK":      3,
	}
)

func (x Placement) Enum() *Placement {
	p := new(Placement)
	*p = x
	return p
}

func (x Placement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Placement) Descriptor() protoreflect.EnumDescriptor {
	return file_api_grpc_feed_service_proto_enumTypes[0].Descriptor()
}

func (Placement) Type() protoreflect.EnumType {
	return &file_api_grpc_feed_service_proto_enumTypes[0]
}

func (x Placement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Placement.Descriptor instead.
func (Placement) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{0}
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpreadsheetId string `protobuf:"bytes,1,opt,name=spreadsheet_id,json=spreadsheetId,proto3" json:"spreadsheet_id,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetSpreadsheetId() string {
	if x != nil {
		return x.SpreadsheetId
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetResult() []string {
	if x != nil {
		return x.Result
	}
	return nil
}

type ValidateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeedUrl   string    `protobuf:"bytes,1,opt,name=feed_url,json=feedUrl,proto3" json:"feed_url,omitempty"`
	Placement Placement `protobuf:"varint,2,opt,name=placement,proto3,enum=Placement" json:"placement,omitempty"`
}

func (x *ValidateFeedRequest) Reset() {
	*x = ValidateFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedRequest) ProtoMessage() {}

func (x *ValidateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateFeedRequest.ProtoReflect.Descriptor instead.
func (*ValidateFeedRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateFeedRequest) GetFeedUrl() string {
	if x != nil {
		return x.FeedUrl
	}
	return ""
}

func (x *ValidateFeedRequest) GetPlacement() Placement {
	if x != nil {
		return x.Placement
	}
	return Placement_PLACEMENT_YANDEX_REALTY
}

type ValidateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ValidateFeedResponse) Reset() {
	*x = ValidateFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedResponse) ProtoMessage() {}

func (x *ValidateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateFeedResponse.ProtoReflect.Descriptor instead.
func (*ValidateFeedResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateFeedResponse) GetResult() []string {
	if x != nil {
		return x.Result
	}
	return nil
}

type ValidateFeedAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpreadsheetId string `protobuf:"bytes,1,opt,name=spreadsheet_id,json=spreadsheetId,proto3" json:"spreadsheet_id,omitempty"`
}

func (x *ValidateFeedAllRequest) Reset() {
	*x = ValidateFeedAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedAllRequest) ProtoMessage() {}

func (x *ValidateFeedAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateFeedAllRequest.ProtoReflect.Descriptor instead.
func (*ValidateFeedAllRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateFeedAllRequest) GetSpreadsheetId() string {
	if x != nil {
		return x.SpreadsheetId
	}
	return ""
}

type ValidateFeedAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ValidateFeedAllResponse) Reset() {
	*x = ValidateFeedAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedAllResponse) ProtoMessage() {}

func (x *ValidateFeedAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateFeedAllResponse.ProtoReflect.Descriptor instead.
func (*ValidateFeedAllResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateFeedAllResponse) GetResult() []string {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_grpc_feed_service_proto protoreflect.FileDescriptor

var file_api_grpc_feed_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5a, 0x0a,
	0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12,
	0x28, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x14, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3f, 0x0a, 0x16, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x72,
	0x65, 0x61, 0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x17, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x69, 0x0a,
	0x09, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4c,
	0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x59, 0x41, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x52,
	0x45, 0x41, 0x4c, 0x54, 0x59, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x49, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x56, 0x49, 0x54, 0x4f, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x4f,
	0x4d, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x03, 0x32, 0x90, 0x03, 0x0a, 0x0b, 0x46, 0x65, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x0d, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x11, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x0d,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x43, 0x69, 0x61, 0x6e,
	0x12, 0x0d, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x41, 0x76,
	0x69, 0x74, 0x6f, 0x12, 0x0d, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x73, 0x44, 0x6f, 0x6d, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_feed_service_proto_rawDescOnce sync.Once
	file_api_grpc_feed_service_proto_rawDescData = file_api_grpc_feed_service_proto_rawDesc
)

func file_api_grpc_feed_service_proto_rawDescGZIP() []byte {
	file_api_grpc_feed_service_proto_rawDescOnce.Do(func() {
		file_api_grpc_feed_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_feed_service_proto_rawDescData)
	})
	return file_api_grpc_feed_service_proto_rawDescData
}

var file_api_grpc_feed_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_grpc_feed_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_grpc_feed_service_proto_goTypes = []interface{}{
	(Placement)(0),                  // 0: Placement
	(*CheckRequest)(nil),            // 1: CheckRequest
	(*CheckResponse)(nil),           // 2: CheckResponse
	(*ValidateFeedRequest)(nil),     // 3: ValidateFeedRequest
	(*ValidateFeedResponse)(nil),    // 4: ValidateFeedResponse
	(*ValidateFeedAllRequest)(nil),  // 5: ValidateFeedAllRequest
	(*ValidateFeedAllResponse)(nil), // 6: ValidateFeedAllResponse
}
var file_api_grpc_feed_service_proto_depIdxs = []int32{
	0, // 0: ValidateFeedRequest.placement:type_name -> Placement
	1, // 1: FeedService.CheckPhonesAll:input_type -> CheckRequest
	1, // 2: FeedService.CheckPhonesRealty:input_type -> CheckRequest
	1, // 3: FeedService.CheckPhonesCian:input_type -> CheckRequest
	1, // 4: FeedService.CheckPhonesAvito:input_type -> CheckRequest
	1, // 5: FeedService.CheckPhonesDomclick:input_type -> CheckRequest
	3, // 6: FeedService.ValidateFeed:input_type -> ValidateFeedRequest
	5, // 7: FeedService.ValidateFeedAll:input_type -> ValidateFeedAllRequest
	2, // 8: FeedService.CheckPhonesAll:output_type -> CheckResponse
	2, // 9: FeedService.CheckPhonesRealty:output_type -> CheckResponse
	2, // 10: FeedService.CheckPhonesCian:output_type -> CheckResponse
	2, // 11: FeedService.CheckPhonesAvito:output_type -> CheckResponse
	2, // 12: FeedService.CheckPhonesDomclick:output_type -> CheckResponse
	4, // 13: FeedService.ValidateFeed:output_type -> ValidateFeedResponse
	6, // 14: FeedService.ValidateFeedAll:output_type -> ValidateFeedAllResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpc_feed_service_proto_init() }
func file_api_grpc_feed_service_proto_init() {
	if File_api_grpc_feed_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_feed_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_api_grpc_feed_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_api_grpc_feed_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateFeedRequest); i {
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
		file_api_grpc_feed_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateFeedResponse); i {
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
		file_api_grpc_feed_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateFeedAllRequest); i {
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
		file_api_grpc_feed_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateFeedAllResponse); i {
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
			RawDescriptor: file_api_grpc_feed_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_feed_service_proto_goTypes,
		DependencyIndexes: file_api_grpc_feed_service_proto_depIdxs,
		EnumInfos:         file_api_grpc_feed_service_proto_enumTypes,
		MessageInfos:      file_api_grpc_feed_service_proto_msgTypes,
	}.Build()
	File_api_grpc_feed_service_proto = out.File
	file_api_grpc_feed_service_proto_rawDesc = nil
	file_api_grpc_feed_service_proto_goTypes = nil
	file_api_grpc_feed_service_proto_depIdxs = nil
}
