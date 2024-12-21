package elastic

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"sync"
)

// ProcessRecordSet - Spawn three workers to marshal query results and write them to stdout.
func ProcessRecordSet(traceId uuid.UUID, hitsChan <-chan any, wg *sync.WaitGroup) {
	const (
		workerCount = 3
	)
	for workerId := 0; workerId < workerCount; workerId++ {
		wg.Add(1)
		go func() {
			var (
				err    error
				output []byte
			)
			defer wg.Done()
			hitId := 0
			for hit := range hitsChan {
				if output, err = json.Marshal(hit); err != nil {
					log.Printf("[workerId:%d][hitId:%d]error marshalling query result: %v", workerId, hitId, err)
				}
				// write the query result to stdout for consumption by splunk
				// ToDo: we need to add the traceId, hitId and workerId to our query results
				fmt.Println(string(output))
				hitId++
			}
		}()
	}
}
