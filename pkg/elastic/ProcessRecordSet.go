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
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			var (
				err    error
				output []byte
			)
			defer wg.Done()
			for hit := range hitsChan {
				if output, err = json.Marshal(hit); err != nil {
					log.Printf("[traceId:%s]error marshalling query result: %v", traceId.String(), err)
				}
				// write the query result to stdout for consumption by splunk
				// ToDo: we need to add the traceId to our query results
				fmt.Println(string(output))
			}
		}()
	}
}
