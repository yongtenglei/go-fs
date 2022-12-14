// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: proto/namenode/namenode.proto

package namenode

import (
	datanode "go-fs/proto/datanode"
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

type NameNodeMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId        string                       `protobuf:"bytes,1,opt,name=blockId,proto3" json:"blockId,omitempty"`
	BlockAddresses []*datanode.DataNodeInstance `protobuf:"bytes,2,rep,name=blockAddresses,proto3" json:"blockAddresses,omitempty"`
}

func (x *NameNodeMetaData) Reset() {
	*x = NameNodeMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_namenode_namenode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameNodeMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameNodeMetaData) ProtoMessage() {}

func (x *NameNodeMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_namenode_namenode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameNodeMetaData.ProtoReflect.Descriptor instead.
func (*NameNodeMetaData) Descriptor() ([]byte, []int) {
	return file_proto_namenode_namenode_proto_rawDescGZIP(), []int{0}
}

func (x *NameNodeMetaData) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *NameNodeMetaData) GetBlockAddresses() []*datanode.DataNodeInstance {
	if x != nil {
		return x.BlockAddresses
	}
	return nil
}

type GetBlockSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request bool `protobuf:"varint,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetBlockSizeRequest) Reset() {
	*x = GetBlockSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_namenode_namenode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockSizeRequest) ProtoMessage() {}

func (x *GetBlockSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_namenode_namenode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockSizeRequest.ProtoReflect.Descriptor instead.
func (*GetBlockSizeRequest) Descriptor() ([]byte, []int) {
	return file_proto_namenode_namenode_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockSizeRequest) GetRequest() bool {
	if x != nil {
		return x.Request
	}
	return false
}

type GetBlockSizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockSize uint64 `protobuf:"varint,1,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
}

func (x *GetBlockSizeResponse) Reset() {
	*x = GetBlockSizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_namenode_namenode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockSizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockSizeResponse) ProtoMessage() {}

func (x *GetBlockSizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_namenode_namenode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockSizeResponse.ProtoReflect.Descriptor instead.
func (*GetBlockSizeResponse) Descriptor() ([]byte, []int) {
	return file_proto_namenode_namenode_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockSizeResponse) GetBlockSize() uint64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_namenode_namenode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_namenode_namenode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_proto_namenode_namenode_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameNodeMetaDataList []*NameNodeMetaData `protobuf:"bytes,1,rep,name=nameNodeMetaDataList,proto3" json:"nameNodeMetaDataList,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_namenode_namenode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_namenode_namenode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_proto_namenode_namenode_proto_rawDescGZIP(), []int{4}
}

func (x *ReadResponse) GetNameNodeMetaDataList() []*NameNodeMetaData {
	if x != nil {
		return x.NameNodeMetaDataList
	}
	return nil
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileSize uint64 `protobuf:"varint,2,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_namenode_namenode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_namenode_namenode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_proto_namenode_namenode_proto_rawDescGZIP(), []int{5}
}

func (x *WriteRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *WriteRequest) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameNodeMetaDataList []*NameNodeMetaData `protobuf:"bytes,1,rep,name=nameNodeMetaDataList,proto3" json:"nameNodeMetaDataList,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_namenode_namenode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_namenode_namenode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_proto_namenode_namenode_proto_rawDescGZIP(), []int{6}
}

func (x *WriteResponse) GetNameNodeMetaDataList() []*NameNodeMetaData {
	if x != nil {
		return x.NameNodeMetaDataList
	}
	return nil
}

var File_proto_namenode_namenode_proto protoreflect.FileDescriptor

var file_proto_namenode_namenode_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x29, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x14,
	0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x14, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x0c,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x5f, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x14, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x14, 0x6e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xd9, 0x01, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x6f, 0x2d, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_namenode_namenode_proto_rawDescOnce sync.Once
	file_proto_namenode_namenode_proto_rawDescData = file_proto_namenode_namenode_proto_rawDesc
)

func file_proto_namenode_namenode_proto_rawDescGZIP() []byte {
	file_proto_namenode_namenode_proto_rawDescOnce.Do(func() {
		file_proto_namenode_namenode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_namenode_namenode_proto_rawDescData)
	})
	return file_proto_namenode_namenode_proto_rawDescData
}

var file_proto_namenode_namenode_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_namenode_namenode_proto_goTypes = []interface{}{
	(*NameNodeMetaData)(nil),          // 0: namenode.NameNodeMetaData
	(*GetBlockSizeRequest)(nil),       // 1: namenode.GetBlockSizeRequest
	(*GetBlockSizeResponse)(nil),      // 2: namenode.GetBlockSizeResponse
	(*ReadRequest)(nil),               // 3: namenode.ReadRequest
	(*ReadResponse)(nil),              // 4: namenode.ReadResponse
	(*WriteRequest)(nil),              // 5: namenode.WriteRequest
	(*WriteResponse)(nil),             // 6: namenode.WriteResponse
	(*datanode.DataNodeInstance)(nil), // 7: datanode.DataNodeInstance
}
var file_proto_namenode_namenode_proto_depIdxs = []int32{
	7, // 0: namenode.NameNodeMetaData.blockAddresses:type_name -> datanode.DataNodeInstance
	0, // 1: namenode.ReadResponse.nameNodeMetaDataList:type_name -> namenode.NameNodeMetaData
	0, // 2: namenode.WriteResponse.nameNodeMetaDataList:type_name -> namenode.NameNodeMetaData
	1, // 3: namenode.NameNodeService.GetBlockSize:input_type -> namenode.GetBlockSizeRequest
	3, // 4: namenode.NameNodeService.ReadData:input_type -> namenode.ReadRequest
	5, // 5: namenode.NameNodeService.WriteData:input_type -> namenode.WriteRequest
	2, // 6: namenode.NameNodeService.GetBlockSize:output_type -> namenode.GetBlockSizeResponse
	4, // 7: namenode.NameNodeService.ReadData:output_type -> namenode.ReadResponse
	6, // 8: namenode.NameNodeService.WriteData:output_type -> namenode.WriteResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_namenode_namenode_proto_init() }
func file_proto_namenode_namenode_proto_init() {
	if File_proto_namenode_namenode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_namenode_namenode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameNodeMetaData); i {
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
		file_proto_namenode_namenode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockSizeRequest); i {
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
		file_proto_namenode_namenode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockSizeResponse); i {
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
		file_proto_namenode_namenode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_proto_namenode_namenode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_proto_namenode_namenode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_proto_namenode_namenode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
			RawDescriptor: file_proto_namenode_namenode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_namenode_namenode_proto_goTypes,
		DependencyIndexes: file_proto_namenode_namenode_proto_depIdxs,
		MessageInfos:      file_proto_namenode_namenode_proto_msgTypes,
	}.Build()
	File_proto_namenode_namenode_proto = out.File
	file_proto_namenode_namenode_proto_rawDesc = nil
	file_proto_namenode_namenode_proto_goTypes = nil
	file_proto_namenode_namenode_proto_depIdxs = nil
}
