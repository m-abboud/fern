/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as FernOpenapiIr from "../../../../api";
import * as core from "../../../../core";

export const SdkGroupInfo: core.serialization.ObjectSchema<serializers.SdkGroupInfo.Raw, FernOpenapiIr.SdkGroupInfo> =
    core.serialization.objectWithoutOptionalProperties({
        summary: core.serialization.string().optional(),
        description: core.serialization.string().optional(),
    });

export declare namespace SdkGroupInfo {
    interface Raw {
        summary?: string | null;
        description?: string | null;
    }
}