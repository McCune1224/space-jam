package types

type FactionSymbols struct {
	Symbols []string `json:"symbols"`
}

type FactionTrait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Faction struct {
	Symbol       string         `json:"symbol"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Headquarters string         `json:"headquarters"`
	Traits       []FactionTrait `json:"traits"`
	IsRecruiting bool           `json:"isRecruiting"`
}
