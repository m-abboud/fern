/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as FernIr from "../../../../api";
import * as core from "../../../../core";

export const ExampleSetContainer: core.serialization.ObjectSchema<
    serializers.ExampleSetContainer.Raw,
    FernIr.ExampleSetContainer
> = core.serialization.objectWithoutOptionalProperties({
    set: core.serialization.list(
        core.serialization.lazyObject(async () => (await import("../../..")).ExampleTypeReference)
    ),
    itemType: core.serialization.lazy(async () => (await import("../../..")).TypeReference).optional(),
});

export declare namespace ExampleSetContainer {
    interface Raw {
        set: serializers.ExampleTypeReference.Raw[];
        itemType?: serializers.TypeReference.Raw | null;
    }
}
