package elastic

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
	"time"
)

// SearchWithRetry - Search the elastic cluster and retry on failure
func SearchWithRetry(itemId int, es *elasticsearch.Client, index,
	query string) (res *esapi.Response, retries int, err error) {

	for attempt := 1; attempt <= maxRetries; attempt++ {
		res, err = es.Search(
			es.Search.WithIndex(index),
			es.Search.WithBody(
				strings.NewReader(
					fmt.Sprintf(`{"query":{"query_string":{"query":"%s"}}}`, query))),
			es.Search.WithScroll(scrollTimeout),
			es.Search.WithSize(batchSize),
		)
		if err == nil {
			return res, retries, nil
		}
		retries++
		log.Printf("[item %d][attempt %d/%d]Retrying initial search: %s", itemId, attempt, maxRetries, err)
		time.Sleep(retryDelay)
	}

	return nil, retries, err

}
