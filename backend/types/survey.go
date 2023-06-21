package types

import "time"

type Survey struct {
	Signature string `json:"signature"`
	Symbol    string `json:"symbol"`
	Deposits  []struct {
		Symbole string `json:"symbole"`
	} `json:"deposits"`
	Expiration time.Time `json:"expiration"`
	Size       string    `json:"size"`
}
