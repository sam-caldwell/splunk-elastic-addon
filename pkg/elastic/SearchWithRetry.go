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
func SearchWithRetry(es *elasticsearch.Client, index, query string) (*esapi.Response, int, error) {

	var res *esapi.Response

	var err error

	retries := 0

	for attempt := 0; attempt < maxRetries; attempt++ {
		res, err = es.Search(
			es.Search.WithIndex(index),
			es.Search.WithBody(strings.NewReader(fmt.Sprintf(`{"query":{"query_string":{"query":"%s"}}}`, query))),
			es.Search.WithScroll(scrollTimeout),
			es.Search.WithSize(batchSize),
		)
		if err == nil {
			return res, retries, nil
		}
		retries++
		log.Printf("Retrying initial search (attempt %d): %s", attempt+1, err)
		time.Sleep(retryDelay)
	}

	return nil, retries, err

}
