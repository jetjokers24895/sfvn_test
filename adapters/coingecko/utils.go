package adapter

import (
	"math"
	"time"
)

func CalcDuringDayForHOLC(startDate int64) int64 {
	now := time.Now().UnixMilli()
	aDay := 3600 * 1000 * 24 // a day in milliseconds
	return int64(math.Round((float64(now) - float64(startDate)) / float64(aDay)))
}

func CalcDuringDay(startDate int64, endDate int64) int64 {
	aDay := 3600 * 1000 * 24 // a day in milliseconds
	return int64(math.Round((float64(endDate) - float64(startDate)) / float64(aDay)))
}
