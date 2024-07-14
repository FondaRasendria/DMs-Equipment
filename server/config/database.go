package config

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "host=localhost user=companion password=fondarasendria dbname=5eCompanion port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Connected to Database")
}

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=localhost port=5432 user=companion password=fondarasendria dbname=5eCompanion sslmode=disable")

	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Connected to Database")

	return db
}
