"use strict";
/**
 * This file was auto-generated by Fern from our API Definition.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.SeedErrorPropertyError = void 0;
class SeedErrorPropertyError extends Error {
    constructor(errorName) {
        super();
        this.errorName = errorName;
        Object.setPrototypeOf(this, SeedErrorPropertyError.prototype);
    }
}
exports.SeedErrorPropertyError = SeedErrorPropertyError;
