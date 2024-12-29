package api

import (
	"database/sql"
	"log"
	"net/http"
	"to-do-api/services/user"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func CreateServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}
func (server *ApiServer) RunServer() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Print("Server started at port 8080")

	return http.ListenAndServe(server.addr, router)
}
