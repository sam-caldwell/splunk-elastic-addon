package elastic

import "time"

const (

	// maxRetries - the number of times we will retry interactions with elastic
	maxRetries = 3

	// retryDelay - number of seconds between retries when interacting with elastic
	retryDelay = 2 * time.Second

	// scrollTimeout - number of minutes before the query context will be kept alive when querying elastic
	scrollTimeout = 2 * time.Minute

	// batchSize - the number of records which will be processed in each page of query results.
	batchSize = 500

	// hitQueueSize - the number of records ingested from elastic which will be processed to splunk
	hitQueueSize = 1000
)
