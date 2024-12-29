package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/signup", h.CreateUser).Methods("POST")
	router.HandleFunc("/login", h.LoginUser).Methods("POST")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("testing")
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

}
