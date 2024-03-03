package adapter

import "fmt"

var DefaultCurrency = "usd"
var DefaultPrecision = 2

var ApiPath = map[string]string{
	"holc": "coins/bitcoin/ohlc",
}

type AutoPeriodElement struct {
	From      int
	To        int
	Period    string
	NumPeriod float64
}

var PeriodThirstyMinutes = "30m"
var PeriodFourHours = "4h"
var PeriodFoursDays = "4d"

// ratio on a day
var NumPeriodThirstyMinutes = 1.0 / 48 // 30m = 1/48 day (30m/1440m)
var NumPeriodFourHours = 1.0 / 6       // 4h/24h
var NumPeriodFoursDays = 4.0           // 4d/1d

var AutoPeriodMap = map[string]AutoPeriodElement{

	PeriodThirstyMinutes: {From: 1, To: 2, NumPeriod: NumPeriodThirstyMinutes},
	PeriodFourHours:      {From: 3, To: 30, NumPeriod: NumPeriodFourHours},
	PeriodFoursDays:      {From: 31, To: -1, NumPeriod: NumPeriodFoursDays},
}

var RedisPeriodLabel = "period"

func (a *AutoPeriodElement) FindPeriodWithDays(days int64) string {
	for key, value := range AutoPeriodMap {
		if days >= int64(value.From) && key == PeriodFoursDays {
			return key
		}

		if days >= int64(value.From) && days <= int64(value.To) {
			return key
		}
	}
	return PeriodThirstyMinutes
}

var HOLCAvailDays = []int{1, 7, 14, 30, 90}

func FindHOLCAvailDays(days int64) string {
	for _, el := range HOLCAvailDays {
		if days < int64(el) {
			return fmt.Sprintf("%d", el)
		}
	}
	return "max"
}
