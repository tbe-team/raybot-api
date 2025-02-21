from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SetLiftMotorConfigurationRequest(_message.Message):
    __slots__ = ("distance", "direction")
    class Direction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        DIRECTION_UNSPECIFIED: _ClassVar[SetLiftMotorConfigurationRequest.Direction]
        DIRECTION_UP: _ClassVar[SetLiftMotorConfigurationRequest.Direction]
        DIRECTION_DOWN: _ClassVar[SetLiftMotorConfigurationRequest.Direction]
    DIRECTION_UNSPECIFIED: SetLiftMotorConfigurationRequest.Direction
    DIRECTION_UP: SetLiftMotorConfigurationRequest.Direction
    DIRECTION_DOWN: SetLiftMotorConfigurationRequest.Direction
    DISTANCE_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    distance: float
    direction: SetLiftMotorConfigurationRequest.Direction
    def __init__(self, distance: _Optional[float] = ..., direction: _Optional[_Union[SetLiftMotorConfigurationRequest.Direction, str]] = ...) -> None: ...

class SetLiftMotorConfigurationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
