/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as errors from "../../../../../../errors";
import * as SeedExhaustive from "../../../../..";

export class ErrorWithUnionBody extends errors.SeedExhaustiveError {
    constructor(body: SeedExhaustive.types.Animal) {
        super({
            message: "ErrorWithUnionBody",
            statusCode: 400,
            body: body,
        });
        Object.setPrototypeOf(this, ErrorWithUnionBody.prototype);
    }
}