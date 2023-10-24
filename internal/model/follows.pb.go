// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: follows.proto

package model

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

type Follow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerId  string `protobuf:"bytes,1,opt,name=followerId,proto3" json:"followerId,omitempty"`
	FolloweeId  string `protobuf:"bytes,2,opt,name=followeeId,proto3" json:"followeeId,omitempty"`
	FollowTime  int64  `protobuf:"varint,3,opt,name=followTime,proto3" json:"followTime,omitempty"`
	Isfollowing bool   `protobuf:"varint,4,opt,name=Isfollowing,proto3" json:"Isfollowing,omitempty"`
}

func (x *Follow) Reset() {
	*x = Follow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follows_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Follow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Follow) ProtoMessage() {}

func (x *Follow) ProtoReflect() protoreflect.Message {
	mi := &file_follows_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Follow.ProtoReflect.Descriptor instead.
func (*Follow) Descriptor() ([]byte, []int) {
	return file_follows_proto_rawDescGZIP(), []int{0}
}

func (x *Follow) GetFollowerId() string {
	if x != nil {
		return x.FollowerId
	}
	return ""
}

func (x *Follow) GetFolloweeId() string {
	if x != nil {
		return x.FolloweeId
	}
	return ""
}

func (x *Follow) GetFollowTime() int64 {
	if x != nil {
		return x.FollowTime
	}
	return 0
}

func (x *Follow) GetIsfollowing() bool {
	if x != nil {
		return x.Isfollowing
	}
	return false
}

type Follower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Status          bool   `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Time            int64  `protobuf:"varint,3,opt,name=Time,proto3" json:"Time,omitempty"`
	FollowAccount   string `protobuf:"bytes,4,opt,name=FollowAccount,proto3" json:"FollowAccount,omitempty"`
	UnfollowAccount string `protobuf:"bytes,5,opt,name=unfollowAccount,proto3" json:"unfollowAccount,omitempty"`
}

func (x *Follower) Reset() {
	*x = Follower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follows_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Follower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Follower) ProtoMessage() {}

func (x *Follower) ProtoReflect() protoreflect.Message {
	mi := &file_follows_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Follower.ProtoReflect.Descriptor instead.
func (*Follower) Descriptor() ([]byte, []int) {
	return file_follows_proto_rawDescGZIP(), []int{1}
}

func (x *Follower) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Follower) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Follower) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Follower) GetFollowAccount() string {
	if x != nil {
		return x.FollowAccount
	}
	return ""
}

func (x *Follower) GetUnfollowAccount() string {
	if x != nil {
		return x.UnfollowAccount
	}
	return ""
}

type Followee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	BlockAccountLink string `protobuf:"bytes,2,opt,name=BlockAccountLink,proto3" json:"BlockAccountLink,omitempty"`
	BlockAccountName string `protobuf:"bytes,3,opt,name=BlockAccountName,proto3" json:"BlockAccountName,omitempty"`
	Time             int64  `protobuf:"varint,4,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *Followee) Reset() {
	*x = Followee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follows_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Followee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Followee) ProtoMessage() {}

func (x *Followee) ProtoReflect() protoreflect.Message {
	mi := &file_follows_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Followee.ProtoReflect.Descriptor instead.
func (*Followee) Descriptor() ([]byte, []int) {
	return file_follows_proto_rawDescGZIP(), []int{2}
}

func (x *Followee) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Followee) GetBlockAccountLink() string {
	if x != nil {
		return x.BlockAccountLink
	}
	return ""
}

func (x *Followee) GetBlockAccountName() string {
	if x != nil {
		return x.BlockAccountName
	}
	return ""
}

func (x *Followee) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_follows_proto protoreflect.FileDescriptor

var file_follows_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x65, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x73, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x22, 0x9e, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x6e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follows_proto_rawDescOnce sync.Once
	file_follows_proto_rawDescData = file_follows_proto_rawDesc
)

func file_follows_proto_rawDescGZIP() []byte {
	file_follows_proto_rawDescOnce.Do(func() {
		file_follows_proto_rawDescData = protoimpl.X.CompressGZIP(file_follows_proto_rawDescData)
	})
	return file_follows_proto_rawDescData
}

var file_follows_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_follows_proto_goTypes = []interface{}{
	(*Follow)(nil),   // 0: model.Follow
	(*Follower)(nil), // 1: model.Follower
	(*Followee)(nil), // 2: model.Followee
}
var file_follows_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_follows_proto_init() }
func file_follows_proto_init() {
	if File_follows_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follows_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Follow); i {
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
		file_follows_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Follower); i {
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
		file_follows_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Followee); i {
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
			RawDescriptor: file_follows_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_follows_proto_goTypes,
		DependencyIndexes: file_follows_proto_depIdxs,
		MessageInfos:      file_follows_proto_msgTypes,
	}.Build()
	File_follows_proto = out.File
	file_follows_proto_rawDesc = nil
	file_follows_proto_goTypes = nil
	file_follows_proto_depIdxs = nil
}
