package simulator

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

const MIN_FUZZY_STRENGTH_PERCENT float64 = 0.6

var config *SimulatorConfig
var tempPath string = "testSchema.json" // todo: configurable

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
	var bestResponse fuzzyMatchStrength
	for _, endpoint := range config.Endpoints {
		currMatchStrength := fuzzyMatchResponse(requestPath, &endpoint)
		if currMatchStrength.matchStrength > bestResponse.matchStrength {
			bestResponse = currMatchStrength
		}
	}
	if bestResponse.matchStrength < MIN_FUZZY_STRENGTH_PERCENT {
		log.Warn(fmt.Sprintf("Unable to find config entry for request %s", requestPath))
		return resp
	}
	resp.Endpoint = bestResponse.endpoint
	return resp
}

type fuzzyMatchStrength struct {
	matchStrength float64
	endpoint      *Endpoint
}

// fuzzyMatchResponse matches content within a slash delimiter char-by-char.
// Returns 0 match strength if the number of slashes differ, else returns
// the matchStrength [0, 1.0] and considered endpoint.
// Match strength is calculated as the sum of matched characters divided by
// the number of non-slash characters in endpoint.Path.
func fuzzyMatchResponse(requestPath string, endpoint *Endpoint) fuzzyMatchStrength {
	reqSplit := strings.Split(requestPath, "/")
	endSplit := strings.Split(endpoint.Path, "/")
	if len(reqSplit) != len(endSplit) {
		return fuzzyMatchStrength{
			0,
			endpoint,
		}
	}
	matchCount := 0.0
	for i := range reqSplit {
		for j := 0; j < min(len(reqSplit[i]), len(endSplit[i])); j++ {
			if reqSplit[i][j] == endSplit[i][j] {
				matchCount++
			}
		}
	}
	return fuzzyMatchStrength{
		matchCount / float64(len(endpoint.Path)-(len(endSplit)-1)),
		endpoint,
	}
}
