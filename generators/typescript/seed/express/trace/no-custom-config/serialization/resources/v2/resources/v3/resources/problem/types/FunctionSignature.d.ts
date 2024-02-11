/**
 * This file was auto-generated by Fern from our API Definition.
 */
import * as serializers from "../../../../../../..";
import * as SeedTrace from "../../../../../../../../api";
import * as core from "../../../../../../../../core";
export declare const FunctionSignature: core.serialization.Schema<serializers.v2.v3.FunctionSignature.Raw, SeedTrace.v2.v3.FunctionSignature>;
export declare namespace FunctionSignature {
    type Raw = FunctionSignature.Void | FunctionSignature.NonVoid | FunctionSignature.VoidThatTakesActualResult;
    interface Void extends serializers.v2.v3.VoidFunctionSignature.Raw {
        type: "void";
    }
    interface NonVoid extends serializers.v2.v3.NonVoidFunctionSignature.Raw {
        type: "nonVoid";
    }
    interface VoidThatTakesActualResult extends serializers.v2.v3.VoidFunctionSignatureThatTakesActualResult.Raw {
        type: "voidThatTakesActualResult";
    }
}
