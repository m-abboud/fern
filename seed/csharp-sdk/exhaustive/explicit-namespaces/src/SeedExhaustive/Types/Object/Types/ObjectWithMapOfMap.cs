using System.Text.Json.Serialization;

#nullable enable

namespace SeedExhaustive.Types.Object;

public record ObjectWithMapOfMap
{
    [JsonPropertyName("map")]
    public Dictionary<string, Dictionary<string, string>> Map { get; init; } =
        new Dictionary<string, Dictionary<string, string>>();
}
