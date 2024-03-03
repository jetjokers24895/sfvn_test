package DTOS

import (
	"errors"
	"fmt"
	"time"
)

type GetHistories struct {
	StartDate string `json:"start_date" query:"start_date"`
	EndDate   string `json:"end_date" query:"end_date"`
	Period    string `json:"period" query:"period"`
	Symbol    string `json:"symbol" query:"symbol"`
}

func (d *GetHistories) Validate() error {
	if d.Symbol == "" {
		return errors.New("symbol is required")
	}

	//validate startDate
	_, err := time.Parse("2006-01-02", d.StartDate)
	if err != nil {
		fmt.Println(err)
		return errors.New("start_date was not right format. It should be `yyyy-mm-dd`, example: 2006-01-02")
	}

	//validate endDate
	_, err = time.Parse("2006-01-02", d.EndDate)
	if err != nil {
		fmt.Println(err)
		return errors.New("end_date was not right format. It should be `yyyy-mm-dd` format, example: 2006-01-02")
	}

	return nil
}
