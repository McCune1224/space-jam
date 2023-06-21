package types

import "time"

type Cooldown struct {
	ShipSymbol       string    `json:"shipSymbol"`
	totalSeconds     int       `json:"totalSeconds"`
	remainingSeconds int       `json:"remainingSeconds"`
	Expiration       time.Time `json:"expiration"`
}
