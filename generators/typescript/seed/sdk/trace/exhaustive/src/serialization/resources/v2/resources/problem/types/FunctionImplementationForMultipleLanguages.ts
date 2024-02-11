/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../../../..";
import * as SeedTrace from "../../../../../../api";
import * as core from "../../../../../../core";

export const FunctionImplementationForMultipleLanguages: core.serialization.ObjectSchema<
    serializers.v2.FunctionImplementationForMultipleLanguages.Raw,
    SeedTrace.v2.FunctionImplementationForMultipleLanguages
> = core.serialization.object({
    codeByLanguage: core.serialization.record(
        core.serialization.lazy(async () => (await import("../../../../..")).Language),
        core.serialization.lazyObject(async () => (await import("../../../../..")).v2.FunctionImplementation).optional()
    ),
});

export declare namespace FunctionImplementationForMultipleLanguages {
    interface Raw {
        codeByLanguage: Record<serializers.Language.Raw, serializers.v2.FunctionImplementation.Raw | null | undefined>;
    }
}
