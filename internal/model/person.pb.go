// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: person.proto

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Attachment        string   `protobuf:"bytes,4,opt,name=attachment,proto3" json:"attachment,omitempty"`
	AttributedTo      string   `protobuf:"bytes,5,opt,name=attributedTo,proto3" json:"attributedTo,omitempty"`
	Context           string   `protobuf:"bytes,8,opt,name=context,proto3" json:"context,omitempty"`
	MediaType         string   `protobuf:"bytes,9,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	EndTime           int64    `protobuf:"varint,10,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Generator         string   `protobuf:"bytes,11,opt,name=generator,proto3" json:"generator,omitempty"`
	Icon              string   `protobuf:"bytes,12,opt,name=icon,proto3" json:"icon,omitempty"`
	Image             string   `protobuf:"bytes,13,opt,name=image,proto3" json:"image,omitempty"`
	InReplyTo         string   `protobuf:"bytes,14,opt,name=inReplyTo,proto3" json:"inReplyTo,omitempty"`
	Location          string   `protobuf:"bytes,15,opt,name=location,proto3" json:"location,omitempty"`
	Preview           string   `protobuf:"bytes,16,opt,name=preview,proto3" json:"preview,omitempty"`
	PublishedTime     int64    `protobuf:"varint,17,opt,name=publishedTime,proto3" json:"publishedTime,omitempty"`
	Replies           []string `protobuf:"bytes,18,rep,name=replies,proto3" json:"replies,omitempty"`
	StartTime         int64    `protobuf:"varint,19,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Summary           string   `protobuf:"bytes,20,opt,name=summary,proto3" json:"summary,omitempty"`
	Tag               []string `protobuf:"bytes,21,rep,name=tag,proto3" json:"tag,omitempty"`
	UpdatedTime       int64    `protobuf:"varint,22,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	Url               []string `protobuf:"bytes,23,rep,name=url,proto3" json:"url,omitempty"`
	Too               []string `protobuf:"bytes,24,rep,name=too,proto3" json:"too,omitempty"`
	Bto               []string `protobuf:"bytes,25,rep,name=bto,proto3" json:"bto,omitempty"`
	Cc                []string `protobuf:"bytes,26,rep,name=cc,proto3" json:"cc,omitempty"`
	Bcc               []string `protobuf:"bytes,27,rep,name=bcc,proto3" json:"bcc,omitempty"`
	Likes             string   `protobuf:"bytes,29,opt,name=likes,proto3" json:"likes,omitempty"`
	Shares            string   `protobuf:"bytes,30,opt,name=shares,proto3" json:"shares,omitempty"`
	Inbox             string   `protobuf:"bytes,32,opt,name=inbox,proto3" json:"inbox,omitempty"`
	Outbox            string   `protobuf:"bytes,33,opt,name=outbox,proto3" json:"outbox,omitempty"`
	Following         []string `protobuf:"bytes,34,rep,name=following,proto3" json:"following,omitempty"`
	Followers         []string `protobuf:"bytes,35,rep,name=followers,proto3" json:"followers,omitempty"`
	Liked             []string `protobuf:"bytes,36,rep,name=liked,proto3" json:"liked,omitempty"`
	PreferredUsername string   `protobuf:"bytes,37,opt,name=preferredUsername,proto3" json:"preferredUsername,omitempty"`
	Endpoints         string   `protobuf:"bytes,38,opt,name=endpoints,proto3" json:"endpoints,omitempty"` //no need
	Streams           string   `protobuf:"bytes,39,opt,name=streams,proto3" json:"streams,omitempty"`
	PublicKey         string   `protobuf:"bytes,40,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	FragmentationKey  string   `protobuf:"bytes,41,opt,name=fragmentationKey,proto3" json:"fragmentationKey,omitempty"`
	Username          string   `protobuf:"bytes,42,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetAttachment() string {
	if x != nil {
		return x.Attachment
	}
	return ""
}

func (x *Person) GetAttributedTo() string {
	if x != nil {
		return x.AttributedTo
	}
	return ""
}

func (x *Person) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *Person) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *Person) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Person) GetGenerator() string {
	if x != nil {
		return x.Generator
	}
	return ""
}

func (x *Person) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Person) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Person) GetInReplyTo() string {
	if x != nil {
		return x.InReplyTo
	}
	return ""
}

func (x *Person) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Person) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *Person) GetPublishedTime() int64 {
	if x != nil {
		return x.PublishedTime
	}
	return 0
}

func (x *Person) GetReplies() []string {
	if x != nil {
		return x.Replies
	}
	return nil
}

func (x *Person) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Person) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Person) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *Person) GetUpdatedTime() int64 {
	if x != nil {
		return x.UpdatedTime
	}
	return 0
}

func (x *Person) GetUrl() []string {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *Person) GetToo() []string {
	if x != nil {
		return x.Too
	}
	return nil
}

func (x *Person) GetBto() []string {
	if x != nil {
		return x.Bto
	}
	return nil
}

func (x *Person) GetCc() []string {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *Person) GetBcc() []string {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *Person) GetLikes() string {
	if x != nil {
		return x.Likes
	}
	return ""
}

func (x *Person) GetShares() string {
	if x != nil {
		return x.Shares
	}
	return ""
}

func (x *Person) GetInbox() string {
	if x != nil {
		return x.Inbox
	}
	return ""
}

func (x *Person) GetOutbox() string {
	if x != nil {
		return x.Outbox
	}
	return ""
}

func (x *Person) GetFollowing() []string {
	if x != nil {
		return x.Following
	}
	return nil
}

func (x *Person) GetFollowers() []string {
	if x != nil {
		return x.Followers
	}
	return nil
}

func (x *Person) GetLiked() []string {
	if x != nil {
		return x.Liked
	}
	return nil
}

func (x *Person) GetPreferredUsername() string {
	if x != nil {
		return x.PreferredUsername
	}
	return ""
}

func (x *Person) GetEndpoints() string {
	if x != nil {
		return x.Endpoints
	}
	return ""
}

func (x *Person) GetStreams() string {
	if x != nil {
		return x.Streams
	}
	return ""
}

func (x *Person) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Person) GetFragmentationKey() string {
	if x != nil {
		return x.FragmentationKey
	}
	return ""
}

func (x *Person) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person       *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	Password     string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email        string  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"` //optional
	Phone        string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"` //optional
	Created_Time int64   `protobuf:"varint,5,opt,name=created_Time,json=createdTime,proto3" json:"created_Time,omitempty"`
	Updated_Time int64   `protobuf:"varint,6,opt,name=updated_Time,json=updatedTime,proto3" json:"updated_Time,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Account) GetCreated_Time() int64 {
	if x != nil {
		return x.Created_Time
	}
	return 0
}

func (x *Account) GetUpdated_Time() int64 {
	if x != nil {
		return x.Updated_Time
	}
	return 0
}

type UserPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon             string `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	Url              string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Reactions        string `protobuf:"bytes,7,opt,name=reactions,proto3" json:"reactions,omitempty"`
	PublicKey        string `protobuf:"bytes,13,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	FragmentationKey string `protobuf:"bytes,14,opt,name=fragmentationKey,proto3" json:"fragmentationKey,omitempty"`
}

func (x *UserPreview) Reset() {
	*x = UserPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPreview) ProtoMessage() {}

func (x *UserPreview) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPreview.ProtoReflect.Descriptor instead.
func (*UserPreview) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{2}
}

func (x *UserPreview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPreview) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserPreview) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UserPreview) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UserPreview) GetReactions() string {
	if x != nil {
		return x.Reactions
	}
	return ""
}

func (x *UserPreview) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *UserPreview) GetFragmentationKey() string {
	if x != nil {
		return x.FragmentationKey
	}
	return ""
}

// collection of Persons
type Persons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=Persons,proto3" json:"Persons,omitempty"`
}

func (x *Persons) Reset() {
	*x = Persons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Persons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Persons) ProtoMessage() {}

func (x *Persons) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Persons.ProtoReflect.Descriptor instead.
func (*Persons) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{3}
}

func (x *Persons) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

// collection of Accounts
type Accounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *Accounts) Reset() {
	*x = Accounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accounts) ProtoMessage() {}

func (x *Accounts) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accounts.ProtoReflect.Descriptor instead.
func (*Accounts) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{4}
}

func (x *Accounts) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xc8, 0x07, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x54, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x64, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x20, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x17, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x6f, 0x18, 0x18, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x74, 0x6f, 0x18, 0x19, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x62, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x63, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x1b, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x63, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75,
	0x74, 0x62, 0x6f, 0x78, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x62,
	0x6f, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18,
	0x22, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x23, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x24, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x69, 0x6b, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x27, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x29, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xbe, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0xbf, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x22, 0x32, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x27,
	0x0a, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x36, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_person_proto_goTypes = []interface{}{
	(*Person)(nil),      // 0: model.Person
	(*Account)(nil),     // 1: model.Account
	(*UserPreview)(nil), // 2: model.UserPreview
	(*Persons)(nil),     // 3: model.Persons
	(*Accounts)(nil),    // 4: model.Accounts
}
var file_person_proto_depIdxs = []int32{
	0, // 0: model.Account.person:type_name -> model.Person
	0, // 1: model.Persons.Persons:type_name -> model.Person
	1, // 2: model.Accounts.accounts:type_name -> model.Account
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPreview); i {
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
		file_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Persons); i {
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
		file_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accounts); i {
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
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}
