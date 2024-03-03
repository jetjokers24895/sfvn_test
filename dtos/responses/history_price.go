package dtos

import "github.com/shopspring/decimal"

type DTOGetHistoryPriceResponse struct {
	Time   int64           `json:"time"` // unix timestamp in milliseconds
	High   decimal.Decimal `json:"high"`
	Open   decimal.Decimal `json:"open"`
	Close  decimal.Decimal `json:"close"`
	Low    decimal.Decimal `json:"low"`
	Change float64         `json:"change"`
}
