from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class UserRequest(_message.Message):
    __slots__ = ("name", "age")
    NAME_FIELD_NUMBER: _ClassVar[int]
    AGE_FIELD_NUMBER: _ClassVar[int]
    name: str
    age: int
    def __init__(self, name: _Optional[str] = ..., age: _Optional[int] = ...) -> None: ...

class UserResponse(_message.Message):
    __slots__ = ("msg", "code")
    MSG_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    msg: str
    code: int
    def __init__(self, msg: _Optional[str] = ..., code: _Optional[int] = ...) -> None: ...

class GetUserRequest(_message.Message):
    __slots__ = ("name",)
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class GetUserResponse(_message.Message):
    __slots__ = ("name", "age")
    NAME_FIELD_NUMBER: _ClassVar[int]
    AGE_FIELD_NUMBER: _ClassVar[int]
    name: str
    age: str
    def __init__(self, name: _Optional[str] = ..., age: _Optional[str] = ...) -> None: ...
