package database

import (
	"app/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Connection *gorm.DB
}

func Connect(config config.Config) (*Database, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=Europe/Moscow",
		config.PostgreSQL.User,
		config.PostgreSQL.Password,
		config.PostgreSQL.DBname,
		config.PostgreSQL.Host,
		config.PostgreSQL.Port,
		config.PostgreSQL.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return &Database{
		Connection: db,
	}, nil
}
