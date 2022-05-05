package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func CreateDBConnection(host string, port int, username string, password string, dbname string) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error on opening db", err)
	}

	dberr := db.Ping()
	if dberr != nil {
		log.Fatal("error on connecting to database: ", dberr)
	}

	log.Println("successfully connected to postgres database")

	return db
}
