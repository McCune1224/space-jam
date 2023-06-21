public class Deliver
{
    public string? TradeSymbol { get; set; }
    public string? DestinationSymbol { get; set; }
    public int UnitsRequired { get; set; }
    public int UnitsFulfilled { get; set; }
}

public class ContractTerms
{

    public DateTime Deadline { get; set; }
    public ContractPayment? Payment { get; set; }
    public List<Deliver>? Deliver { get; set; }
}


