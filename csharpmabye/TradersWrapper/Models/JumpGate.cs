// {
//   "jumpRange": 0,
//   "factionSymbol": "string",
//   "connectedSystems": [
//     {
//       "symbol": "string",
//       "sectorSymbol": "string",
//       "type": "NEUTRON_STAR",
//       "factionSymbol": "string",
//       "x": 0,
//       "y": 0,
//       "distance": 0
//     }
//   ]
// }

public class JumpGate
{
    public int JumpRange { get; set; }
    public string? FactionSymbol { get; set; }
    public List<ConnectedSystem>? ConnectedSystems { get; set; }
}
