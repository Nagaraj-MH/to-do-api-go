package main

import (
	"database/sql"
	"log"
	"to-do-api/cmd/api"
	"to-do-api/db"
)

func main() {
	db, ok := db.NewPostgre(map[string]string{})
	if ok != nil {
		log.Fatal(ok)
	}
	ConnectDB(db)
	server := api.CreateServer(":8080", nil)
	err := server.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(db *sql.DB) {
	db.Ping()
	db.Close()
}
