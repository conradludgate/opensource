//*
// API to manage token service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: priv/tokenservice/v1alpha1/tokenservice.proto

package tokenservicev1alpha1

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

type GetAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetAccessTokenRequest) Reset() {
	*x = GetAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenRequest) ProtoMessage() {}

func (x *GetAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*GetAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccessTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *GetAccessTokenResponse) Reset() {
	*x = GetAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenResponse) ProtoMessage() {}

func (x *GetAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*GetAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccessTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type CreatePATRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePATRequest) Reset() {
	*x = CreatePATRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePATRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePATRequest) ProtoMessage() {}

func (x *CreatePATRequest) ProtoReflect() protoreflect.Message {
	mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePATRequest.ProtoReflect.Descriptor instead.
func (*CreatePATRequest) Descriptor() ([]byte, []int) {
	return file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescGZIP(), []int{2}
}

type CreatePATResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pat string `protobuf:"bytes,2,opt,name=pat,proto3" json:"pat,omitempty"`
}

func (x *CreatePATResponse) Reset() {
	*x = CreatePATResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePATResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePATResponse) ProtoMessage() {}

func (x *CreatePATResponse) ProtoReflect() protoreflect.Message {
	mi := &file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePATResponse.ProtoReflect.Descriptor instead.
func (*CreatePATResponse) Descriptor() ([]byte, []int) {
	return file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePATResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreatePATResponse) GetPat() string {
	if x != nil {
		return x.Pat
	}
	return ""
}

var File_priv_tokenservice_v1alpha1_tokenservice_proto protoreflect.FileDescriptor

var file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x2d, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x41, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x41, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x61, 0x74, 0x32, 0xf8, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02,
	0x01, 0x12, 0x6a, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x41, 0x54, 0x12, 0x2c,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x41, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70,
	0x72, 0x69, 0x76, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x41, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x88, 0x02,
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x2e, 0x6a, 0x65, 0x74, 0x70, 0x61, 0x63,
	0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x50, 0x54, 0x58, 0xaa, 0x02, 0x1a, 0x50, 0x72, 0x69, 0x76, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x50, 0x72, 0x69, 0x76, 0x5c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x26, 0x50, 0x72, 0x69, 0x76, 0x5c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x50, 0x72, 0x69, 0x76,
	0x3a, 0x3a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescOnce sync.Once
	file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescData = file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDesc
)

func file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescGZIP() []byte {
	file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescOnce.Do(func() {
		file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescData)
	})
	return file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDescData
}

var file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_priv_tokenservice_v1alpha1_tokenservice_proto_goTypes = []interface{}{
	(*GetAccessTokenRequest)(nil),  // 0: priv.tokenservice.v1alpha1.GetAccessTokenRequest
	(*GetAccessTokenResponse)(nil), // 1: priv.tokenservice.v1alpha1.GetAccessTokenResponse
	(*CreatePATRequest)(nil),       // 2: priv.tokenservice.v1alpha1.CreatePATRequest
	(*CreatePATResponse)(nil),      // 3: priv.tokenservice.v1alpha1.CreatePATResponse
}
var file_priv_tokenservice_v1alpha1_tokenservice_proto_depIdxs = []int32{
	0, // 0: priv.tokenservice.v1alpha1.TokenService.GetAccessToken:input_type -> priv.tokenservice.v1alpha1.GetAccessTokenRequest
	2, // 1: priv.tokenservice.v1alpha1.TokenService.CreatePAT:input_type -> priv.tokenservice.v1alpha1.CreatePATRequest
	1, // 2: priv.tokenservice.v1alpha1.TokenService.GetAccessToken:output_type -> priv.tokenservice.v1alpha1.GetAccessTokenResponse
	3, // 3: priv.tokenservice.v1alpha1.TokenService.CreatePAT:output_type -> priv.tokenservice.v1alpha1.CreatePATResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_priv_tokenservice_v1alpha1_tokenservice_proto_init() }
func file_priv_tokenservice_v1alpha1_tokenservice_proto_init() {
	if File_priv_tokenservice_v1alpha1_tokenservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenRequest); i {
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
		file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessTokenResponse); i {
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
		file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePATRequest); i {
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
		file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePATResponse); i {
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
			RawDescriptor: file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_priv_tokenservice_v1alpha1_tokenservice_proto_goTypes,
		DependencyIndexes: file_priv_tokenservice_v1alpha1_tokenservice_proto_depIdxs,
		MessageInfos:      file_priv_tokenservice_v1alpha1_tokenservice_proto_msgTypes,
	}.Build()
	File_priv_tokenservice_v1alpha1_tokenservice_proto = out.File
	file_priv_tokenservice_v1alpha1_tokenservice_proto_rawDesc = nil
	file_priv_tokenservice_v1alpha1_tokenservice_proto_goTypes = nil
	file_priv_tokenservice_v1alpha1_tokenservice_proto_depIdxs = nil
}