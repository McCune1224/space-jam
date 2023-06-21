package types

import "time"

type (
	ShipyardTransaction struct {
		WaypointSymbol string    `json:"waypointSymbol"`
		ShipSymbol     string    `json:"shipSymbol"`
		Price          int       `json:"price"`
		AgentSymbol    string    `json:"agentSymbol"`
		Timestamp      time.Time `json:"timestamp"`
	}

	ShipyardShip struct {
		Type          string       `json:"type"`
		Name          string       `json:"name"`
		Description   string       `json:"description"`
		PurchasePrice int          `json:"purchasePrice"`
		Frame         ShipFrame    `json:"frame"`
		Reactor       ShipReactor  `json:"reactor"`
		Engine        ShipEngine   `json:"engine"`
		Modules       []ShipModule `json:"modules"`
		Mounts        []ShipMount  `json:"mounts"`
	}

	Shipyard struct {
		Symbol    string `json:"symbol"`
		ShipTypes []struct {
			Type string `json:"type"`
		} `json:"shipTypes"`
	}
)
