package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ItemHandlers interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	CreateItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	GetItem(w http.ResponseWriter, r *http.Request)
}

type AppRouter struct {
	Router      *chi.Mux
	itemHandler ItemHandlers
}

func (a *AppRouter) MountHandlers() {
	a.Router.Get("/create", a.itemHandler.CreateItem)
	a.Router.Get("/update", a.itemHandler.UpdateItem)
	a.Router.Get("/delete", a.itemHandler.DeleteItem)
	a.Router.Get("/get", a.itemHandler.GetItem)
	a.Router.Get("/", a.itemHandler.HealthCheck)
}

func NewAppRouter(r *chi.Mux, handler ItemHandlers) AppRouter {
	return AppRouter{
		Router:      r,
		itemHandler: handler,
	}
}
