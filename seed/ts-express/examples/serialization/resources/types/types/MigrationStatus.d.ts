/**
 * This file was auto-generated by Fern from our API Definition.
 */
import * as serializers from "../../..";
import * as SeedExamples from "../../../../api";
import * as core from "../../../../core";
export declare const MigrationStatus: core.serialization.Schema<serializers.MigrationStatus.Raw, SeedExamples.MigrationStatus>;
export declare namespace MigrationStatus {
    type Raw = "RUNNING" | "FAILED" | "FINISHED";
}