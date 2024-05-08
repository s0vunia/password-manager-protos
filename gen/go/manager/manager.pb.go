// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: manager/manager.proto

package manager

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

type ItemType int32

const (
	ItemType_Login ItemType = 0
)

// Enum value maps for ItemType.
var (
	ItemType_name = map[int32]string{
		0: "Login",
	}
	ItemType_value = map[string]int32{
		"Login": 0,
	}
)

func (x ItemType) Enum() *ItemType {
	p := new(ItemType)
	*p = x
	return p
}

func (x ItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_manager_manager_proto_enumTypes[0].Descriptor()
}

func (ItemType) Type() protoreflect.EnumType {
	return &file_manager_manager_proto_enumTypes[0]
}

func (x ItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemType.Descriptor instead.
func (ItemType) EnumDescriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{0}
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CreateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FolderId   *UUID  `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	UserId     *UUID  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsFavorite bool   `protobuf:"varint,4,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
}

func (x *CreateItemRequest) Reset() {
	*x = CreateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemRequest) ProtoMessage() {}

func (x *CreateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemRequest.ProtoReflect.Descriptor instead.
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{1}
}

func (x *CreateItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateItemRequest) GetFolderId() *UUID {
	if x != nil {
		return x.FolderId
	}
	return nil
}

func (x *CreateItemRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *CreateItemRequest) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type CreateItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateItemResponse) Reset() {
	*x = CreateItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemResponse) ProtoMessage() {}

func (x *CreateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemResponse.ProtoReflect.Descriptor instead.
func (*CreateItemResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{2}
}

func (x *CreateItemResponse) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type CreateLoginItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item            *CreateItemRequest  `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	ItemId          *CreateItemResponse `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Login           string              `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
	EncryptPassword string              `protobuf:"bytes,4,opt,name=encrypt_password,json=encryptPassword,proto3" json:"encrypt_password,omitempty"`
}

func (x *CreateLoginItemRequest) Reset() {
	*x = CreateLoginItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginItemRequest) ProtoMessage() {}

func (x *CreateLoginItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginItemRequest.ProtoReflect.Descriptor instead.
func (*CreateLoginItemRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLoginItemRequest) GetItem() *CreateItemRequest {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *CreateLoginItemRequest) GetItemId() *CreateItemResponse {
	if x != nil {
		return x.ItemId
	}
	return nil
}

func (x *CreateLoginItemRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateLoginItemRequest) GetEncryptPassword() string {
	if x != nil {
		return x.EncryptPassword
	}
	return ""
}

type CreateLoginItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *CreateItemResponse `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateLoginItemResponse) Reset() {
	*x = CreateLoginItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginItemResponse) ProtoMessage() {}

func (x *CreateLoginItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginItemResponse.ProtoReflect.Descriptor instead.
func (*CreateLoginItemResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLoginItemResponse) GetItem() *CreateItemResponse {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{5}
}

func (x *GetItemRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       ItemType `protobuf:"varint,2,opt,name=type,proto3,enum=auth.ItemType" json:"type,omitempty"`
	FolderId   *UUID    `protobuf:"bytes,3,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	IsFavorite bool     `protobuf:"varint,4,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{6}
}

func (x *GetItemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetItemResponse) GetType() ItemType {
	if x != nil {
		return x.Type
	}
	return ItemType_Login
}

func (x *GetItemResponse) GetFolderId() *UUID {
	if x != nil {
		return x.FolderId
	}
	return nil
}

func (x *GetItemResponse) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type GetLoginItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *GetItemRequest `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetLoginItemRequest) Reset() {
	*x = GetLoginItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginItemRequest) ProtoMessage() {}

func (x *GetLoginItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginItemRequest.ProtoReflect.Descriptor instead.
func (*GetLoginItemRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{7}
}

func (x *GetLoginItemRequest) GetItem() *GetItemRequest {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetLoginItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item            *GetItemResponse `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Login           string           `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	EncryptPassword string           `protobuf:"bytes,3,opt,name=encrypt_password,json=encryptPassword,proto3" json:"encrypt_password,omitempty"`
}

func (x *GetLoginItemResponse) Reset() {
	*x = GetLoginItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginItemResponse) ProtoMessage() {}

func (x *GetLoginItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginItemResponse.ProtoReflect.Descriptor instead.
func (*GetLoginItemResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{8}
}

func (x *GetLoginItemResponse) GetItem() *GetItemResponse {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *GetLoginItemResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetLoginItemResponse) GetEncryptPassword() string {
	if x != nil {
		return x.EncryptPassword
	}
	return ""
}

type GetItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetItemsRequest) Reset() {
	*x = GetItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsRequest) ProtoMessage() {}

func (x *GetItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsRequest.ProtoReflect.Descriptor instead.
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{9}
}

type GetItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfItems []*GetItemResponse `protobuf:"bytes,1,rep,name=list_of_items,json=listOfItems,proto3" json:"list_of_items,omitempty"`
}

func (x *GetItemsResponse) Reset() {
	*x = GetItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsResponse) ProtoMessage() {}

func (x *GetItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsResponse.ProtoReflect.Descriptor instead.
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{10}
}

func (x *GetItemsResponse) GetListOfItems() []*GetItemResponse {
	if x != nil {
		return x.ListOfItems
	}
	return nil
}

type GetLoginItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLoginItemsRequest) Reset() {
	*x = GetLoginItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginItemsRequest) ProtoMessage() {}

func (x *GetLoginItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginItemsRequest.ProtoReflect.Descriptor instead.
func (*GetLoginItemsRequest) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{11}
}

type GetLoginItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfItems []*GetLoginItemResponse `protobuf:"bytes,1,rep,name=list_of_items,json=listOfItems,proto3" json:"list_of_items,omitempty"`
}

func (x *GetLoginItemsResponse) Reset() {
	*x = GetLoginItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginItemsResponse) ProtoMessage() {}

func (x *GetLoginItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginItemsResponse.ProtoReflect.Descriptor instead.
func (*GetLoginItemsResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{12}
}

func (x *GetLoginItemsResponse) GetListOfItems() []*GetLoginItemResponse {
	if x != nil {
		return x.ListOfItems
	}
	return nil
}

var File_manager_manager_proto protoreflect.FileDescriptor

var file_manager_manager_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x1c, 0x0a,
	0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x22, 0x30, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x31,
	0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x47, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x2c, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x22,
	0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0x82, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x4f, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x57, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x6c, 0x69, 0x73,
	0x74, 0x4f, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x15, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x00, 0x32,
	0xdd, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x30,
	0x76, 0x75, 0x6e, 0x69, 0x61, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_manager_proto_rawDescOnce sync.Once
	file_manager_manager_proto_rawDescData = file_manager_manager_proto_rawDesc
)

func file_manager_manager_proto_rawDescGZIP() []byte {
	file_manager_manager_proto_rawDescOnce.Do(func() {
		file_manager_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_manager_proto_rawDescData)
	})
	return file_manager_manager_proto_rawDescData
}

var file_manager_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_manager_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_manager_manager_proto_goTypes = []interface{}{
	(ItemType)(0),                   // 0: auth.ItemType
	(*UUID)(nil),                    // 1: auth.UUID
	(*CreateItemRequest)(nil),       // 2: auth.CreateItemRequest
	(*CreateItemResponse)(nil),      // 3: auth.CreateItemResponse
	(*CreateLoginItemRequest)(nil),  // 4: auth.CreateLoginItemRequest
	(*CreateLoginItemResponse)(nil), // 5: auth.CreateLoginItemResponse
	(*GetItemRequest)(nil),          // 6: auth.GetItemRequest
	(*GetItemResponse)(nil),         // 7: auth.GetItemResponse
	(*GetLoginItemRequest)(nil),     // 8: auth.GetLoginItemRequest
	(*GetLoginItemResponse)(nil),    // 9: auth.GetLoginItemResponse
	(*GetItemsRequest)(nil),         // 10: auth.GetItemsRequest
	(*GetItemsResponse)(nil),        // 11: auth.GetItemsResponse
	(*GetLoginItemsRequest)(nil),    // 12: auth.GetLoginItemsRequest
	(*GetLoginItemsResponse)(nil),   // 13: auth.GetLoginItemsResponse
}
var file_manager_manager_proto_depIdxs = []int32{
	1,  // 0: auth.CreateItemRequest.folder_id:type_name -> auth.UUID
	1,  // 1: auth.CreateItemRequest.user_id:type_name -> auth.UUID
	1,  // 2: auth.CreateItemResponse.id:type_name -> auth.UUID
	2,  // 3: auth.CreateLoginItemRequest.item:type_name -> auth.CreateItemRequest
	3,  // 4: auth.CreateLoginItemRequest.item_id:type_name -> auth.CreateItemResponse
	3,  // 5: auth.CreateLoginItemResponse.item:type_name -> auth.CreateItemResponse
	1,  // 6: auth.GetItemRequest.id:type_name -> auth.UUID
	0,  // 7: auth.GetItemResponse.type:type_name -> auth.ItemType
	1,  // 8: auth.GetItemResponse.folder_id:type_name -> auth.UUID
	6,  // 9: auth.GetLoginItemRequest.item:type_name -> auth.GetItemRequest
	7,  // 10: auth.GetLoginItemResponse.item:type_name -> auth.GetItemResponse
	7,  // 11: auth.GetItemsResponse.list_of_items:type_name -> auth.GetItemResponse
	9,  // 12: auth.GetLoginItemsResponse.list_of_items:type_name -> auth.GetLoginItemResponse
	4,  // 13: auth.Manager.CreateLoginItem:input_type -> auth.CreateLoginItemRequest
	6,  // 14: auth.Manager.GetItem:input_type -> auth.GetItemRequest
	10, // 15: auth.Manager.GetItems:input_type -> auth.GetItemsRequest
	8,  // 16: auth.Manager.GetLoginItem:input_type -> auth.GetLoginItemRequest
	12, // 17: auth.Manager.GetLoginItems:input_type -> auth.GetLoginItemsRequest
	5,  // 18: auth.Manager.CreateLoginItem:output_type -> auth.CreateLoginItemResponse
	7,  // 19: auth.Manager.GetItem:output_type -> auth.GetItemResponse
	11, // 20: auth.Manager.GetItems:output_type -> auth.GetItemsResponse
	9,  // 21: auth.Manager.GetLoginItem:output_type -> auth.GetLoginItemResponse
	13, // 22: auth.Manager.GetLoginItems:output_type -> auth.GetLoginItemsResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_manager_manager_proto_init() }
func file_manager_manager_proto_init() {
	if File_manager_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_manager_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItemRequest); i {
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
		file_manager_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItemResponse); i {
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
		file_manager_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginItemRequest); i {
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
		file_manager_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginItemResponse); i {
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
		file_manager_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
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
		file_manager_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemResponse); i {
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
		file_manager_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginItemRequest); i {
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
		file_manager_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginItemResponse); i {
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
		file_manager_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsRequest); i {
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
		file_manager_manager_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsResponse); i {
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
		file_manager_manager_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginItemsRequest); i {
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
		file_manager_manager_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginItemsResponse); i {
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
			RawDescriptor: file_manager_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_manager_proto_goTypes,
		DependencyIndexes: file_manager_manager_proto_depIdxs,
		EnumInfos:         file_manager_manager_proto_enumTypes,
		MessageInfos:      file_manager_manager_proto_msgTypes,
	}.Build()
	File_manager_manager_proto = out.File
	file_manager_manager_proto_rawDesc = nil
	file_manager_manager_proto_goTypes = nil
	file_manager_manager_proto_depIdxs = nil
}
