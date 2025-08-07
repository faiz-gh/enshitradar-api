package channel

import (
	"fmt"
	"net/http"

	"github.com/faiz-gh/enshitradar-api/types"
	"github.com/faiz-gh/enshitradar-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ChannelStore
}

func NewHandler(store types.ChannelStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/channel", h.HandleGetChannels).Methods(http.MethodGet)
	router.HandleFunc("/channel/{channel_id}", h.HandleGetChannelByID).Methods(http.MethodGet)
	router.HandleFunc("/channel/name/{name}", h.HandleGetChannelByName).Methods(http.MethodGet)
	router.HandleFunc("/channel", h.HandleAddChannel).Methods(http.MethodPost)
}

func (h *Handler) HandleGetChannels(w http.ResponseWriter, r *http.Request) {
	channels, err := h.store.GetChannels()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, channels)
}

func (h *Handler) HandleGetChannelByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars["channel_id"]

	channel, err := h.store.GetChannelByID(channelID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if channel == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("channel not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, channel)
}

func (h *Handler) HandleGetChannelByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	channel, err := h.store.GetChannelByName(name)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if channel == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("channel not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, channel)
}

func (h *Handler) HandleAddChannel(w http.ResponseWriter, r *http.Request) {
	var channelPayload types.AddChannelPayload
	if err := utils.ParseJSON(r, &channelPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var channel *types.Channel
	channel, err := h.store.AddChannel(channelPayload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, channel)
}
