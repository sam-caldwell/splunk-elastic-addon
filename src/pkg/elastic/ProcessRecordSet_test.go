package elastic

import (
	"github.com/google/uuid"
	"github.com/sam-caldwell/splunk-elastic-addon/pkg/data"
	"sync"
	"testing"
)

// TestProcessRecordSet tests the ProcessRecordSet function.
func TestProcessRecordSet(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		traceId := uuid.New()
		hitsChan := make(chan data.RecordSet, 10)
		var wg sync.WaitGroup

		// Simulate input data
		go func() {
			for i := 0; i < 5; i++ {
				hitsChan <- data.RecordSet{
					ItemId:  i,
					BatchId: 0,
					Hit:     map[string]any{"id": i},
				}
			}
			close(hitsChan)
		}()

		// Process the data
		ProcessRecordSet(traceId, hitsChan, &wg)
		wg.Wait()
	})

	t.Run("EmptyChannel", func(t *testing.T) {
		traceId := uuid.New()
		hitsChan := make(chan data.RecordSet)
		var wg sync.WaitGroup

		// Close channel immediately to simulate no input
		close(hitsChan)

		// Process the data
		ProcessRecordSet(traceId, hitsChan, &wg)
		wg.Wait()
	})

	t.Run("InvalidData", func(t *testing.T) {
		traceId := uuid.New()
		hitsChan := make(chan data.RecordSet, 1)
		var wg sync.WaitGroup

		// Send invalid data
		hitsChan <- data.RecordSet{
			ItemId:  -1,
			BatchId: -1,
			Hit:     nil,
		}
		close(hitsChan)

		// Process the data
		ProcessRecordSet(traceId, hitsChan, &wg)
		wg.Wait()
	})
}
