# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: raybot/v1/auth.proto
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
    'raybot/v1/auth.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14raybot/v1/auth.proto\x12\traybot.v1\"K\n\x11GetSessionRequest\x12\x1a\n\x08username\x18\x01 \x01(\tR\x08username\x12\x1a\n\x08password\x18\x02 \x01(\tR\x08password\"3\n\x12GetSessionResponse\x12\x1d\n\nsession_id\x18\x01 \x01(\tR\tsessionId2Z\n\x0b\x41uthService\x12K\n\nGetSession\x12\x1c.raybot.v1.GetSessionRequest\x1a\x1d.raybot.v1.GetSessionResponse\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'raybot.v1.auth_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_GETSESSIONREQUEST']._serialized_start=35
  _globals['_GETSESSIONREQUEST']._serialized_end=110
  _globals['_GETSESSIONRESPONSE']._serialized_start=112
  _globals['_GETSESSIONRESPONSE']._serialized_end=163
  _globals['_AUTHSERVICE']._serialized_start=165
  _globals['_AUTHSERVICE']._serialized_end=255
# @@protoc_insertion_point(module_scope)
