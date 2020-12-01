// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: github.com/hashtagchris/twirp-example/rpc/haberdasher/haberdasher.proto

package haberdasher

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

type Hat_Fabric int32

const (
	Hat_UNKNOWN       Hat_Fabric = 0
	Hat_SATIN         Hat_Fabric = 1
	Hat_COTTON        Hat_Fabric = 2
	Hat_FELT          Hat_Fabric = 3
	Hat_VINYL         Hat_Fabric = 4
	Hat_FRUIT_LEATHER Hat_Fabric = 5
)

// Enum value maps for Hat_Fabric.
var (
	Hat_Fabric_name = map[int32]string{
		0: "UNKNOWN",
		1: "SATIN",
		2: "COTTON",
		3: "FELT",
		4: "VINYL",
		5: "FRUIT_LEATHER",
	}
	Hat_Fabric_value = map[string]int32{
		"UNKNOWN":       0,
		"SATIN":         1,
		"COTTON":        2,
		"FELT":          3,
		"VINYL":         4,
		"FRUIT_LEATHER": 5,
	}
)

func (x Hat_Fabric) Enum() *Hat_Fabric {
	p := new(Hat_Fabric)
	*p = x
	return p
}

func (x Hat_Fabric) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Hat_Fabric) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_enumTypes[0].Descriptor()
}

func (Hat_Fabric) Type() protoreflect.EnumType {
	return &file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_enumTypes[0]
}

func (x Hat_Fabric) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Hat_Fabric.Descriptor instead.
func (Hat_Fabric) EnumDescriptor() ([]byte, []int) {
	return file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescGZIP(), []int{0, 0}
}

// A Hat is a piece of headwear made by a Haberdasher.
type Hat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The size of a hat should always be in inches.
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// The color of a hat will never be 'invisible', but other than
	// that, anything is fair game.
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	// The name of a hat is it's type. Like, 'bowler', or something.
	Name    string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Fabric  Hat_Fabric `protobuf:"varint,4,opt,name=fabric,proto3,enum=twitch.twirp.example.haberdasher.Hat_Fabric" json:"fabric,omitempty"`
	Festive bool       `protobuf:"varint,5,opt,name=festive,proto3" json:"festive,omitempty"`
}

func (x *Hat) Reset() {
	*x = Hat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hat) ProtoMessage() {}

func (x *Hat) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hat.ProtoReflect.Descriptor instead.
func (*Hat) Descriptor() ([]byte, []int) {
	return file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescGZIP(), []int{0}
}

func (x *Hat) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Hat) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Hat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hat) GetFabric() Hat_Fabric {
	if x != nil {
		return x.Fabric
	}
	return Hat_UNKNOWN
}

func (x *Hat) GetFestive() bool {
	if x != nil {
		return x.Festive
	}
	return false
}

// Size is passed when requesting a new hat to be made. It's always measured in
// inches.
type Size struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inches int32 `protobuf:"varint,1,opt,name=inches,proto3" json:"inches,omitempty"`
}

func (x *Size) Reset() {
	*x = Size{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size) ProtoMessage() {}

func (x *Size) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size.ProtoReflect.Descriptor instead.
func (*Size) Descriptor() ([]byte, []int) {
	return file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescGZIP(), []int{1}
}

func (x *Size) GetInches() int32 {
	if x != nil {
		return x.Inches
	}
	return 0
}

var File_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto protoreflect.FileDescriptor

var file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x74, 0x61, 0x67, 0x63, 0x68, 0x72, 0x69, 0x73, 0x2f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2d,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x61, 0x62, 0x65,
	0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61, 0x73,
	0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x74, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x68, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x22, 0xf9, 0x01, 0x0a, 0x03,
	0x48, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x44, 0x0a, 0x06, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61,
	0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x74, 0x2e, 0x46, 0x61, 0x62, 0x72, 0x69, 0x63, 0x52,
	0x06, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x65, 0x73, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x65, 0x73, 0x74, 0x69, 0x76,
	0x65, 0x22, 0x54, 0x0a, 0x06, 0x46, 0x61, 0x62, 0x72, 0x69, 0x63, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x41, 0x54, 0x49,
	0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x54, 0x54, 0x4f, 0x4e, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x45, 0x4c, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x4e,
	0x59, 0x4c, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x52, 0x55, 0x49, 0x54, 0x5f, 0x4c, 0x45,
	0x41, 0x54, 0x48, 0x45, 0x52, 0x10, 0x05, 0x22, 0x1e, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x32, 0x67, 0x0a, 0x0b, 0x48, 0x61, 0x62, 0x65, 0x72,
	0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x07, 0x4d, 0x61, 0x6b, 0x65, 0x48, 0x61,
	0x74, 0x12, 0x26, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61,
	0x73, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x25, 0x2e, 0x74, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x2e, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x74,
	0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x63, 0x68, 0x72, 0x69, 0x73, 0x2f, 0x74, 0x77, 0x69, 0x72,
	0x70, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x61,
	0x62, 0x65, 0x72, 0x64, 0x61, 0x73, 0x68, 0x65, 0x72, 0x3b, 0x68, 0x61, 0x62, 0x65, 0x72, 0x64,
	0x61, 0x73, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescOnce sync.Once
	file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescData = file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDesc
)

func file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescGZIP() []byte {
	file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescOnce.Do(func() {
		file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescData)
	})
	return file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDescData
}

var file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_goTypes = []interface{}{
	(Hat_Fabric)(0), // 0: twitch.twirp.example.haberdasher.Hat.Fabric
	(*Hat)(nil),     // 1: twitch.twirp.example.haberdasher.Hat
	(*Size)(nil),    // 2: twitch.twirp.example.haberdasher.Size
}
var file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_depIdxs = []int32{
	0, // 0: twitch.twirp.example.haberdasher.Hat.fabric:type_name -> twitch.twirp.example.haberdasher.Hat.Fabric
	2, // 1: twitch.twirp.example.haberdasher.Haberdasher.MakeHat:input_type -> twitch.twirp.example.haberdasher.Size
	1, // 2: twitch.twirp.example.haberdasher.Haberdasher.MakeHat:output_type -> twitch.twirp.example.haberdasher.Hat
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_init() }
func file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_init() {
	if File_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hat); i {
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
		file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Size); i {
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
			RawDescriptor: file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_goTypes,
		DependencyIndexes: file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_depIdxs,
		EnumInfos:         file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_enumTypes,
		MessageInfos:      file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_msgTypes,
	}.Build()
	File_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto = out.File
	file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_rawDesc = nil
	file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_goTypes = nil
	file_github_com_hashtagchris_twirp_example_rpc_haberdasher_haberdasher_proto_depIdxs = nil
}
