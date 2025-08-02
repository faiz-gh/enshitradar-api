package user

import (
	"fmt"
	"net/http"

	"github.com/faiz-gh/enshitradar-api/types"
	"github.com/faiz-gh/enshitradar-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user", h.handleAddUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", h.handleGetUserByID).Methods(http.MethodGet)
}

func (h *Handler) handleAddUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.store.AddUser()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}

func (h *Handler) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.store.GetUserByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
