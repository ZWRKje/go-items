package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-api/internal/api/router"
	"simple-api/internal/config"
	"simple-api/internal/handlers"
	"simple-api/internal/repository"
	"simple-api/internal/service"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
	Config config.Config
}

func CreateServer() *Server {
	config, err := config.Init()
	if err != nil {
		log.Fatal("Init config err")
	}

	server := &Server{
		Router: chi.NewRouter(),
		Config: *config,
	}

	return server
}

func main() {
	server := CreateServer()
	itemRepo := repository.NewItemsRepository(server.Config)
	itemSer := service.NewItemService(itemRepo)
	handler := handlers.NewAppHandler(itemSer)

	r := router.NewAppRouter(server.Router, handler)
	r.MountHandlers()
	fmt.Println("server running on port:8080")
	http.ListenAndServe(":8080", r.Router)
}
