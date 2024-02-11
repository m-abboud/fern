/**
 * This file was auto-generated by Fern from our API Definition.
 */
import * as serializers from "../../../../../../..";
import * as SeedTrace from "../../../../../../../../api";
import * as core from "../../../../../../../../core";
export declare const TestCaseImplementationDescriptionBoard: core.serialization.Schema<serializers.v2.v3.TestCaseImplementationDescriptionBoard.Raw, SeedTrace.v2.v3.TestCaseImplementationDescriptionBoard>;
export declare namespace TestCaseImplementationDescriptionBoard {
    type Raw = TestCaseImplementationDescriptionBoard.Html | TestCaseImplementationDescriptionBoard.ParamId;
    interface Html {
        type: "html";
        value: string;
    }
    interface ParamId {
        type: "paramId";
        value: serializers.v2.v3.ParameterId.Raw;
    }
}
