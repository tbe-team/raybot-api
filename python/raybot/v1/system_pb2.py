# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: raybot/v1/system.proto
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
    'raybot/v1/system.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16raybot/v1/system.proto\x12\traybot.v1\"\x16\n\x14\x45mergencyStopRequest\"\x17\n\x15\x45mergencyStopResponse2e\n\rSystemService\x12T\n\rEmergencyStop\x12\x1f.raybot.v1.EmergencyStopRequest\x1a .raybot.v1.EmergencyStopResponse\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'raybot.v1.system_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_EMERGENCYSTOPREQUEST']._serialized_start=37
  _globals['_EMERGENCYSTOPREQUEST']._serialized_end=59
  _globals['_EMERGENCYSTOPRESPONSE']._serialized_start=61
  _globals['_EMERGENCYSTOPRESPONSE']._serialized_end=84
  _globals['_SYSTEMSERVICE']._serialized_start=86
  _globals['_SYSTEMSERVICE']._serialized_end=187
# @@protoc_insertion_point(module_scope)
