package data

import (
	"encoding/json"
	"github.com/google/uuid"
	"testing"
)

// TestQueryResult tests marshaling and unmarshaling of QueryResult using subtests.
func TestQueryResult(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		traceId := uuid.New()
		query := QueryResult{
			Metadata: MetaData{
				Time:     1625254082,
				TraceId:  traceId,
				ItemId:   1,
				BatchId:  10,
				ResultId: 100,
				WorkerId: 42,
			},
			Results: []map[string]interface{}{{"id": 1}, {"id": 2}},
		}

		// Marshal to JSON
		data, err := json.Marshal(query)
		if err != nil {
			t.Fatalf("Failed to marshal QueryResult: %v", err)
		}

		// Unmarshal back to struct
		var result QueryResult
		err = json.Unmarshal(data, &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal QueryResult: %v", err)
		}

		// Validate content
		if result.Metadata.TraceId != query.Metadata.TraceId || result.Metadata.ItemId != query.Metadata.ItemId {
			t.Errorf("Expected %v, got %v", query, result)
		}
	})

	t.Run("SadPath", func(t *testing.T) {
		invalidJSON := []byte(`{"metadata": {"time": "invalid", "traceId": "invalid", "itemId": 1, "batchId": 10, "resultId": 100, "workerId": 42}, "data": [{"id": 1}]}`)

		var result QueryResult
		err := json.Unmarshal(invalidJSON, &result)

		// Expect an error due to invalid data type
		if err == nil {
			t.Error("Expected error but got none")
		}
	})
}
