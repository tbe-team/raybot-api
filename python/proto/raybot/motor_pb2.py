# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: proto/raybot/motor.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'proto/raybot/motor.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18proto/raybot/motor.proto\x12\x05motor\"\xc0\x01\n\x10SetConfigRequest\x12\x14\n\x05speed\x18\x01 \x01(\rR\x05speed\x12?\n\tdirection\x18\x02 \x01(\x0e\x32!.motor.SetConfigRequest.DirectionR\tdirection\"U\n\tDirection\x12\x19\n\x15\x44IRECTION_UNSPECIFIED\x10\x00\x12\x15\n\x11\x44IRECTION_FORWARD\x10\x01\x12\x16\n\x12\x44IRECTION_BACKWARD\x10\x02\"\x13\n\x11SetConfigResponse\"6\n\x13StreamStatusRequest\x12\x1f\n\x0binterval_ms\x18\x01 \x01(\rR\nintervalMs\"D\n\x14StreamStatusResponse\x12\x14\n\x05speed\x18\x01 \x01(\rR\x05speed\x12\x16\n\x06status\x18\x02 \x01(\tR\x06status2\x9e\x01\n\x11\x44riveMotorService\x12>\n\tSetConfig\x12\x17.motor.SetConfigRequest\x1a\x18.motor.SetConfigResponse\x12I\n\x0cStreamStatus\x12\x1a.motor.StreamStatusRequest\x1a\x1b.motor.StreamStatusResponse0\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.raybot.motor_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_SETCONFIGREQUEST']._serialized_start=36
  _globals['_SETCONFIGREQUEST']._serialized_end=228
  _globals['_SETCONFIGREQUEST_DIRECTION']._serialized_start=143
  _globals['_SETCONFIGREQUEST_DIRECTION']._serialized_end=228
  _globals['_SETCONFIGRESPONSE']._serialized_start=230
  _globals['_SETCONFIGRESPONSE']._serialized_end=249
  _globals['_STREAMSTATUSREQUEST']._serialized_start=251
  _globals['_STREAMSTATUSREQUEST']._serialized_end=305
  _globals['_STREAMSTATUSRESPONSE']._serialized_start=307
  _globals['_STREAMSTATUSRESPONSE']._serialized_end=375
  _globals['_DRIVEMOTORSERVICE']._serialized_start=378
  _globals['_DRIVEMOTORSERVICE']._serialized_end=536
# @@protoc_insertion_point(module_scope)
