/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as serializers from "../../..";
import * as SeedTrace from "../../../../api";
import * as core from "../../../../core";

export const ExceptionV2: core.serialization.Schema<serializers.ExceptionV2.Raw, SeedTrace.ExceptionV2> =
    core.serialization
        .union("type", {
            generic: core.serialization.lazyObject(async () => (await import("../../..")).ExceptionInfo),
            timeout: core.serialization.object({}),
        })
        .transform<SeedTrace.ExceptionV2>({
            transform: (value) => {
                switch (value.type) {
                    case "generic":
                        return SeedTrace.ExceptionV2.generic(value);
                    case "timeout":
                        return SeedTrace.ExceptionV2.timeout();
                    default:
                        return SeedTrace.ExceptionV2._unknown(value);
                }
            },
            untransform: ({ _visit, ...value }) => value as any,
        });

export declare namespace ExceptionV2 {
    type Raw = ExceptionV2.Generic | ExceptionV2.Timeout;

    interface Generic extends serializers.ExceptionInfo.Raw {
        type: "generic";
    }

    interface Timeout {
        type: "timeout";
    }
}
