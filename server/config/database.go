package config

import (
	"database/sql"
	"fmt"
	"golang-crud/helper"

	"github.com/rs/zerolog/log"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=localhost port=5432 user=companion password=fondarasendria dbname=5eCompanion sslmode=disable")

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to Database")

	return db
}
