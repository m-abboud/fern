/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../..";
import * as SeedAudiences from "../../../../../api";
import * as core from "../../../../../core";

export const FindRequest: core.serialization.Schema<
    serializers.FindRequest.Raw,
    Omit<SeedAudiences.FindRequest, "optionalString">
> = core.serialization.object({
    publicProperty: core.serialization.string().optional(),
    privateProperty: core.serialization.number().optional(),
});

export declare namespace FindRequest {
    interface Raw {
        publicProperty?: string | null;
        privateProperty?: number | null;
    }
}