// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: distancesensor/v1/service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_distancesensor_v1_service_proto protoreflect.FileDescriptor

const file_distancesensor_v1_service_proto_rawDesc = "" +
	"\n" +
	"\x1fdistancesensor/v1/service.proto\x12\x11distancesensor.v1\x1a'distancesensor/v1/distance_sensor.proto2\x89\x01\n" +
	"\x15DistanceSensorService\x12p\n" +
	"\x11GetDistanceSensor\x12+.distancesensor.v1.GetDistanceSensorRequest\x1a,.distancesensor.v1.GetDistanceSensorResponse\"\x00B2Z0github.com/tbe-team/raybot-api/distancesensor/v1b\x06proto3"

var file_distancesensor_v1_service_proto_goTypes = []any{
	(*GetDistanceSensorRequest)(nil),  // 0: distancesensor.v1.GetDistanceSensorRequest
	(*GetDistanceSensorResponse)(nil), // 1: distancesensor.v1.GetDistanceSensorResponse
}
var file_distancesensor_v1_service_proto_depIdxs = []int32{
	0, // 0: distancesensor.v1.DistanceSensorService.GetDistanceSensor:input_type -> distancesensor.v1.GetDistanceSensorRequest
	1, // 1: distancesensor.v1.DistanceSensorService.GetDistanceSensor:output_type -> distancesensor.v1.GetDistanceSensorResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_distancesensor_v1_service_proto_init() }
func file_distancesensor_v1_service_proto_init() {
	if File_distancesensor_v1_service_proto != nil {
		return
	}
	file_distancesensor_v1_distance_sensor_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_distancesensor_v1_service_proto_rawDesc), len(file_distancesensor_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_distancesensor_v1_service_proto_goTypes,
		DependencyIndexes: file_distancesensor_v1_service_proto_depIdxs,
	}.Build()
	File_distancesensor_v1_service_proto = out.File
	file_distancesensor_v1_service_proto_goTypes = nil
	file_distancesensor_v1_service_proto_depIdxs = nil
}
