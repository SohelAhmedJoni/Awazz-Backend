// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: notifications.proto

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

type Notifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body          string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Source        string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Image         string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Sound         string `protobuf:"bytes,5,opt,name=sound,proto3" json:"sound,omitempty"`
	Time          int64  `protobuf:"varint,6,opt,name=time,proto3" json:"time,omitempty"`
	Channel       string `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	PriorityLevel int64  `protobuf:"varint,8,opt,name=priority_level,json=priorityLevel,proto3" json:"priority_level,omitempty"`
	ReadStatus    bool   `protobuf:"varint,9,opt,name=readStatus,proto3" json:"readStatus,omitempty"`
	Created       int64  `protobuf:"varint,10,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Notifications) Reset() {
	*x = Notifications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifications) ProtoMessage() {}

func (x *Notifications) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifications.ProtoReflect.Descriptor instead.
func (*Notifications) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *Notifications) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notifications) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Notifications) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Notifications) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Notifications) GetSound() string {
	if x != nil {
		return x.Sound
	}
	return ""
}

func (x *Notifications) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Notifications) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Notifications) GetPriorityLevel() int64 {
	if x != nil {
		return x.PriorityLevel
	}
	return 0
}

func (x *Notifications) GetReadStatus() bool {
	if x != nil {
		return x.ReadStatus
	}
	return false
}

func (x *Notifications) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

var File_notifications_proto protoreflect.FileDescriptor

var file_notifications_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x8c, 0x02, 0x0a,
	0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_proto_rawDescOnce sync.Once
	file_notifications_proto_rawDescData = file_notifications_proto_rawDesc
)

func file_notifications_proto_rawDescGZIP() []byte {
	file_notifications_proto_rawDescOnce.Do(func() {
		file_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_proto_rawDescData)
	})
	return file_notifications_proto_rawDescData
}

var file_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notifications_proto_goTypes = []interface{}{
	(*Notifications)(nil), // 0: model.Notifications
}
var file_notifications_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notifications_proto_init() }
func file_notifications_proto_init() {
	if File_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notifications); i {
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
			RawDescriptor: file_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notifications_proto_goTypes,
		DependencyIndexes: file_notifications_proto_depIdxs,
		MessageInfos:      file_notifications_proto_msgTypes,
	}.Build()
	File_notifications_proto = out.File
	file_notifications_proto_rawDesc = nil
	file_notifications_proto_goTypes = nil
	file_notifications_proto_depIdxs = nil
}
