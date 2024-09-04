# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations
from ......core.pydantic_utilities import UniversalBaseModel
import typing_extensions
from .object_with_optional_field import ObjectWithOptionalField
from ......core.serialization import FieldMetadata
import typing
from ......core.pydantic_utilities import universal_root_validator
from ......core.pydantic_utilities import universal_field_validator
from ......core.pydantic_utilities import IS_PYDANTIC_V2
import pydantic


class NestedObjectWithRequiredField(UniversalBaseModel):
    string: str
    nested_object: typing_extensions.Annotated[
        ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
    ]

    class Partial(typing.TypedDict):
        string: typing_extensions.NotRequired[str]
        nested_object: typing_extensions.NotRequired[
            typing_extensions.Annotated[
                ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
            ]
        ]

    class Validators:
        """
        Use this class to add validators to the Pydantic model.

            @NestedObjectWithRequiredField.Validators.root()
            def validate(values: NestedObjectWithRequiredField.Partial) -> NestedObjectWithRequiredField.Partial:
                ...

            @NestedObjectWithRequiredField.Validators.field("string")
            def validate_string(string: str, values: NestedObjectWithRequiredField.Partial) -> str:
                ...

            @NestedObjectWithRequiredField.Validators.field("nested_object")
            def validate_nested_object(nested_object: typing_extensions.Annotated[ObjectWithOptionalField, FieldMetadata(alias="NestedObject")], values: NestedObjectWithRequiredField.Partial) -> typing_extensions.Annotated[ObjectWithOptionalField, FieldMetadata(alias="NestedObject")]:
                ...
        """

        _pre_validators: typing.ClassVar[
            typing.List[NestedObjectWithRequiredField.Validators._PreRootValidator]
        ] = []
        _post_validators: typing.ClassVar[
            typing.List[NestedObjectWithRequiredField.Validators._RootValidator]
        ] = []
        _string_pre_validators: typing.ClassVar[
            typing.List[NestedObjectWithRequiredField.Validators.PreStringValidator]
        ] = []
        _string_post_validators: typing.ClassVar[
            typing.List[NestedObjectWithRequiredField.Validators.StringValidator]
        ] = []
        _nested_object_pre_validators: typing.ClassVar[
            typing.List[
                NestedObjectWithRequiredField.Validators.PreNestedObjectValidator
            ]
        ] = []
        _nested_object_post_validators: typing.ClassVar[
            typing.List[NestedObjectWithRequiredField.Validators.NestedObjectValidator]
        ] = []

        @typing.overload
        @classmethod
        def root(
            cls, *, pre: typing.Literal[False] = False
        ) -> typing.Callable[
            [NestedObjectWithRequiredField.Validators._RootValidator],
            NestedObjectWithRequiredField.Validators._RootValidator,
        ]: ...
        @typing.overload
        @classmethod
        def root(
            cls, *, pre: typing.Literal[True]
        ) -> typing.Callable[
            [NestedObjectWithRequiredField.Validators._PreRootValidator],
            NestedObjectWithRequiredField.Validators._PreRootValidator,
        ]: ...
        @classmethod
        def root(cls, *, pre: bool = False) -> typing.Any:
            def decorator(validator: typing.Any) -> typing.Any:
                if pre:
                    cls._pre_validators.append(validator)
                else:
                    cls._post_validators.append(validator)
                return validator

            return decorator

        @typing.overload
        @classmethod
        def field(
            cls, field_name: typing.Literal["string"], *, pre: typing.Literal[True]
        ) -> typing.Callable[
            [NestedObjectWithRequiredField.Validators.PreStringValidator],
            NestedObjectWithRequiredField.Validators.PreStringValidator,
        ]: ...
        @typing.overload
        @classmethod
        def field(
            cls,
            field_name: typing.Literal["string"],
            *,
            pre: typing.Literal[False] = False,
        ) -> typing.Callable[
            [NestedObjectWithRequiredField.Validators.StringValidator],
            NestedObjectWithRequiredField.Validators.StringValidator,
        ]: ...
        @typing.overload
        @classmethod
        def field(
            cls,
            field_name: typing.Literal["nested_object"],
            *,
            pre: typing.Literal[True],
        ) -> typing.Callable[
            [NestedObjectWithRequiredField.Validators.PreNestedObjectValidator],
            NestedObjectWithRequiredField.Validators.PreNestedObjectValidator,
        ]: ...
        @typing.overload
        @classmethod
        def field(
            cls,
            field_name: typing.Literal["nested_object"],
            *,
            pre: typing.Literal[False] = False,
        ) -> typing.Callable[
            [NestedObjectWithRequiredField.Validators.NestedObjectValidator],
            NestedObjectWithRequiredField.Validators.NestedObjectValidator,
        ]: ...
        @classmethod
        def field(cls, field_name: str, *, pre: bool = False) -> typing.Any:
            def decorator(validator: typing.Any) -> typing.Any:
                if field_name == "string":
                    if pre:
                        cls._string_pre_validators.append(validator)
                    else:
                        cls._string_post_validators.append(validator)
                if field_name == "nested_object":
                    if pre:
                        cls._nested_object_pre_validators.append(validator)
                    else:
                        cls._nested_object_post_validators.append(validator)
                return validator

            return decorator

        class PreStringValidator(typing.Protocol):
            def __call__(
                self, __v: typing.Any, __values: NestedObjectWithRequiredField.Partial
            ) -> typing.Any: ...

        class StringValidator(typing.Protocol):
            def __call__(
                self, __v: str, __values: NestedObjectWithRequiredField.Partial
            ) -> str: ...

        class PreNestedObjectValidator(typing.Protocol):
            def __call__(
                self, __v: typing.Any, __values: NestedObjectWithRequiredField.Partial
            ) -> typing.Any: ...

        class NestedObjectValidator(typing.Protocol):
            def __call__(
                self,
                __v: typing_extensions.Annotated[
                    ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
                ],
                __values: NestedObjectWithRequiredField.Partial,
            ) -> typing_extensions.Annotated[
                ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
            ]: ...

        class _PreRootValidator(typing.Protocol):
            def __call__(self, __values: typing.Any) -> typing.Any: ...

        class _RootValidator(typing.Protocol):
            def __call__(
                self, __values: NestedObjectWithRequiredField.Partial
            ) -> NestedObjectWithRequiredField.Partial: ...

    @universal_root_validator(pre=True)
    def _pre_validate_types_nested_object_with_required_field(
        cls, values: NestedObjectWithRequiredField.Partial
    ) -> NestedObjectWithRequiredField.Partial:
        for validator in NestedObjectWithRequiredField.Validators._pre_validators:
            values = validator(values)
        return values

    @universal_root_validator(pre=False)
    def _post_validate_types_nested_object_with_required_field(
        cls, values: NestedObjectWithRequiredField.Partial
    ) -> NestedObjectWithRequiredField.Partial:
        for validator in NestedObjectWithRequiredField.Validators._post_validators:
            values = validator(values)
        return values

    @universal_field_validator("string", pre=True)
    def _pre_validate_string(
        cls, v: str, values: NestedObjectWithRequiredField.Partial
    ) -> str:
        for (
            validator
        ) in NestedObjectWithRequiredField.Validators._string_pre_validators:
            v = validator(v, values)
        return v

    @universal_field_validator("string", pre=False)
    def _post_validate_string(
        cls, v: str, values: NestedObjectWithRequiredField.Partial
    ) -> str:
        for (
            validator
        ) in NestedObjectWithRequiredField.Validators._string_post_validators:
            v = validator(v, values)
        return v

    @universal_field_validator("nested_object", pre=True)
    def _pre_validate_nested_object(
        cls,
        v: typing_extensions.Annotated[
            ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
        ],
        values: NestedObjectWithRequiredField.Partial,
    ) -> typing_extensions.Annotated[
        ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
    ]:
        for (
            validator
        ) in NestedObjectWithRequiredField.Validators._nested_object_pre_validators:
            v = validator(v, values)
        return v

    @universal_field_validator("nested_object", pre=False)
    def _post_validate_nested_object(
        cls,
        v: typing_extensions.Annotated[
            ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
        ],
        values: NestedObjectWithRequiredField.Partial,
    ) -> typing_extensions.Annotated[
        ObjectWithOptionalField, FieldMetadata(alias="NestedObject")
    ]:
        for (
            validator
        ) in NestedObjectWithRequiredField.Validators._nested_object_post_validators:
            v = validator(v, values)
        return v

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(
            extra="forbid"
        )  # type: ignore # Pydantic v2
    else:

        class Config:
            extra = pydantic.Extra.forbid
