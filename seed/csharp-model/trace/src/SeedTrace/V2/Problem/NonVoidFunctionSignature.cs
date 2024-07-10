using System.Text.Json.Serialization;
using SeedTrace.V2;

#nullable enable

namespace SeedTrace.V2;

public record NonVoidFunctionSignature
{
    [JsonPropertyName("parameters")]
    public IEnumerable<Parameter> Parameters { get; init; } = new List<Parameter>();

    [JsonPropertyName("returnType")]
    public required object ReturnType { get; init; }
}
