// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: comments/service.proto

package comments_pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleID     uint32  `protobuf:"varint,1,opt,name=articleID,proto3" json:"articleID,omitempty"`
	AuthorID      uint32  `protobuf:"varint,2,opt,name=authorID,proto3" json:"authorID,omitempty"`
	Body          string  `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	ReplyTo       *uint32 `protobuf:"varint,4,opt,name=replyTo,proto3,oneof" json:"replyTo,omitempty"`
	MentionUserID *uint32 `protobuf:"varint,5,opt,name=mentionUserID,proto3,oneof" json:"mentionUserID,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCommentRequest) GetArticleID() uint32 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

func (x *CreateCommentRequest) GetAuthorID() uint32 {
	if x != nil {
		return x.AuthorID
	}
	return 0
}

func (x *CreateCommentRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CreateCommentRequest) GetReplyTo() uint32 {
	if x != nil && x.ReplyTo != nil {
		return *x.ReplyTo
	}
	return 0
}

func (x *CreateCommentRequest) GetMentionUserID() uint32 {
	if x != nil && x.MentionUserID != nil {
		return *x.MentionUserID
	}
	return 0
}

type DeleteCommnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentID uint32 `protobuf:"varint,1,opt,name=commentID,proto3" json:"commentID,omitempty"`
	ArticleID uint32 `protobuf:"varint,2,opt,name=articleID,proto3" json:"articleID,omitempty"`
	UserID    uint32 `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *DeleteCommnetRequest) Reset() {
	*x = DeleteCommnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommnetRequest) ProtoMessage() {}

func (x *DeleteCommnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommnetRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommnetRequest) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteCommnetRequest) GetCommentID() uint32 {
	if x != nil {
		return x.CommentID
	}
	return 0
}

func (x *DeleteCommnetRequest) GetArticleID() uint32 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

func (x *DeleteCommnetRequest) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetCommentsByArticleIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleID uint32  `protobuf:"varint,1,opt,name=articleID,proto3" json:"articleID,omitempty"`
	UserID    *uint32 `protobuf:"varint,2,opt,name=userID,proto3,oneof" json:"userID,omitempty"`
}

func (x *GetCommentsByArticleIDRequest) Reset() {
	*x = GetCommentsByArticleIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByArticleIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByArticleIDRequest) ProtoMessage() {}

func (x *GetCommentsByArticleIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByArticleIDRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsByArticleIDRequest) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCommentsByArticleIDRequest) GetArticleID() uint32 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

func (x *GetCommentsByArticleIDRequest) GetUserID() uint32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

type CommentElem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorID   uint32               `protobuf:"varint,1,opt,name=authorID,proto3" json:"authorID,omitempty"`
	Body       string               `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ReplyCount uint32               `protobuf:"varint,4,opt,name=replyCount,proto3" json:"replyCount,omitempty"`
}

func (x *CommentElem) Reset() {
	*x = CommentElem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentElem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentElem) ProtoMessage() {}

func (x *CommentElem) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentElem.ProtoReflect.Descriptor instead.
func (*CommentElem) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{3}
}

func (x *CommentElem) GetAuthorID() uint32 {
	if x != nil {
		return x.AuthorID
	}
	return 0
}

func (x *CommentElem) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CommentElem) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CommentElem) GetReplyCount() uint32 {
	if x != nil {
		return x.ReplyCount
	}
	return 0
}

type GetCommentsByArticleIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*CommentElem `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetCommentsByArticleIDResponse) Reset() {
	*x = GetCommentsByArticleIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsByArticleIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByArticleIDResponse) ProtoMessage() {}

func (x *GetCommentsByArticleIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByArticleIDResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsByArticleIDResponse) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentsByArticleIDResponse) GetComments() []*CommentElem {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetRepliesByCommentIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentID uint32  `protobuf:"varint,1,opt,name=commentID,proto3" json:"commentID,omitempty"`
	ArticleID uint32  `protobuf:"varint,2,opt,name=articleID,proto3" json:"articleID,omitempty"`
	UserID    *uint32 `protobuf:"varint,3,opt,name=userID,proto3,oneof" json:"userID,omitempty"`
}

func (x *GetRepliesByCommentIDRequest) Reset() {
	*x = GetRepliesByCommentIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepliesByCommentIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepliesByCommentIDRequest) ProtoMessage() {}

func (x *GetRepliesByCommentIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepliesByCommentIDRequest.ProtoReflect.Descriptor instead.
func (*GetRepliesByCommentIDRequest) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetRepliesByCommentIDRequest) GetCommentID() uint32 {
	if x != nil {
		return x.CommentID
	}
	return 0
}

func (x *GetRepliesByCommentIDRequest) GetArticleID() uint32 {
	if x != nil {
		return x.ArticleID
	}
	return 0
}

func (x *GetRepliesByCommentIDRequest) GetUserID() uint32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

type ReplyElem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorID      uint32               `protobuf:"varint,1,opt,name=authorID,proto3" json:"authorID,omitempty"`
	Body          string               `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	MentionUserID *uint32              `protobuf:"varint,4,opt,name=mentionUserID,proto3,oneof" json:"mentionUserID,omitempty"`
}

func (x *ReplyElem) Reset() {
	*x = ReplyElem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyElem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyElem) ProtoMessage() {}

func (x *ReplyElem) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyElem.ProtoReflect.Descriptor instead.
func (*ReplyElem) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{6}
}

func (x *ReplyElem) GetAuthorID() uint32 {
	if x != nil {
		return x.AuthorID
	}
	return 0
}

func (x *ReplyElem) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ReplyElem) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ReplyElem) GetMentionUserID() uint32 {
	if x != nil && x.MentionUserID != nil {
		return *x.MentionUserID
	}
	return 0
}

type GetRepliesByCommentIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replies []*ReplyElem `protobuf:"bytes,1,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *GetRepliesByCommentIDResponse) Reset() {
	*x = GetRepliesByCommentIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepliesByCommentIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepliesByCommentIDResponse) ProtoMessage() {}

func (x *GetRepliesByCommentIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepliesByCommentIDResponse.ProtoReflect.Descriptor instead.
func (*GetRepliesByCommentIDResponse) Descriptor() ([]byte, []int) {
	return file_comments_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetRepliesByCommentIDResponse) GetReplies() []*ReplyElem {
	if x != nil {
		return x.Replies
	}
	return nil
}

var File_comments_service_proto protoreflect.FileDescriptor

var file_comments_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x65,
	0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcc, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52,
	0x0d, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x6a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x65, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6c,
	0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x1e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x82, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x42, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1b,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xb2, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x45, 0x6c, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29,
	0x0a, 0x0d, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0d, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4e, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x45, 0x6c,
	0x65, 0x6d, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x32, 0xf9, 0x02, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e,
	0x65, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x6b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x55, 0x4c, 0x4c, 0x2d, 0x53, 0x6f, 0x66, 0x74, 0x77,
	0x61, 0x72, 0x65, 0x4d, 0x65, 0x69, 0x73, 0x74, 0x65, 0x72, 0x48, 0x69, 0x67, 0x68, 0x53, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comments_service_proto_rawDescOnce sync.Once
	file_comments_service_proto_rawDescData = file_comments_service_proto_rawDesc
)

func file_comments_service_proto_rawDescGZIP() []byte {
	file_comments_service_proto_rawDescOnce.Do(func() {
		file_comments_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_service_proto_rawDescData)
	})
	return file_comments_service_proto_rawDescData
}

var file_comments_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_comments_service_proto_goTypes = []interface{}{
	(*CreateCommentRequest)(nil),           // 0: commnets.CreateCommentRequest
	(*DeleteCommnetRequest)(nil),           // 1: commnets.DeleteCommnetRequest
	(*GetCommentsByArticleIDRequest)(nil),  // 2: commnets.GetCommentsByArticleIDRequest
	(*CommentElem)(nil),                    // 3: commnets.CommentElem
	(*GetCommentsByArticleIDResponse)(nil), // 4: commnets.GetCommentsByArticleIDResponse
	(*GetRepliesByCommentIDRequest)(nil),   // 5: commnets.GetRepliesByCommentIDRequest
	(*ReplyElem)(nil),                      // 6: commnets.ReplyElem
	(*GetRepliesByCommentIDResponse)(nil),  // 7: commnets.GetRepliesByCommentIDResponse
	(*timestamp.Timestamp)(nil),            // 8: google.protobuf.Timestamp
	(*empty.Empty)(nil),                    // 9: google.protobuf.Empty
}
var file_comments_service_proto_depIdxs = []int32{
	8, // 0: commnets.CommentElem.createdAt:type_name -> google.protobuf.Timestamp
	3, // 1: commnets.GetCommentsByArticleIDResponse.comments:type_name -> commnets.CommentElem
	8, // 2: commnets.ReplyElem.createdAt:type_name -> google.protobuf.Timestamp
	6, // 3: commnets.GetRepliesByCommentIDResponse.replies:type_name -> commnets.ReplyElem
	0, // 4: commnets.CommentService.CreateComment:input_type -> commnets.CreateCommentRequest
	1, // 5: commnets.CommentService.DeleteCommnet:input_type -> commnets.DeleteCommnetRequest
	2, // 6: commnets.CommentService.GetCommentsByArticleID:input_type -> commnets.GetCommentsByArticleIDRequest
	5, // 7: commnets.CommentService.GetRepliesByCommentID:input_type -> commnets.GetRepliesByCommentIDRequest
	9, // 8: commnets.CommentService.CreateComment:output_type -> google.protobuf.Empty
	9, // 9: commnets.CommentService.DeleteCommnet:output_type -> google.protobuf.Empty
	4, // 10: commnets.CommentService.GetCommentsByArticleID:output_type -> commnets.GetCommentsByArticleIDResponse
	7, // 11: commnets.CommentService.GetRepliesByCommentID:output_type -> commnets.GetRepliesByCommentIDResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_comments_service_proto_init() }
func file_comments_service_proto_init() {
	if File_comments_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comments_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
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
		file_comments_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommnetRequest); i {
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
		file_comments_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByArticleIDRequest); i {
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
		file_comments_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentElem); i {
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
		file_comments_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsByArticleIDResponse); i {
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
		file_comments_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepliesByCommentIDRequest); i {
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
		file_comments_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyElem); i {
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
		file_comments_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepliesByCommentIDResponse); i {
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
	file_comments_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_comments_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_comments_service_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_comments_service_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comments_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comments_service_proto_goTypes,
		DependencyIndexes: file_comments_service_proto_depIdxs,
		MessageInfos:      file_comments_service_proto_msgTypes,
	}.Build()
	File_comments_service_proto = out.File
	file_comments_service_proto_rawDesc = nil
	file_comments_service_proto_goTypes = nil
	file_comments_service_proto_depIdxs = nil
}
