package handlers

import (
	"net/http"
	"strconv"

	"github.com/NethuNe/ditto/library"

	log "github.com/sirupsen/logrus"
)

const NUM_BYTES string = "numBytes"

// todo: GenerateKiloBytes, GenerateMegaBytes, GenerateBytes for ease, perhaps

// GenerateBytes expects a request with query param `numBytes`, representing
// the integer number of bytes to generate.
// It generates that many random bytes and returns a 200 response with those bytes.
func GenerateBytes(w http.ResponseWriter, r *http.Request) {
	numBytesRequest := r.URL.Query().Get(NUM_BYTES)
	if numBytesRequest == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status": "/generateBytes expects query param numBytes to be a non-zero integer"}`))
		return
	}
	numBytes, err := strconv.Atoi(numBytesRequest)
	if err != nil {
		log.Warnf("generateBytes called with invalid numBytes value %s", numBytesRequest[0:min(10, len(numBytesRequest))])
		return
	}
	rawBytes, err := library.GenerateRawByteData(numBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(rawBytes)
}
