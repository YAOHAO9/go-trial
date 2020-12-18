// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: rpcMsg.proto

package message

import (
	reflect "reflect"
	sync "sync"

	"github.com/YAOHAO9/pine/rpc/session"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RPCMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string           `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	Handler   string           `protobuf:"bytes,2,opt,name=Handler,proto3" json:"Handler,omitempty"`
	Type      int32            `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	RawData   []byte           `protobuf:"bytes,4,opt,name=RawData,proto3" json:"RawData,omitempty"`
	RequestID *int32           `protobuf:"varint,5,opt,name=RequestID,proto3,oneof" json:"RequestID,omitempty"`
	Session   *session.Session `protobuf:"bytes,6,opt,name=Session,proto3" json:"Session,omitempty"`
}

func (x *RPCMsg) Reset() {
	*x = RPCMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcMsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCMsg) ProtoMessage() {}

func (x *RPCMsg) ProtoReflect() protoreflect.Message {
	mi := &file_rpcMsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCMsg.ProtoReflect.Descriptor instead.
func (*RPCMsg) Descriptor() ([]byte, []int) {
	return file_rpcMsg_proto_rawDescGZIP(), []int{0}
}

func (x *RPCMsg) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *RPCMsg) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

func (x *RPCMsg) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *RPCMsg) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

func (x *RPCMsg) GetRequestID() int32 {
	if x != nil && x.RequestID != nil {
		return *x.RequestID
	}
	return 0
}

func (x *RPCMsg) GetSession() *session.Session {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_rpcMsg_proto protoreflect.FileDescriptor

var file_rpcMsg_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x70, 0x63, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x06, 0x52, 0x50, 0x43, 0x4d, 0x73,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x2a, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcMsg_proto_rawDescOnce sync.Once
	file_rpcMsg_proto_rawDescData = file_rpcMsg_proto_rawDesc
)

func file_rpcMsg_proto_rawDescGZIP() []byte {
	file_rpcMsg_proto_rawDescOnce.Do(func() {
		file_rpcMsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcMsg_proto_rawDescData)
	})
	return file_rpcMsg_proto_rawDescData
}

var file_rpcMsg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpcMsg_proto_goTypes = []interface{}{
	(*RPCMsg)(nil),          // 0: message.RPCMsg
	(*session.Session)(nil), // 1: message.Session
}
var file_rpcMsg_proto_depIdxs = []int32{
	1, // 0: message.RPCMsg.Session:type_name -> message.Session
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpcMsg_proto_init() }
func file_rpcMsg_proto_init() {
	if File_rpcMsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcMsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCMsg); i {
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
	file_rpcMsg_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpcMsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpcMsg_proto_goTypes,
		DependencyIndexes: file_rpcMsg_proto_depIdxs,
		MessageInfos:      file_rpcMsg_proto_msgTypes,
	}.Build()
	File_rpcMsg_proto = out.File
	file_rpcMsg_proto_rawDesc = nil
	file_rpcMsg_proto_goTypes = nil
	file_rpcMsg_proto_depIdxs = nil
}
