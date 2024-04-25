# frozen_string_literal: true

require_relative "../../requests"
require "async"

module SeedExhaustiveClient
  class ReqWithHeadersClient
    # @return [SeedExhaustiveClient::RequestClient]
    attr_reader :request_client

    # @param request_client [SeedExhaustiveClient::RequestClient]
    # @return [SeedExhaustiveClient::ReqWithHeadersClient]
    def initialize(request_client:)
      @request_client = request_client
    end

    # @param x_test_endpoint_header [String]
    # @param request [String]
    # @param request_options [SeedExhaustiveClient::RequestOptions]
    # @return [Void]
    def get_with_custom_header(x_test_endpoint_header:, request:, request_options: nil)
      @request_client.conn.post do |req|
        req.options.timeout = request_options.timeout_in_seconds unless request_options&.timeout_in_seconds.nil?
        req.headers["Authorization"] = request_options.token unless request_options&.token.nil?
        req.headers = {
          **req.headers,
          **(request_options&.additional_headers || {}),
          "X-TEST-ENDPOINT-HEADER": x_test_endpoint_header
        }.compact
        req.body = { **(request || {}), **(request_options&.additional_body_parameters || {}) }.compact
        req.url "#{@request_client.get_url(request_options: request_options)}/test-headers/custom-header"
      end
    end
  end

  class AsyncReqWithHeadersClient
    # @return [SeedExhaustiveClient::AsyncRequestClient]
    attr_reader :request_client

    # @param request_client [SeedExhaustiveClient::AsyncRequestClient]
    # @return [SeedExhaustiveClient::AsyncReqWithHeadersClient]
    def initialize(request_client:)
      @request_client = request_client
    end

    # @param x_test_endpoint_header [String]
    # @param request [String]
    # @param request_options [SeedExhaustiveClient::RequestOptions]
    # @return [Void]
    def get_with_custom_header(x_test_endpoint_header:, request:, request_options: nil)
      Async do
        @request_client.conn.post do |req|
          req.options.timeout = request_options.timeout_in_seconds unless request_options&.timeout_in_seconds.nil?
          req.headers["Authorization"] = request_options.token unless request_options&.token.nil?
          req.headers = {
            **req.headers,
            **(request_options&.additional_headers || {}),
            "X-TEST-ENDPOINT-HEADER": x_test_endpoint_header
          }.compact
          req.body = { **(request || {}), **(request_options&.additional_body_parameters || {}) }.compact
          req.url "#{@request_client.get_url(request_options: request_options)}/test-headers/custom-header"
        end
      end
    end
  end
end