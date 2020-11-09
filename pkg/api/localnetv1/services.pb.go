// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: pkg/api/localnetv1/services.proto

package localnetv1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Protocol int32

const (
	Protocol_UnknownProtocol Protocol = 0
	Protocol_TCP             Protocol = 1
	Protocol_UDP             Protocol = 2
	Protocol_SCTP            Protocol = 3
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "UnknownProtocol",
		1: "TCP",
		2: "UDP",
		3: "SCTP",
	}
	Protocol_value = map[string]int32{
		"UnknownProtocol": 0,
		"TCP":             1,
		"UDP":             2,
		"SCTP":            3,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_api_localnetv1_services_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_pkg_api_localnetv1_services_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{0}
}

type NextFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique instance ID to manage proxy restarts
	InstanceID uint64 `protobuf:"varint,1,opt,name=InstanceID,proto3" json:"InstanceID,omitempty"`
	// The latest revision we're aware of (0 at first)
	Rev uint64 `protobuf:"varint,2,opt,name=Rev,proto3" json:"Rev,omitempty"`
	// Filter endpoints by requiring conditions
	RequiredEndpointConditions *EndpointConditions `protobuf:"bytes,3,opt,name=RequiredEndpointConditions,proto3" json:"RequiredEndpointConditions,omitempty"`
}

func (x *NextFilter) Reset() {
	*x = NextFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextFilter) ProtoMessage() {}

func (x *NextFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextFilter.ProtoReflect.Descriptor instead.
func (*NextFilter) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{0}
}

func (x *NextFilter) GetInstanceID() uint64 {
	if x != nil {
		return x.InstanceID
	}
	return 0
}

func (x *NextFilter) GetRev() uint64 {
	if x != nil {
		return x.Rev
	}
	return 0
}

func (x *NextFilter) GetRequiredEndpointConditions() *EndpointConditions {
	if x != nil {
		return x.RequiredEndpointConditions
	}
	return nil
}

type NextItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filter to use to get the next notification (first item in stream)
	Next *NextFilter `protobuf:"bytes,1,opt,name=Next,proto3" json:"Next,omitempty"`
	// A service endpoints item (any item after the first)
	Endpoints *ServiceEndpoints `protobuf:"bytes,2,opt,name=Endpoints,proto3" json:"Endpoints,omitempty"`
}

func (x *NextItem) Reset() {
	*x = NextItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextItem) ProtoMessage() {}

func (x *NextItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextItem.ProtoReflect.Descriptor instead.
func (*NextItem) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{1}
}

func (x *NextItem) GetNext() *NextFilter {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *NextItem) GetEndpoints() *ServiceEndpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

type ServiceEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string      `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type      string      `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	IPs       *ServiceIPs `protobuf:"bytes,4,opt,name=IPs,proto3" json:"IPs,omitempty"`
	// true if the service maps the whole IP, not just individual ports.
	MapAll bool `protobuf:"varint,5,opt,name=MapAll,proto3" json:"MapAll,omitempty"`
	// Individual ports mapped for the this service
	Ports []*PortMapping `protobuf:"bytes,6,rep,name=Ports,proto3" json:"Ports,omitempty"`
	// All possible endpoints for this service
	Endpoints              []*Endpoint `protobuf:"bytes,7,rep,name=Endpoints,proto3" json:"Endpoints,omitempty"`
	ExternalTrafficToLocal bool        `protobuf:"varint,8,opt,name=ExternalTrafficToLocal,proto3" json:"ExternalTrafficToLocal,omitempty"`
}

func (x *ServiceEndpoints) Reset() {
	*x = ServiceEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceEndpoints) ProtoMessage() {}

func (x *ServiceEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceEndpoints.ProtoReflect.Descriptor instead.
func (*ServiceEndpoints) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceEndpoints) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ServiceEndpoints) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceEndpoints) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServiceEndpoints) GetIPs() *ServiceIPs {
	if x != nil {
		return x.IPs
	}
	return nil
}

func (x *ServiceEndpoints) GetMapAll() bool {
	if x != nil {
		return x.MapAll
	}
	return false
}

func (x *ServiceEndpoints) GetPorts() []*PortMapping {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *ServiceEndpoints) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ServiceEndpoints) GetExternalTrafficToLocal() bool {
	if x != nil {
		return x.ExternalTrafficToLocal
	}
	return false
}

type ServiceIPs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterIP   string `protobuf:"bytes,1,opt,name=ClusterIP,proto3" json:"ClusterIP,omitempty"`
	ExternalIPs *IPSet `protobuf:"bytes,2,opt,name=ExternalIPs,proto3" json:"ExternalIPs,omitempty"`
}

func (x *ServiceIPs) Reset() {
	*x = ServiceIPs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceIPs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceIPs) ProtoMessage() {}

func (x *ServiceIPs) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceIPs.ProtoReflect.Descriptor instead.
func (*ServiceIPs) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceIPs) GetClusterIP() string {
	if x != nil {
		return x.ClusterIP
	}
	return ""
}

func (x *ServiceIPs) GetExternalIPs() *IPSet {
	if x != nil {
		return x.ExternalIPs
	}
	return nil
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname   string              `protobuf:"bytes,1,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	Conditions *EndpointConditions `protobuf:"bytes,2,opt,name=Conditions,proto3" json:"Conditions,omitempty"`
	IPs        *IPSet              `protobuf:"bytes,3,opt,name=IPs,proto3" json:"IPs,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{4}
}

func (x *Endpoint) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Endpoint) GetConditions() *EndpointConditions {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Endpoint) GetIPs() *IPSet {
	if x != nil {
		return x.IPs
	}
	return nil
}

type IPSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V4 []string `protobuf:"bytes,1,rep,name=V4,proto3" json:"V4,omitempty"`
	V6 []string `protobuf:"bytes,2,rep,name=V6,proto3" json:"V6,omitempty"`
}

func (x *IPSet) Reset() {
	*x = IPSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPSet) ProtoMessage() {}

func (x *IPSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPSet.ProtoReflect.Descriptor instead.
func (*IPSet) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{5}
}

func (x *IPSet) GetV4() []string {
	if x != nil {
		return x.V4
	}
	return nil
}

func (x *IPSet) GetV6() []string {
	if x != nil {
		return x.V6
	}
	return nil
}

type EndpointConditions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready    bool `protobuf:"varint,1,opt,name=Ready,proto3" json:"Ready,omitempty"`
	Selected bool `protobuf:"varint,2,opt,name=Selected,proto3" json:"Selected,omitempty"`
	Local    bool `protobuf:"varint,3,opt,name=Local,proto3" json:"Local,omitempty"`
}

func (x *EndpointConditions) Reset() {
	*x = EndpointConditions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointConditions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointConditions) ProtoMessage() {}

func (x *EndpointConditions) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointConditions.ProtoReflect.Descriptor instead.
func (*EndpointConditions) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{6}
}

func (x *EndpointConditions) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

func (x *EndpointConditions) GetSelected() bool {
	if x != nil {
		return x.Selected
	}
	return false
}

func (x *EndpointConditions) GetLocal() bool {
	if x != nil {
		return x.Local
	}
	return false
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Protocol Protocol `protobuf:"varint,2,opt,name=Protocol,proto3,enum=localnetv1.Protocol" json:"Protocol,omitempty"`
	Port     int32    `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{7}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_UnknownProtocol
}

func (x *Port) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type PortMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Protocol   Protocol `protobuf:"varint,2,opt,name=Protocol,proto3,enum=localnetv1.Protocol" json:"Protocol,omitempty"`
	Port       int32    `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
	NodePort   int32    `protobuf:"varint,4,opt,name=NodePort,proto3" json:"NodePort,omitempty"`
	TargetPort int32    `protobuf:"varint,5,opt,name=TargetPort,proto3" json:"TargetPort,omitempty"`
}

func (x *PortMapping) Reset() {
	*x = PortMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_localnetv1_services_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMapping) ProtoMessage() {}

func (x *PortMapping) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_localnetv1_services_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMapping.ProtoReflect.Descriptor instead.
func (*PortMapping) Descriptor() ([]byte, []int) {
	return file_pkg_api_localnetv1_services_proto_rawDescGZIP(), []int{8}
}

func (x *PortMapping) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PortMapping) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_UnknownProtocol
}

func (x *PortMapping) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PortMapping) GetNodePort() int32 {
	if x != nil {
		return x.NodePort
	}
	return 0
}

func (x *PortMapping) GetTargetPort() int32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

var File_pkg_api_localnetv1_services_proto protoreflect.FileDescriptor

var file_pkg_api_localnetv1_services_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e,
	0x65, 0x74, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x22,
	0x9e, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x10,
	0x0a, 0x03, 0x52, 0x65, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x52, 0x65, 0x76,
	0x12, 0x5e, 0x0a, 0x1a, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x1a, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x72, 0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2a, 0x0a, 0x04,
	0x4e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x09, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x28, 0x0a, 0x03, 0x49, 0x50, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x50, 0x73, 0x52, 0x03, 0x49, 0x50, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x70,
	0x41, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4d, 0x61, 0x70, 0x41, 0x6c,
	0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x32, 0x0a, 0x09, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0x5f, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x50, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x50, 0x12, 0x33, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x53, 0x65, 0x74,
	0x52, 0x0b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x73, 0x22, 0x8b, 0x01,
	0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x49, 0x50, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31,
	0x2e, 0x49, 0x50, 0x53, 0x65, 0x74, 0x52, 0x03, 0x49, 0x50, 0x73, 0x22, 0x27, 0x0a, 0x05, 0x49,
	0x50, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x56, 0x34, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x56, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x56, 0x36, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x56, 0x36, 0x22, 0x5c, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65,
	0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x22, 0x60, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x2a, 0x3b, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54,
	0x43, 0x50, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x43, 0x54, 0x50, 0x10, 0x03, 0x32, 0x43, 0x0a, 0x09, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x16, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x6c, 0x75, 0x73,
	0x65, 0x61, 0x75, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x32, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6e, 0x65, 0x74,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_localnetv1_services_proto_rawDescOnce sync.Once
	file_pkg_api_localnetv1_services_proto_rawDescData = file_pkg_api_localnetv1_services_proto_rawDesc
)

func file_pkg_api_localnetv1_services_proto_rawDescGZIP() []byte {
	file_pkg_api_localnetv1_services_proto_rawDescOnce.Do(func() {
		file_pkg_api_localnetv1_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_localnetv1_services_proto_rawDescData)
	})
	return file_pkg_api_localnetv1_services_proto_rawDescData
}

var file_pkg_api_localnetv1_services_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_api_localnetv1_services_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_api_localnetv1_services_proto_goTypes = []interface{}{
	(Protocol)(0),              // 0: localnetv1.Protocol
	(*NextFilter)(nil),         // 1: localnetv1.NextFilter
	(*NextItem)(nil),           // 2: localnetv1.NextItem
	(*ServiceEndpoints)(nil),   // 3: localnetv1.ServiceEndpoints
	(*ServiceIPs)(nil),         // 4: localnetv1.ServiceIPs
	(*Endpoint)(nil),           // 5: localnetv1.Endpoint
	(*IPSet)(nil),              // 6: localnetv1.IPSet
	(*EndpointConditions)(nil), // 7: localnetv1.EndpointConditions
	(*Port)(nil),               // 8: localnetv1.Port
	(*PortMapping)(nil),        // 9: localnetv1.PortMapping
}
var file_pkg_api_localnetv1_services_proto_depIdxs = []int32{
	7,  // 0: localnetv1.NextFilter.RequiredEndpointConditions:type_name -> localnetv1.EndpointConditions
	1,  // 1: localnetv1.NextItem.Next:type_name -> localnetv1.NextFilter
	3,  // 2: localnetv1.NextItem.Endpoints:type_name -> localnetv1.ServiceEndpoints
	4,  // 3: localnetv1.ServiceEndpoints.IPs:type_name -> localnetv1.ServiceIPs
	9,  // 4: localnetv1.ServiceEndpoints.Ports:type_name -> localnetv1.PortMapping
	5,  // 5: localnetv1.ServiceEndpoints.Endpoints:type_name -> localnetv1.Endpoint
	6,  // 6: localnetv1.ServiceIPs.ExternalIPs:type_name -> localnetv1.IPSet
	7,  // 7: localnetv1.Endpoint.Conditions:type_name -> localnetv1.EndpointConditions
	6,  // 8: localnetv1.Endpoint.IPs:type_name -> localnetv1.IPSet
	0,  // 9: localnetv1.Port.Protocol:type_name -> localnetv1.Protocol
	0,  // 10: localnetv1.PortMapping.Protocol:type_name -> localnetv1.Protocol
	1,  // 11: localnetv1.Endpoints.Next:input_type -> localnetv1.NextFilter
	2,  // 12: localnetv1.Endpoints.Next:output_type -> localnetv1.NextItem
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pkg_api_localnetv1_services_proto_init() }
func file_pkg_api_localnetv1_services_proto_init() {
	if File_pkg_api_localnetv1_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_localnetv1_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextFilter); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextItem); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceEndpoints); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceIPs); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPSet); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointConditions); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_pkg_api_localnetv1_services_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMapping); i {
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
			RawDescriptor: file_pkg_api_localnetv1_services_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_localnetv1_services_proto_goTypes,
		DependencyIndexes: file_pkg_api_localnetv1_services_proto_depIdxs,
		EnumInfos:         file_pkg_api_localnetv1_services_proto_enumTypes,
		MessageInfos:      file_pkg_api_localnetv1_services_proto_msgTypes,
	}.Build()
	File_pkg_api_localnetv1_services_proto = out.File
	file_pkg_api_localnetv1_services_proto_rawDesc = nil
	file_pkg_api_localnetv1_services_proto_goTypes = nil
	file_pkg_api_localnetv1_services_proto_depIdxs = nil
}
