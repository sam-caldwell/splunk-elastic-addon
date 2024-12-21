package elastic

import (
	"encoding/json"
	"fmt"
	"sync"
)

func ProcessRecordSet(hitsChan <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for hit := range hitsChan {
		output, _ := json.Marshal(hit)
		fmt.Println(string(output))
	}
}
