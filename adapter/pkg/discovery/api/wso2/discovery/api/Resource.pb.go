//  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
//  WSO2 Inc. licenses this file to you under the Apache License,
//  Version 2.0 (the "License"); you may not use this file except
//  in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//  KIND, either express or implied.  See the License for the
//  specific language governing permissions and limitations
//  under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/api/Resource.proto

package api

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

// Resource config model
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                string            `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Methods             []*Operation      `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
	Summary             string            `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	Description         string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	ProductionEndpoints *EndpointCluster  `protobuf:"bytes,6,opt,name=productionEndpoints,proto3" json:"productionEndpoints,omitempty"`
	SandboxEndpoints    *EndpointCluster  `protobuf:"bytes,7,opt,name=sandboxEndpoints,proto3" json:"sandboxEndpoints,omitempty"`
	Security            map[string]string `protobuf:"bytes,8,rep,name=security,proto3" json:"security,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Consumes            []string          `protobuf:"bytes,9,rep,name=consumes,proto3" json:"consumes,omitempty"`
	Schemes             []string          `protobuf:"bytes,10,rep,name=schemes,proto3" json:"schemes,omitempty"`
	Tags                []string          `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_Resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_Resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_Resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Resource) GetMethods() []*Operation {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Resource) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Resource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Resource) GetProductionEndpoints() *EndpointCluster {
	if x != nil {
		return x.ProductionEndpoints
	}
	return nil
}

func (x *Resource) GetSandboxEndpoints() *EndpointCluster {
	if x != nil {
		return x.SandboxEndpoints
	}
	return nil
}

func (x *Resource) GetSecurity() map[string]string {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Resource) GetConsumes() []string {
	if x != nil {
		return x.Consumes
	}
	return nil
}

func (x *Resource) GetSchemes() []string {
	if x != nil {
		return x.Schemes
	}
	return nil
}

func (x *Resource) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Operation model which maps to a particular http methods
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method          string             `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Security        []*SecurityList    `protobuf:"bytes,2,rep,name=security,proto3" json:"security,omitempty"`
	Tier            string             `protobuf:"bytes,3,opt,name=tier,proto3" json:"tier,omitempty"`
	DisableSecurity bool               `protobuf:"varint,4,opt,name=disableSecurity,proto3" json:"disableSecurity,omitempty"`
	Policies        *OperationPolicies `protobuf:"bytes,5,opt,name=policies,proto3" json:"policies,omitempty"`
	MockedApiConfig *MockedApiConfig   `protobuf:"bytes,6,opt,name=mockedApiConfig,proto3" json:"mockedApiConfig,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_Resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_Resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_Resource_proto_rawDescGZIP(), []int{1}
}

func (x *Operation) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Operation) GetSecurity() []*SecurityList {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Operation) GetTier() string {
	if x != nil {
		return x.Tier
	}
	return ""
}

func (x *Operation) GetDisableSecurity() bool {
	if x != nil {
		return x.DisableSecurity
	}
	return false
}

func (x *Operation) GetPolicies() *OperationPolicies {
	if x != nil {
		return x.Policies
	}
	return nil
}

func (x *Operation) GetMockedApiConfig() *MockedApiConfig {
	if x != nil {
		return x.MockedApiConfig
	}
	return nil
}

// OperationPolicies holds policies of the APIM operations
type OperationPolicies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  []*Policy `protobuf:"bytes,1,rep,name=request,proto3" json:"request,omitempty"`
	Response []*Policy `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
	Fault    []*Policy `protobuf:"bytes,3,rep,name=fault,proto3" json:"fault,omitempty"`
}

func (x *OperationPolicies) Reset() {
	*x = OperationPolicies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_Resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationPolicies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationPolicies) ProtoMessage() {}

func (x *OperationPolicies) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_Resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationPolicies.ProtoReflect.Descriptor instead.
func (*OperationPolicies) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_Resource_proto_rawDescGZIP(), []int{2}
}

func (x *OperationPolicies) GetRequest() []*Policy {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *OperationPolicies) GetResponse() []*Policy {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *OperationPolicies) GetFault() []*Policy {
	if x != nil {
		return x.Fault
	}
	return nil
}

// Policy holds APIM policies
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyName string            `protobuf:"bytes,1,opt,name=policyName,proto3" json:"policyName,omitempty"`
	Action     string            `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Order      uint32            `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
	Parameters map[string]string `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_Resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_Resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_Resource_proto_rawDescGZIP(), []int{3}
}

func (x *Policy) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *Policy) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Policy) GetOrder() uint32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *Policy) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_wso2_discovery_api_Resource_proto protoreflect.FileDescriptor

var file_wso2_discovery_api_Resource_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x29, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x77, 0x73,
	0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x04, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x07, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55,
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77, 0x73,
	0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x10, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x10, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb1, 0x02, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x28, 0x0a,
	0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x77, 0x73, 0x6f, 0x32,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0f, 0x6d, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41,
	0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x6d, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xb3, 0x01, 0x0a, 0x11, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22,
	0xe1, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x77, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x63, 0x68, 0x6f, 0x72, 0x65, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0d, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_api_Resource_proto_rawDescOnce sync.Once
	file_wso2_discovery_api_Resource_proto_rawDescData = file_wso2_discovery_api_Resource_proto_rawDesc
)

func file_wso2_discovery_api_Resource_proto_rawDescGZIP() []byte {
	file_wso2_discovery_api_Resource_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_api_Resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_api_Resource_proto_rawDescData)
	})
	return file_wso2_discovery_api_Resource_proto_rawDescData
}

var file_wso2_discovery_api_Resource_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wso2_discovery_api_Resource_proto_goTypes = []interface{}{
	(*Resource)(nil),          // 0: wso2.discovery.api.Resource
	(*Operation)(nil),         // 1: wso2.discovery.api.Operation
	(*OperationPolicies)(nil), // 2: wso2.discovery.api.OperationPolicies
	(*Policy)(nil),            // 3: wso2.discovery.api.Policy
	nil,                       // 4: wso2.discovery.api.Resource.SecurityEntry
	nil,                       // 5: wso2.discovery.api.Policy.ParametersEntry
	(*EndpointCluster)(nil),   // 6: wso2.discovery.api.EndpointCluster
	(*SecurityList)(nil),      // 7: wso2.discovery.api.SecurityList
	(*MockedApiConfig)(nil),   // 8: wso2.discovery.api.MockedApiConfig
}
var file_wso2_discovery_api_Resource_proto_depIdxs = []int32{
	1,  // 0: wso2.discovery.api.Resource.methods:type_name -> wso2.discovery.api.Operation
	6,  // 1: wso2.discovery.api.Resource.productionEndpoints:type_name -> wso2.discovery.api.EndpointCluster
	6,  // 2: wso2.discovery.api.Resource.sandboxEndpoints:type_name -> wso2.discovery.api.EndpointCluster
	4,  // 3: wso2.discovery.api.Resource.security:type_name -> wso2.discovery.api.Resource.SecurityEntry
	7,  // 4: wso2.discovery.api.Operation.security:type_name -> wso2.discovery.api.SecurityList
	2,  // 5: wso2.discovery.api.Operation.policies:type_name -> wso2.discovery.api.OperationPolicies
	8,  // 6: wso2.discovery.api.Operation.mockedApiConfig:type_name -> wso2.discovery.api.MockedApiConfig
	3,  // 7: wso2.discovery.api.OperationPolicies.request:type_name -> wso2.discovery.api.Policy
	3,  // 8: wso2.discovery.api.OperationPolicies.response:type_name -> wso2.discovery.api.Policy
	3,  // 9: wso2.discovery.api.OperationPolicies.fault:type_name -> wso2.discovery.api.Policy
	5,  // 10: wso2.discovery.api.Policy.parameters:type_name -> wso2.discovery.api.Policy.ParametersEntry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_wso2_discovery_api_Resource_proto_init() }
func file_wso2_discovery_api_Resource_proto_init() {
	if File_wso2_discovery_api_Resource_proto != nil {
		return
	}
	file_wso2_discovery_api_endpoint_cluster_proto_init()
	file_wso2_discovery_api_security_scheme_proto_init()
	file_wso2_discovery_api_mocked_api_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_api_Resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_wso2_discovery_api_Resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
		file_wso2_discovery_api_Resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationPolicies); i {
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
		file_wso2_discovery_api_Resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
			RawDescriptor: file_wso2_discovery_api_Resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_api_Resource_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_api_Resource_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_api_Resource_proto_msgTypes,
	}.Build()
	File_wso2_discovery_api_Resource_proto = out.File
	file_wso2_discovery_api_Resource_proto_rawDesc = nil
	file_wso2_discovery_api_Resource_proto_goTypes = nil
	file_wso2_discovery_api_Resource_proto_depIdxs = nil
}
