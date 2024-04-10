/**
 * This file was auto-generated by Fern from our API Definition.
 */

import * as core from "../../../../core";
import * as SeedPagination from "../../..";
import urlJoin from "url-join";
import * as serializers from "../../../../serialization";
import * as errors from "../../../../errors";

export declare namespace Users {
    interface Options {
        environment: core.Supplier<string>;
        token?: core.Supplier<core.BearerToken | undefined>;
    }

    interface RequestOptions {
        timeoutInSeconds?: number;
        maxRetries?: number;
    }
}

export class Users {
    constructor(protected readonly _options: Users.Options) {}

    public async listWithCursorPagination(
        request: SeedPagination.ListUsersCursorPaginationRequest = {},
        requestOptions?: Users.RequestOptions
    ): Promise<SeedPagination.ListUsersPaginationResponse> {
        const { page, perPage, order, startingAfter } = request;
        const _queryParams: Record<string, string | string[] | object | object[]> = {};
        if (page != null) {
            _queryParams["page"] = page.toString();
        }

        if (perPage != null) {
            _queryParams["per_page"] = perPage.toString();
        }

        if (order != null) {
            _queryParams["order"] = order;
        }

        if (startingAfter != null) {
            _queryParams["starting_after"] = startingAfter;
        }

        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "/users"),
            method: "GET",
            headers: {
                Authorization: await this._getAuthorizationHeader(),
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/pagination",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version,
            },
            contentType: "application/json",
            queryParameters: _queryParams,
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
        });
        if (_response.ok) {
            return await serializers.ListUsersPaginationResponse.parseOrThrow(_response.body, {
                unrecognizedObjectKeys: "passthrough",
                allowUnrecognizedUnionMembers: true,
                allowUnrecognizedEnumValues: true,
                breadcrumbsPrefix: ["response"],
            });
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedPaginationError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedPaginationError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedPaginationTimeoutError();
            case "unknown":
                throw new errors.SeedPaginationError({
                    message: _response.error.errorMessage,
                });
        }
    }

    public async listWithOffsetPagination(
        request: SeedPagination.ListUsersOffsetPaginationRequest = {},
        requestOptions?: Users.RequestOptions
    ): Promise<SeedPagination.ListUsersPaginationResponse> {
        const { page, perPage, order, startingAfter } = request;
        const _queryParams: Record<string, string | string[] | object | object[]> = {};
        if (page != null) {
            _queryParams["page"] = page.toString();
        }

        if (perPage != null) {
            _queryParams["per_page"] = perPage.toString();
        }

        if (order != null) {
            _queryParams["order"] = order;
        }

        if (startingAfter != null) {
            _queryParams["starting_after"] = startingAfter;
        }

        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "/users"),
            method: "GET",
            headers: {
                Authorization: await this._getAuthorizationHeader(),
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/pagination",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version,
            },
            contentType: "application/json",
            queryParameters: _queryParams,
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
        });
        if (_response.ok) {
            return await serializers.ListUsersPaginationResponse.parseOrThrow(_response.body, {
                unrecognizedObjectKeys: "passthrough",
                allowUnrecognizedUnionMembers: true,
                allowUnrecognizedEnumValues: true,
                breadcrumbsPrefix: ["response"],
            });
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedPaginationError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedPaginationError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedPaginationTimeoutError();
            case "unknown":
                throw new errors.SeedPaginationError({
                    message: _response.error.errorMessage,
                });
        }
    }

    public async listWithExtendedResults(
        request: SeedPagination.ListUsersExtendedRequest = {},
        requestOptions?: Users.RequestOptions
    ): Promise<SeedPagination.ListUsersExtendedResponse> {
        const { cursor } = request;
        const _queryParams: Record<string, string | string[] | object | object[]> = {};
        if (cursor != null) {
            _queryParams["cursor"] = cursor;
        }

        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "/users"),
            method: "GET",
            headers: {
                Authorization: await this._getAuthorizationHeader(),
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/pagination",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version,
            },
            contentType: "application/json",
            queryParameters: _queryParams,
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
        });
        if (_response.ok) {
            return await serializers.ListUsersExtendedResponse.parseOrThrow(_response.body, {
                unrecognizedObjectKeys: "passthrough",
                allowUnrecognizedUnionMembers: true,
                allowUnrecognizedEnumValues: true,
                breadcrumbsPrefix: ["response"],
            });
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedPaginationError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedPaginationError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedPaginationTimeoutError();
            case "unknown":
                throw new errors.SeedPaginationError({
                    message: _response.error.errorMessage,
                });
        }
    }

    public async listUsernames(
        request: SeedPagination.ListUsernamesRequest = {},
        requestOptions?: Users.RequestOptions
    ): Promise<SeedPagination.UsernameCursor> {
        const { startingAfter } = request;
        const _queryParams: Record<string, string | string[] | object | object[]> = {};
        if (startingAfter != null) {
            _queryParams["starting_after"] = startingAfter;
        }

        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "/users"),
            method: "GET",
            headers: {
                Authorization: await this._getAuthorizationHeader(),
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/pagination",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version,
            },
            contentType: "application/json",
            queryParameters: _queryParams,
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
        });
        if (_response.ok) {
            return await serializers.UsernameCursor.parseOrThrow(_response.body, {
                unrecognizedObjectKeys: "passthrough",
                allowUnrecognizedUnionMembers: true,
                allowUnrecognizedEnumValues: true,
                breadcrumbsPrefix: ["response"],
            });
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedPaginationError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedPaginationError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedPaginationTimeoutError();
            case "unknown":
                throw new errors.SeedPaginationError({
                    message: _response.error.errorMessage,
                });
        }
    }

    public async listWithGlobalConfig(
        request: SeedPagination.ListWithGlobalConfigRequest = {},
        requestOptions?: Users.RequestOptions
    ): Promise<SeedPagination.UsernameContainer> {
        const { offset } = request;
        const _queryParams: Record<string, string | string[] | object | object[]> = {};
        if (offset != null) {
            _queryParams["offset"] = offset.toString();
        }

        const _response = await core.fetcher({
            url: urlJoin(await core.Supplier.get(this._options.environment), "/users"),
            method: "GET",
            headers: {
                Authorization: await this._getAuthorizationHeader(),
                "X-Fern-Language": "JavaScript",
                "X-Fern-SDK-Name": "@fern/pagination",
                "X-Fern-SDK-Version": "0.0.1",
                "X-Fern-Runtime": core.RUNTIME.type,
                "X-Fern-Runtime-Version": core.RUNTIME.version,
            },
            contentType: "application/json",
            queryParameters: _queryParams,
            timeoutMs: requestOptions?.timeoutInSeconds != null ? requestOptions.timeoutInSeconds * 1000 : 60000,
            maxRetries: requestOptions?.maxRetries,
        });
        if (_response.ok) {
            return await serializers.UsernameContainer.parseOrThrow(_response.body, {
                unrecognizedObjectKeys: "passthrough",
                allowUnrecognizedUnionMembers: true,
                allowUnrecognizedEnumValues: true,
                breadcrumbsPrefix: ["response"],
            });
        }

        if (_response.error.reason === "status-code") {
            throw new errors.SeedPaginationError({
                statusCode: _response.error.statusCode,
                body: _response.error.body,
            });
        }

        switch (_response.error.reason) {
            case "non-json":
                throw new errors.SeedPaginationError({
                    statusCode: _response.error.statusCode,
                    body: _response.error.rawBody,
                });
            case "timeout":
                throw new errors.SeedPaginationTimeoutError();
            case "unknown":
                throw new errors.SeedPaginationError({
                    message: _response.error.errorMessage,
                });
        }
    }

    protected async _getAuthorizationHeader() {
        const bearer = await core.Supplier.get(this._options.token);
        if (bearer != null) {
            return `Bearer ${bearer}`;
        }

        return undefined;
    }
}