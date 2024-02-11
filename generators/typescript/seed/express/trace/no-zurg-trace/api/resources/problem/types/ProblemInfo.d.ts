/**
 * This file was auto-generated by Fern from our API Definition.
 */
import * as SeedTrace from "../../..";
export interface ProblemInfo {
    problemId: SeedTrace.ProblemId;
    problemDescription: SeedTrace.ProblemDescription;
    problemName: string;
    problemVersion: number;
    files: Record<SeedTrace.Language, SeedTrace.ProblemFiles | undefined>;
    inputParams: SeedTrace.VariableTypeAndName[];
    outputType: SeedTrace.VariableType;
    testcases: SeedTrace.TestCaseWithExpectedResult[];
    methodName: string;
    supportsCustomTestCases: boolean;
}
