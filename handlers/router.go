package handlers

import (
	"github.com/gorilla/mux"
	"github.com/wadeling/registry-ctl/api"
	"github.com/wadeling/registry-ctl/api/registry/gc"
	"github.com/wadeling/registry-ctl/config"
	"net/http"
)

func newRouter(conf config.Configuration) http.Handler {
	// create the root rooter
	rootRouter := mux.NewRouter()
	rootRouter.StrictSlash(true)
	rootRouter.HandleFunc("/api/health", api.Health).Methods("GET")

	rootRouter.Path("/api/registry/gc").Methods(http.MethodPost).Handler(gc.NewHandler(conf.RegistryConfig))
	return rootRouter
}
