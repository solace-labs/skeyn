// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/rule.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RuleBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddress string               `protobuf:"bytes,1,opt,name=walletAddress,proto3" json:"walletAddress,omitempty"`
	OwnerAddress  string               `protobuf:"bytes,2,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	Rules         []*AccessControlRule `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	SpendingCaps  []*SpendingCap       `protobuf:"bytes,4,rep,name=spendingCaps,proto3" json:"spendingCaps,omitempty"`
}

func (x *RuleBook) Reset() {
	*x = RuleBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleBook) ProtoMessage() {}

func (x *RuleBook) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleBook.ProtoReflect.Descriptor instead.
func (*RuleBook) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{0}
}

func (x *RuleBook) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

func (x *RuleBook) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *RuleBook) GetRules() []*AccessControlRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *RuleBook) GetSpendingCaps() []*SpendingCap {
	if x != nil {
		return x.SpendingCaps
	}
	return nil
}

// Consider for base token transfers
// TODO: Make fromAddress into a group
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparator  string `protobuf:"bytes,1,opt,name=comparator,proto3" json:"comparator,omitempty"`
	FromAddress string `protobuf:"bytes,2,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress   string `protobuf:"bytes,3,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	TargetValue int32  `protobuf:"varint,4,opt,name=targetValue,proto3" json:"targetValue,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{1}
}

func (x *Rule) GetComparator() string {
	if x != nil {
		return x.Comparator
	}
	return ""
}

func (x *Rule) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *Rule) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *Rule) GetTargetValue() int32 {
	if x != nil {
		return x.TargetValue
	}
	return 0
}

type SpendingCap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: Bytes vs String
	Sender       string                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	TokenAddress string                 `protobuf:"bytes,2,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	CurrentValue int32                  `protobuf:"varint,3,opt,name=currentValue,proto3" json:"currentValue,omitempty"`
	LastUpdated  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	Cap          int32                  `protobuf:"varint,5,opt,name=cap,proto3" json:"cap,omitempty"`
}

func (x *SpendingCap) Reset() {
	*x = SpendingCap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendingCap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendingCap) ProtoMessage() {}

func (x *SpendingCap) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendingCap.ProtoReflect.Descriptor instead.
func (*SpendingCap) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{2}
}

func (x *SpendingCap) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *SpendingCap) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

func (x *SpendingCap) GetCurrentValue() int32 {
	if x != nil {
		return x.CurrentValue
	}
	return 0
}

func (x *SpendingCap) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *SpendingCap) GetCap() int32 {
	if x != nil {
		return x.Cap
	}
	return 0
}

type AccessControlRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddr string `protobuf:"bytes,1,opt,name=walletAddr,proto3" json:"walletAddr,omitempty"`
	// Auto generated ID which takes into consideration the
	// primary keys involved in identifyng the rules - such that rules don't overlap
	GeneratedIds     []string          `protobuf:"bytes,2,rep,name=generatedIds,proto3" json:"generatedIds,omitempty"`
	Namespace        string            `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ValueRangeClause *ValueRangeClause `protobuf:"bytes,4,opt,name=valueRangeClause,proto3" json:"valueRangeClause,omitempty"`
	TimeWindowClause *TimeWindowClause `protobuf:"bytes,5,opt,name=timeWindowClause,proto3" json:"timeWindowClause,omitempty"`
	EscalationClause *EscalationClause `protobuf:"bytes,6,opt,name=escalationClause,proto3" json:"escalationClause,omitempty"`
	SenderGroup      *SenderGroup      `protobuf:"bytes,7,opt,name=senderGroup,proto3" json:"senderGroup,omitempty"`
	RecipientAddress string            `protobuf:"bytes,8,opt,name=recipientAddress,proto3" json:"recipientAddress,omitempty"`
	TokenAddress     string            `protobuf:"bytes,9,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
}

func (x *AccessControlRule) Reset() {
	*x = AccessControlRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessControlRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessControlRule) ProtoMessage() {}

func (x *AccessControlRule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessControlRule.ProtoReflect.Descriptor instead.
func (*AccessControlRule) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{3}
}

func (x *AccessControlRule) GetWalletAddr() string {
	if x != nil {
		return x.WalletAddr
	}
	return ""
}

func (x *AccessControlRule) GetGeneratedIds() []string {
	if x != nil {
		return x.GeneratedIds
	}
	return nil
}

func (x *AccessControlRule) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AccessControlRule) GetValueRangeClause() *ValueRangeClause {
	if x != nil {
		return x.ValueRangeClause
	}
	return nil
}

func (x *AccessControlRule) GetTimeWindowClause() *TimeWindowClause {
	if x != nil {
		return x.TimeWindowClause
	}
	return nil
}

func (x *AccessControlRule) GetEscalationClause() *EscalationClause {
	if x != nil {
		return x.EscalationClause
	}
	return nil
}

func (x *AccessControlRule) GetSenderGroup() *SenderGroup {
	if x != nil {
		return x.SenderGroup
	}
	return nil
}

func (x *AccessControlRule) GetRecipientAddress() string {
	if x != nil {
		return x.RecipientAddress
	}
	return ""
}

func (x *AccessControlRule) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

type SenderGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *SenderGroup) Reset() {
	*x = SenderGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SenderGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenderGroup) ProtoMessage() {}

func (x *SenderGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenderGroup.ProtoReflect.Descriptor instead.
func (*SenderGroup) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{4}
}

func (x *SenderGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SenderGroup) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type ValueRangeClause struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinVal uint64 `protobuf:"varint,1,opt,name=minVal,proto3" json:"minVal,omitempty"`
	MaxVal uint64 `protobuf:"varint,2,opt,name=maxVal,proto3" json:"maxVal,omitempty"`
}

func (x *ValueRangeClause) Reset() {
	*x = ValueRangeClause{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueRangeClause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueRangeClause) ProtoMessage() {}

func (x *ValueRangeClause) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueRangeClause.ProtoReflect.Descriptor instead.
func (*ValueRangeClause) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{5}
}

func (x *ValueRangeClause) GetMinVal() uint64 {
	if x != nil {
		return x.MinVal
	}
	return 0
}

func (x *ValueRangeClause) GetMaxVal() uint64 {
	if x != nil {
		return x.MaxVal
	}
	return 0
}

type TimeWindowClause struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TimeWindowClause) Reset() {
	*x = TimeWindowClause{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeWindowClause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeWindowClause) ProtoMessage() {}

func (x *TimeWindowClause) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeWindowClause.ProtoReflect.Descriptor instead.
func (*TimeWindowClause) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{6}
}

type EscalationClause struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Logic     string   `protobuf:"bytes,2,opt,name=logic,proto3" json:"logic,omitempty"` // ([0] OR [1]) AND ([2] OR [3])
}

func (x *EscalationClause) Reset() {
	*x = EscalationClause{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EscalationClause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EscalationClause) ProtoMessage() {}

func (x *EscalationClause) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EscalationClause.ProtoReflect.Descriptor instead.
func (*EscalationClause) Descriptor() ([]byte, []int) {
	return file_proto_rule_proto_rawDescGZIP(), []int{7}
}

func (x *EscalationClause) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *EscalationClause) GetLogic() string {
	if x != nil {
		return x.Logic
	}
	return ""
}

var File_proto_rule_proto protoreflect.FileDescriptor

var file_proto_rule_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x08, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x24, 0x0a, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x43, 0x61, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x70, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x43, 0x61, 0x70, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x61,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x61,
	0x70, 0x22, 0xb2, 0x03, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65,
	0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x43,
	0x6c, 0x61, 0x75, 0x73, 0x65, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x65, 0x73, 0x63, 0x61, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x45, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c,
	0x61, 0x75, 0x73, 0x65, 0x52, 0x10, 0x65, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x69, 0x6e, 0x56, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x69, 0x6e,
	0x56, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x54,
	0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x22,
	0x46, 0x0a, 0x10, 0x45, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x61,
	0x75, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rule_proto_rawDescOnce sync.Once
	file_proto_rule_proto_rawDescData = file_proto_rule_proto_rawDesc
)

func file_proto_rule_proto_rawDescGZIP() []byte {
	file_proto_rule_proto_rawDescOnce.Do(func() {
		file_proto_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rule_proto_rawDescData)
	})
	return file_proto_rule_proto_rawDescData
}

var file_proto_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_rule_proto_goTypes = []interface{}{
	(*RuleBook)(nil),              // 0: RuleBook
	(*Rule)(nil),                  // 1: Rule
	(*SpendingCap)(nil),           // 2: SpendingCap
	(*AccessControlRule)(nil),     // 3: AccessControlRule
	(*SenderGroup)(nil),           // 4: SenderGroup
	(*ValueRangeClause)(nil),      // 5: ValueRangeClause
	(*TimeWindowClause)(nil),      // 6: TimeWindowClause
	(*EscalationClause)(nil),      // 7: EscalationClause
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_proto_rule_proto_depIdxs = []int32{
	3, // 0: RuleBook.rules:type_name -> AccessControlRule
	2, // 1: RuleBook.spendingCaps:type_name -> SpendingCap
	8, // 2: SpendingCap.lastUpdated:type_name -> google.protobuf.Timestamp
	5, // 3: AccessControlRule.valueRangeClause:type_name -> ValueRangeClause
	6, // 4: AccessControlRule.timeWindowClause:type_name -> TimeWindowClause
	7, // 5: AccessControlRule.escalationClause:type_name -> EscalationClause
	4, // 6: AccessControlRule.senderGroup:type_name -> SenderGroup
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_rule_proto_init() }
func file_proto_rule_proto_init() {
	if File_proto_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleBook); i {
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
		file_proto_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_proto_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpendingCap); i {
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
		file_proto_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessControlRule); i {
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
		file_proto_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SenderGroup); i {
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
		file_proto_rule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueRangeClause); i {
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
		file_proto_rule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeWindowClause); i {
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
		file_proto_rule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EscalationClause); i {
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
			RawDescriptor: file_proto_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_rule_proto_goTypes,
		DependencyIndexes: file_proto_rule_proto_depIdxs,
		MessageInfos:      file_proto_rule_proto_msgTypes,
	}.Build()
	File_proto_rule_proto = out.File
	file_proto_rule_proto_rawDesc = nil
	file_proto_rule_proto_goTypes = nil
	file_proto_rule_proto_depIdxs = nil
}
