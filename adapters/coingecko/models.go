package adapter

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type HOLCEntity struct {
	gorm.Model

	Symbol    string `gorm:"uniqueIndex:idx_searchkey,"`
	TimeStamp int64  `gorm:"uniqueIndex:idx_searchkey,"`
	Period    string `gorm:"uniqueIndex:idx_searchkey,"`

	High  decimal.Decimal `gorm:"type:decimal(16,2);"`
	Low   decimal.Decimal `gorm:"type:decimal(16,2);"`
	Open  decimal.Decimal `gorm:"type:decimal(16,2);"`
	Close decimal.Decimal `gorm:"type:decimal(16,2);"`
}

func (a *HOLCEntity) TableName() string {
	return "holc"
}
