// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: kuscia/api/filters/http/kuscia_crypt/v3/crypt.proto

package v3

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

type CryptRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source            string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination       string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Algorithm         string `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	SecretKey         string `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	SecretKeyVersion  string `protobuf:"bytes,5,opt,name=secret_key_version,json=secretKeyVersion,proto3" json:"secret_key_version,omitempty"`
	ReserveKey        string `protobuf:"bytes,6,opt,name=reserve_key,json=reserveKey,proto3" json:"reserve_key,omitempty"`
	ReserveKeyVersion string `protobuf:"bytes,7,opt,name=reserve_key_version,json=reserveKeyVersion,proto3" json:"reserve_key_version,omitempty"`
}

func (x *CryptRule) Reset() {
	*x = CryptRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptRule) ProtoMessage() {}

func (x *CryptRule) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptRule.ProtoReflect.Descriptor instead.
func (*CryptRule) Descriptor() ([]byte, []int) {
	return file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescGZIP(), []int{0}
}

func (x *CryptRule) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CryptRule) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *CryptRule) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *CryptRule) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *CryptRule) GetSecretKeyVersion() string {
	if x != nil {
		return x.SecretKeyVersion
	}
	return ""
}

func (x *CryptRule) GetReserveKey() string {
	if x != nil {
		return x.ReserveKey
	}
	return ""
}

func (x *CryptRule) GetReserveKeyVersion() string {
	if x != nil {
		return x.ReserveKeyVersion
	}
	return ""
}

type Crypt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfNamespace string       `protobuf:"bytes,1,opt,name=self_namespace,json=selfNamespace,proto3" json:"self_namespace,omitempty"`
	EncryptRules  []*CryptRule `protobuf:"bytes,2,rep,name=encrypt_rules,json=encryptRules,proto3" json:"encrypt_rules,omitempty"`
	DecryptRules  []*CryptRule `protobuf:"bytes,3,rep,name=decrypt_rules,json=decryptRules,proto3" json:"decrypt_rules,omitempty"`
}

func (x *Crypt) Reset() {
	*x = Crypt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crypt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crypt) ProtoMessage() {}

func (x *Crypt) ProtoReflect() protoreflect.Message {
	mi := &file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crypt.ProtoReflect.Descriptor instead.
func (*Crypt) Descriptor() ([]byte, []int) {
	return file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescGZIP(), []int{1}
}

func (x *Crypt) GetSelfNamespace() string {
	if x != nil {
		return x.SelfNamespace
	}
	return ""
}

func (x *Crypt) GetEncryptRules() []*CryptRule {
	if x != nil {
		return x.EncryptRules
	}
	return nil
}

func (x *Crypt) GetDecryptRules() []*CryptRule {
	if x != nil {
		return x.DecryptRules
	}
	return nil
}

var File_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto protoreflect.FileDescriptor

var file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61,
	0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x5f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x2e, 0x76, 0x33, 0x22, 0x81, 0x02, 0x0a, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x4b, 0x65,
	0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xec, 0x01, 0x0a, 0x05, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x6c, 0x66,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0d, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x5d, 0x0a, 0x0d, 0x64, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2d, 0x69, 0x6e, 0x63, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6b, 0x75, 0x73, 0x63,
	0x69, 0x61, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x6b, 0x75, 0x73, 0x63, 0x69, 0x61, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescOnce sync.Once
	file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescData = file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDesc
)

func file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescGZIP() []byte {
	file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescOnce.Do(func() {
		file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescData)
	})
	return file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDescData
}

var file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_goTypes = []interface{}{
	(*CryptRule)(nil), // 0: envoy.extensions.filters.http.kuscia_crypt.v3.CryptRule
	(*Crypt)(nil),     // 1: envoy.extensions.filters.http.kuscia_crypt.v3.Crypt
}
var file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.filters.http.kuscia_crypt.v3.Crypt.encrypt_rules:type_name -> envoy.extensions.filters.http.kuscia_crypt.v3.CryptRule
	0, // 1: envoy.extensions.filters.http.kuscia_crypt.v3.Crypt.decrypt_rules:type_name -> envoy.extensions.filters.http.kuscia_crypt.v3.CryptRule
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_init() }
func file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_init() {
	if File_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptRule); i {
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
		file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crypt); i {
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
			RawDescriptor: file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_goTypes,
		DependencyIndexes: file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_depIdxs,
		MessageInfos:      file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_msgTypes,
	}.Build()
	File_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto = out.File
	file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_rawDesc = nil
	file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_goTypes = nil
	file_kuscia_api_filters_http_kuscia_crypt_v3_crypt_proto_depIdxs = nil
}
