/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "../../../../core";
import * as serializers from "../../../index";
import * as SeedUnions from "../../../../api/index";

const _Base = core.serialization.object({
    id: core.serialization.string(),
});
export const UnionWithBaseProperties: core.serialization.Schema<
    serializers.UnionWithBaseProperties.Raw,
    SeedUnions.UnionWithBaseProperties
> = core.serialization
    .union("type", {
        integer: core.serialization
            .object({
                value: core.serialization.number(),
            })
            .extend(_Base),
        string: core.serialization
            .object({
                value: core.serialization.string(),
            })
            .extend(_Base),
        foo: core.serialization.lazyObject(() => serializers.Foo).extend(_Base),
    })
    .transform<SeedUnions.UnionWithBaseProperties>({
        transform: (value) => value,
        untransform: (value) => value,
    });

export declare namespace UnionWithBaseProperties {
    type Raw = UnionWithBaseProperties.Integer | UnionWithBaseProperties.String | UnionWithBaseProperties.Foo;

    interface Integer extends _Base {
        type: "integer";
        value: number;
    }

    interface String extends _Base {
        type: "string";
        value: string;
    }

    interface Foo extends _Base, serializers.Foo.Raw {
        type: "foo";
    }

    interface _Base {
        id: string;
    }
}
