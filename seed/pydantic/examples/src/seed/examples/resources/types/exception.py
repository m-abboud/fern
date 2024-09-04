# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations
from ...core.pydantic_utilities import UniversalBaseModel
import typing
import typing_extensions
from ...core.serialization import FieldMetadata
from ...core.pydantic_utilities import IS_PYDANTIC_V2
import pydantic


class Exception_Generic(UniversalBaseModel):
    """
    Examples
    --------
    from seed.examples.resources import Exception_Generic

    Exception_Generic(
        exception_type="Unavailable",
        exception_message="This component is unavailable!",
        exception_stacktrace="<logs>",
    )
    """

    type: typing.Literal["generic"] = "generic"
    exception_type: typing_extensions.Annotated[str, FieldMetadata(alias="exceptionType")]
    exception_message: typing_extensions.Annotated[str, FieldMetadata(alias="exceptionMessage")]
    exception_stacktrace: typing_extensions.Annotated[str, FieldMetadata(alias="exceptionStacktrace")]

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(extra="allow")  # type: ignore # Pydantic v2
    else:

        class Config:
            extra = pydantic.Extra.allow


class Exception_Timeout(UniversalBaseModel):
    """
    Examples
    --------
    from seed.examples.resources import Exception_Generic

    Exception_Generic(
        exception_type="Unavailable",
        exception_message="This component is unavailable!",
        exception_stacktrace="<logs>",
    )
    """

    type: typing.Literal["timeout"] = "timeout"

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(extra="allow")  # type: ignore # Pydantic v2
    else:

        class Config:
            extra = pydantic.Extra.allow


"""
from seed.examples.resources import Exception_Generic

Exception_Generic(
    exception_type="Unavailable",
    exception_message="This component is unavailable!",
    exception_stacktrace="<logs>",
)
"""
Exception = typing.Union[Exception_Generic, Exception_Timeout]
