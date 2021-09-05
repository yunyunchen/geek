package data

import (
	"geek/04/configs"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"time"
)


var Provider = wire.NewSet(NewData, NewDatabase, NewUserRepo)

type Data struct {
	database *gorm.DB
}

func NewDatabase(config *configs.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DatabaseDsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func NewData(database *gorm.DB) *Data {
	return &Data{
		database: database,
	}
}
