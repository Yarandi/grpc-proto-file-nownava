// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1-devel
// 	protoc        v5.26.1
// source: voteupdater.proto

package voteupdater

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

type VoteupdaterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songid      string `protobuf:"bytes,1,opt,name=songid,proto3" json:"songid,omitempty"`
	Stationid   string `protobuf:"bytes,2,opt,name=stationid,proto3" json:"stationid,omitempty"`
	Stationname string `protobuf:"bytes,3,opt,name=stationname,proto3" json:"stationname,omitempty"`
	Text        string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Artist      string `protobuf:"bytes,5,opt,name=artist,proto3" json:"artist,omitempty"`
	Vote        string `protobuf:"bytes,6,opt,name=vote,proto3" json:"vote,omitempty"`
	Token       string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	Mobile      string `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *VoteupdaterRequest) Reset() {
	*x = VoteupdaterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteupdater_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteupdaterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteupdaterRequest) ProtoMessage() {}

func (x *VoteupdaterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voteupdater_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteupdaterRequest.ProtoReflect.Descriptor instead.
func (*VoteupdaterRequest) Descriptor() ([]byte, []int) {
	return file_voteupdater_proto_rawDescGZIP(), []int{0}
}

func (x *VoteupdaterRequest) GetSongid() string {
	if x != nil {
		return x.Songid
	}
	return ""
}

func (x *VoteupdaterRequest) GetStationid() string {
	if x != nil {
		return x.Stationid
	}
	return ""
}

func (x *VoteupdaterRequest) GetStationname() string {
	if x != nil {
		return x.Stationname
	}
	return ""
}

func (x *VoteupdaterRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *VoteupdaterRequest) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *VoteupdaterRequest) GetVote() string {
	if x != nil {
		return x.Vote
	}
	return ""
}

func (x *VoteupdaterRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *VoteupdaterRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type VoteupdaterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success        bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TotalVoteups   int32  `protobuf:"varint,2,opt,name=total_voteups,json=totalVoteups,proto3" json:"total_voteups,omitempty"`
	TotalVotedowns int32  `protobuf:"varint,3,opt,name=total_votedowns,json=totalVotedowns,proto3" json:"total_votedowns,omitempty"`
	Mobile         string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *VoteupdaterResponse) Reset() {
	*x = VoteupdaterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteupdater_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteupdaterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteupdaterResponse) ProtoMessage() {}

func (x *VoteupdaterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voteupdater_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteupdaterResponse.ProtoReflect.Descriptor instead.
func (*VoteupdaterResponse) Descriptor() ([]byte, []int) {
	return file_voteupdater_proto_rawDescGZIP(), []int{1}
}

func (x *VoteupdaterResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *VoteupdaterResponse) GetTotalVoteups() int32 {
	if x != nil {
		return x.TotalVoteups
	}
	return 0
}

func (x *VoteupdaterResponse) GetTotalVotedowns() int32 {
	if x != nil {
		return x.TotalVotedowns
	}
	return 0
}

func (x *VoteupdaterResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

var File_voteupdater_proto protoreflect.FileDescriptor

var file_voteupdater_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x6f, 0x74, 0x65, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x12, 0x56, 0x6f, 0x74, 0x65, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x6e, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6f,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x22, 0x95, 0x01, 0x0a, 0x13, 0x56, 0x6f, 0x74, 0x65, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x65,
	0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x56, 0x6f, 0x74, 0x65, 0x75, 0x70, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x64, 0x6f, 0x77, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x32, 0x47, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x59, 0x61, 0x72, 0x61, 0x6e, 0x64, 0x69, 0x2f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x67, 0x6f, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voteupdater_proto_rawDescOnce sync.Once
	file_voteupdater_proto_rawDescData = file_voteupdater_proto_rawDesc
)

func file_voteupdater_proto_rawDescGZIP() []byte {
	file_voteupdater_proto_rawDescOnce.Do(func() {
		file_voteupdater_proto_rawDescData = protoimpl.X.CompressGZIP(file_voteupdater_proto_rawDescData)
	})
	return file_voteupdater_proto_rawDescData
}

var file_voteupdater_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_voteupdater_proto_goTypes = []interface{}{
	(*VoteupdaterRequest)(nil),  // 0: VoteupdaterRequest
	(*VoteupdaterResponse)(nil), // 1: VoteupdaterResponse
}
var file_voteupdater_proto_depIdxs = []int32{
	0, // 0: Voteupdater.Voteupdater:input_type -> VoteupdaterRequest
	1, // 1: Voteupdater.Voteupdater:output_type -> VoteupdaterResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_voteupdater_proto_init() }
func file_voteupdater_proto_init() {
	if File_voteupdater_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voteupdater_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteupdaterRequest); i {
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
		file_voteupdater_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteupdaterResponse); i {
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
			RawDescriptor: file_voteupdater_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voteupdater_proto_goTypes,
		DependencyIndexes: file_voteupdater_proto_depIdxs,
		MessageInfos:      file_voteupdater_proto_msgTypes,
	}.Build()
	File_voteupdater_proto = out.File
	file_voteupdater_proto_rawDesc = nil
	file_voteupdater_proto_goTypes = nil
	file_voteupdater_proto_depIdxs = nil
}
