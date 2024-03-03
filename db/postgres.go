package db

import (
	"fmt"
	"sfvn_test/config"

	// _ "github.com/go-sql-driver/mysql" // nolint
	// _ "gorm.io/gorm/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDNS(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Ho_Chi_Minh",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port, cfg.DB.SSLMode)
}

func InitPostgres(cfg *config.Config) *gorm.DB {
	dataSourceName := getDNS(cfg)
	db, err := gorm.Open(postgres.Open(dataSourceName))
	if err != nil {
		panic(err.Error())
	}
	return db
}
