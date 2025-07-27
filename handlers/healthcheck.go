package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Info("Replying to healthcheck with 200 status ok")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
