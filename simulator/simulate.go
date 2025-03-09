package simulator

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var config *SimulatorConfig
var tempPath string = "simulator/testSchema.json"

func init() {
	var err error
	config, err = ParseConfig(tempPath)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to initialize config with inputted path %s and error: %s",
			tempPath, err.Error()))
		log.Exit(1)
	}
}

func GetConfiguredResponse(requestPath string) ConfiguredResponse {
	resp := ConfiguredResponse{}
	if config == nil {
		log.Error("Config not initialized, unable to simulate")
		return resp
	}
	for _, endpoint := range config.Endpoints {
		if endpoint.Path == requestPath {
			resp.Endpoint = &endpoint
			return resp
		}
	}
	log.Warn(fmt.Sprintf("Unable to find config entry for request %s", requestPath))
	return resp
}
