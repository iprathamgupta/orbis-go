// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: orbis/transport/v1alpha1/transport.proto

package transportv1alpha1

import (
	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transport string `protobuf:"bytes,1,opt,name=transport,proto3" json:"transport,omitempty"`
}

func (x *GetPeersRequest) Reset() {
	*x = GetPeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeersRequest) ProtoMessage() {}

func (x *GetPeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeersRequest.ProtoReflect.Descriptor instead.
func (*GetPeersRequest) Descriptor() ([]byte, []int) {
	return file_orbis_transport_v1alpha1_transport_proto_rawDescGZIP(), []int{0}
}

func (x *GetPeersRequest) GetTransport() string {
	if x != nil {
		return x.Transport
	}
	return ""
}

type GetPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *GetPeersResponse) Reset() {
	*x = GetPeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeersResponse) ProtoMessage() {}

func (x *GetPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeersResponse.ProtoReflect.Descriptor instead.
func (*GetPeersResponse) Descriptor() ([]byte, []int) {
	return file_orbis_transport_v1alpha1_transport_proto_rawDescGZIP(), []int{1}
}

func (x *GetPeersResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type GetHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transport string `protobuf:"bytes,1,opt,name=transport,proto3" json:"transport,omitempty"`
}

func (x *GetHostRequest) Reset() {
	*x = GetHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostRequest) ProtoMessage() {}

func (x *GetHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostRequest.ProtoReflect.Descriptor instead.
func (*GetHostRequest) Descriptor() ([]byte, []int) {
	return file_orbis_transport_v1alpha1_transport_proto_rawDescGZIP(), []int{2}
}

func (x *GetHostRequest) GetTransport() string {
	if x != nil {
		return x.Transport
	}
	return ""
}

type GetHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GetHostResponse) Reset() {
	*x = GetHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostResponse) ProtoMessage() {}

func (x *GetHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostResponse.ProtoReflect.Descriptor instead.
func (*GetHostResponse) Descriptor() ([]byte, []int) {
	return file_orbis_transport_v1alpha1_transport_proto_rawDescGZIP(), []int{3}
}

func (x *GetHostResponse) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // multiaddress
	PublicKey *pb.PublicKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_orbis_transport_v1alpha1_transport_proto_rawDescGZIP(), []int{4}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetPublicKey() *pb.PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp    int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                             // unix time
	Id           string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                            // message id
	Type         string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`                                        // message type
	Payload      []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`                                  // generic payload
	Gossip       bool   `protobuf:"varint,5,opt,name=gossip,proto3" json:"gossip,omitempty"`                                   // gossip message over pubsub
	NodeId       string `protobuf:"bytes,6,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`                      // author node id (peer.ID)
	NodePubKey   []byte `protobuf:"bytes,7,opt,name=node_pub_key,json=nodePubKey,proto3" json:"node_pub_key,omitempty"`        // authoring node pubkey (32bytes)
	RingId       string `protobuf:"bytes,8,opt,name=ring_id,json=ringId,proto3" json:"ring_id,omitempty"`                      //
	Signature    []byte `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`                              // signature of the message (including payload)
	TargetId     string `protobuf:"bytes,10,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`               // id of target
	TargetPubKey []byte `protobuf:"bytes,11,opt,name=target_pub_key,json=targetPubKey,proto3" json:"target_pub_key,omitempty"` // pubkey of target
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_transport_v1alpha1_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_orbis_transport_v1alpha1_transport_proto_rawDescGZIP(), []int{5}
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Message) GetGossip() bool {
	if x != nil {
		return x.Gossip
	}
	return false
}

func (x *Message) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Message) GetNodePubKey() []byte {
	if x != nil {
		return x.NodePubKey
	}
	return nil
}

func (x *Message) GetRingId() string {
	if x != nil {
		return x.RingId
	}
	return ""
}

func (x *Message) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Message) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Message) GetTargetPubKey() []byte {
	if x != nil {
		return x.TargetPubKey
	}
	return nil
}

var File_orbis_transport_v1alpha1_transport_proto protoreflect.FileDescriptor

var file_orbis_transport_v1alpha1_transport_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x72, 0x62, 0x69,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x45, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x22, 0x6c, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x70,
	0x32, 0x70, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x22, 0xb2, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x67, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x32, 0xb4, 0x02, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x7b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x7d, 0x12, 0x90, 0x01, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x29, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x7d, 0x42, 0x88, 0x02,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x73,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72,
	0x62, 0x69, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x54, 0x58, 0xaa, 0x02,
	0x18, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x18, 0x4f, 0x72, 0x62, 0x69,
	0x73, 0x5c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x24, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x5c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x4f, 0x72,
	0x62, 0x69, 0x73, 0x3a, 0x3a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orbis_transport_v1alpha1_transport_proto_rawDescOnce sync.Once
	file_orbis_transport_v1alpha1_transport_proto_rawDescData = file_orbis_transport_v1alpha1_transport_proto_rawDesc
)

func file_orbis_transport_v1alpha1_transport_proto_rawDescGZIP() []byte {
	file_orbis_transport_v1alpha1_transport_proto_rawDescOnce.Do(func() {
		file_orbis_transport_v1alpha1_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_orbis_transport_v1alpha1_transport_proto_rawDescData)
	})
	return file_orbis_transport_v1alpha1_transport_proto_rawDescData
}

var file_orbis_transport_v1alpha1_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orbis_transport_v1alpha1_transport_proto_goTypes = []interface{}{
	(*GetPeersRequest)(nil),  // 0: orbis.transport.v1alpha1.GetPeersRequest
	(*GetPeersResponse)(nil), // 1: orbis.transport.v1alpha1.GetPeersResponse
	(*GetHostRequest)(nil),   // 2: orbis.transport.v1alpha1.GetHostRequest
	(*GetHostResponse)(nil),  // 3: orbis.transport.v1alpha1.GetHostResponse
	(*Node)(nil),             // 4: orbis.transport.v1alpha1.Node
	(*Message)(nil),          // 5: orbis.transport.v1alpha1.Message
	(*pb.PublicKey)(nil),     // 6: libp2p.crypto.v1.PublicKey
}
var file_orbis_transport_v1alpha1_transport_proto_depIdxs = []int32{
	4, // 0: orbis.transport.v1alpha1.GetPeersResponse.nodes:type_name -> orbis.transport.v1alpha1.Node
	4, // 1: orbis.transport.v1alpha1.GetHostResponse.node:type_name -> orbis.transport.v1alpha1.Node
	6, // 2: orbis.transport.v1alpha1.Node.public_key:type_name -> libp2p.crypto.v1.PublicKey
	2, // 3: orbis.transport.v1alpha1.TransportService.GetHost:input_type -> orbis.transport.v1alpha1.GetHostRequest
	0, // 4: orbis.transport.v1alpha1.TransportService.GetPeers:input_type -> orbis.transport.v1alpha1.GetPeersRequest
	3, // 5: orbis.transport.v1alpha1.TransportService.GetHost:output_type -> orbis.transport.v1alpha1.GetHostResponse
	1, // 6: orbis.transport.v1alpha1.TransportService.GetPeers:output_type -> orbis.transport.v1alpha1.GetPeersResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_orbis_transport_v1alpha1_transport_proto_init() }
func file_orbis_transport_v1alpha1_transport_proto_init() {
	if File_orbis_transport_v1alpha1_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orbis_transport_v1alpha1_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeersRequest); i {
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
		file_orbis_transport_v1alpha1_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeersResponse); i {
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
		file_orbis_transport_v1alpha1_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostRequest); i {
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
		file_orbis_transport_v1alpha1_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostResponse); i {
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
		file_orbis_transport_v1alpha1_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_orbis_transport_v1alpha1_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_orbis_transport_v1alpha1_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orbis_transport_v1alpha1_transport_proto_goTypes,
		DependencyIndexes: file_orbis_transport_v1alpha1_transport_proto_depIdxs,
		MessageInfos:      file_orbis_transport_v1alpha1_transport_proto_msgTypes,
	}.Build()
	File_orbis_transport_v1alpha1_transport_proto = out.File
	file_orbis_transport_v1alpha1_transport_proto_rawDesc = nil
	file_orbis_transport_v1alpha1_transport_proto_goTypes = nil
	file_orbis_transport_v1alpha1_transport_proto_depIdxs = nil
}
