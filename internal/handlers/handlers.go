package handlers

import (
	"fmt"
	"net/http"
)

type AppHandler struct {
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}
func (a *AppHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World1!"))
}

func (a *AppHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create Item")
}

func (a *AppHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update Item")
}

func (a *AppHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete Item")
}

func (a *AppHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get Item")
}
