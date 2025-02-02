package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgre(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
