package handler

import (
	"github.com/soyhouston256/go-api/middleware"
	"net/http"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	p := newPerson(storage)
	mux.HandleFunc("/v1/persons/create", middleware.Log(p.create))

	mux.HandleFunc("/v1/persons/get-all", middleware.Log(p.getAll))

	mux.HandleFunc("/v1/persons/update", middleware.Log(p.update))

	mux.HandleFunc("/v1/persons/delete", middleware.Log(p.delete))

	mux.HandleFunc("/v1/persons/get-by-id", middleware.Log(middleware.Authenticated(p.getById)))

}

func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)
	mux.HandleFunc("/v1/login", h.login)
}
