from __future__ import annotations

from dataclasses import dataclass
from typing import Optional

from .module import Module
from .qualfied_name import QualifiedName


@dataclass(frozen=True)
class ReferenceImport:
    module: Module
    named_import: Optional[str] = None
    alias: Optional[str] = None


@dataclass(frozen=True)
class Reference:
    qualified_name_excluding_import: QualifiedName
    """
    the qualfied name of the reference "inside" the import.

    example:
        import typing

        l: typing.List # qualified_name_excluding_import == (List,)

    example:
        from typing import List

        l: List # qualified_name_excluding_import == ()
    """

    import_: Optional[ReferenceImport] = None
    """
    not required for built-ins, like str
    """

    must_import_after_current_declaration: bool = False
    """
    in Python 3.7+, annotations can be imported after they're used, with:
      from __future__ import annotations.
    for built-ins and non-annotation references, this field is ignored.
    """
