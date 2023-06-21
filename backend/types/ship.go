package types

import "time"

type ShipModificationTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TradeSymbol    string    `json:"tradeSymbol"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}

type ShipRegistration struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}

type ShipFrame struct {
	Symbol         string           `json:"symbol"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	Condition      int              `json:"condition"`
	ModuleSlots    int              `json:"moduleSlots"`
	MountingPoints int              `json:"mountingPoints"`
	FuelCapacity   int              `json:"fuelCapacity"`
	Requirements   ShipRequirements `json:"requirements"`
}

type ShipNavRouteDestination struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type ShipNavRouteDeparture struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type ShipNavRoute struct {
	Destination   ShipNavRouteDestination `json:"destination"`
	Departure     ShipNavRouteDeparture   `json:"departure"`
	DepartureTime time.Time               `json:"departureTime"`
	ArrivalTime   time.Time               `json:"arrivalTime"`
}

type ShipNav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type ShipCrew struct {
	Current  int    `json:"current"`
	Required int    `json:"required"`
	Capacity int    `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int    `json:"morale"`
	Wages    int    `json:"wages"`
}

type ShipRequirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
	Slots int `json:"slots"`
}

type ShipReactor struct {
	Symbol       string           `json:"symbol"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Condition    int              `json:"condition"`
	PowerOutput  int              `json:"powerOutput"`
	Requirements ShipRequirements `json:"requirements"`
}
type ShipEngine struct {
	Symbol       string           `json:"symbol"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Condition    int              `json:"condition"`
	Speed        int              `json:"speed"`
	Requirements ShipRequirements `json:"requirements"`
}

type ShipModule struct {
	Symbol       string           `json:"symbol"`
	Capacity     int              `json:"capacity"`
	Range        int              `json:"range"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Requirements ShipRequirements `json:"requirements"`
}

type ShipMount struct {
	Symbol       string           `json:"symbol"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Strength     int              `json:"strength"`
	Deposits     []string         `json:"deposits"`
	Requirements ShipRequirements `json:"requirements"`
}

type ShipCargoInventory struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       int    `json:"units"`
}

type ShipCargo struct {
	Capacity  int                  `json:"capacity"`
	Units     int                  `json:"units"`
	Inventory []ShipCargoInventory `json:"inventory"`
}

type ShipFuel struct {
	Current  int `json:"current"`
	Capacity int `json:"capacity"`
	Consumed struct {
		Amount    int       `json:"amount"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"consumed"`
}

type Ship struct {
	Symbol       string           `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav          ShipNav          `json:"nav"`
	Crew         ShipCrew         `json:"crew"`
	Frame        ShipFrame        `json:"frame"`
	Reactor      ShipReactor      `json:"reactor"`
	Engine       ShipEngine       `json:"engine"`
	Modules      []ShipModule     `json:"modules"`
	Mounts       []ShipMount      `json:"mounts"`
	Cargo        ShipCargo        `json:"cargo"`
	Fuel         ShipFuel         `json:"fuel"`
}
