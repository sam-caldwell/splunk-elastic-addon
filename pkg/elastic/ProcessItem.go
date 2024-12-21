package elastic

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	"github.com/sam-caldwell/splunk-elastic-addon/pkg/data"
	"log"
	"sync"
	"time"
)

// ProcessItem - Run a query from Splunk, query elastic and return the results back to Splunk.
func ProcessItem(traceId uuid.UUID, itemId int, item data.Item, streamWorkingGroup *sync.WaitGroup) {
	var (
		err           error
		es            *elasticsearch.Client
		res           *esapi.Response
		workerWg      sync.WaitGroup
		searchRetries int
		scrollRetries int
	)
	defer streamWorkingGroup.Done()

	startTime := time.Now().Unix()

	if es, err = CreateClient(item.ElasticHost, item.Username, item.Password, item.APIKey, item.CACertPath); err != nil {
		log.Printf("[item_%d]Error creating Elasticsearch client: %s", itemId, err)
		return
	}

	if res, searchRetries, err = SearchWithRetry(itemId, es, item.ElasticIndex, item.QueryString); err != nil {
		log.Printf("[item_%d]Initial search failed after %d retries: %s", itemId, searchRetries, err)
		return
	}

	defer closeReader(res.Body)

	hitsChan := make(chan any, hitQueueSize)
	ProcessRecordSet(traceId, hitsChan, &workerWg)

	scrollId := ""
	batchId := 0

	for {
		var (
			result map[string]any
			hits   []any
		)

		if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
			log.Printf("[item_%d][batch_%d]Error decoding response body: %s", itemId, batchId, err)
			return
		}

		if hits = result["hits"].(map[string]any)["hits"].([]any); len(hits) == 0 {
			break
		}

		for _, hit := range hits {
			hitsChan <- hit
		}

		scrollId = result["_scroll_id"].(string)

		res, scrollRetries, err = ScrollWithRetry(itemId, batchId, es, &scrollId)

		if err != nil {
			log.Printf("[item_%d][batch_%d]Scroll request failed after %d retries: %s",
				itemId, batchId, scrollRetries, err)
			break
		}

		log.Printf("[item_%d][batch_%d]Batch processed successfully", itemId, batchId)

		batchId++
	}

	close(hitsChan)
	workerWg.Wait()

	if _, err = es.ClearScroll(es.ClearScroll.WithScrollID(scrollId)); err != nil {
		log.Printf("[item_%d][batch_%d]Failed to clear scroll: %s", itemId, batchId, err)
	}

	log.Printf("[item_%d]Query completed - Start: %d, End: %d, totalBatches: %d, searchRetries: %d, "+
		"scrollRetries: %d", itemId, startTime, time.Now().Unix(), batchId+1, searchRetries, scrollRetries)
}
