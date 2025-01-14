package db

import (
	"FasuhaStats/initializers"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCon *gorm.DB

func ConnectToDb() {
	cfg := initializers.LoadEnvDbConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode,
	)

	var err error
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
