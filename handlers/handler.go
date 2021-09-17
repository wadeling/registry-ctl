package handlers

import (
	gorilla_handlers "github.com/gorilla/handlers"
	"github.com/wadeling/registry-ctl/config"
	"net/http"
	"os"
)

// NewHandlerChain returns a gorilla router which is wrapped by  authenticate handler
// and logging handler
func NewHandlerChain(conf config.Configuration) http.Handler {
	h := newRouter(conf)
	h = gorilla_handlers.LoggingHandler(os.Stdout, h)
	return h
}
