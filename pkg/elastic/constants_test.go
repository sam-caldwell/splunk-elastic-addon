package elastic

import (
	"testing"
	"time"
)

func TestConstants(t *testing.T) {
	t.Run("maxRetries", func(t *testing.T) {
		if maxRetries != 3 {
			t.Fatal("expected 3 got ", maxRetries)
		}
	})
	t.Run("retryDelay", func(t *testing.T) {
		if retryDelay != 2*time.Second {
			t.Fatal("expected 2 seconds got ", retryDelay)
		}
	})
	t.Run("scrollTimeout", func(t *testing.T) {
		if scrollTimeout != 2*time.Minute {
			t.Fatal("expected 2 minutes got ", scrollTimeout)
		}
	})
	t.Run("batchSize", func(t *testing.T) {
		if batchSize != 500 {
			t.Fatal("expected 500 got ", batchSize)
		}
	})
	t.Run("hitQueueSize", func(t *testing.T) {
		if hitQueueSize != 1000 {
			t.Fatal("expected 1000 got ", hitQueueSize)
		}
	})
}
