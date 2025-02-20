from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SetConfigRequest(_message.Message):
    __slots__ = ("speed", "direction")
    class Direction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        DIRECTION_UNSPECIFIED: _ClassVar[SetConfigRequest.Direction]
        DIRECTION_FORWARD: _ClassVar[SetConfigRequest.Direction]
        DIRECTION_BACKWARD: _ClassVar[SetConfigRequest.Direction]
    DIRECTION_UNSPECIFIED: SetConfigRequest.Direction
    DIRECTION_FORWARD: SetConfigRequest.Direction
    DIRECTION_BACKWARD: SetConfigRequest.Direction
    SPEED_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    speed: int
    direction: SetConfigRequest.Direction
    def __init__(self, speed: _Optional[int] = ..., direction: _Optional[_Union[SetConfigRequest.Direction, str]] = ...) -> None: ...

class SetConfigResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class StreamStatusRequest(_message.Message):
    __slots__ = ("interval_ms",)
    INTERVAL_MS_FIELD_NUMBER: _ClassVar[int]
    interval_ms: int
    def __init__(self, interval_ms: _Optional[int] = ...) -> None: ...

class StreamStatusResponse(_message.Message):
    __slots__ = ("speed", "status")
    SPEED_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    speed: int
    status: str
    def __init__(self, speed: _Optional[int] = ..., status: _Optional[str] = ...) -> None: ...
