// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: cargo/v1/service.proto

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

var File_cargo_v1_service_proto protoreflect.FileDescriptor

const file_cargo_v1_service_proto_rawDesc = "" +
	"\n" +
	"\x16cargo/v1/service.proto\x12\bcargo.v1\x1a\x14cargo/v1/cargo.proto2S\n" +
	"\fCargoService\x12C\n" +
	"\bGetCargo\x12\x19.cargo.v1.GetCargoRequest\x1a\x1a.cargo.v1.GetCargoResponse\"\x00B)Z'github.com/tbe-team/raybot-api/cargo/v1b\x06proto3"

var file_cargo_v1_service_proto_goTypes = []any{
	(*GetCargoRequest)(nil),  // 0: cargo.v1.GetCargoRequest
	(*GetCargoResponse)(nil), // 1: cargo.v1.GetCargoResponse
}
var file_cargo_v1_service_proto_depIdxs = []int32{
	0, // 0: cargo.v1.CargoService.GetCargo:input_type -> cargo.v1.GetCargoRequest
	1, // 1: cargo.v1.CargoService.GetCargo:output_type -> cargo.v1.GetCargoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cargo_v1_service_proto_init() }
func file_cargo_v1_service_proto_init() {
	if File_cargo_v1_service_proto != nil {
		return
	}
	file_cargo_v1_cargo_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cargo_v1_service_proto_rawDesc), len(file_cargo_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cargo_v1_service_proto_goTypes,
		DependencyIndexes: file_cargo_v1_service_proto_depIdxs,
	}.Build()
	File_cargo_v1_service_proto = out.File
	file_cargo_v1_service_proto_goTypes = nil
	file_cargo_v1_service_proto_depIdxs = nil
}
