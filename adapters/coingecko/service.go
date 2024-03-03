package adapter

import (
	"errors"
	"fmt"

	DTOS "sfvn_test/adapters/coingecko/dtos"

	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

type CoingeckoService interface {
	GetHOLC(symbol string, period string, startDate int64, endDate int64) ([]*DTOS.DTOGetHOLCResponse, error)
	ValidatePeriod(period string, startDate int64, endDate int64) error
}

type coingecko struct {
	repo Repository
}

func CoingeckoServiceProvider(domain string, apikey string, db *gorm.DB, redisClient rueidis.Client) CoingeckoService {
	return &coingecko{
		repo: RepositoryProvider(db, redisClient, domain, apikey),
	}
}

func (c *coingecko) ValidatePeriod(period string, startDate int64, endDate int64) error {
	// check period format
	el, ok := AutoPeriodMap[period]
	if !ok {
		return fmt.Errorf("invalid period. Period should be: %s, %s, %s", PeriodThirstyMinutes,
			PeriodFourHours, PeriodFoursDays)
	}

	days := int(CalcDuringDayForHOLC(startDate))
	// check period with days
	if days < el.From {
		return errors.New("invalid period")
	}

	if days > el.To && el.To != -1 {
		return errors.New("invalid period with start_date")
	}

	// time get data must larger than or equal period
	inTime := CalcDuringDay(startDate, endDate)
	if inTime < int64(el.NumPeriod) {
		return errors.New("end_date - start_date must larger than or equal period")
	}

	return nil
}

func (c *coingecko) GetHOLC(symbol string, period string, startDate int64, endDate int64) ([]*DTOS.DTOGetHOLCResponse, error) {
	if err := c.ValidatePeriod(period, startDate, endDate); err != nil {
		return nil, err
	}
	res, err := c.repo.GetHOLCData(symbol, period, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return res, nil
}
