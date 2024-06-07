/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.literal.core;

/**
 * This class serves as the base exception for all errors in the SDK.
 */
public class SeedLiteralError extends RuntimeException {
    public SeedLiteralError(String message) {
        super(message);
    }

    public SeedLiteralError(String message, Exception e) {
        super(message, e);
    }
}