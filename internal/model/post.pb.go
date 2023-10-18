// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: post.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Post represents a post in a social media platform.
type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the post.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Community to which the post belongs.
	Community string `protobuf:"bytes,2,opt,name=community,proto3" json:"community,omitempty"`
	// Content of the post.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Timestamp when the post was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Timestamp when the post was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Timestamp when the post was deleted.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// Number of likes the post has received.
	Likes int64 `protobuf:"varint,7,opt,name=likes,proto3" json:"likes,omitempty"`
	// Number of shares the post has received.
	Shares int64 `protobuf:"varint,8,opt,name=shares,proto3" json:"shares,omitempty"`
	// Number of comments the post has received.
	Comments int64 `protobuf:"varint,9,opt,name=comments,proto3" json:"comments,omitempty"`
	// Author of the post.
	Author string `protobuf:"bytes,10,opt,name=author,proto3" json:"author,omitempty"`
	// Parent post of the current post.
	Parent string `protobuf:"bytes,11,opt,name=parent,proto3" json:"parent,omitempty"`
	// Rank of the post.
	Rank int64 `protobuf:"varint,12,opt,name=rank,proto3" json:"rank,omitempty"`
	// Child posts of the current post.
	Children []string `protobuf:"bytes,13,rep,name=children,proto3" json:"children,omitempty"`
	// Tags associated with the post.
	Tags []string `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
	// Users mentioned in the post.
	Mentions []string `protobuf:"bytes,15,rep,name=mentions,proto3" json:"mentions,omitempty"`
	// Indicates if the post is sensitive.
	IsSensitive bool `protobuf:"varint,16,opt,name=is_sensitive,json=isSensitive,proto3" json:"is_sensitive,omitempty"`
	// Indicates if the post is not safe for work.
	IsNsfw bool `protobuf:"varint,17,opt,name=is_nsfw,json=isNsfw,proto3" json:"is_nsfw,omitempty"`
	// Indicates if the post has been deleted.
	IsDeleted bool `protobuf:"varint,18,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// Indicates if the post is pinned.
	IsPinned bool `protobuf:"varint,19,opt,name=is_pinned,json=isPinned,proto3" json:"is_pinned,omitempty"`
	// Indicates if the post has been edited.
	IsEdited bool `protobuf:"varint,20,opt,name=is_edited,json=isEdited,proto3" json:"is_edited,omitempty"`
	// Indicates if the post has been liked by the user.
	IsLiked bool `protobuf:"varint,21,opt,name=is_liked,json=isLiked,proto3" json:"is_liked,omitempty"`
	// Indicates if the post has been shared by the user.
	IsShared bool `protobuf:"varint,22,opt,name=is_shared,json=isShared,proto3" json:"is_shared,omitempty"`
	// Indicates if the post has been commented on by the user.
	IsCommented bool `protobuf:"varint,23,opt,name=is_commented,json=isCommented,proto3" json:"is_commented,omitempty"`
	// Indicates if the user is subscribed to the post.
	IsSubscribed bool `protobuf:"varint,24,opt,name=is_subscribed,json=isSubscribed,proto3" json:"is_subscribed,omitempty"`
	// Indicates if the post has been bookmarked by the user.
	IsBookmarked bool `protobuf:"varint,25,opt,name=is_bookmarked,json=isBookmarked,proto3" json:"is_bookmarked,omitempty"`
	// Indicates if the post has been reblogged by the user.
	IsReblogged bool `protobuf:"varint,26,opt,name=is_reblogged,json=isReblogged,proto3" json:"is_reblogged,omitempty"`
	// Indicates if the user has been mentioned in the post.
	IsMentioned bool `protobuf:"varint,27,opt,name=is_mentioned,json=isMentioned,proto3" json:"is_mentioned,omitempty"`
	// Indicates if the post is a poll.
	IsPoll bool `protobuf:"varint,28,opt,name=is_poll,json=isPoll,proto3" json:"is_poll,omitempty"`
	// Indicates if the user has voted in the poll.
	IsPollVoted bool `protobuf:"varint,29,opt,name=is_poll_voted,json=isPollVoted,proto3" json:"is_poll_voted,omitempty"`
	// Indicates if the poll has expired.
	IsPollExpired bool `protobuf:"varint,30,opt,name=is_poll_expired,json=isPollExpired,proto3" json:"is_poll_expired,omitempty"`
	// Indicates if the poll has been closed.
	IsPollClosed bool `protobuf:"varint,31,opt,name=is_poll_closed,json=isPollClosed,proto3" json:"is_poll_closed,omitempty"`
	// Indicates if the poll allows multiple choices.
	IsPollMultiple bool `protobuf:"varint,32,opt,name=is_poll_multiple,json=isPollMultiple,proto3" json:"is_poll_multiple,omitempty"`
	// Indicates if the poll hides the total number of votes.
	IsPollHideTotals bool `protobuf:"varint,33,opt,name=is_poll_hide_totals,json=isPollHideTotals,proto3" json:"is_poll_hide_totals,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetCommunity() string {
	if x != nil {
		return x.Community
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Post) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Post) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *Post) GetShares() int64 {
	if x != nil {
		return x.Shares
	}
	return 0
}

func (x *Post) GetComments() int64 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *Post) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Post) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *Post) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Post) GetChildren() []string {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Post) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Post) GetMentions() []string {
	if x != nil {
		return x.Mentions
	}
	return nil
}

func (x *Post) GetIsSensitive() bool {
	if x != nil {
		return x.IsSensitive
	}
	return false
}

func (x *Post) GetIsNsfw() bool {
	if x != nil {
		return x.IsNsfw
	}
	return false
}

func (x *Post) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *Post) GetIsPinned() bool {
	if x != nil {
		return x.IsPinned
	}
	return false
}

func (x *Post) GetIsEdited() bool {
	if x != nil {
		return x.IsEdited
	}
	return false
}

func (x *Post) GetIsLiked() bool {
	if x != nil {
		return x.IsLiked
	}
	return false
}

func (x *Post) GetIsShared() bool {
	if x != nil {
		return x.IsShared
	}
	return false
}

func (x *Post) GetIsCommented() bool {
	if x != nil {
		return x.IsCommented
	}
	return false
}

func (x *Post) GetIsSubscribed() bool {
	if x != nil {
		return x.IsSubscribed
	}
	return false
}

func (x *Post) GetIsBookmarked() bool {
	if x != nil {
		return x.IsBookmarked
	}
	return false
}

func (x *Post) GetIsReblogged() bool {
	if x != nil {
		return x.IsReblogged
	}
	return false
}

func (x *Post) GetIsMentioned() bool {
	if x != nil {
		return x.IsMentioned
	}
	return false
}

func (x *Post) GetIsPoll() bool {
	if x != nil {
		return x.IsPoll
	}
	return false
}

func (x *Post) GetIsPollVoted() bool {
	if x != nil {
		return x.IsPollVoted
	}
	return false
}

func (x *Post) GetIsPollExpired() bool {
	if x != nil {
		return x.IsPollExpired
	}
	return false
}

func (x *Post) GetIsPollClosed() bool {
	if x != nil {
		return x.IsPollClosed
	}
	return false
}

func (x *Post) GetIsPollMultiple() bool {
	if x != nil {
		return x.IsPollMultiple
	}
	return false
}

func (x *Post) GetIsPollHideTotals() bool {
	if x != nil {
		return x.IsPollHideTotals
	}
	return false
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x08, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6e, 0x73, 0x66, 0x77, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4e, 0x73, 0x66, 0x77, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65,
	0x64, 0x69, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x45,
	0x64, 0x69, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65,
	0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73,
	0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x52, 0x65, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x1b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x1c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f,
	0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x6c,
	0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x69,
	0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x6c,
	0x5f, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x73, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x50, 0x6f, 0x6c, 0x6c, 0x48, 0x69, 0x64, 0x65, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData = file_post_proto_rawDesc
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_proto_rawDescData)
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_post_proto_goTypes = []interface{}{
	(*Post)(nil),                  // 0: model.Post
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_post_proto_depIdxs = []int32{
	1, // 0: model.Post.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: model.Post.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: model.Post.deleted_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
			RawDescriptor: file_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_rawDesc = nil
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
