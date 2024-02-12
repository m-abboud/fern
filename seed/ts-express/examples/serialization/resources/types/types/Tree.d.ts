/**
 * This file was auto-generated by Fern from our API Definition.
 */
import * as serializers from "../../..";
import * as SeedExamples from "../../../../api";
import * as core from "../../../../core";
export declare const Tree: core.serialization.ObjectSchema<serializers.Tree.Raw, SeedExamples.Tree>;
export declare namespace Tree {
    interface Raw {
        nodes?: serializers.Node.Raw[] | null;
    }
}