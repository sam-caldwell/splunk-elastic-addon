package elastic

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"os"
)

// CreateClient - Create an elastic client
func CreateClient(host, username, password, apiKey, caCertPath string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{Addresses: []string{host}}

	if trim(username) != "" && trim(password) != "" {
		cfg.Username = username
		cfg.Password = password
	}

	if trim(apiKey) != "" {
		cfg.APIKey = apiKey
	}

	if trim(caCertPath) != "" {
		caCert, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA certificate: %w", err)
		}
		cfg.CACert = caCert
	}

	return elasticsearch.NewClient(cfg)
}
