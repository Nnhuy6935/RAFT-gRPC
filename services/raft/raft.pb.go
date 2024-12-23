// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: services/raft/raft.proto

package raft

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

// Các thông điệp
type RequestVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId int32 `protobuf:"varint,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	mi := &file_services_raft_raft_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVoteRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteRequest) GetCandidateId() int32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	mi := &file_services_raft_raft_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{1}
}

func (x *RequestVoteResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type AppendEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term     int32    `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId int32    `protobuf:"varint,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	Entries  []string `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"` // log entries
}

func (x *AppendEntriesRequest) Reset() {
	*x = AppendEntriesRequest{}
	mi := &file_services_raft_raft_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppendEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesRequest) ProtoMessage() {}

func (x *AppendEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesRequest.ProtoReflect.Descriptor instead.
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{2}
}

func (x *AppendEntriesRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesRequest) GetLeaderId() int32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *AppendEntriesRequest) GetEntries() []string {
	if x != nil {
		return x.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AppendEntriesResponse) Reset() {
	*x = AppendEntriesResponse{}
	mi := &file_services_raft_raft_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppendEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesResponse) ProtoMessage() {}

func (x *AppendEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesResponse.ProtoReflect.Descriptor instead.
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{3}
}

func (x *AppendEntriesResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// đăng ký một node mới vào mạng
type RegisterNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId  int32  `protobuf:"varint,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // Địa chỉ của node
}

func (x *RegisterNodeRequest) Reset() {
	*x = RegisterNodeRequest{}
	mi := &file_services_raft_raft_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeRequest) ProtoMessage() {}

func (x *RegisterNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeRequest.ProtoReflect.Descriptor instead.
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterNodeRequest) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *RegisterNodeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegisterNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RegisterNodeResponse) Reset() {
	*x = RegisterNodeResponse{}
	mi := &file_services_raft_raft_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterNodeResponse) ProtoMessage() {}

func (x *RegisterNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterNodeResponse.ProtoReflect.Descriptor instead.
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterNodeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateId int32 `protobuf:"varint,1,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	Term        int32 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	mi := &file_services_raft_raft_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{6}
}

func (x *VoteRequest) GetCandidateId() int32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *VoteRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoteGranted bool  `protobuf:"varint,1,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
	Term        int32 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	mi := &file_services_raft_raft_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{7}
}

func (x *VoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

func (x *VoteResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

// Thêm các thông điệp cho chức năng requestLeader
type RequestLeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term     int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`         // Thời gian hiện tại của node
	LeaderId int32 `protobuf:"varint,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"` // id của leader
	Value    int32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RequestLeaderRequest) Reset() {
	*x = RequestLeaderRequest{}
	mi := &file_services_raft_raft_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestLeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLeaderRequest) ProtoMessage() {}

func (x *RequestLeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLeaderRequest.ProtoReflect.Descriptor instead.
func (*RequestLeaderRequest) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{8}
}

func (x *RequestLeaderRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestLeaderRequest) GetLeaderId() int32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *RequestLeaderRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type RequestLeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term     int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`         // Thời gian hiện tại của leader
	LeaderId int32 `protobuf:"varint,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"` // ID của leader
}

func (x *RequestLeaderResponse) Reset() {
	*x = RequestLeaderResponse{}
	mi := &file_services_raft_raft_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLeaderResponse) ProtoMessage() {}

func (x *RequestLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLeaderResponse.ProtoReflect.Descriptor instead.
func (*RequestLeaderResponse) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{9}
}

func (x *RequestLeaderResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestLeaderResponse) GetLeaderId() int32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

// Danh sách các node
type NodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"` // Danh sách các node
}

func (x *NodeList) Reset() {
	*x = NodeList{}
	mi := &file_services_raft_raft_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeList) ProtoMessage() {}

func (x *NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeList.ProtoReflect.Descriptor instead.
func (*NodeList) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{10}
}

func (x *NodeList) GetNodes() []*NodeInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// Thông điệp trống
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_services_raft_raft_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{11}
}

// Thông tin về node
type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          // ID của node
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // Địa chỉ IP và cổng của node
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	mi := &file_services_raft_raft_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{12}
}

func (x *NodeInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Phản hồi từ server
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Trạng thái thành công
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Thông điệp chi tiết
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_services_raft_raft_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{13}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// thông tin về giá trị request tới leader
type RequestValueToLeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeRequestId int32 `protobuf:"varint,1,opt,name=nodeRequestId,proto3" json:"nodeRequestId,omitempty"`
	RequestValue  int32 `protobuf:"varint,2,opt,name=requestValue,proto3" json:"requestValue,omitempty"`
	Term          int32 `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *RequestValueToLeaderRequest) Reset() {
	*x = RequestValueToLeaderRequest{}
	mi := &file_services_raft_raft_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestValueToLeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestValueToLeaderRequest) ProtoMessage() {}

func (x *RequestValueToLeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestValueToLeaderRequest.ProtoReflect.Descriptor instead.
func (*RequestValueToLeaderRequest) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{14}
}

func (x *RequestValueToLeaderRequest) GetNodeRequestId() int32 {
	if x != nil {
		return x.NodeRequestId
	}
	return 0
}

func (x *RequestValueToLeaderRequest) GetRequestValue() int32 {
	if x != nil {
		return x.RequestValue
	}
	return 0
}

func (x *RequestValueToLeaderRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

// thông tin về request leader gửi tới các node để change value
type RequestVoteChangeValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderId     int32 `protobuf:"varint,1,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	Term         int32 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	RequestValue int32 `protobuf:"varint,3,opt,name=requestValue,proto3" json:"requestValue,omitempty"`
}

func (x *RequestVoteChangeValueRequest) Reset() {
	*x = RequestVoteChangeValueRequest{}
	mi := &file_services_raft_raft_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVoteChangeValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteChangeValueRequest) ProtoMessage() {}

func (x *RequestVoteChangeValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteChangeValueRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteChangeValueRequest) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{15}
}

func (x *RequestVoteChangeValueRequest) GetLeaderId() int32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *RequestVoteChangeValueRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteChangeValueRequest) GetRequestValue() int32 {
	if x != nil {
		return x.RequestValue
	}
	return 0
}

// phản hồi các node về việc change value
type RequestVoteChangeValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderId    int32 `protobuf:"varint,3,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	Term        int32 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,1,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *RequestVoteChangeValueResponse) Reset() {
	*x = RequestVoteChangeValueResponse{}
	mi := &file_services_raft_raft_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVoteChangeValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteChangeValueResponse) ProtoMessage() {}

func (x *RequestVoteChangeValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_raft_raft_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteChangeValueResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteChangeValueResponse) Descriptor() ([]byte, []int) {
	return file_services_raft_raft_proto_rawDescGZIP(), []int{16}
}

func (x *RequestVoteChangeValueResponse) GetLeaderId() int32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *RequestVoteChangeValueResponse) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteChangeValueResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

var File_services_raft_raft_proto protoreflect.FileDescriptor

var file_services_raft_raft_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f,
	0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x66, 0x74,
	0x22, 0x4a, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f,
	0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0x60, 0x0a, 0x14, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x15, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x47, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x43, 0x0a,
	0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x22, 0x44, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x5c, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x47, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x30, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x7b, 0x0a, 0x1b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x6f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x73, 0x0a,
	0x1d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x72, 0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x32, 0xea, 0x03, 0x0a, 0x04, 0x52, 0x61, 0x66, 0x74, 0x12,
	0x42, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x61, 0x66,
	0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x6f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54,
	0x6f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63,
	0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x72, 0x61, 0x66, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x72, 0x61, 0x66, 0x74, 0x3b, 0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_raft_raft_proto_rawDescOnce sync.Once
	file_services_raft_raft_proto_rawDescData = file_services_raft_raft_proto_rawDesc
)

func file_services_raft_raft_proto_rawDescGZIP() []byte {
	file_services_raft_raft_proto_rawDescOnce.Do(func() {
		file_services_raft_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_raft_raft_proto_rawDescData)
	})
	return file_services_raft_raft_proto_rawDescData
}

var file_services_raft_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_services_raft_raft_proto_goTypes = []any{
	(*RequestVoteRequest)(nil),             // 0: raft.RequestVoteRequest
	(*RequestVoteResponse)(nil),            // 1: raft.RequestVoteResponse
	(*AppendEntriesRequest)(nil),           // 2: raft.AppendEntriesRequest
	(*AppendEntriesResponse)(nil),          // 3: raft.AppendEntriesResponse
	(*RegisterNodeRequest)(nil),            // 4: raft.RegisterNodeRequest
	(*RegisterNodeResponse)(nil),           // 5: raft.RegisterNodeResponse
	(*VoteRequest)(nil),                    // 6: raft.VoteRequest
	(*VoteResponse)(nil),                   // 7: raft.VoteResponse
	(*RequestLeaderRequest)(nil),           // 8: raft.RequestLeaderRequest
	(*RequestLeaderResponse)(nil),          // 9: raft.RequestLeaderResponse
	(*NodeList)(nil),                       // 10: raft.NodeList
	(*Empty)(nil),                          // 11: raft.Empty
	(*NodeInfo)(nil),                       // 12: raft.NodeInfo
	(*Response)(nil),                       // 13: raft.Response
	(*RequestValueToLeaderRequest)(nil),    // 14: raft.RequestValueToLeaderRequest
	(*RequestVoteChangeValueRequest)(nil),  // 15: raft.RequestVoteChangeValueRequest
	(*RequestVoteChangeValueResponse)(nil), // 16: raft.RequestVoteChangeValueResponse
}
var file_services_raft_raft_proto_depIdxs = []int32{
	12, // 0: raft.NodeList.nodes:type_name -> raft.NodeInfo
	0,  // 1: raft.Raft.RequestVote:input_type -> raft.RequestVoteRequest
	2,  // 2: raft.Raft.AppendEntries:input_type -> raft.AppendEntriesRequest
	8,  // 3: raft.Raft.RequestLeader:input_type -> raft.RequestLeaderRequest
	12, // 4: raft.Raft.RegisterNode:input_type -> raft.NodeInfo
	11, // 5: raft.Raft.GetNodeList:input_type -> raft.Empty
	14, // 6: raft.Raft.RequestValueToLeader:input_type -> raft.RequestValueToLeaderRequest
	15, // 7: raft.Raft.RequestVoteChangeValue:input_type -> raft.RequestVoteChangeValueRequest
	1,  // 8: raft.Raft.RequestVote:output_type -> raft.RequestVoteResponse
	3,  // 9: raft.Raft.AppendEntries:output_type -> raft.AppendEntriesResponse
	9,  // 10: raft.Raft.RequestLeader:output_type -> raft.RequestLeaderResponse
	13, // 11: raft.Raft.RegisterNode:output_type -> raft.Response
	10, // 12: raft.Raft.GetNodeList:output_type -> raft.NodeList
	13, // 13: raft.Raft.RequestValueToLeader:output_type -> raft.Response
	16, // 14: raft.Raft.RequestVoteChangeValue:output_type -> raft.RequestVoteChangeValueResponse
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_services_raft_raft_proto_init() }
func file_services_raft_raft_proto_init() {
	if File_services_raft_raft_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_raft_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_raft_raft_proto_goTypes,
		DependencyIndexes: file_services_raft_raft_proto_depIdxs,
		MessageInfos:      file_services_raft_raft_proto_msgTypes,
	}.Build()
	File_services_raft_raft_proto = out.File
	file_services_raft_raft_proto_rawDesc = nil
	file_services_raft_raft_proto_goTypes = nil
	file_services_raft_raft_proto_depIdxs = nil
}
