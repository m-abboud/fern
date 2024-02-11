/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as SeedApi from "../../../../api";
import * as core from "../../../../core";

export const CreateMovieRequest: core.serialization.ObjectSchema<
    serializers.CreateMovieRequest.Raw,
    SeedApi.CreateMovieRequest
> = core.serialization.object({
    title: core.serialization.string(),
    rating: core.serialization.number(),
});

export declare namespace CreateMovieRequest {
    interface Raw {
        title: string;
        rating: number;
    }
}
