package gc

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/wadeling/registry-ctl/api"
	"net/http"
	"os/exec"
	"time"
)

func NewHandler(registryConf string) http.Handler {
	return &handler{
		registryConf: registryConf,
	}
}

type handler struct {
	registryConf string
}

// ServeHTTP ...
func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		h.start(w, req)
	default:
		api.HandleNotMethodAllowed(w)
	}
}

// Result ...
type Result struct {
	Status    bool      `json:"status"`
	Msg       string    `json:"msg"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
}

// start ...
func (h *handler) start(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Start to execute garbage collection...")
	cmd := exec.Command("sh", "-c", "registry garbage-collect --delete-untagged=false "+h.registryConf)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	start := time.Now()
	if err := cmd.Run(); err != nil {
		logrus.Errorf("Fail to execute GC: %v, command err: %s", err, errBuf.String())
		api.HandleInternalServerError(w, err)
		return
	}

	gcr := Result{true, outBuf.String(), start, time.Now()}
	if err := api.WriteJSON(w, gcr); err != nil {
		logrus.Errorf("failed to write response: %v", err)
		return
	}
	logrus.Info("Successful to execute garbage collection...")
}
