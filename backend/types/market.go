package types

import "time"

type Export struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Import struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Exchange struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MarketTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	TradeSymbol    string    `json:"tradeSymbol"`
	Type           string    `json:"type"`
	Units          int       `json:"units"`
	PricePerUnit   int       `json:"pricePerUnit"`
	TotalPrice     int       `json:"totalPrice"`
	Timestamp      time.Time `json:"timestamp"`
}

type MarketTradeGood struct {
	Symbol        string `json:"symbol"`
	TradeVolume   int    `json:"tradeVolume"`
	Supply        string `json:"supply"`
	PurchasePrice int    `json:"purchasePrice"`
	SellPrice     int    `json:"sellPrice"`
}

type Market struct {
	Symbol       string              `json:"symbol"`
	Exports      []Export            `json:"exports"`
	Imports      []Import            `json:"imports"`
	Exchange     []Exchange          `json:"exchange"`
	Transactions []MarketTransaction `json:"transactions"`
	TradeGoods   []MarketTradeGood   `json:"tradeGoods"`
}
