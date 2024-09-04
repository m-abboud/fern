# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations
from .....core.pydantic_utilities import UniversalBaseModel
import typing
import typing_extensions
from .....core.serialization import FieldMetadata
from .....core.pydantic_utilities import IS_PYDANTIC_V2
import pydantic


class Animal_Dog(UniversalBaseModel):
    animal: typing.Literal["dog"] = "dog"
    name: str
    likes_to_woof: typing_extensions.Annotated[bool, FieldMetadata(alias="likesToWoof")]

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(extra="allow")  # type: ignore # Pydantic v2
    else:

        class Config:
            extra = pydantic.Extra.allow


class Animal_Cat(UniversalBaseModel):
    animal: typing.Literal["cat"] = "cat"
    name: str
    likes_to_meow: typing_extensions.Annotated[bool, FieldMetadata(alias="likesToMeow")]

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(extra="allow")  # type: ignore # Pydantic v2
    else:

        class Config:
            extra = pydantic.Extra.allow


Animal = typing.Union[Animal_Dog, Animal_Cat]
