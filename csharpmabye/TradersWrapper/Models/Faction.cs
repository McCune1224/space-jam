public class FactionTrait
{
    public string? Name { get; set; }
    public string? Description { get; set; }
    public string? Symbol { get; set; }
}
public class Faction
{
    public string? Symbol { get; set; }
    public string? Name { get; set; }
    public string? Description { get; set; }
    public string? Headquarters { get; set; }
    public FactionTrait[]? Traits { get; set; }
    public bool IsRecuiting { get; set; }
}
