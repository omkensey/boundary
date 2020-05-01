// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: controller/api/resources/hosts/v1/static_host.proto

package hosts

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// StaticHost contains all fields related to a StaticHost resource
type StaticHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Canonical path of the resource from the API's base URI
	// Output only.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The type of the resource, to help differentiate schemas
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Friendly name, if set
	FriendlyName *wrappers.StringValue `protobuf:"bytes,3,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	// The time this host was created
	// Ouput only.
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// The time this host was last updated
	// Output only.
	UpdatedTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// Whether the host is disabled
	Disabled *wrappers.BoolValue `protobuf:"bytes,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// The address (DNS or IP name) used to reach the host
	Address *wrappers.StringValue `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *StaticHost) Reset() {
	*x = StaticHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_hosts_v1_static_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticHost) ProtoMessage() {}

func (x *StaticHost) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_hosts_v1_static_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticHost.ProtoReflect.Descriptor instead.
func (*StaticHost) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_hosts_v1_static_host_proto_rawDescGZIP(), []int{0}
}

func (x *StaticHost) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *StaticHost) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StaticHost) GetFriendlyName() *wrappers.StringValue {
	if x != nil {
		return x.FriendlyName
	}
	return nil
}

func (x *StaticHost) GetCreatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *StaticHost) GetUpdatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *StaticHost) GetDisabled() *wrappers.BoolValue {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (x *StaticHost) GetAddress() *wrappers.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_controller_api_resources_hosts_v1_static_host_proto protoreflect.FileDescriptor

var file_controller_api_resources_hosts_v1_static_host_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x41, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x74,
	0x6f, 0x77, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x3b, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resources_hosts_v1_static_host_proto_rawDescOnce sync.Once
	file_controller_api_resources_hosts_v1_static_host_proto_rawDescData = file_controller_api_resources_hosts_v1_static_host_proto_rawDesc
)

func file_controller_api_resources_hosts_v1_static_host_proto_rawDescGZIP() []byte {
	file_controller_api_resources_hosts_v1_static_host_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_hosts_v1_static_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_hosts_v1_static_host_proto_rawDescData)
	})
	return file_controller_api_resources_hosts_v1_static_host_proto_rawDescData
}

var file_controller_api_resources_hosts_v1_static_host_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_api_resources_hosts_v1_static_host_proto_goTypes = []interface{}{
	(*StaticHost)(nil),           // 0: controller.api.resources.hosts.v1.StaticHost
	(*wrappers.StringValue)(nil), // 1: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*wrappers.BoolValue)(nil),   // 3: google.protobuf.BoolValue
}
var file_controller_api_resources_hosts_v1_static_host_proto_depIdxs = []int32{
	1, // 0: controller.api.resources.hosts.v1.StaticHost.friendly_name:type_name -> google.protobuf.StringValue
	2, // 1: controller.api.resources.hosts.v1.StaticHost.created_time:type_name -> google.protobuf.Timestamp
	2, // 2: controller.api.resources.hosts.v1.StaticHost.updated_time:type_name -> google.protobuf.Timestamp
	3, // 3: controller.api.resources.hosts.v1.StaticHost.disabled:type_name -> google.protobuf.BoolValue
	1, // 4: controller.api.resources.hosts.v1.StaticHost.address:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_controller_api_resources_hosts_v1_static_host_proto_init() }
func file_controller_api_resources_hosts_v1_static_host_proto_init() {
	if File_controller_api_resources_hosts_v1_static_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_hosts_v1_static_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticHost); i {
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
			RawDescriptor: file_controller_api_resources_hosts_v1_static_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_hosts_v1_static_host_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_hosts_v1_static_host_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_hosts_v1_static_host_proto_msgTypes,
	}.Build()
	File_controller_api_resources_hosts_v1_static_host_proto = out.File
	file_controller_api_resources_hosts_v1_static_host_proto_rawDesc = nil
	file_controller_api_resources_hosts_v1_static_host_proto_goTypes = nil
	file_controller_api_resources_hosts_v1_static_host_proto_depIdxs = nil
}
