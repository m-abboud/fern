using System.Net.Http;

namespace SeedMixedCase.Core;

public static class HttpMethodExtensions
{
    public static readonly HttpMethod Patch = new("PATCH");
}
