// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: protos/bank.proto

package main

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CustomerInterface int32

const (
	CustomerInterface_unknown  CustomerInterface = 0
	CustomerInterface_query    CustomerInterface = 1
	CustomerInterface_deposit  CustomerInterface = 2
	CustomerInterface_withdraw CustomerInterface = 3
)

// Enum value maps for CustomerInterface.
var (
	CustomerInterface_name = map[int32]string{
		0: "unknown",
		1: "query",
		2: "deposit",
		3: "withdraw",
	}
	CustomerInterface_value = map[string]int32{
		"unknown":  0,
		"query":    1,
		"deposit":  2,
		"withdraw": 3,
	}
)

func (x CustomerInterface) Enum() *CustomerInterface {
	p := new(CustomerInterface)
	*p = x
	return p
}

func (x CustomerInterface) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerInterface) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_bank_proto_enumTypes[0].Descriptor()
}

func (CustomerInterface) Type() protoreflect.EnumType {
	return &file_protos_bank_proto_enumTypes[0]
}

func (x CustomerInterface) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerInterface.Descriptor instead.
func (CustomerInterface) EnumDescriptor() ([]byte, []int) {
	return file_protos_bank_proto_rawDescGZIP(), []int{0}
}

type BranchInterface int32

const (
	BranchInterface_propogate_withdraw BranchInterface = 0
	BranchInterface_propogate_deposit  BranchInterface = 1
)

// Enum value maps for BranchInterface.
var (
	BranchInterface_name = map[int32]string{
		0: "propogate_withdraw",
		1: "propogate_deposit",
	}
	BranchInterface_value = map[string]int32{
		"propogate_withdraw": 0,
		"propogate_deposit":  1,
	}
)

func (x BranchInterface) Enum() *BranchInterface {
	p := new(BranchInterface)
	*p = x
	return p
}

func (x BranchInterface) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BranchInterface) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_bank_proto_enumTypes[1].Descriptor()
}

func (BranchInterface) Type() protoreflect.EnumType {
	return &file_protos_bank_proto_enumTypes[1]
}

func (x BranchInterface) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BranchInterface.Descriptor instead.
func (BranchInterface) EnumDescriptor() ([]byte, []int) {
	return file_protos_bank_proto_rawDescGZIP(), []int{1}
}

type Result int32

const (
	Result_success Result = 0
	Result_error   Result = 1
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "success",
		1: "error",
	}
	Result_value = map[string]int32{
		"success": 0,
		"error":   1,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_bank_proto_enumTypes[2].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_protos_bank_proto_enumTypes[2]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_protos_bank_proto_rawDescGZIP(), []int{2}
}

type MsgDeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id_       int64             `protobuf:"varint,1,opt,name=id_,json=id,proto3" json:"id_,omitempty"`
	Interface CustomerInterface `protobuf:"varint,2,opt,name=interface,proto3,enum=bank.CustomerInterface" json:"interface,omitempty"`
	Money     float64           `protobuf:"fixed64,3,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *MsgDeliveryRequest) Reset() {
	*x = MsgDeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeliveryRequest) ProtoMessage() {}

func (x *MsgDeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeliveryRequest.ProtoReflect.Descriptor instead.
func (*MsgDeliveryRequest) Descriptor() ([]byte, []int) {
	return file_protos_bank_proto_rawDescGZIP(), []int{0}
}

func (x *MsgDeliveryRequest) GetId_() int64 {
	if x != nil {
		return x.Id_
	}
	return 0
}

func (x *MsgDeliveryRequest) GetInterface() CustomerInterface {
	if x != nil {
		return x.Interface
	}
	return CustomerInterface_unknown
}

func (x *MsgDeliveryRequest) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

type PropogateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id_       int64           `protobuf:"varint,1,opt,name=id_,json=id,proto3" json:"id_,omitempty"`
	Interface BranchInterface `protobuf:"varint,2,opt,name=interface,proto3,enum=bank.BranchInterface" json:"interface,omitempty"`
	Money     float64         `protobuf:"fixed64,3,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *PropogateRequest) Reset() {
	*x = PropogateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropogateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropogateRequest) ProtoMessage() {}

func (x *PropogateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropogateRequest.ProtoReflect.Descriptor instead.
func (*PropogateRequest) Descriptor() ([]byte, []int) {
	return file_protos_bank_proto_rawDescGZIP(), []int{1}
}

func (x *PropogateRequest) GetId_() int64 {
	if x != nil {
		return x.Id_
	}
	return 0
}

func (x *PropogateRequest) GetInterface() BranchInterface {
	if x != nil {
		return x.Interface
	}
	return BranchInterface_propogate_withdraw
}

func (x *PropogateRequest) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

type MsgDeliveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id_       int64             `protobuf:"varint,1,opt,name=id_,json=id,proto3" json:"id_,omitempty"`
	Result    Result            `protobuf:"varint,2,opt,name=result,proto3,enum=bank.Result" json:"result,omitempty"`
	Money     float64           `protobuf:"fixed64,3,opt,name=money,proto3" json:"money,omitempty"`
	Interface CustomerInterface `protobuf:"varint,4,opt,name=interface,proto3,enum=bank.CustomerInterface" json:"interface,omitempty"`
}

func (x *MsgDeliveryResponse) Reset() {
	*x = MsgDeliveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDeliveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDeliveryResponse) ProtoMessage() {}

func (x *MsgDeliveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDeliveryResponse.ProtoReflect.Descriptor instead.
func (*MsgDeliveryResponse) Descriptor() ([]byte, []int) {
	return file_protos_bank_proto_rawDescGZIP(), []int{2}
}

func (x *MsgDeliveryResponse) GetId_() int64 {
	if x != nil {
		return x.Id_
	}
	return 0
}

func (x *MsgDeliveryResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_success
}

func (x *MsgDeliveryResponse) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *MsgDeliveryResponse) GetInterface() CustomerInterface {
	if x != nil {
		return x.Interface
	}
	return CustomerInterface_unknown
}

var File_protos_bank_proto protoreflect.FileDescriptor

var file_protos_bank_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x22, 0x72, 0x0a, 0x12, 0x4d, 0x73, 0x67,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0f, 0x0a, 0x03, 0x69, 0x64, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x35, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x6e, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0f, 0x0a, 0x03, 0x69, 0x64, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x99, 0x01,
	0x0a, 0x13, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x69, 0x64, 0x5f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x35, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x09,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2a, 0x46, 0x0a, 0x11, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10,
	0x03, 0x2a, 0x40, 0x0a, 0x0f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x10, 0x01, 0x2a, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x01, 0x32, 0x91, 0x01, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x44,
	0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x18, 0x2e,
	0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x67, 0x61, 0x74,
	0x65, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x6d, 0x61,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_bank_proto_rawDescOnce sync.Once
	file_protos_bank_proto_rawDescData = file_protos_bank_proto_rawDesc
)

func file_protos_bank_proto_rawDescGZIP() []byte {
	file_protos_bank_proto_rawDescOnce.Do(func() {
		file_protos_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_bank_proto_rawDescData)
	})
	return file_protos_bank_proto_rawDescData
}

var file_protos_bank_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protos_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_bank_proto_goTypes = []interface{}{
	(CustomerInterface)(0),      // 0: bank.CustomerInterface
	(BranchInterface)(0),        // 1: bank.BranchInterface
	(Result)(0),                 // 2: bank.Result
	(*MsgDeliveryRequest)(nil),  // 3: bank.MsgDeliveryRequest
	(*PropogateRequest)(nil),    // 4: bank.PropogateRequest
	(*MsgDeliveryResponse)(nil), // 5: bank.MsgDeliveryResponse
}
var file_protos_bank_proto_depIdxs = []int32{
	0, // 0: bank.MsgDeliveryRequest.interface:type_name -> bank.CustomerInterface
	1, // 1: bank.PropogateRequest.interface:type_name -> bank.BranchInterface
	2, // 2: bank.MsgDeliveryResponse.result:type_name -> bank.Result
	0, // 3: bank.MsgDeliveryResponse.interface:type_name -> bank.CustomerInterface
	3, // 4: bank.Bank.MsgDelivery:input_type -> bank.MsgDeliveryRequest
	4, // 5: bank.Bank.PropogateMsg:input_type -> bank.PropogateRequest
	5, // 6: bank.Bank.MsgDelivery:output_type -> bank.MsgDeliveryResponse
	5, // 7: bank.Bank.PropogateMsg:output_type -> bank.MsgDeliveryResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_bank_proto_init() }
func file_protos_bank_proto_init() {
	if File_protos_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeliveryRequest); i {
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
		file_protos_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropogateRequest); i {
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
		file_protos_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDeliveryResponse); i {
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
			RawDescriptor: file_protos_bank_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_bank_proto_goTypes,
		DependencyIndexes: file_protos_bank_proto_depIdxs,
		EnumInfos:         file_protos_bank_proto_enumTypes,
		MessageInfos:      file_protos_bank_proto_msgTypes,
	}.Build()
	File_protos_bank_proto = out.File
	file_protos_bank_proto_rawDesc = nil
	file_protos_bank_proto_goTypes = nil
	file_protos_bank_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BankClient is the client API for Bank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BankClient interface {
	MsgDelivery(ctx context.Context, in *MsgDeliveryRequest, opts ...grpc.CallOption) (*MsgDeliveryResponse, error)
	PropogateMsg(ctx context.Context, in *PropogateRequest, opts ...grpc.CallOption) (*MsgDeliveryResponse, error)
}

type bankClient struct {
	cc grpc.ClientConnInterface
}

func NewBankClient(cc grpc.ClientConnInterface) BankClient {
	return &bankClient{cc}
}

func (c *bankClient) MsgDelivery(ctx context.Context, in *MsgDeliveryRequest, opts ...grpc.CallOption) (*MsgDeliveryResponse, error) {
	out := new(MsgDeliveryResponse)
	err := c.cc.Invoke(ctx, "/bank.Bank/MsgDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) PropogateMsg(ctx context.Context, in *PropogateRequest, opts ...grpc.CallOption) (*MsgDeliveryResponse, error) {
	out := new(MsgDeliveryResponse)
	err := c.cc.Invoke(ctx, "/bank.Bank/PropogateMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServer is the server API for Bank service.
type BankServer interface {
	MsgDelivery(context.Context, *MsgDeliveryRequest) (*MsgDeliveryResponse, error)
	PropogateMsg(context.Context, *PropogateRequest) (*MsgDeliveryResponse, error)
}

// UnimplementedBankServer can be embedded to have forward compatible implementations.
type UnimplementedBankServer struct {
}

func (*UnimplementedBankServer) MsgDelivery(context.Context, *MsgDeliveryRequest) (*MsgDeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgDelivery not implemented")
}
func (*UnimplementedBankServer) PropogateMsg(context.Context, *PropogateRequest) (*MsgDeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PropogateMsg not implemented")
}

func RegisterBankServer(s *grpc.Server, srv BankServer) {
	s.RegisterService(&_Bank_serviceDesc, srv)
}

func _Bank_MsgDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeliveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).MsgDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.Bank/MsgDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).MsgDelivery(ctx, req.(*MsgDeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_PropogateMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropogateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).PropogateMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank.Bank/PropogateMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).PropogateMsg(ctx, req.(*PropogateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bank_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bank.Bank",
	HandlerType: (*BankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgDelivery",
			Handler:    _Bank_MsgDelivery_Handler,
		},
		{
			MethodName: "PropogateMsg",
			Handler:    _Bank_PropogateMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/bank.proto",
}
