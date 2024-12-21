package elastic

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sam-caldwell/splunk-elastic-addon/pkg/data"
	"github.com/sam-caldwell/splunk-elastic-addon/pkg/log"
	"io"
	"sync"
	"time"
)

// ProcessItem - Run a query from Splunk, query elastic and return the results back to Splunk.
func ProcessItem(item data.Item, wg *sync.WaitGroup) {
	var (
		err       error
		es        *elasticsearch.Client
		res       *esapi.Response
		retries   int
		startTime time.Time = time.Now()
		workerWg  sync.WaitGroup
	)

	defer wg.Done()

	if es, err = CreateClient(item.ElasticHost, item.Username, item.Password, item.APIKey, item.CACertPath); err != nil {
		log.Printf("Error creating Elasticsearch client: %s", err)
		return
	}

	if res, retries, err = SearchWithRetry(es, item.ElasticIndex, item.QueryString); err != nil {
		log.Printf("Initial search failed after retries: %s", err)
		return
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("Error closing Body: %s", err)
		}
	}(res.Body)

	hitsChan := make(chan interface{}, hitQueueSize)
	for i := 0; i < 3; i++ {
		workerWg.Add(1)
		go ProcessRecordSet(hitsChan, &workerWg)
	}

	scrollID := ""
	batchID := 0
	totalRetries := retries

	for {
		batchID++
		var result map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			fmt.Printf("Error decoding response body: %s", err)
			return
		}

		hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
		if len(hits) == 0 {
			break
		}

		for _, hit := range hits {
			hitsChan <- hit
		}

		scrollID = result["_scroll_id"].(string)

		res, retries, err = ScrollWithRetry(es, scrollID)
		totalRetries += retries
		if err != nil {
			log.Printf("Scroll request failed after retries: %s", err)
			break
		}

		log.Printf("Batch %d processed successfully", batchID)
	}

	close(hitsChan)
	workerWg.Wait()

	_, err = es.ClearScroll(es.ClearScroll.WithScrollID(scrollID))
	if err != nil {
		log.Printf("Failed to clear scroll: %s", err)
	}

	endTime := time.Now()
	log.Printf("Query completed - Start: %s, End: %s, Total Batches: %d, Total Retries: %d", startTime, endTime, batchID, totalRetries)
}
