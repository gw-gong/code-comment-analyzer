from google.protobuf import struct_pb2 as _struct_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AnalyzeFileContentReq(_message.Message):
    __slots__ = ("language", "file_content")
    LANGUAGE_FIELD_NUMBER: _ClassVar[int]
    FILE_CONTENT_FIELD_NUMBER: _ClassVar[int]
    language: str
    file_content: str
    def __init__(self, language: _Optional[str] = ..., file_content: _Optional[str] = ...) -> None: ...

class AnalyzeFileContentRes(_message.Message):
    __slots__ = ("analyzed_data",)
    ANALYZED_DATA_FIELD_NUMBER: _ClassVar[int]
    analyzed_data: _struct_pb2.Struct
    def __init__(self, analyzed_data: _Optional[_Union[_struct_pb2.Struct, _Mapping]] = ...) -> None: ...
