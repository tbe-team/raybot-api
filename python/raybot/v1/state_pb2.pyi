from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetRobotStateRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetRobotStateResponse(_message.Message):
    __slots__ = ("robot_state",)
    ROBOT_STATE_FIELD_NUMBER: _ClassVar[int]
    robot_state: RobotState
    def __init__(self, robot_state: _Optional[_Union[RobotState, _Mapping]] = ...) -> None: ...

class RobotState(_message.Message):
    __slots__ = ("battery_state", "wifi_state", "sensor_data", "position_data")
    BATTERY_STATE_FIELD_NUMBER: _ClassVar[int]
    WIFI_STATE_FIELD_NUMBER: _ClassVar[int]
    SENSOR_DATA_FIELD_NUMBER: _ClassVar[int]
    POSITION_DATA_FIELD_NUMBER: _ClassVar[int]
    battery_state: BatteryState
    wifi_state: WifiState
    sensor_data: SensorData
    position_data: PositionData
    def __init__(self, battery_state: _Optional[_Union[BatteryState, _Mapping]] = ..., wifi_state: _Optional[_Union[WifiState, _Mapping]] = ..., sensor_data: _Optional[_Union[SensorData, _Mapping]] = ..., position_data: _Optional[_Union[PositionData, _Mapping]] = ...) -> None: ...

class BatteryState(_message.Message):
    __slots__ = ("charge_percentage", "current", "voltage", "temperature", "status")
    class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        STATUS_UNSPECIFIED: _ClassVar[BatteryState.Status]
        STATUS_CHARGING: _ClassVar[BatteryState.Status]
        STATUS_DISCHARGING: _ClassVar[BatteryState.Status]
    STATUS_UNSPECIFIED: BatteryState.Status
    STATUS_CHARGING: BatteryState.Status
    STATUS_DISCHARGING: BatteryState.Status
    CHARGE_PERCENTAGE_FIELD_NUMBER: _ClassVar[int]
    CURRENT_FIELD_NUMBER: _ClassVar[int]
    VOLTAGE_FIELD_NUMBER: _ClassVar[int]
    TEMPERATURE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    charge_percentage: float
    current: float
    voltage: float
    temperature: float
    status: BatteryState.Status
    def __init__(self, charge_percentage: _Optional[float] = ..., current: _Optional[float] = ..., voltage: _Optional[float] = ..., temperature: _Optional[float] = ..., status: _Optional[_Union[BatteryState.Status, str]] = ...) -> None: ...

class WifiState(_message.Message):
    __slots__ = ("current_mode", "ssid")
    class Mode(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        MODE_UNSPECIFIED: _ClassVar[WifiState.Mode]
        MODE_STA: _ClassVar[WifiState.Mode]
    MODE_UNSPECIFIED: WifiState.Mode
    MODE_STA: WifiState.Mode
    CURRENT_MODE_FIELD_NUMBER: _ClassVar[int]
    SSID_FIELD_NUMBER: _ClassVar[int]
    current_mode: WifiState.Mode
    ssid: str
    def __init__(self, current_mode: _Optional[_Union[WifiState.Mode, str]] = ..., ssid: _Optional[str] = ...) -> None: ...

class SensorData(_message.Message):
    __slots__ = ("front_distance", "rear_distance", "drop_distance")
    FRONT_DISTANCE_FIELD_NUMBER: _ClassVar[int]
    REAR_DISTANCE_FIELD_NUMBER: _ClassVar[int]
    DROP_DISTANCE_FIELD_NUMBER: _ClassVar[int]
    front_distance: float
    rear_distance: float
    drop_distance: float
    def __init__(self, front_distance: _Optional[float] = ..., rear_distance: _Optional[float] = ..., drop_distance: _Optional[float] = ...) -> None: ...

class PositionData(_message.Message):
    __slots__ = ("current_location", "target_location")
    CURRENT_LOCATION_FIELD_NUMBER: _ClassVar[int]
    TARGET_LOCATION_FIELD_NUMBER: _ClassVar[int]
    current_location: str
    target_location: str
    def __init__(self, current_location: _Optional[str] = ..., target_location: _Optional[str] = ...) -> None: ...
