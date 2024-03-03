package user

import (
	"context"
	"errors"
	"fmt"

	"sfvn_test/config"
	dtosReq "sfvn_test/dtos/requests"
	dtosRes "sfvn_test/dtos/responses"
	"sfvn_test/repositories"
	"sfvn_test/utils"

	adapter "sfvn_test/adapters/coingecko"

	"github.com/redis/rueidis"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ServiceWrapper interface {
	GetHistoriesOfSymbol(context.Context, *dtosReq.GetHistories, string) ([]*dtosRes.DTOGetHistoryPriceResponse, error)
}

type service struct {
	holcRepo          repositories.HOLCRepo
	coingeckorService adapter.CoingeckoService
}

func ProviderService(db *gorm.DB, config *config.Config, redisClient rueidis.Client) ServiceWrapper {
	return &service{
		holcRepo: repositories.HOLCRepoProvider(db, redisClient),
		coingeckorService: adapter.CoingeckoServiceProvider(
			config.Coingeckor.Domain,
			config.Coingeckor.ApiKey, db, redisClient),
	}
}

func (s *service) CalculateChangePricePercent(latestPrice, closePrice decimal.Decimal) float64 {
	if latestPrice.Equal(decimal.Zero) {
		return 0.0
	}

	return closePrice.Sub(latestPrice).DivRound(closePrice, 2).InexactFloat64() * 100.0
}

func (s *service) GetHistoriesOfSymbol(ctx context.Context, queryInput *dtosReq.GetHistories, uid string) ([]*dtosRes.DTOGetHistoryPriceResponse, error) {
	startTimestamp, err := utils.GetTimeStampStartOfTheDay(queryInput.StartDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	endTimestamp, err := utils.GetTimeStampEndOfTheDay(queryInput.EndDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	holcData, err := s.coingeckorService.GetHOLC(queryInput.Symbol, queryInput.Period, startTimestamp, endTimestamp)
	if err != nil {
		return nil, err
	}

	var res = make([]*dtosRes.DTOGetHistoryPriceResponse, 0)
	latestPrice, err := s.holcRepo.GetLatestPriceUserGot(uid, queryInput.Symbol)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal error")
	}

	for _, el := range holcData {
		res = append(res, &dtosRes.DTOGetHistoryPriceResponse{
			Time:   el.Time,
			High:   el.High,
			Open:   el.Open,
			Close:  el.Close,
			Low:    el.Low,
			Change: s.CalculateChangePricePercent(latestPrice, el.Close),
		})
	}

	//cache latest price
	if len(holcData) > 0 {
		if err := s.holcRepo.CacheLatestPriceUserGot(uid, queryInput.Symbol, holcData[len(holcData)-1].Close); err != nil {
			fmt.Println(err)
		}
	}

	return res, nil
}
