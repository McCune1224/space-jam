package types



type ExtractionYield struct {
	Symbol string `json:"symbol"`
	Units  int    `json:"units"`
}

type Extraction struct {
	ShipSymbol string            `json:"shipSymbol"`
	Yield      []ExtractionYield `json:"yield"`
}
