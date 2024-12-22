package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"time"
)

// ScrollWithRetry - iterate over the paginated elastic query results (with retries if needed)
func ScrollWithRetry(itemId, batchId int, es *elasticsearch.Client, scrollId *string) (res *esapi.Response, retries int, err error) {

	for attempt := 1; attempt <= maxRetries; attempt++ {

		if res, err = es.Scroll(es.Scroll.WithScrollID(*scrollId), es.Scroll.WithScroll(scrollTimeout)); err != nil {
			return res, retries, nil
		}

		retries++

		log.Printf("[item %d][batch %d][attempt %d/%d]Retrying scroll request: %s",
			itemId, batchId, attempt, maxRetries, err)

		time.Sleep(retryDelay)

	}

	return nil, retries, err

}
