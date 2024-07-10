using System.Text.Json.Serialization;

#nullable enable

namespace SeedTrace.V2;

public record FunctionImplementation
{
    [JsonPropertyName("impl")]
    public required string Impl { get; init; }

    [JsonPropertyName("imports")]
    public string? Imports { get; init; }
}
