// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: blxr-mev.proto

package grpc_testing

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

type TxHashListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthHeader string   `protobuf:"bytes,1,opt,name=auth_header,json=authHeader,proto3" json:"auth_header,omitempty"`
	TxHashes   [][]byte `protobuf:"bytes,2,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
}

func (x *TxHashListRequest) Reset() {
	*x = TxHashListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxHashListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxHashListRequest) ProtoMessage() {}

func (x *TxHashListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxHashListRequest.ProtoReflect.Descriptor instead.
func (*TxHashListRequest) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{0}
}

func (x *TxHashListRequest) GetAuthHeader() string {
	if x != nil {
		return x.AuthHeader
	}
	return ""
}

func (x *TxHashListRequest) GetTxHashes() [][]byte {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

type ShortIDListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortIDs []uint32 `protobuf:"varint,1,rep,packed,name=shortIDs,proto3" json:"shortIDs,omitempty"`
}

func (x *ShortIDListReply) Reset() {
	*x = ShortIDListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortIDListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortIDListReply) ProtoMessage() {}

func (x *ShortIDListReply) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortIDListReply.ProtoReflect.Descriptor instead.
func (*ShortIDListReply) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{1}
}

func (x *ShortIDListReply) GetShortIDs() []uint32 {
	if x != nil {
		return x.ShortIDs
	}
	return nil
}

type BidTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot                 uint64 `protobuf:"varint,1,opt,name=Slot,proto3" json:"Slot,omitempty"`
	ParentHash           []byte `protobuf:"bytes,2,opt,name=ParentHash,proto3" json:"ParentHash,omitempty"`
	BlockHash            []byte `protobuf:"bytes,3,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	BuilderPubkey        []byte `protobuf:"bytes,4,opt,name=BuilderPubkey,proto3" json:"BuilderPubkey,omitempty"`
	ProposerPubkey       []byte `protobuf:"bytes,5,opt,name=ProposerPubkey,proto3" json:"ProposerPubkey,omitempty"`
	ProposerFeeRecipient []byte `protobuf:"bytes,6,opt,name=ProposerFeeRecipient,proto3" json:"ProposerFeeRecipient,omitempty"`
	GasLimit             uint64 `protobuf:"varint,7,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
	GasUsed              uint64 `protobuf:"varint,8,opt,name=GasUsed,proto3" json:"GasUsed,omitempty"`
	Value                string `protobuf:"bytes,9,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *BidTrace) Reset() {
	*x = BidTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidTrace) ProtoMessage() {}

func (x *BidTrace) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidTrace.ProtoReflect.Descriptor instead.
func (*BidTrace) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{2}
}

func (x *BidTrace) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *BidTrace) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *BidTrace) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *BidTrace) GetBuilderPubkey() []byte {
	if x != nil {
		return x.BuilderPubkey
	}
	return nil
}

func (x *BidTrace) GetProposerPubkey() []byte {
	if x != nil {
		return x.ProposerPubkey
	}
	return nil
}

func (x *BidTrace) GetProposerFeeRecipient() []byte {
	if x != nil {
		return x.ProposerFeeRecipient
	}
	return nil
}

func (x *BidTrace) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *BidTrace) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *BidTrace) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Withdrawal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index          uint64 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	ValidatorIndex uint64 `protobuf:"varint,2,opt,name=ValidatorIndex,proto3" json:"ValidatorIndex,omitempty"`
	Address        []byte `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Amount         uint64 `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *Withdrawal) Reset() {
	*x = Withdrawal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Withdrawal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdrawal) ProtoMessage() {}

func (x *Withdrawal) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdrawal.ProtoReflect.Descriptor instead.
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{3}
}

func (x *Withdrawal) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Withdrawal) GetValidatorIndex() uint64 {
	if x != nil {
		return x.ValidatorIndex
	}
	return 0
}

func (x *Withdrawal) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Withdrawal) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ExecutionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash    []byte        `protobuf:"bytes,1,opt,name=ParentHash,proto3" json:"ParentHash,omitempty"`
	StateRoot     []byte        `protobuf:"bytes,2,opt,name=StateRoot,proto3" json:"StateRoot,omitempty"`
	ReceiptsRoot  []byte        `protobuf:"bytes,3,opt,name=ReceiptsRoot,proto3" json:"ReceiptsRoot,omitempty"`
	LogsBloom     []byte        `protobuf:"bytes,4,opt,name=LogsBloom,proto3" json:"LogsBloom,omitempty"`
	PrevRandao    []byte        `protobuf:"bytes,5,opt,name=PrevRandao,proto3" json:"PrevRandao,omitempty"`
	ExtraData     []byte        `protobuf:"bytes,6,opt,name=ExtraData,proto3" json:"ExtraData,omitempty"`
	BaseFeePerGas []byte        `protobuf:"bytes,7,opt,name=BaseFeePerGas,proto3" json:"BaseFeePerGas,omitempty"`
	FeeRecipient  []byte        `protobuf:"bytes,8,opt,name=FeeRecipient,proto3" json:"FeeRecipient,omitempty"`
	BlockHash     []byte        `protobuf:"bytes,9,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Transactions  []*CompressTx `protobuf:"bytes,10,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
	Withdrawals   []*Withdrawal `protobuf:"bytes,11,rep,name=Withdrawals,proto3" json:"Withdrawals,omitempty"`
	BlockNumber   uint64        `protobuf:"varint,12,opt,name=BlockNumber,proto3" json:"BlockNumber,omitempty"`
	GasLimit      uint64        `protobuf:"varint,13,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
	Timestamp     uint64        `protobuf:"varint,14,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	GasUsed       uint64        `protobuf:"varint,15,opt,name=GasUsed,proto3" json:"GasUsed,omitempty"`
}

func (x *ExecutionPayload) Reset() {
	*x = ExecutionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionPayload) ProtoMessage() {}

func (x *ExecutionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionPayload.ProtoReflect.Descriptor instead.
func (*ExecutionPayload) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{4}
}

func (x *ExecutionPayload) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *ExecutionPayload) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

func (x *ExecutionPayload) GetReceiptsRoot() []byte {
	if x != nil {
		return x.ReceiptsRoot
	}
	return nil
}

func (x *ExecutionPayload) GetLogsBloom() []byte {
	if x != nil {
		return x.LogsBloom
	}
	return nil
}

func (x *ExecutionPayload) GetPrevRandao() []byte {
	if x != nil {
		return x.PrevRandao
	}
	return nil
}

func (x *ExecutionPayload) GetExtraData() []byte {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

func (x *ExecutionPayload) GetBaseFeePerGas() []byte {
	if x != nil {
		return x.BaseFeePerGas
	}
	return nil
}

func (x *ExecutionPayload) GetFeeRecipient() []byte {
	if x != nil {
		return x.FeeRecipient
	}
	return nil
}

func (x *ExecutionPayload) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *ExecutionPayload) GetTransactions() []*CompressTx {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ExecutionPayload) GetWithdrawals() []*Withdrawal {
	if x != nil {
		return x.Withdrawals
	}
	return nil
}

func (x *ExecutionPayload) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *ExecutionPayload) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *ExecutionPayload) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ExecutionPayload) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

type CompressTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawData []byte `protobuf:"bytes,1,opt,name=rawData,proto3" json:"rawData,omitempty"`
	ShortID uint32 `protobuf:"varint,2,opt,name=shortID,proto3" json:"shortID,omitempty"`
}

func (x *CompressTx) Reset() {
	*x = CompressTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompressTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompressTx) ProtoMessage() {}

func (x *CompressTx) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompressTx.ProtoReflect.Descriptor instead.
func (*CompressTx) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{5}
}

func (x *CompressTx) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

func (x *CompressTx) GetShortID() uint32 {
	if x != nil {
		return x.ShortID
	}
	return 0
}

type SubmitBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidTrace         *BidTrace         `protobuf:"bytes,1,opt,name=bidTrace,proto3" json:"bidTrace,omitempty"`
	ExecutionPayload *ExecutionPayload `protobuf:"bytes,2,opt,name=executionPayload,proto3" json:"executionPayload,omitempty"`
	Signature        []byte            `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	AuthHeader       string            `protobuf:"bytes,4,opt,name=auth_header,json=authHeader,proto3" json:"auth_header,omitempty"`
}

func (x *SubmitBlockRequest) Reset() {
	*x = SubmitBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitBlockRequest) ProtoMessage() {}

func (x *SubmitBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitBlockRequest.ProtoReflect.Descriptor instead.
func (*SubmitBlockRequest) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{6}
}

func (x *SubmitBlockRequest) GetBidTrace() *BidTrace {
	if x != nil {
		return x.BidTrace
	}
	return nil
}

func (x *SubmitBlockRequest) GetExecutionPayload() *ExecutionPayload {
	if x != nil {
		return x.ExecutionPayload
	}
	return nil
}

func (x *SubmitBlockRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SubmitBlockRequest) GetAuthHeader() string {
	if x != nil {
		return x.AuthHeader
	}
	return ""
}

type SubmitBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubmitBlockResponse) Reset() {
	*x = SubmitBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blxr_mev_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitBlockResponse) ProtoMessage() {}

func (x *SubmitBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blxr_mev_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitBlockResponse.ProtoReflect.Descriptor instead.
func (*SubmitBlockResponse) Descriptor() ([]byte, []int) {
	return file_blxr_mev_proto_rawDescGZIP(), []int{7}
}

func (x *SubmitBlockResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitBlockResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_blxr_mev_proto protoreflect.FileDescriptor

var file_blxr_mev_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6c, 0x78, 0x72, 0x2d, 0x6d, 0x65, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x50, 0x0a, 0x11, 0x54, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x10, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x73, 0x22, 0xaa, 0x02, 0x0a, 0x08,
	0x42, 0x69, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x73, 0x55,
	0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x47, 0x61, 0x73, 0x55, 0x73,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7c, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9e, 0x04, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x72, 0x65, 0x76, 0x52, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x50, 0x72, 0x65, 0x76, 0x52, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x61, 0x73,
	0x65, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x42, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x46, 0x65, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x37, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x78, 0x52, 0x0c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x61, 0x6c, 0x52, 0x0b, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x22, 0x40, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x22, 0xc9, 0x01, 0x0a, 0x12, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x42, 0x69, 0x64,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x08, 0x62, 0x69, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9a, 0x01, 0x0a, 0x07, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x43, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49,
	0x44, 0x73, 0x12, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x54, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x44,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x53, 0x0a, 0x05, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x68, 0x65,
	0x6e, 0x72, 0x79, 0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blxr_mev_proto_rawDescOnce sync.Once
	file_blxr_mev_proto_rawDescData = file_blxr_mev_proto_rawDesc
)

func file_blxr_mev_proto_rawDescGZIP() []byte {
	file_blxr_mev_proto_rawDescOnce.Do(func() {
		file_blxr_mev_proto_rawDescData = protoimpl.X.CompressGZIP(file_blxr_mev_proto_rawDescData)
	})
	return file_blxr_mev_proto_rawDescData
}

var file_blxr_mev_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_blxr_mev_proto_goTypes = []interface{}{
	(*TxHashListRequest)(nil),   // 0: gateway.TxHashListRequest
	(*ShortIDListReply)(nil),    // 1: gateway.ShortIDListReply
	(*BidTrace)(nil),            // 2: gateway.BidTrace
	(*Withdrawal)(nil),          // 3: gateway.Withdrawal
	(*ExecutionPayload)(nil),    // 4: gateway.ExecutionPayload
	(*CompressTx)(nil),          // 5: gateway.compressTx
	(*SubmitBlockRequest)(nil),  // 6: gateway.SubmitBlockRequest
	(*SubmitBlockResponse)(nil), // 7: gateway.SubmitBlockResponse
}
var file_blxr_mev_proto_depIdxs = []int32{
	5, // 0: gateway.ExecutionPayload.Transactions:type_name -> gateway.compressTx
	3, // 1: gateway.ExecutionPayload.Withdrawals:type_name -> gateway.Withdrawal
	2, // 2: gateway.SubmitBlockRequest.bidTrace:type_name -> gateway.BidTrace
	4, // 3: gateway.SubmitBlockRequest.executionPayload:type_name -> gateway.ExecutionPayload
	0, // 4: gateway.Gateway.ShortIDs:input_type -> gateway.TxHashListRequest
	6, // 5: gateway.Gateway.SubmitBlock:input_type -> gateway.SubmitBlockRequest
	6, // 6: gateway.Relay.SubmitBlock:input_type -> gateway.SubmitBlockRequest
	1, // 7: gateway.Gateway.ShortIDs:output_type -> gateway.ShortIDListReply
	7, // 8: gateway.Gateway.SubmitBlock:output_type -> gateway.SubmitBlockResponse
	7, // 9: gateway.Relay.SubmitBlock:output_type -> gateway.SubmitBlockResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_blxr_mev_proto_init() }
func file_blxr_mev_proto_init() {
	if File_blxr_mev_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blxr_mev_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxHashListRequest); i {
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
		file_blxr_mev_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortIDListReply); i {
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
		file_blxr_mev_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidTrace); i {
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
		file_blxr_mev_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Withdrawal); i {
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
		file_blxr_mev_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionPayload); i {
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
		file_blxr_mev_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompressTx); i {
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
		file_blxr_mev_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitBlockRequest); i {
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
		file_blxr_mev_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitBlockResponse); i {
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
			RawDescriptor: file_blxr_mev_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_blxr_mev_proto_goTypes,
		DependencyIndexes: file_blxr_mev_proto_depIdxs,
		MessageInfos:      file_blxr_mev_proto_msgTypes,
	}.Build()
	File_blxr_mev_proto = out.File
	file_blxr_mev_proto_rawDesc = nil
	file_blxr_mev_proto_goTypes = nil
	file_blxr_mev_proto_depIdxs = nil
}
