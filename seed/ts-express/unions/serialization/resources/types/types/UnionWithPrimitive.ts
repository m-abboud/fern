/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../index";
import * as SeedUnions from "../../../../api/index";
import * as core from "../../../../core";

export const UnionWithPrimitive: core.serialization.Schema<
    serializers.UnionWithPrimitive.Raw,
    SeedUnions.UnionWithPrimitive
> = core.serialization
    .union("type", {
        integer: core.serialization.object({
            value: core.serialization.number(),
        }),
        string: core.serialization.object({
            value: core.serialization.string(),
        }),
    })
    .transform<SeedUnions.UnionWithPrimitive>({
        transform: (value) => value,
        untransform: (value) => value,
    });

export declare namespace UnionWithPrimitive {
    type Raw = UnionWithPrimitive.Integer | UnionWithPrimitive.String;

    interface Integer {
        type: "integer";
        value: number;
    }

    interface String {
        type: "string";
        value: string;
    }
}