// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: double_spend.proto

package network_pb

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

type DoubleSpendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
}

func (x *DoubleSpendRequest) Reset() {
	*x = DoubleSpendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_double_spend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleSpendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleSpendRequest) ProtoMessage() {}

func (x *DoubleSpendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_double_spend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleSpendRequest.ProtoReflect.Descriptor instead.
func (*DoubleSpendRequest) Descriptor() ([]byte, []int) {
	return file_double_spend_proto_rawDescGZIP(), []int{0}
}

func (x *DoubleSpendRequest) GetStart() []byte {
	if x != nil {
		return x.Start
	}
	return nil
}

type DoubleSpendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs []*DoubleSpend `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *DoubleSpendResponse) Reset() {
	*x = DoubleSpendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_double_spend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleSpendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleSpendResponse) ProtoMessage() {}

func (x *DoubleSpendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_double_spend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleSpendResponse.ProtoReflect.Descriptor instead.
func (*DoubleSpendResponse) Descriptor() ([]byte, []int) {
	return file_double_spend_proto_rawDescGZIP(), []int{1}
}

func (x *DoubleSpendResponse) GetTxs() []*DoubleSpend {
	if x != nil {
		return x.Txs
	}
	return nil
}

type DoubleSpend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *DoubleSpend) Reset() {
	*x = DoubleSpend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_double_spend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleSpend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleSpend) ProtoMessage() {}

func (x *DoubleSpend) ProtoReflect() protoreflect.Message {
	mi := &file_double_spend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleSpend.ProtoReflect.Descriptor instead.
func (*DoubleSpend) Descriptor() ([]byte, []int) {
	return file_double_spend_proto_rawDescGZIP(), []int{2}
}

func (x *DoubleSpend) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

var File_double_spend_proto protoreflect.FileDescriptor

var file_double_spend_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x62,
	0x22, 0x2a, 0x0a, 0x12, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x13,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x62, 0x2e, 0x44, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0x1d,
	0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x78, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x6d, 0x6f,
	0x63, 0x61, 0x73, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x66, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_double_spend_proto_rawDescOnce sync.Once
	file_double_spend_proto_rawDescData = file_double_spend_proto_rawDesc
)

func file_double_spend_proto_rawDescGZIP() []byte {
	file_double_spend_proto_rawDescOnce.Do(func() {
		file_double_spend_proto_rawDescData = protoimpl.X.CompressGZIP(file_double_spend_proto_rawDescData)
	})
	return file_double_spend_proto_rawDescData
}

var file_double_spend_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_double_spend_proto_goTypes = []interface{}{
	(*DoubleSpendRequest)(nil),  // 0: network_pb.DoubleSpendRequest
	(*DoubleSpendResponse)(nil), // 1: network_pb.DoubleSpendResponse
	(*DoubleSpend)(nil),         // 2: network_pb.DoubleSpend
}
var file_double_spend_proto_depIdxs = []int32{
	2, // 0: network_pb.DoubleSpendResponse.txs:type_name -> network_pb.DoubleSpend
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_double_spend_proto_init() }
func file_double_spend_proto_init() {
	if File_double_spend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_double_spend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleSpendRequest); i {
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
		file_double_spend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleSpendResponse); i {
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
		file_double_spend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleSpend); i {
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
			RawDescriptor: file_double_spend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_double_spend_proto_goTypes,
		DependencyIndexes: file_double_spend_proto_depIdxs,
		MessageInfos:      file_double_spend_proto_msgTypes,
	}.Build()
	File_double_spend_proto = out.File
	file_double_spend_proto_rawDesc = nil
	file_double_spend_proto_goTypes = nil
	file_double_spend_proto_depIdxs = nil
}
