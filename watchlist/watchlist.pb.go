// protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative watchlist.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: watchlist.proto

package watchlist

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

// Элемент списка просмотра
type WatchlistItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MediaId   int64  `protobuf:"varint,2,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	UserId    int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *WatchlistItem) Reset() {
	*x = WatchlistItem{}
	mi := &file_watchlist_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchlistItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchlistItem) ProtoMessage() {}

func (x *WatchlistItem) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchlistItem.ProtoReflect.Descriptor instead.
func (*WatchlistItem) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{0}
}

func (x *WatchlistItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WatchlistItem) GetMediaId() int64 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *WatchlistItem) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WatchlistItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// Запрос на добавление медиа в список просмотра
type AddToWatchlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaId int64 `protobuf:"varint,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AddToWatchlistRequest) Reset() {
	*x = AddToWatchlistRequest{}
	mi := &file_watchlist_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToWatchlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToWatchlistRequest) ProtoMessage() {}

func (x *AddToWatchlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToWatchlistRequest.ProtoReflect.Descriptor instead.
func (*AddToWatchlistRequest) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{1}
}

func (x *AddToWatchlistRequest) GetMediaId() int64 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *AddToWatchlistRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// Ответ на добавление медиа в список просмотра
type AddToWatchlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddToWatchlistResponse) Reset() {
	*x = AddToWatchlistResponse{}
	mi := &file_watchlist_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToWatchlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToWatchlistResponse) ProtoMessage() {}

func (x *AddToWatchlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToWatchlistResponse.ProtoReflect.Descriptor instead.
func (*AddToWatchlistResponse) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{2}
}

func (x *AddToWatchlistResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Запрос на удаление медиа из списка просмотра
type RemoveFromWatchlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaId int64 `protobuf:"varint,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveFromWatchlistRequest) Reset() {
	*x = RemoveFromWatchlistRequest{}
	mi := &file_watchlist_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromWatchlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromWatchlistRequest) ProtoMessage() {}

func (x *RemoveFromWatchlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromWatchlistRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromWatchlistRequest) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFromWatchlistRequest) GetMediaId() int64 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *RemoveFromWatchlistRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// Ответ на удаление медиа из списка просмотра
type RemoveFromWatchlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveFromWatchlistResponse) Reset() {
	*x = RemoveFromWatchlistResponse{}
	mi := &file_watchlist_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromWatchlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromWatchlistResponse) ProtoMessage() {}

func (x *RemoveFromWatchlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromWatchlistResponse.ProtoReflect.Descriptor instead.
func (*RemoveFromWatchlistResponse) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveFromWatchlistResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Запрос на получение списка просмотра пользователя
type GetWatchlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetWatchlistRequest) Reset() {
	*x = GetWatchlistRequest{}
	mi := &file_watchlist_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWatchlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWatchlistRequest) ProtoMessage() {}

func (x *GetWatchlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWatchlistRequest.ProtoReflect.Descriptor instead.
func (*GetWatchlistRequest) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{5}
}

func (x *GetWatchlistRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// Ответ на получение списка просмотра пользователя
type GetWatchlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Watchlists []*WatchlistItem `protobuf:"bytes,1,rep,name=watchlists,proto3" json:"watchlists,omitempty"`
}

func (x *GetWatchlistResponse) Reset() {
	*x = GetWatchlistResponse{}
	mi := &file_watchlist_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWatchlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWatchlistResponse) ProtoMessage() {}

func (x *GetWatchlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWatchlistResponse.ProtoReflect.Descriptor instead.
func (*GetWatchlistResponse) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{6}
}

func (x *GetWatchlistResponse) GetWatchlists() []*WatchlistItem {
	if x != nil {
		return x.Watchlists
	}
	return nil
}

// Запрос на проверку наличия медиа в списке просмотра
type CheckInWatchlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaId int64 `protobuf:"varint,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CheckInWatchlistRequest) Reset() {
	*x = CheckInWatchlistRequest{}
	mi := &file_watchlist_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckInWatchlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInWatchlistRequest) ProtoMessage() {}

func (x *CheckInWatchlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInWatchlistRequest.ProtoReflect.Descriptor instead.
func (*CheckInWatchlistRequest) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{7}
}

func (x *CheckInWatchlistRequest) GetMediaId() int64 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *CheckInWatchlistRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// Ответ на проверку наличия медиа в списке просмотра
type CheckInWatchlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InWatchlist bool `protobuf:"varint,1,opt,name=in_watchlist,json=inWatchlist,proto3" json:"in_watchlist,omitempty"`
}

func (x *CheckInWatchlistResponse) Reset() {
	*x = CheckInWatchlistResponse{}
	mi := &file_watchlist_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckInWatchlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInWatchlistResponse) ProtoMessage() {}

func (x *CheckInWatchlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_watchlist_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInWatchlistResponse.ProtoReflect.Descriptor instead.
func (*CheckInWatchlistResponse) Descriptor() ([]byte, []int) {
	return file_watchlist_proto_rawDescGZIP(), []int{8}
}

func (x *CheckInWatchlistResponse) GetInWatchlist() bool {
	if x != nil {
		return x.InWatchlist
	}
	return false
}

var File_watchlist_proto protoreflect.FileDescriptor

var file_watchlist_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x72, 0x0a, 0x0d,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x4b, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a,
	0x16, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x50, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2e, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x4d,
	0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3d, 0x0a,
	0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x5f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x85, 0x03, 0x0a,
	0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x57, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x13, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x25, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x1e, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x6b, 0x61, 0x74,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_watchlist_proto_rawDescOnce sync.Once
	file_watchlist_proto_rawDescData = file_watchlist_proto_rawDesc
)

func file_watchlist_proto_rawDescGZIP() []byte {
	file_watchlist_proto_rawDescOnce.Do(func() {
		file_watchlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_watchlist_proto_rawDescData)
	})
	return file_watchlist_proto_rawDescData
}

var file_watchlist_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_watchlist_proto_goTypes = []any{
	(*WatchlistItem)(nil),               // 0: watchlist.WatchlistItem
	(*AddToWatchlistRequest)(nil),       // 1: watchlist.AddToWatchlistRequest
	(*AddToWatchlistResponse)(nil),      // 2: watchlist.AddToWatchlistResponse
	(*RemoveFromWatchlistRequest)(nil),  // 3: watchlist.RemoveFromWatchlistRequest
	(*RemoveFromWatchlistResponse)(nil), // 4: watchlist.RemoveFromWatchlistResponse
	(*GetWatchlistRequest)(nil),         // 5: watchlist.GetWatchlistRequest
	(*GetWatchlistResponse)(nil),        // 6: watchlist.GetWatchlistResponse
	(*CheckInWatchlistRequest)(nil),     // 7: watchlist.CheckInWatchlistRequest
	(*CheckInWatchlistResponse)(nil),    // 8: watchlist.CheckInWatchlistResponse
}
var file_watchlist_proto_depIdxs = []int32{
	0, // 0: watchlist.GetWatchlistResponse.watchlists:type_name -> watchlist.WatchlistItem
	1, // 1: watchlist.WatchlistService.AddToWatchlist:input_type -> watchlist.AddToWatchlistRequest
	3, // 2: watchlist.WatchlistService.RemoveFromWatchlist:input_type -> watchlist.RemoveFromWatchlistRequest
	5, // 3: watchlist.WatchlistService.GetWatchlist:input_type -> watchlist.GetWatchlistRequest
	7, // 4: watchlist.WatchlistService.CheckInWatchlist:input_type -> watchlist.CheckInWatchlistRequest
	2, // 5: watchlist.WatchlistService.AddToWatchlist:output_type -> watchlist.AddToWatchlistResponse
	4, // 6: watchlist.WatchlistService.RemoveFromWatchlist:output_type -> watchlist.RemoveFromWatchlistResponse
	6, // 7: watchlist.WatchlistService.GetWatchlist:output_type -> watchlist.GetWatchlistResponse
	8, // 8: watchlist.WatchlistService.CheckInWatchlist:output_type -> watchlist.CheckInWatchlistResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_watchlist_proto_init() }
func file_watchlist_proto_init() {
	if File_watchlist_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_watchlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_watchlist_proto_goTypes,
		DependencyIndexes: file_watchlist_proto_depIdxs,
		MessageInfos:      file_watchlist_proto_msgTypes,
	}.Build()
	File_watchlist_proto = out.File
	file_watchlist_proto_rawDesc = nil
	file_watchlist_proto_goTypes = nil
	file_watchlist_proto_depIdxs = nil
}
