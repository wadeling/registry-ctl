package api

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if err := WriteJSON(w, "healthy"); err != nil {
		log.Errorf("Failed to write response: %v", err)
		return
	}
}
