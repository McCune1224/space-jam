package types

type ConnectedSystem struct {
	Symbol        string `json:"symbol"`
	SectorSymbol  string `json:"sectorSymbol"`
	Type          string `json:"type"`
	FactionSymbol string `json:"factionSymbol"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Distance      int    `json:"distance"`
}
