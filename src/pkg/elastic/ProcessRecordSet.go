package elastic

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	data2 "github.com/sam-caldwell/splunk-elastic-addon/src/pkg/data"
	"log"
	"os"
	"sync"
	"time"
)

// ProcessRecordSet - Spawn three workers to marshal query results and write them to stdout.
func ProcessRecordSet(traceId uuid.UUID, hitsChan <-chan data2.RecordSet, wg *sync.WaitGroup) {
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
			id := 0
			for hit := range hitsChan {
				result := data2.QueryResult{
					Metadata: data2.MetaData{
						Time:     time.Now().Unix(),
						TraceId:  traceId,
						ItemId:   hit.ItemId,
						BatchId:  hit.BatchId,
						ResultId: id,
						WorkerId: workerId,
					},
					Results: hit.Hit,
				}
				if output, err = json.Marshal(result); err != nil {
					log.Printf("[workerId:%d][hitId:%d][ItemId:%d][BatchId: %d]"+
						"error marshalling query result: %v", workerId, id, hit.ItemId, hit.BatchId, err)
				}
				// write the query result to stdout for consumption by splunk
				// ToDo: we need to add the traceId, hitId and workerId to our query results
				_, _ = fmt.Fprintln(os.Stdout, string(output))
				fmt.Println(string(output))
				id++
			}
		}()
	}
}
