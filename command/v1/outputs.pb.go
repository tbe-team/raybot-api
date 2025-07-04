// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: command/v1/outputs.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommandOutputs struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Outputs:
	//
	//	*CommandOutputs_Stop
	//	*CommandOutputs_MoveForward
	//	*CommandOutputs_MoveBackward
	//	*CommandOutputs_MoveTo
	//	*CommandOutputs_CargoOpen
	//	*CommandOutputs_CargoClose
	//	*CommandOutputs_CargoLift
	//	*CommandOutputs_CargoLower
	//	*CommandOutputs_CargoCheckQr
	//	*CommandOutputs_ScanLocation
	//	*CommandOutputs_Wait
	Outputs       isCommandOutputs_Outputs `protobuf_oneof:"outputs"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommandOutputs) Reset() {
	*x = CommandOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandOutputs) ProtoMessage() {}

func (x *CommandOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandOutputs.ProtoReflect.Descriptor instead.
func (*CommandOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{0}
}

func (x *CommandOutputs) GetOutputs() isCommandOutputs_Outputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *CommandOutputs) GetStop() *StopOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_Stop); ok {
			return x.Stop
		}
	}
	return nil
}

func (x *CommandOutputs) GetMoveForward() *MoveForwardOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_MoveForward); ok {
			return x.MoveForward
		}
	}
	return nil
}

func (x *CommandOutputs) GetMoveBackward() *MoveBackwardOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_MoveBackward); ok {
			return x.MoveBackward
		}
	}
	return nil
}

func (x *CommandOutputs) GetMoveTo() *MoveToOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_MoveTo); ok {
			return x.MoveTo
		}
	}
	return nil
}

func (x *CommandOutputs) GetCargoOpen() *CargoOpenOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_CargoOpen); ok {
			return x.CargoOpen
		}
	}
	return nil
}

func (x *CommandOutputs) GetCargoClose() *CargoCloseOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_CargoClose); ok {
			return x.CargoClose
		}
	}
	return nil
}

func (x *CommandOutputs) GetCargoLift() *CargoLiftOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_CargoLift); ok {
			return x.CargoLift
		}
	}
	return nil
}

func (x *CommandOutputs) GetCargoLower() *CargoLowerOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_CargoLower); ok {
			return x.CargoLower
		}
	}
	return nil
}

func (x *CommandOutputs) GetCargoCheckQr() *CargoCheckQROutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_CargoCheckQr); ok {
			return x.CargoCheckQr
		}
	}
	return nil
}

func (x *CommandOutputs) GetScanLocation() *ScanLocationOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_ScanLocation); ok {
			return x.ScanLocation
		}
	}
	return nil
}

func (x *CommandOutputs) GetWait() *WaitOutputs {
	if x != nil {
		if x, ok := x.Outputs.(*CommandOutputs_Wait); ok {
			return x.Wait
		}
	}
	return nil
}

type isCommandOutputs_Outputs interface {
	isCommandOutputs_Outputs()
}

type CommandOutputs_Stop struct {
	Stop *StopOutputs `protobuf:"bytes,1,opt,name=stop,proto3,oneof"`
}

type CommandOutputs_MoveForward struct {
	MoveForward *MoveForwardOutputs `protobuf:"bytes,2,opt,name=move_forward,json=moveForward,proto3,oneof"`
}

type CommandOutputs_MoveBackward struct {
	MoveBackward *MoveBackwardOutputs `protobuf:"bytes,3,opt,name=move_backward,json=moveBackward,proto3,oneof"`
}

type CommandOutputs_MoveTo struct {
	MoveTo *MoveToOutputs `protobuf:"bytes,4,opt,name=move_to,json=moveTo,proto3,oneof"`
}

type CommandOutputs_CargoOpen struct {
	CargoOpen *CargoOpenOutputs `protobuf:"bytes,5,opt,name=cargo_open,json=cargoOpen,proto3,oneof"`
}

type CommandOutputs_CargoClose struct {
	CargoClose *CargoCloseOutputs `protobuf:"bytes,6,opt,name=cargo_close,json=cargoClose,proto3,oneof"`
}

type CommandOutputs_CargoLift struct {
	CargoLift *CargoLiftOutputs `protobuf:"bytes,7,opt,name=cargo_lift,json=cargoLift,proto3,oneof"`
}

type CommandOutputs_CargoLower struct {
	CargoLower *CargoLowerOutputs `protobuf:"bytes,8,opt,name=cargo_lower,json=cargoLower,proto3,oneof"`
}

type CommandOutputs_CargoCheckQr struct {
	CargoCheckQr *CargoCheckQROutputs `protobuf:"bytes,9,opt,name=cargo_check_qr,json=cargoCheckQr,proto3,oneof"`
}

type CommandOutputs_ScanLocation struct {
	ScanLocation *ScanLocationOutputs `protobuf:"bytes,10,opt,name=scan_location,json=scanLocation,proto3,oneof"`
}

type CommandOutputs_Wait struct {
	Wait *WaitOutputs `protobuf:"bytes,11,opt,name=wait,proto3,oneof"`
}

func (*CommandOutputs_Stop) isCommandOutputs_Outputs() {}

func (*CommandOutputs_MoveForward) isCommandOutputs_Outputs() {}

func (*CommandOutputs_MoveBackward) isCommandOutputs_Outputs() {}

func (*CommandOutputs_MoveTo) isCommandOutputs_Outputs() {}

func (*CommandOutputs_CargoOpen) isCommandOutputs_Outputs() {}

func (*CommandOutputs_CargoClose) isCommandOutputs_Outputs() {}

func (*CommandOutputs_CargoLift) isCommandOutputs_Outputs() {}

func (*CommandOutputs_CargoLower) isCommandOutputs_Outputs() {}

func (*CommandOutputs_CargoCheckQr) isCommandOutputs_Outputs() {}

func (*CommandOutputs_ScanLocation) isCommandOutputs_Outputs() {}

func (*CommandOutputs_Wait) isCommandOutputs_Outputs() {}

type StopOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopOutputs) Reset() {
	*x = StopOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopOutputs) ProtoMessage() {}

func (x *StopOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopOutputs.ProtoReflect.Descriptor instead.
func (*StopOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{1}
}

type MoveForwardOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveForwardOutputs) Reset() {
	*x = MoveForwardOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveForwardOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveForwardOutputs) ProtoMessage() {}

func (x *MoveForwardOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveForwardOutputs.ProtoReflect.Descriptor instead.
func (*MoveForwardOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{2}
}

type MoveBackwardOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveBackwardOutputs) Reset() {
	*x = MoveBackwardOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveBackwardOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveBackwardOutputs) ProtoMessage() {}

func (x *MoveBackwardOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveBackwardOutputs.ProtoReflect.Descriptor instead.
func (*MoveBackwardOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{3}
}

type MoveToOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveToOutputs) Reset() {
	*x = MoveToOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveToOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveToOutputs) ProtoMessage() {}

func (x *MoveToOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveToOutputs.ProtoReflect.Descriptor instead.
func (*MoveToOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{4}
}

type CargoOpenOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CargoOpenOutputs) Reset() {
	*x = CargoOpenOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CargoOpenOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CargoOpenOutputs) ProtoMessage() {}

func (x *CargoOpenOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CargoOpenOutputs.ProtoReflect.Descriptor instead.
func (*CargoOpenOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{5}
}

type CargoCloseOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CargoCloseOutputs) Reset() {
	*x = CargoCloseOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CargoCloseOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CargoCloseOutputs) ProtoMessage() {}

func (x *CargoCloseOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CargoCloseOutputs.ProtoReflect.Descriptor instead.
func (*CargoCloseOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{6}
}

type CargoLiftOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CargoLiftOutputs) Reset() {
	*x = CargoLiftOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CargoLiftOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CargoLiftOutputs) ProtoMessage() {}

func (x *CargoLiftOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CargoLiftOutputs.ProtoReflect.Descriptor instead.
func (*CargoLiftOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{7}
}

type CargoLowerOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CargoLowerOutputs) Reset() {
	*x = CargoLowerOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CargoLowerOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CargoLowerOutputs) ProtoMessage() {}

func (x *CargoLowerOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CargoLowerOutputs.ProtoReflect.Descriptor instead.
func (*CargoLowerOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{8}
}

type CargoCheckQROutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CargoCheckQROutputs) Reset() {
	*x = CargoCheckQROutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CargoCheckQROutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CargoCheckQROutputs) ProtoMessage() {}

func (x *CargoCheckQROutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CargoCheckQROutputs.ProtoReflect.Descriptor instead.
func (*CargoCheckQROutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{9}
}

type Location struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Location      string                 `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	ScannedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=scanned_at,json=scannedAt,proto3" json:"scanned_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_command_v1_outputs_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{10}
}

func (x *Location) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Location) GetScannedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ScannedAt
	}
	return nil
}

type ScanLocationOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Locations     []*Location            `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScanLocationOutputs) Reset() {
	*x = ScanLocationOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScanLocationOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanLocationOutputs) ProtoMessage() {}

func (x *ScanLocationOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanLocationOutputs.ProtoReflect.Descriptor instead.
func (*ScanLocationOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{11}
}

func (x *ScanLocationOutputs) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

type WaitOutputs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WaitOutputs) Reset() {
	*x = WaitOutputs{}
	mi := &file_command_v1_outputs_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WaitOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitOutputs) ProtoMessage() {}

func (x *WaitOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_command_v1_outputs_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitOutputs.ProtoReflect.Descriptor instead.
func (*WaitOutputs) Descriptor() ([]byte, []int) {
	return file_command_v1_outputs_proto_rawDescGZIP(), []int{12}
}

var File_command_v1_outputs_proto protoreflect.FileDescriptor

const file_command_v1_outputs_proto_rawDesc = "" +
	"\n" +
	"\x18command/v1/outputs.proto\x12\n" +
	"command.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xcf\x05\n" +
	"\x0eCommandOutputs\x12-\n" +
	"\x04stop\x18\x01 \x01(\v2\x17.command.v1.StopOutputsH\x00R\x04stop\x12C\n" +
	"\fmove_forward\x18\x02 \x01(\v2\x1e.command.v1.MoveForwardOutputsH\x00R\vmoveForward\x12F\n" +
	"\rmove_backward\x18\x03 \x01(\v2\x1f.command.v1.MoveBackwardOutputsH\x00R\fmoveBackward\x124\n" +
	"\amove_to\x18\x04 \x01(\v2\x19.command.v1.MoveToOutputsH\x00R\x06moveTo\x12=\n" +
	"\n" +
	"cargo_open\x18\x05 \x01(\v2\x1c.command.v1.CargoOpenOutputsH\x00R\tcargoOpen\x12@\n" +
	"\vcargo_close\x18\x06 \x01(\v2\x1d.command.v1.CargoCloseOutputsH\x00R\n" +
	"cargoClose\x12=\n" +
	"\n" +
	"cargo_lift\x18\a \x01(\v2\x1c.command.v1.CargoLiftOutputsH\x00R\tcargoLift\x12@\n" +
	"\vcargo_lower\x18\b \x01(\v2\x1d.command.v1.CargoLowerOutputsH\x00R\n" +
	"cargoLower\x12G\n" +
	"\x0ecargo_check_qr\x18\t \x01(\v2\x1f.command.v1.CargoCheckQROutputsH\x00R\fcargoCheckQr\x12F\n" +
	"\rscan_location\x18\n" +
	" \x01(\v2\x1f.command.v1.ScanLocationOutputsH\x00R\fscanLocation\x12-\n" +
	"\x04wait\x18\v \x01(\v2\x17.command.v1.WaitOutputsH\x00R\x04waitB\t\n" +
	"\aoutputs\"\r\n" +
	"\vStopOutputs\"\x14\n" +
	"\x12MoveForwardOutputs\"\x15\n" +
	"\x13MoveBackwardOutputs\"\x0f\n" +
	"\rMoveToOutputs\"\x12\n" +
	"\x10CargoOpenOutputs\"\x13\n" +
	"\x11CargoCloseOutputs\"\x12\n" +
	"\x10CargoLiftOutputs\"\x13\n" +
	"\x11CargoLowerOutputs\"\x15\n" +
	"\x13CargoCheckQROutputs\"a\n" +
	"\bLocation\x12\x1a\n" +
	"\blocation\x18\x01 \x01(\tR\blocation\x129\n" +
	"\n" +
	"scanned_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tscannedAt\"I\n" +
	"\x13ScanLocationOutputs\x122\n" +
	"\tlocations\x18\x01 \x03(\v2\x14.command.v1.LocationR\tlocations\"\r\n" +
	"\vWaitOutputsB+Z)github.com/tbe-team/raybot-api/command/v1b\x06proto3"

var (
	file_command_v1_outputs_proto_rawDescOnce sync.Once
	file_command_v1_outputs_proto_rawDescData []byte
)

func file_command_v1_outputs_proto_rawDescGZIP() []byte {
	file_command_v1_outputs_proto_rawDescOnce.Do(func() {
		file_command_v1_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_command_v1_outputs_proto_rawDesc), len(file_command_v1_outputs_proto_rawDesc)))
	})
	return file_command_v1_outputs_proto_rawDescData
}

var file_command_v1_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_command_v1_outputs_proto_goTypes = []any{
	(*CommandOutputs)(nil),        // 0: command.v1.CommandOutputs
	(*StopOutputs)(nil),           // 1: command.v1.StopOutputs
	(*MoveForwardOutputs)(nil),    // 2: command.v1.MoveForwardOutputs
	(*MoveBackwardOutputs)(nil),   // 3: command.v1.MoveBackwardOutputs
	(*MoveToOutputs)(nil),         // 4: command.v1.MoveToOutputs
	(*CargoOpenOutputs)(nil),      // 5: command.v1.CargoOpenOutputs
	(*CargoCloseOutputs)(nil),     // 6: command.v1.CargoCloseOutputs
	(*CargoLiftOutputs)(nil),      // 7: command.v1.CargoLiftOutputs
	(*CargoLowerOutputs)(nil),     // 8: command.v1.CargoLowerOutputs
	(*CargoCheckQROutputs)(nil),   // 9: command.v1.CargoCheckQROutputs
	(*Location)(nil),              // 10: command.v1.Location
	(*ScanLocationOutputs)(nil),   // 11: command.v1.ScanLocationOutputs
	(*WaitOutputs)(nil),           // 12: command.v1.WaitOutputs
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_command_v1_outputs_proto_depIdxs = []int32{
	1,  // 0: command.v1.CommandOutputs.stop:type_name -> command.v1.StopOutputs
	2,  // 1: command.v1.CommandOutputs.move_forward:type_name -> command.v1.MoveForwardOutputs
	3,  // 2: command.v1.CommandOutputs.move_backward:type_name -> command.v1.MoveBackwardOutputs
	4,  // 3: command.v1.CommandOutputs.move_to:type_name -> command.v1.MoveToOutputs
	5,  // 4: command.v1.CommandOutputs.cargo_open:type_name -> command.v1.CargoOpenOutputs
	6,  // 5: command.v1.CommandOutputs.cargo_close:type_name -> command.v1.CargoCloseOutputs
	7,  // 6: command.v1.CommandOutputs.cargo_lift:type_name -> command.v1.CargoLiftOutputs
	8,  // 7: command.v1.CommandOutputs.cargo_lower:type_name -> command.v1.CargoLowerOutputs
	9,  // 8: command.v1.CommandOutputs.cargo_check_qr:type_name -> command.v1.CargoCheckQROutputs
	11, // 9: command.v1.CommandOutputs.scan_location:type_name -> command.v1.ScanLocationOutputs
	12, // 10: command.v1.CommandOutputs.wait:type_name -> command.v1.WaitOutputs
	13, // 11: command.v1.Location.scanned_at:type_name -> google.protobuf.Timestamp
	10, // 12: command.v1.ScanLocationOutputs.locations:type_name -> command.v1.Location
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_command_v1_outputs_proto_init() }
func file_command_v1_outputs_proto_init() {
	if File_command_v1_outputs_proto != nil {
		return
	}
	file_command_v1_outputs_proto_msgTypes[0].OneofWrappers = []any{
		(*CommandOutputs_Stop)(nil),
		(*CommandOutputs_MoveForward)(nil),
		(*CommandOutputs_MoveBackward)(nil),
		(*CommandOutputs_MoveTo)(nil),
		(*CommandOutputs_CargoOpen)(nil),
		(*CommandOutputs_CargoClose)(nil),
		(*CommandOutputs_CargoLift)(nil),
		(*CommandOutputs_CargoLower)(nil),
		(*CommandOutputs_CargoCheckQr)(nil),
		(*CommandOutputs_ScanLocation)(nil),
		(*CommandOutputs_Wait)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_command_v1_outputs_proto_rawDesc), len(file_command_v1_outputs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_v1_outputs_proto_goTypes,
		DependencyIndexes: file_command_v1_outputs_proto_depIdxs,
		MessageInfos:      file_command_v1_outputs_proto_msgTypes,
	}.Build()
	File_command_v1_outputs_proto = out.File
	file_command_v1_outputs_proto_goTypes = nil
	file_command_v1_outputs_proto_depIdxs = nil
}
