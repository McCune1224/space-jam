package types

import "time"

type WaypointOrbital struct {
	Symbol string `json:"symbol"`
}

type WaypointFaction struct {
	Symbol string `json:"symbol"`
}

type WaypointTrait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type WaypointChart struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	SubmittedBy    string    `json:"submittedBy"`
	SubmittedOn    time.Time `json:"submittedOn"`
}

type Waypoint struct {
	Symbol       string            `json:"symbol"`
	Type         string            `json:"type"`
	SystemSymbol string            `json:"systemSymbol"`
	X            int               `json:"x"`
	Y            int               `json:"y"`
	Orbitals     []WaypointOrbital `json:"orbitals"`
	Faction      WaypointFaction   `json:"faction"`
	Traits       []WaypointTrait   `json:"traits"`
}
