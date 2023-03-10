package router

import (
	"go-identity-manager/cmd/router/routes"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	return ConfigRouter(router)
}

func ConfigRouter(r *mux.Router) *mux.Router {
	packageRoutes := routes.IdentityRouter

	for _, route := range packageRoutes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
