package data

// Item - XML data object representing a splunk input
type Item struct {
	// APIKey - authentication API key to connect to elastic
	APIKey string `xml:"api_key"`

	// CACertPath - TLS-based connections to elastic
	CACertPath string `xml:"ca_cert_path"`

	// Username - elastic username for connection
	Username string `xml:"username"`

	// Password - elastic password for connection
	Password string `xml:"password"`

	// ElasticHost - The target elastic host:port
	ElasticHost string `xml:"elastic_host"`

	// ElasticIndex - the elastic index to be queried
	ElasticIndex string `xml:"elastic_index"`

	// QueryString - a query string to be sent to elastic
	QueryString string `xml:"query_string"`
}
