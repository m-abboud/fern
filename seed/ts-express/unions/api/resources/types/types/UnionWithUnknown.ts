/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as SeedUnions from "../../../index";

export type UnionWithUnknown = SeedUnions.UnionWithUnknown.Foo | SeedUnions.UnionWithUnknown.Unknown;

export declare namespace UnionWithUnknown {
    interface Foo extends SeedUnions.Foo {
        type: "foo";
    }

    interface Unknown {
        type: "unknown";
    }
}