/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as SeedTrace from "../../..";
import * as core from "../../../../core";

export type Error = SeedTrace.admin.storeTracedWorkspaceV2.Error._Unknown;

export declare namespace Error {
    interface _Unknown {
        errorName: void;
        content: core.Fetcher.Error;
    }

    interface _Visitor<_Result> {
        _other: (value: core.Fetcher.Error) => _Result;
    }
}

export const Error = {
    _unknown: (fetcherError: core.Fetcher.Error): SeedTrace.admin.storeTracedWorkspaceV2.Error._Unknown => {
        return {
            errorName: undefined,
            content: fetcherError,
        };
    },

    _visit: <_Result>(
        value: SeedTrace.admin.storeTracedWorkspaceV2.Error,
        visitor: SeedTrace.admin.storeTracedWorkspaceV2.Error._Visitor<_Result>
    ): _Result => {
        switch (value.errorName) {
            default:
                return visitor._other(value as any);
        }
    },
} as const;
