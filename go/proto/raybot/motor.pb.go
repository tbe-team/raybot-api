// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/raybot/motor.proto

package raybot

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

type SetConfigRequest_Direction int32

const (
	SetConfigRequest_DIRECTION_UNSPECIFIED SetConfigRequest_Direction = 0
	SetConfigRequest_DIRECTION_FORWARD     SetConfigRequest_Direction = 1
	SetConfigRequest_DIRECTION_BACKWARD    SetConfigRequest_Direction = 2
)

// Enum value maps for SetConfigRequest_Direction.
var (
	SetConfigRequest_Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "DIRECTION_FORWARD",
		2: "DIRECTION_BACKWARD",
	}
	SetConfigRequest_Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"DIRECTION_FORWARD":     1,
		"DIRECTION_BACKWARD":    2,
	}
)

func (x SetConfigRequest_Direction) Enum() *SetConfigRequest_Direction {
	p := new(SetConfigRequest_Direction)
	*p = x
	return p
}

func (x SetConfigRequest_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetConfigRequest_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_raybot_motor_proto_enumTypes[0].Descriptor()
}

func (SetConfigRequest_Direction) Type() protoreflect.EnumType {
	return &file_proto_raybot_motor_proto_enumTypes[0]
}

func (x SetConfigRequest_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetConfigRequest_Direction.Descriptor instead.
func (SetConfigRequest_Direction) EnumDescriptor() ([]byte, []int) {
	return file_proto_raybot_motor_proto_rawDescGZIP(), []int{0, 0}
}

type SetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed     uint32                     `protobuf:"varint,1,opt,name=speed,proto3" json:"speed,omitempty"`
	Direction SetConfigRequest_Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=motor.SetConfigRequest_Direction" json:"direction,omitempty"`
}

func (x *SetConfigRequest) Reset() {
	*x = SetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raybot_motor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigRequest) ProtoMessage() {}

func (x *SetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raybot_motor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigRequest.ProtoReflect.Descriptor instead.
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_raybot_motor_proto_rawDescGZIP(), []int{0}
}

func (x *SetConfigRequest) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *SetConfigRequest) GetDirection() SetConfigRequest_Direction {
	if x != nil {
		return x.Direction
	}
	return SetConfigRequest_DIRECTION_UNSPECIFIED
}

type SetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetConfigResponse) Reset() {
	*x = SetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raybot_motor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigResponse) ProtoMessage() {}

func (x *SetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raybot_motor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetConfigResponse.ProtoReflect.Descriptor instead.
func (*SetConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_raybot_motor_proto_rawDescGZIP(), []int{1}
}

type StreamStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntervalMs uint32 `protobuf:"varint,1,opt,name=interval_ms,json=intervalMs,proto3" json:"interval_ms,omitempty"` // Thời gian delay giữa các bản tin (ms)
}

func (x *StreamStatusRequest) Reset() {
	*x = StreamStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raybot_motor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamStatusRequest) ProtoMessage() {}

func (x *StreamStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raybot_motor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamStatusRequest.ProtoReflect.Descriptor instead.
func (*StreamStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_raybot_motor_proto_rawDescGZIP(), []int{2}
}

func (x *StreamStatusRequest) GetIntervalMs() uint32 {
	if x != nil {
		return x.IntervalMs
	}
	return 0
}

type StreamStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed  uint32 `protobuf:"varint,1,opt,name=speed,proto3" json:"speed,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // "RUNNING", "STOPPED", etc.
}

func (x *StreamStatusResponse) Reset() {
	*x = StreamStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raybot_motor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamStatusResponse) ProtoMessage() {}

func (x *StreamStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raybot_motor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamStatusResponse.ProtoReflect.Descriptor instead.
func (*StreamStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_raybot_motor_proto_rawDescGZIP(), []int{3}
}

func (x *StreamStatusResponse) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *StreamStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_raybot_motor_proto protoreflect.FileDescriptor

var file_proto_raybot_motor_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x79, 0x62, 0x6f, 0x74, 0x2f, 0x6d,
	0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x74, 0x6f,
	0x72, 0x22, 0xc0, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a,
	0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x57, 0x41,
	0x52, 0x44, 0x10, 0x02, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d,
	0x73, 0x22, 0x44, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9e, 0x01, 0x0a, 0x11, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x4d, 0x6f, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x09, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x2e, 0x6d, 0x6f, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e,
	0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x78, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x42, 0x0a, 0x4d, 0x6f, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x62, 0x65, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x72, 0x61, 0x79, 0x62, 0x6f, 0x74, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x79, 0x62, 0x6f, 0x74,
	0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x4d, 0x6f, 0x74, 0x6f, 0x72, 0xca, 0x02,
	0x05, 0x4d, 0x6f, 0x74, 0x6f, 0x72, 0xe2, 0x02, 0x11, 0x4d, 0x6f, 0x74, 0x6f, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x4d, 0x6f, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_raybot_motor_proto_rawDescOnce sync.Once
	file_proto_raybot_motor_proto_rawDescData = file_proto_raybot_motor_proto_rawDesc
)

func file_proto_raybot_motor_proto_rawDescGZIP() []byte {
	file_proto_raybot_motor_proto_rawDescOnce.Do(func() {
		file_proto_raybot_motor_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_raybot_motor_proto_rawDescData)
	})
	return file_proto_raybot_motor_proto_rawDescData
}

var file_proto_raybot_motor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_raybot_motor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_raybot_motor_proto_goTypes = []interface{}{
	(SetConfigRequest_Direction)(0), // 0: motor.SetConfigRequest.Direction
	(*SetConfigRequest)(nil),        // 1: motor.SetConfigRequest
	(*SetConfigResponse)(nil),       // 2: motor.SetConfigResponse
	(*StreamStatusRequest)(nil),     // 3: motor.StreamStatusRequest
	(*StreamStatusResponse)(nil),    // 4: motor.StreamStatusResponse
}
var file_proto_raybot_motor_proto_depIdxs = []int32{
	0, // 0: motor.SetConfigRequest.direction:type_name -> motor.SetConfigRequest.Direction
	1, // 1: motor.DriveMotorService.SetConfig:input_type -> motor.SetConfigRequest
	3, // 2: motor.DriveMotorService.StreamStatus:input_type -> motor.StreamStatusRequest
	2, // 3: motor.DriveMotorService.SetConfig:output_type -> motor.SetConfigResponse
	4, // 4: motor.DriveMotorService.StreamStatus:output_type -> motor.StreamStatusResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_raybot_motor_proto_init() }
func file_proto_raybot_motor_proto_init() {
	if File_proto_raybot_motor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_raybot_motor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConfigRequest); i {
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
		file_proto_raybot_motor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetConfigResponse); i {
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
		file_proto_raybot_motor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamStatusRequest); i {
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
		file_proto_raybot_motor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamStatusResponse); i {
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
			RawDescriptor: file_proto_raybot_motor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raybot_motor_proto_goTypes,
		DependencyIndexes: file_proto_raybot_motor_proto_depIdxs,
		EnumInfos:         file_proto_raybot_motor_proto_enumTypes,
		MessageInfos:      file_proto_raybot_motor_proto_msgTypes,
	}.Build()
	File_proto_raybot_motor_proto = out.File
	file_proto_raybot_motor_proto_rawDesc = nil
	file_proto_raybot_motor_proto_goTypes = nil
	file_proto_raybot_motor_proto_depIdxs = nil
}
