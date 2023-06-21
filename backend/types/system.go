package types

type SystemWaypoint struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}
type System struct {
	Symbol       string           `json:"symbol"`
	SectorSymbol string           `json:"sectorSymbol"`
	Type         string           `json:"type"`
	X            int              `json:"x"`
	Y            int              `json:"y"`
	Waypoints    []SystemWaypoint `json:"systemWaypoints"`
	Factions     []struct {
		Symbol string `json:"symbol"`
	} `json:"factions"`
}
