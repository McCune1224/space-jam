package types

type Agent struct {
	AccountID       string `json:"accountId"`
	Symbol          string `json:"symbol"`
	Headquarters    string `json:"headquarters"`
	Credits         int    `json:"credits"`
	StartingFaction string `json:"startingFaction"`
}
