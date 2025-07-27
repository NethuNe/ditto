package handlers

import (
	"fmt"
	"net/http"

	"github.com/NethuNe/ditto/simulator"
	log "github.com/sirupsen/logrus"
)

func Simulate(w http.ResponseWriter, r *http.Request) {
	requestPath := r.RequestURI
	log.Infof("Simulating response to request with path %s", requestPath)
	resp := simulator.GetConfiguredResponse(requestPath)
	if resp.Endpoint == nil {
		w.WriteHeader(http.StatusNotFound)
		statusString := fmt.Sprintf(`{"status": "response not found for configured request %s"}`, requestPath)
		log.Warn(statusString)
		w.Write([]byte(statusString))
		return
	}

	endpoint := resp.Endpoint
	w.WriteHeader(endpoint.Status)
	for k, v := range endpoint.Headers {
		w.Header().Add(k, v)
	}
	_, err := w.Write([]byte(endpoint.Body))
	if err != nil {
		log.Errorf("Error writing body to response in simulation request for path %s", requestPath)
	}
}
