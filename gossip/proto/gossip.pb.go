// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: gossip/proto/gossip.proto

package proto

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

type GossipMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId      uint32 `protobuf:"varint,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	SeqNumber   uint64 `protobuf:"varint,2,opt,name=seq_number,json=seqNumber,proto3" json:"seq_number,omitempty"`
	SeqRollover bool   `protobuf:"varint,3,opt,name=seq_rollover,json=seqRollover,proto3" json:"seq_rollover,omitempty"`
	Payload     []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Ttl         uint32 `protobuf:"varint,5,opt,name=ttl,proto3" json:"ttl,omitempty"`
	SeenBy      uint32 `protobuf:"varint,6,opt,name=seen_by,json=seenBy,proto3" json:"seen_by,omitempty"`
}

func (x *GossipMessage) Reset() {
	*x = GossipMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GossipMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GossipMessage) ProtoMessage() {}

func (x *GossipMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_gossip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GossipMessage.ProtoReflect.Descriptor instead.
func (*GossipMessage) Descriptor() ([]byte, []int) {
	return file_gossip_proto_gossip_proto_rawDescGZIP(), []int{0}
}

func (x *GossipMessage) GetPeerId() uint32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *GossipMessage) GetSeqNumber() uint64 {
	if x != nil {
		return x.SeqNumber
	}
	return 0
}

func (x *GossipMessage) GetSeqRollover() bool {
	if x != nil {
		return x.SeqRollover
	}
	return false
}

func (x *GossipMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *GossipMessage) GetTtl() uint32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *GossipMessage) GetSeenBy() uint32 {
	if x != nil {
		return x.SeenBy
	}
	return 0
}

var File_gossip_proto_gossip_proto protoreflect.FileDescriptor

var file_gossip_proto_gossip_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x0d,
	0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x71, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x71, 0x5f, 0x72, 0x6f, 0x6c,
	0x6c, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x65, 0x71,
	0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x74, 0x74, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x65, 0x65, 0x6e, 0x42, 0x79, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x78, 0x70,
	0x6f, 0x6c, 0x65, 0x74, 0x61, 0x65, 0x76, 0x2f, 0x6b, 0x76, 0x2f, 0x67, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gossip_proto_gossip_proto_rawDescOnce sync.Once
	file_gossip_proto_gossip_proto_rawDescData = file_gossip_proto_gossip_proto_rawDesc
)

func file_gossip_proto_gossip_proto_rawDescGZIP() []byte {
	file_gossip_proto_gossip_proto_rawDescOnce.Do(func() {
		file_gossip_proto_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_gossip_proto_gossip_proto_rawDescData)
	})
	return file_gossip_proto_gossip_proto_rawDescData
}

var file_gossip_proto_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gossip_proto_gossip_proto_goTypes = []interface{}{
	(*GossipMessage)(nil), // 0: GossipMessage
}
var file_gossip_proto_gossip_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gossip_proto_gossip_proto_init() }
func file_gossip_proto_gossip_proto_init() {
	if File_gossip_proto_gossip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gossip_proto_gossip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GossipMessage); i {
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
			RawDescriptor: file_gossip_proto_gossip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gossip_proto_gossip_proto_goTypes,
		DependencyIndexes: file_gossip_proto_gossip_proto_depIdxs,
		MessageInfos:      file_gossip_proto_gossip_proto_msgTypes,
	}.Build()
	File_gossip_proto_gossip_proto = out.File
	file_gossip_proto_gossip_proto_rawDesc = nil
	file_gossip_proto_gossip_proto_goTypes = nil
	file_gossip_proto_gossip_proto_depIdxs = nil
}
