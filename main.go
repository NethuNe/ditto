package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/NethuNe/ditto/handlers"
	log "github.com/sirupsen/logrus"
)

var MIN_PORT int = 1
var MAX_PORT int = 65535
var PORT int = 1985 // bear dahn

func main() {
	portStr := os.Getenv("port")
	port, err := strconv.Atoi(portStr)
	if err != nil || port < MIN_PORT || port > MAX_PORT {
		log.Info("No port or invalid port specified, starting on default ", PORT)
	} else {
		PORT = port
	}

	log.SetLevel(log.DebugLevel)
	log.Info("Starting server on port ", PORT)

	http.HandleFunc("/healthcheck", handlers.Healthcheck)
	// get config

	// set config

	// wildcard fallthru
	http.HandleFunc("/", handlers.Simulate)

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

// flag for file path of input file
