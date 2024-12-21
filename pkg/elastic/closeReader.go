package elastic

import (
	"io"
	"log"
)

// closeReader - close the resource and log any error
func closeReader(Body io.ReadCloser) {
	if err := Body.Close(); err != nil {
		log.Printf("Error closing Body: %s", err)
	}
}
