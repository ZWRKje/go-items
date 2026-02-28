package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-api/internal/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ItemService interface {
	CreateItem(ctx context.Context, params service.CreateParams) (int, error)
	UpdateItem(ctx context.Context, id int, params service.UpdateParams) error
	DeleteItem(ctx context.Context, id int) (int, error)
	GetItem(ctx context.Context, id int) (service.Item, error)
}

type AppHandler struct {
	service ItemService
}

func NewAppHandler(s ItemService) *AppHandler {
	return &AppHandler{service: s}
}

func (h *AppHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World1!"))
}

func (h *AppHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var item service.CreateParams
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var id int
	id, err = h.service.CreateItem(ctx, item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := IdResponse{Message: "Success", Id: id}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(jsonResponse))
	fmt.Println("create Item")
}

func (h *AppHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var item service.UpdateParams
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.service.UpdateItem(ctx, id, item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("ok"))
}

func (h *AppHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	h.service.DeleteItem(ctx, id)

	response := IdResponse{Message: "Success", Id: id}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(jsonResponse))
}

func (h *AppHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	item, err := h.service.GetItem(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(jsonResponse))
}
