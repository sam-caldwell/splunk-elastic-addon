package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"time"
)

func ScrollWithRetry(es *elasticsearch.Client, scrollID string) (*esapi.Response, int, error) {
	var res *esapi.Response
	var err error
	retries := 0
	for attempt := 0; attempt < maxRetries; attempt++ {
		res, err = es.Scroll(
			es.Scroll.WithScrollID(scrollID),
			es.Scroll.WithScroll(scrollTimeout),
		)
		if err == nil {
			return res, retries, nil
		}
		retries++
		log.Printf("Retrying scroll request (attempt %d): %s", attempt+1, err)
		time.Sleep(retryDelay)
	}
	return nil, retries, err
}
