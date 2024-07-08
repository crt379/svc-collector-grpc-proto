// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: appapi/appapi.proto

package appapi

import (
	application "github.com/crt379/svc-collector-grpc-proto/application"
	service "github.com/crt379/svc-collector-grpc-proto/service"
	svcapi "github.com/crt379/svc-collector-grpc-proto/svcapi"
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

type AppapiMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application *application.ApplicationMete `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	Service     *service.ServiceMeta         `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Svcapis     []*svcapi.SvcapiMeta         `protobuf:"bytes,3,rep,name=svcapis,proto3" json:"svcapis,omitempty"`
}

func (x *AppapiMeta) Reset() {
	*x = AppapiMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appapi_appapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppapiMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppapiMeta) ProtoMessage() {}

func (x *AppapiMeta) ProtoReflect() protoreflect.Message {
	mi := &file_appapi_appapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppapiMeta.ProtoReflect.Descriptor instead.
func (*AppapiMeta) Descriptor() ([]byte, []int) {
	return file_appapi_appapi_proto_rawDescGZIP(), []int{0}
}

func (x *AppapiMeta) GetApplication() *application.ApplicationMete {
	if x != nil {
		return x.Application
	}
	return nil
}

func (x *AppapiMeta) GetService() *service.ServiceMeta {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *AppapiMeta) GetSvcapis() []*svcapi.SvcapiMeta {
	if x != nil {
		return x.Svcapis
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit    int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Appid    int32  `protobuf:"varint,3,opt,name=appid,proto3" json:"appid,omitempty"`
	Appname  string `protobuf:"bytes,4,opt,name=appname,proto3" json:"appname,omitempty"`
	Appsvcid int32  `protobuf:"varint,5,opt,name=appsvcid,proto3" json:"appsvcid,omitempty"`
	Svcid    int32  `protobuf:"varint,6,opt,name=svcid,proto3" json:"svcid,omitempty"`
	Svcname  string `protobuf:"bytes,7,opt,name=svcname,proto3" json:"svcname,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appapi_appapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appapi_appapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_appapi_appapi_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRequest) GetAppid() int32 {
	if x != nil {
		return x.Appid
	}
	return 0
}

func (x *GetRequest) GetAppname() string {
	if x != nil {
		return x.Appname
	}
	return ""
}

func (x *GetRequest) GetAppsvcid() int32 {
	if x != nil {
		return x.Appsvcid
	}
	return 0
}

func (x *GetRequest) GetSvcid() int32 {
	if x != nil {
		return x.Svcid
	}
	return 0
}

func (x *GetRequest) GetSvcname() string {
	if x != nil {
		return x.Svcname
	}
	return ""
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Appapis []*AppapiMeta `protobuf:"bytes,3,rep,name=appapis,proto3" json:"appapis,omitempty"`
	Count   int32         `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Total   int32         `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	Page    int32         `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	Limit   int32         `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appapi_appapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_appapi_appapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_appapi_appapi_proto_rawDescGZIP(), []int{2}
}

func (x *GetReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetReply) GetAppapis() []*AppapiMeta {
	if x != nil {
		return x.Appapis
	}
	return nil
}

func (x *GetReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetReply) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetReply) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_appapi_appapi_proto protoreflect.FileDescriptor

var file_appapi_appapi_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x1a,
	0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x76, 0x63, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x76, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0a, 0x41,
	0x70, 0x70, 0x61, 0x70, 0x69, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x50, 0x0a, 0x0b, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x65, 0x52, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x07, 0x73, 0x76, 0x63, 0x61, 0x70, 0x69, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x73, 0x76, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x76, 0x63, 0x61, 0x70, 0x69,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x73, 0x76, 0x63, 0x61, 0x70, 0x69, 0x73, 0x22, 0xb2, 0x01,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x70, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x73, 0x76, 0x63,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x70, 0x70, 0x73, 0x76, 0x63,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x76, 0x63, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x73, 0x76, 0x63, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x76, 0x63, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x76, 0x63, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a,
	0x07, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x61, 0x70, 0x69,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x32, 0xcd, 0x01, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x61, 0x70, 0x69, 0x12, 0xc2,
	0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x71, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x6b, 0x5a, 0x23, 0x12, 0x21, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x73, 0x5a, 0x36, 0x12,
	0x34, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x76,
	0x63, 0x73, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x73, 0x76, 0x63, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70,
	0x70, 0x61, 0x70, 0x69, 0x73, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x61,
	0x70, 0x69, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x63, 0x72, 0x74, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x76, 0x63, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f,
	0x73, 0x76, 0x63, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appapi_appapi_proto_rawDescOnce sync.Once
	file_appapi_appapi_proto_rawDescData = file_appapi_appapi_proto_rawDesc
)

func file_appapi_appapi_proto_rawDescGZIP() []byte {
	file_appapi_appapi_proto_rawDescOnce.Do(func() {
		file_appapi_appapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_appapi_appapi_proto_rawDescData)
	})
	return file_appapi_appapi_proto_rawDescData
}

var file_appapi_appapi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_appapi_appapi_proto_goTypes = []interface{}{
	(*AppapiMeta)(nil),                  // 0: service.collector.appapi.AppapiMeta
	(*GetRequest)(nil),                  // 1: service.collector.appapi.GetRequest
	(*GetReply)(nil),                    // 2: service.collector.appapi.GetReply
	(*application.ApplicationMete)(nil), // 3: service.collector.application.ApplicationMete
	(*service.ServiceMeta)(nil),         // 4: service.collector.service.ServiceMeta
	(*svcapi.SvcapiMeta)(nil),           // 5: service.collector.svcapi.SvcapiMeta
}
var file_appapi_appapi_proto_depIdxs = []int32{
	3, // 0: service.collector.appapi.AppapiMeta.application:type_name -> service.collector.application.ApplicationMete
	4, // 1: service.collector.appapi.AppapiMeta.service:type_name -> service.collector.service.ServiceMeta
	5, // 2: service.collector.appapi.AppapiMeta.svcapis:type_name -> service.collector.svcapi.SvcapiMeta
	0, // 3: service.collector.appapi.GetReply.appapis:type_name -> service.collector.appapi.AppapiMeta
	1, // 4: service.collector.appapi.Appapi.Get:input_type -> service.collector.appapi.GetRequest
	2, // 5: service.collector.appapi.Appapi.Get:output_type -> service.collector.appapi.GetReply
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_appapi_appapi_proto_init() }
func file_appapi_appapi_proto_init() {
	if File_appapi_appapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appapi_appapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppapiMeta); i {
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
		file_appapi_appapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_appapi_appapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
			RawDescriptor: file_appapi_appapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appapi_appapi_proto_goTypes,
		DependencyIndexes: file_appapi_appapi_proto_depIdxs,
		MessageInfos:      file_appapi_appapi_proto_msgTypes,
	}.Build()
	File_appapi_appapi_proto = out.File
	file_appapi_appapi_proto_rawDesc = nil
	file_appapi_appapi_proto_goTypes = nil
	file_appapi_appapi_proto_depIdxs = nil
}