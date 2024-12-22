package elastic

import (
	"github.com/google/uuid"
	"github.com/sam-caldwell/splunk-elastic-addon/src/pkg/data"
	"sync"
	"testing"
)

// TestProcessItem tests the ProcessItem function.
func TestProcessItem(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		traceId := uuid.New()
		item := data.Item{
			ElasticHost:  "http://localhost:9200",
			ElasticIndex: "test-index",
			QueryString:  "query",
		}
		var wg sync.WaitGroup
		wg.Add(1)
		ProcessItem(traceId, 1, item, &wg)
		wg.Wait()
	})

	t.Run("ErrorCreatingClient", func(t *testing.T) {
		traceId := uuid.New()
		item := data.Item{
			ElasticHost:  "",
			ElasticIndex: "test-index",
			QueryString:  "query",
		}
		var wg sync.WaitGroup
		wg.Add(1)
		ProcessItem(traceId, 1, item, &wg)
		wg.Wait()
	})
}
