package database

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"

	"github.com/rs/zerolog/log"
)

func CreateDBConnection(host string, port int, username string, password string, dbname string) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	log.Info().Msg("database dsn " + dsn)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal().Msgf("error on opening db", err)
	}

	dberr := db.Ping()
	if dberr != nil {
		log.Fatal().Msgf("error on connecting to database: ", dberr)
	}

	log.Info().Msg("successfully connected to postgres database")

	return db
}
