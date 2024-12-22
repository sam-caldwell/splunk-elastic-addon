package data

import (
	"encoding/xml"
	"testing"
)

// TestItem_MarshalUnmarshal_HappyPath tests marshaling and unmarshalling of Item with valid data.
func TestData_Item(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		item := Item{
			APIKey:       "key1",
			CACertPath:   "/path/to/cert",
			Username:     "user1",
			Password:     "pass1",
			ElasticHost:  "host1:9200",
			ElasticIndex: "index1",
			QueryString:  "query1",
		}

		// Marshal to XML
		data, err := xml.Marshal(item)
		if err != nil {
			t.Fatalf("Failed to marshal Item: %v", err)
		}

		// Unmarshal back to struct
		var result Item
		err = xml.Unmarshal(data, &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal Item: %v", err)
		}

		// Validate content
		if result != item {
			t.Errorf("Expected %v, got %v", item, result)
		}
	})
	t.Run("sad path invalid data", func(t *testing.T) {
		invalidXML := []byte(`
<Item>
	<api_key>key1</api_key>
	<elastic_host>host1:9200</elastichost>
</Item>`)

		var result Item
		err := xml.Unmarshal(invalidXML, &result)

		// Expect no error but missing fields should be empty
		if err == nil {
			t.Fatalf("expected error not found")
		}
	})
}
