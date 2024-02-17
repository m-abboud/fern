/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as FernIr from "../../../../api";
import * as core from "../../../../core";

export const InlinedWebsocketMessageBodyProperty: core.serialization.ObjectSchema<
    serializers.InlinedWebsocketMessageBodyProperty.Raw,
    FernIr.InlinedWebsocketMessageBodyProperty
> = core.serialization
    .objectWithoutOptionalProperties({
        name: core.serialization.lazyObject(async () => (await import("../../..")).NameAndWireValue),
        valueType: core.serialization.lazy(async () => (await import("../../..")).TypeReference),
    })
    .extend(core.serialization.lazyObject(async () => (await import("../../..")).WithDocs));

export declare namespace InlinedWebsocketMessageBodyProperty {
    interface Raw extends serializers.WithDocs.Raw {
        name: serializers.NameAndWireValue.Raw;
        valueType: serializers.TypeReference.Raw;
    }
}