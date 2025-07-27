package simulator

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Returns the *SimulatorConfig or error if encountered by parsing
// the given input file. Expects that filePath is a valid file in the environment
func ParseConfig(filePath string) (*SimulatorConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	configBytes := make([]byte, fileInfo.Size())
	numRead, err := file.Read(configBytes)
	if numRead != len(configBytes) || (err != nil && err != io.EOF) {
		return nil, fmt.Errorf("unexpected error occurred in reading config; expected %d bytes, got %d and error %s",
			len(configBytes), numRead, err)
	}
	config := &SimulatorConfig{}
	err = json.Unmarshal(configBytes, config)
	return config, err
}
