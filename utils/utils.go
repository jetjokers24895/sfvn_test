package utils

import (
	"fmt"
	"time"
)

// output is timestamp in milliseconds
// timezone: utc
func GetTimeStampStartOfTheDay(dateStr string) (int64, error) {
	d, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:01", dateStr))
	if err != nil {
		return 0, nil
	}

	return d.UnixMilli(), nil
}

// output is timestamp in milliseconds
// timezone: utc
func GetTimeStampEndOfTheDay(dateStr string) (int64, error) {
	d, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s 23:59:59", dateStr))
	if err != nil {
		return 0, nil
	}

	return d.UnixMilli(), nil
}
