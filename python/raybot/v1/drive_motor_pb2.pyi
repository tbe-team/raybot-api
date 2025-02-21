from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SetDriveMotorConfigurationRequest(_message.Message):
    __slots__ = ("speed", "direction")
    class Direction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        DIRECTION_UNSPECIFIED: _ClassVar[SetDriveMotorConfigurationRequest.Direction]
        DIRECTION_FORWARD: _ClassVar[SetDriveMotorConfigurationRequest.Direction]
        DIRECTION_BACKWARD: _ClassVar[SetDriveMotorConfigurationRequest.Direction]
    DIRECTION_UNSPECIFIED: SetDriveMotorConfigurationRequest.Direction
    DIRECTION_FORWARD: SetDriveMotorConfigurationRequest.Direction
    DIRECTION_BACKWARD: SetDriveMotorConfigurationRequest.Direction
    SPEED_FIELD_NUMBER: _ClassVar[int]
    DIRECTION_FIELD_NUMBER: _ClassVar[int]
    speed: int
    direction: SetDriveMotorConfigurationRequest.Direction
    def __init__(self, speed: _Optional[int] = ..., direction: _Optional[_Union[SetDriveMotorConfigurationRequest.Direction, str]] = ...) -> None: ...

class SetDriveMotorConfigurationResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
