package adapter

import (
	"context"
	"errors"
	"fmt"

	DTOS "sfvn_test/adapters/coingecko/dtos"

	"github.com/go-resty/resty/v2"
	"github.com/redis/rueidis"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetHOLCData(symbol string, period string, startDate int64, endDate int64) ([]*DTOS.DTOGetHOLCResponse, error)
	SaveHOLC(symbol string, period string, values []*DTOS.DTOGetHOLCResponse) error
	IsLeakData(period string, startDate int64, startDateRes int64) bool
}

type repository struct {
	//redis
	//db gorm
	db        *gorm.DB
	redis     rueidis.Client
	domainApi string
	apiKey    string
}

func RepositoryProvider(db *gorm.DB, redisClient rueidis.Client, domainApi string, apiKey string) Repository {
	db.AutoMigrate(&HOLCEntity{})
	return &repository{
		db:        db,
		redis:     redisClient,
		domainApi: domainApi,
		apiKey:    apiKey,
	}
}

func (r *repository) MakeKeyCacheHOLC(symbol string, period string) string {
	return fmt.Sprintf("%s_%s", symbol, period)
}

func (r *repository) SaveHOLC(symbol string, period string, values []*DTOS.DTOGetHOLCResponse) error {
	//Cache to redis in sorted and save to db

	ctx := context.Background()

	// add to cache
	holcEntities := make([]*HOLCEntity, 0)
	for _, el := range values {

		_holcValue, err := el.ToJsonString()
		if err != nil {
			fmt.Println(err)
			return errors.New("server internal error")
		}

		err = r.redis.Do(ctx, r.redis.B().Zadd().Key(
			r.MakeKeyCacheHOLC(symbol, period)).
			Nx().ScoreMember().
			ScoreMember(float64(el.Time), _holcValue).
			Build()).Error()
		if err != nil {
			fmt.Println(err)
			return errors.New("server internal error")
		}

		holcEntities = append(holcEntities, &HOLCEntity{
			Symbol:    symbol,
			TimeStamp: el.Time,
			Period:    period,
			High:      el.High,
			Low:       el.Low,
			Open:      el.Open,
			Close:     el.Close,
		})
	}

	// save to db
	err := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&holcEntities).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("server internal error")
	}

	return nil
}

func (r *repository) GetHOLCFromCache(symbol string, period string, from int64, to int64) ([]*DTOS.DTOGetHOLCResponse, error) {
	key := r.MakeKeyCacheHOLC(symbol, period)
	ctx := context.Background()
	rawRs, err := r.redis.Do(ctx, r.redis.B().Zrangebyscore().Key(key).
		Min(fmt.Sprintf("%v", from)).
		Max(fmt.Sprintf("%v", to)).Build()).AsStrSlice()
	if err != nil {
		//TODO, return [], if not exists
		return nil, err
	}

	var res = []*DTOS.DTOGetHOLCResponse{}
	for _, jsonStr := range rawRs {
		holcData, err := (&DTOS.DTOGetHOLCResponse{}).FromJsonString(jsonStr)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal error")
		}
		res = append(res, holcData)
	}
	//TODO cast data,
	return res, nil
}

func (r *repository) GetHOLCFromDB(symbol string, period string, from int64, to int64) ([]*DTOS.DTOGetHOLCResponse, error) {
	var holcEntities = []*HOLCEntity{}
	err := r.db.Debug().Where(`
		symbol = ? AND period = ?
		AND time_stamp >= ?
		AND time_stamp <= ?
	`, symbol, period, from, to).Limit(-1).Find(&holcEntities).Error
	if err != nil {
		return nil, err
	}

	var rs = []*DTOS.DTOGetHOLCResponse{}
	for _, entity := range holcEntities {
		rs = append(rs, &DTOS.DTOGetHOLCResponse{
			Time:  entity.TimeStamp,
			High:  entity.High,
			Open:  entity.Open,
			Close: entity.Close,
			Low:   entity.Low,
		})
	}
	return rs, nil
}

func (r *repository) FilterHOLCInTime(values []*DTOS.DTOGetHOLCResponse, startDate int64, endDate int64) []*DTOS.DTOGetHOLCResponse {
	filtered := make([]*DTOS.DTOGetHOLCResponse, 0)
	for _, el := range values {
		if el.Time < startDate || el.Time > endDate {
			continue
		}
		filtered = append(filtered, el)
	}
	return filtered
}

func (r *repository) IsLeakData(period string, startDate int64, startDateRes int64) bool {
	aday := 3600 * 24 * 1000.0 //in miliseconds
	ratioOnDay := float64(startDateRes-startDate) / aday
	return ratioOnDay > AutoPeriodMap[period].NumPeriod
}

func (r *repository) GetHOLCData(symbol string, period string, startDate int64, endDate int64) ([]*DTOS.DTOGetHOLCResponse, error) {
	//	get cache
	// 	-> get db
	// 	-> call api then save to db and set cache 15m

	var rs = []*DTOS.DTOGetHOLCResponse{}
	var err error
	rs, err = r.GetHOLCFromCache(symbol, period, startDate, endDate)
	if err != nil {
		return nil, err
	}

	if len(rs) > 0 {
		filtered := r.FilterHOLCInTime(rs, startDate, endDate)
		// check Leaked data, if data is leak, next step
		if !r.IsLeakData(period, startDate, filtered[0].Time) {
			return filtered, nil
		}
	}

	// get from db

	rs, err = r.GetHOLCFromDB(symbol, period, startDate, endDate)
	if err != nil {
		return nil, err
	}

	if len(rs) > 0 && !r.IsLeakData(period, startDate, rs[0].Time) {
		filtered := r.FilterHOLCInTime(rs, startDate, endDate)
		// check Leaked data, if data is leak, next step
		if !r.IsLeakData(period, startDate, filtered[0].Time) {
			return filtered, nil
		}
	}

	// get from api
	rs, err = r.FetchHOLCFromApi(&DTOS.DTOGetHOLCRequest{
		ID:        symbol,
		Currency:  DefaultCurrency,
		Days:      CalcDuringDayForHOLC(startDate),
		Precision: DefaultPrecision,
	})

	if FindHOLCAvailDays(CalcDuringDayForHOLC(startDate)) == "max" && startDate < rs[0].Time {
		return nil, errors.New("out of range")
	}

	if err != nil {
		return nil, err
	}

	return r.FilterHOLCInTime(rs, startDate, endDate), nil
}

func (r *repository) FetchHOLCFromApi(params *DTOS.DTOGetHOLCRequest) ([]*DTOS.DTOGetHOLCResponse, error) {

	var rawRes [][]interface{}
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"id":          params.ID,
			"vs_currency": DefaultCurrency,
			"days":        FindHOLCAvailDays(params.Days),
			"precision":   fmt.Sprintf("%d", params.Precision),
		}).
		SetHeader("Accept", "application/json").
		SetHeader("x-cg-api-key", r.apiKey).
		SetResult(&rawRes).
		Get(fmt.Sprintf("%s%s", r.domainApi, ApiPath["holc"]))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		fmt.Println(string(resp.Body()))
		return nil, errors.New("call api was failed")
	}

	var res = []*DTOS.DTOGetHOLCResponse{}
	for _, rawElement := range rawRes {
		var el = &DTOS.DTOGetHOLCResponse{}
		el.Time = int64(rawElement[0].(float64))
		el.Close = decimal.NewFromFloat(rawElement[4].(float64))
		el.High = decimal.NewFromFloat(rawElement[2].(float64))
		el.Low = decimal.NewFromFloat(rawElement[3].(float64))
		el.Open = decimal.NewFromFloat(rawElement[1].(float64))
		res = append(res, el)
	}

	//save to in-house
	r.SaveHOLC(params.ID, (&AutoPeriodElement{}).FindPeriodWithDays(params.Days), res)

	return res, nil
}
