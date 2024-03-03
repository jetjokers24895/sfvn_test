package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/rueidis"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type HOLCRepo interface {
	CacheLatestPriceUserGot(uid, symbol string, closePrice decimal.Decimal) error
	GetLatestPriceUserGot(uid, symbol string) (decimal.Decimal, error)
}

type holcRepo struct {
	db    *gorm.DB
	redis rueidis.Client
}

func HOLCRepoProvider(db *gorm.DB,
	redis rueidis.Client) HOLCRepo {
	return &holcRepo{
		db:    db,
		redis: redis,
	}
}

func (r *holcRepo) CacheLatestPriceUserGot(uid, symbol string, closePrice decimal.Decimal) error {
	err := r.redis.Do(context.Background(), r.redis.B().Set().Key(
		r.MakeKeyCachePriceUserGot(uid, symbol)).Value(closePrice.String()).Build()).Error()
	if err != nil {
		fmt.Println(err)
		return errors.New("internal error")
	}
	return nil
}

func (r *holcRepo) GetLatestPriceUserGot(uid, symbol string) (decimal.Decimal, error) {
	value, err := r.redis.Do(context.Background(), r.redis.B().Get().Key(
		r.MakeKeyCachePriceUserGot(uid, symbol)).Build()).AsFloat64()
	if err != nil {
		fmt.Println(err)
		return decimal.Zero, nil
	}

	return decimal.NewFromFloat(value), nil
}

func (r *holcRepo) MakeKeyCachePriceUserGot(uid string, symbol string) string {
	return fmt.Sprintf("%s:%s", uid, symbol)
}
