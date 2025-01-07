package main

import (
	"database/sql"
	"log"
	"os"
	"to-do-api/cmd/api"
	"to-do-api/db"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	log.Print(os.Getenv("CONNSTRING"))
	db, ok := db.NewPostgre(os.Getenv("CONNSTRING"))
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
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}
