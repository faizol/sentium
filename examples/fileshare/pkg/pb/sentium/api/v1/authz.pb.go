// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: sentium/api/v1/authz.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthzCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrincipalId  string `protobuf:"bytes,1,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
	ResourceId   string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
}

func (x *AuthzCheckRequest) Reset() {
	*x = AuthzCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sentium_api_v1_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthzCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthzCheckRequest) ProtoMessage() {}

func (x *AuthzCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sentium_api_v1_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthzCheckRequest.ProtoReflect.Descriptor instead.
func (*AuthzCheckRequest) Descriptor() ([]byte, []int) {
	return file_sentium_api_v1_authz_proto_rawDescGZIP(), []int{0}
}

func (x *AuthzCheckRequest) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

func (x *AuthzCheckRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AuthzCheckRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

type AuthzCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool             `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Attrs *structpb.Struct `protobuf:"bytes,2,opt,name=attrs,proto3,oneof" json:"attrs,omitempty"`
}

func (x *AuthzCheckResponse) Reset() {
	*x = AuthzCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sentium_api_v1_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthzCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthzCheckResponse) ProtoMessage() {}

func (x *AuthzCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sentium_api_v1_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthzCheckResponse.ProtoReflect.Descriptor instead.
func (*AuthzCheckResponse) Descriptor() ([]byte, []int) {
	return file_sentium_api_v1_authz_proto_rawDescGZIP(), []int{1}
}

func (x *AuthzCheckResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *AuthzCheckResponse) GetAttrs() *structpb.Struct {
	if x != nil {
		return x.Attrs
	}
	return nil
}

type AuthzGrantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrincipalId  string           `protobuf:"bytes,1,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
	ResourceId   string           `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceType string           `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Attrs        *structpb.Struct `protobuf:"bytes,4,opt,name=attrs,proto3,oneof" json:"attrs,omitempty"`
}

func (x *AuthzGrantRequest) Reset() {
	*x = AuthzGrantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sentium_api_v1_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthzGrantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthzGrantRequest) ProtoMessage() {}

func (x *AuthzGrantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sentium_api_v1_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthzGrantRequest.ProtoReflect.Descriptor instead.
func (*AuthzGrantRequest) Descriptor() ([]byte, []int) {
	return file_sentium_api_v1_authz_proto_rawDescGZIP(), []int{2}
}

func (x *AuthzGrantRequest) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

func (x *AuthzGrantRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AuthzGrantRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *AuthzGrantRequest) GetAttrs() *structpb.Struct {
	if x != nil {
		return x.Attrs
	}
	return nil
}

type AuthzGrantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthzGrantResponse) Reset() {
	*x = AuthzGrantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sentium_api_v1_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthzGrantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthzGrantResponse) ProtoMessage() {}

func (x *AuthzGrantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sentium_api_v1_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthzGrantResponse.ProtoReflect.Descriptor instead.
func (*AuthzGrantResponse) Descriptor() ([]byte, []int) {
	return file_sentium_api_v1_authz_proto_rawDescGZIP(), []int{3}
}

type AuthzRevokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrincipalId  string `protobuf:"bytes,1,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
	ResourceId   string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
}

func (x *AuthzRevokeRequest) Reset() {
	*x = AuthzRevokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sentium_api_v1_authz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthzRevokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthzRevokeRequest) ProtoMessage() {}

func (x *AuthzRevokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sentium_api_v1_authz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthzRevokeRequest.ProtoReflect.Descriptor instead.
func (*AuthzRevokeRequest) Descriptor() ([]byte, []int) {
	return file_sentium_api_v1_authz_proto_rawDescGZIP(), []int{4}
}

func (x *AuthzRevokeRequest) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

func (x *AuthzRevokeRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AuthzRevokeRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

type AuthzRevokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthzRevokeResponse) Reset() {
	*x = AuthzRevokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sentium_api_v1_authz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthzRevokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthzRevokeResponse) ProtoMessage() {}

func (x *AuthzRevokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sentium_api_v1_authz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthzRevokeResponse.ProtoReflect.Descriptor instead.
func (*AuthzRevokeResponse) Descriptor() ([]byte, []int) {
	return file_sentium_api_v1_authz_proto_rawDescGZIP(), []int{5}
}

var File_sentium_api_v1_authz_proto protoreflect.FileDescriptor

var file_sentium_api_v1_authz_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x65,
	0x6e, 0x74, 0x69, 0x75, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68,
	0x7a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x62, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x32, 0x0a, 0x05,
	0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x11, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x61, 0x74, 0x74,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x48, 0x00, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7d, 0x0a,
	0x12, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x15, 0x0a, 0x13,
	0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xcf, 0x02, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x6a, 0x0a,
	0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x75, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x6e, 0x74,
	0x69, 0x75, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x3a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x6a, 0x0a, 0x05, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x75, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x75, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x3a,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x6e, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12,
	0x22, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x75, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x75, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x72,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sentium_api_v1_authz_proto_rawDescOnce sync.Once
	file_sentium_api_v1_authz_proto_rawDescData = file_sentium_api_v1_authz_proto_rawDesc
)

func file_sentium_api_v1_authz_proto_rawDescGZIP() []byte {
	file_sentium_api_v1_authz_proto_rawDescOnce.Do(func() {
		file_sentium_api_v1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_sentium_api_v1_authz_proto_rawDescData)
	})
	return file_sentium_api_v1_authz_proto_rawDescData
}

var file_sentium_api_v1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sentium_api_v1_authz_proto_goTypes = []interface{}{
	(*AuthzCheckRequest)(nil),   // 0: sentium.api.v1.AuthzCheckRequest
	(*AuthzCheckResponse)(nil),  // 1: sentium.api.v1.AuthzCheckResponse
	(*AuthzGrantRequest)(nil),   // 2: sentium.api.v1.AuthzGrantRequest
	(*AuthzGrantResponse)(nil),  // 3: sentium.api.v1.AuthzGrantResponse
	(*AuthzRevokeRequest)(nil),  // 4: sentium.api.v1.AuthzRevokeRequest
	(*AuthzRevokeResponse)(nil), // 5: sentium.api.v1.AuthzRevokeResponse
	(*structpb.Struct)(nil),     // 6: google.protobuf.Struct
}
var file_sentium_api_v1_authz_proto_depIdxs = []int32{
	6, // 0: sentium.api.v1.AuthzCheckResponse.attrs:type_name -> google.protobuf.Struct
	6, // 1: sentium.api.v1.AuthzGrantRequest.attrs:type_name -> google.protobuf.Struct
	0, // 2: sentium.api.v1.Authz.Check:input_type -> sentium.api.v1.AuthzCheckRequest
	2, // 3: sentium.api.v1.Authz.Grant:input_type -> sentium.api.v1.AuthzGrantRequest
	4, // 4: sentium.api.v1.Authz.Revoke:input_type -> sentium.api.v1.AuthzRevokeRequest
	1, // 5: sentium.api.v1.Authz.Check:output_type -> sentium.api.v1.AuthzCheckResponse
	3, // 6: sentium.api.v1.Authz.Grant:output_type -> sentium.api.v1.AuthzGrantResponse
	5, // 7: sentium.api.v1.Authz.Revoke:output_type -> sentium.api.v1.AuthzRevokeResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sentium_api_v1_authz_proto_init() }
func file_sentium_api_v1_authz_proto_init() {
	if File_sentium_api_v1_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sentium_api_v1_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthzCheckRequest); i {
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
		file_sentium_api_v1_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthzCheckResponse); i {
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
		file_sentium_api_v1_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthzGrantRequest); i {
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
		file_sentium_api_v1_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthzGrantResponse); i {
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
		file_sentium_api_v1_authz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthzRevokeRequest); i {
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
		file_sentium_api_v1_authz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthzRevokeResponse); i {
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
	file_sentium_api_v1_authz_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_sentium_api_v1_authz_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sentium_api_v1_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sentium_api_v1_authz_proto_goTypes,
		DependencyIndexes: file_sentium_api_v1_authz_proto_depIdxs,
		MessageInfos:      file_sentium_api_v1_authz_proto_msgTypes,
	}.Build()
	File_sentium_api_v1_authz_proto = out.File
	file_sentium_api_v1_authz_proto_rawDesc = nil
	file_sentium_api_v1_authz_proto_goTypes = nil
	file_sentium_api_v1_authz_proto_depIdxs = nil
}
