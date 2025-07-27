package library

import (
	"crypto/rand"
	"errors"

	log "github.com/sirupsen/logrus"
)

const MAX_BYTE_COUNT int = 1e7 // 10mb

// GenerateRawByteData generates numBytes number of raw bytes.
// Caps at MAX_BYTE_COUNT.
func GenerateRawByteData(numBytes int) ([]byte, error) {
	if numBytes <= 0 {
		return []byte{}, errors.New("numBytes must be > 0")
	}
	if numBytes > MAX_BYTE_COUNT {
		numBytes = MAX_BYTE_COUNT
	}
	rawBytes := make([]byte, numBytes)
	written, err := rand.Read(rawBytes)
	if written != numBytes || err != nil {
		log.Errorf("Failed to write %d raw bytes. Wrote %d with err %s", numBytes, written, err.Error())
		return []byte{}, err
	}
	return rawBytes, nil
}
