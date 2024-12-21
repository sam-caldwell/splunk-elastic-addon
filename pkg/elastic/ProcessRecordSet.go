package elastic

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// ProcessRecordSet - Spawn three workers to marshal query results and write them to stdout.
func ProcessRecordSet(hitsChan <-chan any, wg *sync.WaitGroup) {
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
					log.Printf("error marshalling query result: %v", err)
				}
				// write the query result to stdout for consumption by splunk
				fmt.Println(string(output))
			}
		}()
	}
}
