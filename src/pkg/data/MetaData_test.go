package data

import (
	"encoding/json"
	"github.com/google/uuid"
	"testing"
)

// TestMetaData tests marshaling and unmarshaling of MetaData using subtests.
func TestMetaData(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		traceId := uuid.New()
		meta := MetaData{
			Time:     1625254082,
			TraceId:  traceId,
			ItemId:   1,
			BatchId:  10,
			ResultId: 100,
			WorkerId: 42,
		}

		// Marshal to JSON
		data, err := json.Marshal(meta)
		if err != nil {
			t.Fatalf("Failed to marshal MetaData: %v", err)
		}

		// Unmarshal back to struct
		var result MetaData
		err = json.Unmarshal(data, &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal MetaData: %v", err)
		}

		// Validate content
		if result.TraceId != meta.TraceId || result.ItemId != meta.ItemId {
			t.Errorf("Expected %v, got %v", meta, result)
		}
	})

	t.Run("SadPath", func(t *testing.T) {
		invalidJSON := []byte(`{"time": "invalid", "traceId": "invalid", "itemId": 1, "batchId": 10, "resultId": 100, "workerId": 42}`)

		var result MetaData
		err := json.Unmarshal(invalidJSON, &result)

		// Expect an error due to invalid data type
		if err == nil {
			t.Error("Expected error but got none")
		}
	})
}
