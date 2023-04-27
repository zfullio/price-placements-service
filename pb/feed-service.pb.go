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
	Placement_PLACEMENT_UNKNOWN       Placement = 0
	Placement_PLACEMENT_CIAN          Placement = 1
	Placement_PLACEMENT_AVITO         Placement = 2
	Placement_PLACEMENT_DOMCLICK      Placement = 3
	Placement_PLACEMENT_M2            Placement = 4
	Placement_PLACEMENT_JCAT          Placement = 5
	Placement_PLACEMENT_GDETOTDOM     Placement = 6
	Placement_PLACEMENT_CALUGA_HOUSE  Placement = 7
	Placement_PLACEMENT_FLATOUTLET    Placement = 8
	Placement_PLACEMENT_GDEETOTDOM    Placement = 9
	Placement_PLACEMENT_HYBRID        Placement = 10
	Placement_PLACEMENT_AVAHO         Placement = 11
	Placement_PLACEMENT_NOVOSTROY_M   Placement = 12
	Placement_PLACEMENT_YANDEX_REALTY Placement = 13
)

// Enum value maps for Placement.
var (
	Placement_name = map[int32]string{
		0:  "PLACEMENT_UNKNOWN",
		1:  "PLACEMENT_CIAN",
		2:  "PLACEMENT_AVITO",
		3:  "PLACEMENT_DOMCLICK",
		4:  "PLACEMENT_M2",
		5:  "PLACEMENT_JCAT",
		6:  "PLACEMENT_GDETOTDOM",
		7:  "PLACEMENT_CALUGA_HOUSE",
		8:  "PLACEMENT_FLATOUTLET",
		9:  "PLACEMENT_GDEETOTDOM",
		10: "PLACEMENT_HYBRID",
		11: "PLACEMENT_AVAHO",
		12: "PLACEMENT_NOVOSTROY_M",
		13: "PLACEMENT_YANDEX_REALTY",
	}
	Placement_value = map[string]int32{
		"PLACEMENT_UNKNOWN":       0,
		"PLACEMENT_CIAN":          1,
		"PLACEMENT_AVITO":         2,
		"PLACEMENT_DOMCLICK":      3,
		"PLACEMENT_M2":            4,
		"PLACEMENT_JCAT":          5,
		"PLACEMENT_GDETOTDOM":     6,
		"PLACEMENT_CALUGA_HOUSE":  7,
		"PLACEMENT_FLATOUTLET":    8,
		"PLACEMENT_GDEETOTDOM":    9,
		"PLACEMENT_HYBRID":        10,
		"PLACEMENT_AVAHO":         11,
		"PLACEMENT_NOVOSTROY_M":   12,
		"PLACEMENT_YANDEX_REALTY": 13,
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

type Status int32

const (
	Status_UNKNOWN Status = 0
	Status_WARNING Status = 1
	Status_ERROR   Status = 2
	Status_OK      Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "WARNING",
		2: "ERROR",
		3: "OK",
	}
	Status_value = map[string]int32{
		"UNKNOWN": 0,
		"WARNING": 1,
		"ERROR":   2,
		"OK":      3,
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
	return file_api_grpc_feed_service_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_grpc_feed_service_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{1}
}

type CheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Developer string    `protobuf:"bytes,1,opt,name=developer,proto3" json:"developer,omitempty"`
	Placement Placement `protobuf:"varint,2,opt,name=placement,proto3,enum=Placement" json:"placement,omitempty"`
	Base      Placement `protobuf:"varint,3,opt,name=base,proto3,enum=Placement" json:"base,omitempty"`
	Url       string    `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Message   string    `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Result    Status    `protobuf:"varint,6,opt,name=result,proto3,enum=Status" json:"result,omitempty"`
}

func (x *CheckResult) Reset() {
	*x = CheckResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResult) ProtoMessage() {}

func (x *CheckResult) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CheckResult.ProtoReflect.Descriptor instead.
func (*CheckResult) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{0}
}

func (x *CheckResult) GetDeveloper() string {
	if x != nil {
		return x.Developer
	}
	return ""
}

func (x *CheckResult) GetPlacement() Placement {
	if x != nil {
		return x.Placement
	}
	return Placement_PLACEMENT_UNKNOWN
}

func (x *CheckResult) GetBase() Placement {
	if x != nil {
		return x.Base
	}
	return Placement_PLACEMENT_UNKNOWN
}

func (x *CheckResult) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CheckResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CheckResult) GetResult() Status {
	if x != nil {
		return x.Result
	}
	return Status_UNKNOWN
}

type CheckPhonesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpreadsheetId string    `protobuf:"bytes,1,opt,name=spreadsheet_id,json=spreadsheetId,proto3" json:"spreadsheet_id,omitempty"`
	Developer     string    `protobuf:"bytes,2,opt,name=developer,proto3" json:"developer,omitempty"`
	Placement     Placement `protobuf:"varint,3,opt,name=placement,proto3,enum=Placement" json:"placement,omitempty"`
}

func (x *CheckPhonesRequest) Reset() {
	*x = CheckPhonesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPhonesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPhonesRequest) ProtoMessage() {}

func (x *CheckPhonesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CheckPhonesRequest.ProtoReflect.Descriptor instead.
func (*CheckPhonesRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPhonesRequest) GetSpreadsheetId() string {
	if x != nil {
		return x.SpreadsheetId
	}
	return ""
}

func (x *CheckPhonesRequest) GetDeveloper() string {
	if x != nil {
		return x.Developer
	}
	return ""
}

func (x *CheckPhonesRequest) GetPlacement() Placement {
	if x != nil {
		return x.Placement
	}
	return Placement_PLACEMENT_UNKNOWN
}

type CheckPhonesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CheckResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *CheckPhonesResponse) Reset() {
	*x = CheckPhonesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPhonesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPhonesResponse) ProtoMessage() {}

func (x *CheckPhonesResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CheckPhonesResponse.ProtoReflect.Descriptor instead.
func (*CheckPhonesResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckPhonesResponse) GetResult() []*CheckResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type CheckPhonesAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpreadsheetId string `protobuf:"bytes,1,opt,name=spreadsheet_id,json=spreadsheetId,proto3" json:"spreadsheet_id,omitempty"`
	Developer     string `protobuf:"bytes,2,opt,name=developer,proto3" json:"developer,omitempty"`
}

func (x *CheckPhonesAllRequest) Reset() {
	*x = CheckPhonesAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPhonesAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPhonesAllRequest) ProtoMessage() {}

func (x *CheckPhonesAllRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CheckPhonesAllRequest.ProtoReflect.Descriptor instead.
func (*CheckPhonesAllRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{3}
}

func (x *CheckPhonesAllRequest) GetSpreadsheetId() string {
	if x != nil {
		return x.SpreadsheetId
	}
	return ""
}

func (x *CheckPhonesAllRequest) GetDeveloper() string {
	if x != nil {
		return x.Developer
	}
	return ""
}

type CheckPhonesAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CheckResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *CheckPhonesAllResponse) Reset() {
	*x = CheckPhonesAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPhonesAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPhonesAllResponse) ProtoMessage() {}

func (x *CheckPhonesAllResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CheckPhonesAllResponse.ProtoReflect.Descriptor instead.
func (*CheckPhonesAllResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{4}
}

func (x *CheckPhonesAllResponse) GetResult() []*CheckResult {
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
		mi := &file_api_grpc_feed_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedRequest) ProtoMessage() {}

func (x *ValidateFeedRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ValidateFeedRequest.ProtoReflect.Descriptor instead.
func (*ValidateFeedRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{5}
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
	return Placement_PLACEMENT_UNKNOWN
}

type ValidateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CheckResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ValidateFeedResponse) Reset() {
	*x = ValidateFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedResponse) ProtoMessage() {}

func (x *ValidateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[6]
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
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateFeedResponse) GetResult() []*CheckResult {
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
	Developer     string `protobuf:"bytes,2,opt,name=developer,proto3" json:"developer,omitempty"`
}

func (x *ValidateFeedAllRequest) Reset() {
	*x = ValidateFeedAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedAllRequest) ProtoMessage() {}

func (x *ValidateFeedAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[7]
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
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateFeedAllRequest) GetSpreadsheetId() string {
	if x != nil {
		return x.SpreadsheetId
	}
	return ""
}

func (x *ValidateFeedAllRequest) GetDeveloper() string {
	if x != nil {
		return x.Developer
	}
	return ""
}

type ValidateFeedAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CheckResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ValidateFeedAllResponse) Reset() {
	*x = ValidateFeedAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_feed_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateFeedAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateFeedAllResponse) ProtoMessage() {}

func (x *ValidateFeedAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_feed_service_proto_msgTypes[8]
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
	return file_api_grpc_feed_service_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateFeedAllResponse) GetResult() []*CheckResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_grpc_feed_service_proto protoreflect.FileDescriptor

var file_api_grpc_feed_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01,
	0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x70, 0x72,
	0x65, 0x61, 0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5c, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x5a, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65,
	0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x3c, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5d, 0x0a,
	0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x70, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x73, 0x68, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x17,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0xd5, 0x02,
	0x0a, 0x09, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x43, 0x49, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x41, 0x56, 0x49, 0x54, 0x4f, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x4f, 0x4d, 0x43, 0x4c, 0x49, 0x43,
	0x4b, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x4d, 0x32, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x4a, 0x43, 0x41, 0x54, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x44, 0x45, 0x54, 0x4f, 0x54, 0x44, 0x4f, 0x4d,
	0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x43, 0x41, 0x4c, 0x55, 0x47, 0x41, 0x5f, 0x48, 0x4f, 0x55, 0x53, 0x45, 0x10, 0x07, 0x12, 0x18,
	0x0a, 0x14, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4c, 0x41, 0x54,
	0x4f, 0x55, 0x54, 0x4c, 0x45, 0x54, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x44, 0x45, 0x45, 0x54, 0x4f, 0x54, 0x44, 0x4f, 0x4d,
	0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x48, 0x59, 0x42, 0x52, 0x49, 0x44, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x48, 0x4f, 0x10, 0x0b, 0x12, 0x19, 0x0a,
	0x15, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x56, 0x4f, 0x53,
	0x54, 0x52, 0x4f, 0x59, 0x5f, 0x4d, 0x10, 0x0c, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4c, 0x41, 0x43,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x59, 0x41, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x52, 0x45, 0x41,
	0x4c, 0x54, 0x59, 0x10, 0x0d, 0x2a, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x03, 0x32, 0xd2, 0x03, 0x0a,
	0x0b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x61, 0x6c, 0x74, 0x79, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x43, 0x69,
	0x61, 0x6e, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x41, 0x76, 0x69, 0x74,
	0x6f, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x13,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x44, 0x6f, 0x6d, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x14,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x17,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_api_grpc_feed_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_grpc_feed_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_grpc_feed_service_proto_goTypes = []interface{}{
	(Placement)(0),                  // 0: Placement
	(Status)(0),                     // 1: Status
	(*CheckResult)(nil),             // 2: CheckResult
	(*CheckPhonesRequest)(nil),      // 3: CheckPhonesRequest
	(*CheckPhonesResponse)(nil),     // 4: CheckPhonesResponse
	(*CheckPhonesAllRequest)(nil),   // 5: CheckPhonesAllRequest
	(*CheckPhonesAllResponse)(nil),  // 6: CheckPhonesAllResponse
	(*ValidateFeedRequest)(nil),     // 7: ValidateFeedRequest
	(*ValidateFeedResponse)(nil),    // 8: ValidateFeedResponse
	(*ValidateFeedAllRequest)(nil),  // 9: ValidateFeedAllRequest
	(*ValidateFeedAllResponse)(nil), // 10: ValidateFeedAllResponse
}
var file_api_grpc_feed_service_proto_depIdxs = []int32{
	0,  // 0: CheckResult.placement:type_name -> Placement
	0,  // 1: CheckResult.base:type_name -> Placement
	1,  // 2: CheckResult.result:type_name -> Status
	0,  // 3: CheckPhonesRequest.placement:type_name -> Placement
	2,  // 4: CheckPhonesResponse.result:type_name -> CheckResult
	2,  // 5: CheckPhonesAllResponse.result:type_name -> CheckResult
	0,  // 6: ValidateFeedRequest.placement:type_name -> Placement
	2,  // 7: ValidateFeedResponse.result:type_name -> CheckResult
	2,  // 8: ValidateFeedAllResponse.result:type_name -> CheckResult
	5,  // 9: FeedService.CheckPhonesAll:input_type -> CheckPhonesAllRequest
	3,  // 10: FeedService.CheckPhonesRealty:input_type -> CheckPhonesRequest
	3,  // 11: FeedService.CheckPhonesCian:input_type -> CheckPhonesRequest
	3,  // 12: FeedService.CheckPhonesAvito:input_type -> CheckPhonesRequest
	3,  // 13: FeedService.CheckPhonesDomclick:input_type -> CheckPhonesRequest
	7,  // 14: FeedService.ValidateFeed:input_type -> ValidateFeedRequest
	9,  // 15: FeedService.ValidateFeedAll:input_type -> ValidateFeedAllRequest
	6,  // 16: FeedService.CheckPhonesAll:output_type -> CheckPhonesAllResponse
	4,  // 17: FeedService.CheckPhonesRealty:output_type -> CheckPhonesResponse
	4,  // 18: FeedService.CheckPhonesCian:output_type -> CheckPhonesResponse
	4,  // 19: FeedService.CheckPhonesAvito:output_type -> CheckPhonesResponse
	4,  // 20: FeedService.CheckPhonesDomclick:output_type -> CheckPhonesResponse
	8,  // 21: FeedService.ValidateFeed:output_type -> ValidateFeedResponse
	10, // 22: FeedService.ValidateFeedAll:output_type -> ValidateFeedAllResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_grpc_feed_service_proto_init() }
func file_api_grpc_feed_service_proto_init() {
	if File_api_grpc_feed_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_feed_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResult); i {
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
			switch v := v.(*CheckPhonesRequest); i {
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
			switch v := v.(*CheckPhonesResponse); i {
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
			switch v := v.(*CheckPhonesAllRequest); i {
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
			switch v := v.(*CheckPhonesAllResponse); i {
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
		file_api_grpc_feed_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_grpc_feed_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_grpc_feed_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
			NumEnums:      2,
			NumMessages:   9,
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
