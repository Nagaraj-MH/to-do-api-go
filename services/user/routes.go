package user

import (
	"net/http"
	"to-do-api/types"
	"to-do-api/utils"

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
	// get json payload
	var payload types.RegisterUsers
	err := utils.ParseJSON(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

}
