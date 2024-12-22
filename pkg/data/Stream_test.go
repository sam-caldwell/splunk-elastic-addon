package data

import (
	"encoding/xml"
	"testing"
)

// TestStream_MarshalUnmarshal_HappyPath tests marshaling and unmarshaling of Stream with valid data.
func TestStream_MarshalUnmarshal_HappyPath(t *testing.T) {
	stream := Stream{
		Items: []Item{
			{APIKey: "key1", CACertPath: "/path/to/cert", Username: "user1", Password: "pass1", ElasticHost: "host1:9200", ElasticIndex: "index1", QueryString: "query1"},
			{APIKey: "key2", CACertPath: "/path/to/cert2", Username: "user2", Password: "pass2", ElasticHost: "host2:9200", ElasticIndex: "index2", QueryString: "query2"},
		},
	}

	// Marshal to XML
	data, err := xml.Marshal(stream)
	if err != nil {
		t.Fatalf("Failed to marshal Stream: %v", err)
	}

	// Unmarshal back to struct
	var result Stream
	err = xml.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal Stream: %v", err)
	}

	// Validate length
	if len(result.Items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result.Items))
	}

	// Validate content
	for i, item := range result.Items {
		if item.APIKey != stream.Items[i].APIKey || item.ElasticHost != stream.Items[i].ElasticHost {
			t.Errorf("Item mismatch at index %d", i)
		}
	}
}

// TestStream_MarshalUnmarshal_SadPath tests unmarshalling of Stream with invalid data.
func TestStream_MarshalUnmarshal_SadPath(t *testing.T) {
	invalidXML := []byte(`
<Stream>
	<item>
		<api_key>key1</api_key>
		<elastic_host>host1:9200</elastic_host>
	</item>
	<item>
		<api_key></api_key>
		<elastic_host>host2:9200</elastic_host>
</Stream>`)

	var result Stream
	err := xml.Unmarshal(invalidXML, &result)

	// Expect an error due to missing required fields
	if err == nil {
		t.Error("Expected error but got none")
	}
}
