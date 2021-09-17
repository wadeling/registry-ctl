package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleNotMethodAllowed ...
func HandleNotMethodAllowed(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusMethodNotAllowed)
	errPayload := fmt.Sprintf("%s", "method not allow")
	fmt.Fprintln(w, errPayload)
}

func HandleInternalServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	errPayload := fmt.Sprintf("server error:%v", err)
	fmt.Fprintln(w, errPayload)
}

func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	errPayload := fmt.Errorf("err:%v", err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, errPayload)
}

// WriteJSON response status code will be written automatically if there is an error
func WriteJSON(w http.ResponseWriter, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		SendError(w, err)
		return err
	}

	if _, err = w.Write(b); err != nil {
		return err
	}
	return nil
}
