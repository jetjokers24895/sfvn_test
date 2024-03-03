package adapter

import (
	"encoding/json"
	"errors"

	"github.com/shopspring/decimal"
)

type DTOGetHOLCRequest struct {
	ID        string `json:"id"`
	Currency  string `json:"vs_currency"`
	Days      int64  `json:"days"`
	Interval  string `json:"interval"`
	Precision int    `json:"precision"`
}

type DTOGetHOLCResponse struct {
	Time  int64           `json:"time"` // unix timestamp in milliseconds
	High  decimal.Decimal `json:"high"`
	Open  decimal.Decimal `json:"open"`
	Close decimal.Decimal `json:"close"`
	Low   decimal.Decimal `json:"low"`
}

func (dts *DTOGetHOLCResponse) ToJsonString() (string, error) {
	str, err := json.Marshal(dts)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return string(str), nil
}

func (dts *DTOGetHOLCResponse) FromJsonString(jsonStr string) (*DTOGetHOLCResponse, error) {
	err := json.Unmarshal([]byte(jsonStr), &dts)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return dts, nil
}
