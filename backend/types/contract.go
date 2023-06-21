package types

import "time"

// Different types of contracts that can be created
const (
	CONTRACT_PROCUREMENT = "PROCUREMENT"
	CONTRACT_TRANSPORT   = "TRANSPORT"
	CONTRACT_SHUTTLE     = "SHUTTLE"
)

type ContractPayment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}

type ContractDeliverGood struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFufilled     int    `json:"unitsFufilled"`
}

type ContractTerms struct {
	Deadline     time.Time             `json:"deadline"`
	Payment      ContractPayment       `json:"payment"`
	DeliverGoods []ContractDeliverGood `json:"deliver"`
}

type Contract struct {
	ID               string        `json:"id"`
	FactionSymbol    string        `json:"factionSymbol"`
	Type             string        `json:"type"`
	Terms            ContractTerms `json:"terms"`
	Accepted         bool          `json:"accepted"`
	Fulfilled        bool          `json:"fulfilled"`
	Expiration       time.Time     `json:"expiration"`
	DeadlineToAccept time.Time     `json:"deadlineToAccept"`
}
