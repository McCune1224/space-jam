package types

type JumpGate struct {
	JumpRange        int               `json:"jumpRange"`
	FactionSymbol    string            `json:"factionSymbol"`
	ConnectedSystems []ConnectedSystem `json:"connectedSystems"`
}
