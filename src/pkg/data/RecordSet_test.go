package data

import (
	"encoding/json"
	"testing"
)

// TestRecordSet tests marshaling and unmarshaling of RecordSet using subtests.
func TestRecordSet(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		record := RecordSet{
			ItemId:  1,
			BatchId: 100,
			Hit:     map[string]interface{}{"key": "value"},
		}

		// Marshal to JSON
		data, err := json.Marshal(record)
		if err != nil {
			t.Fatalf("Failed to marshal RecordSet: %v", err)
		}

		// Unmarshal back to struct
		var result RecordSet
		err = json.Unmarshal(data, &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal RecordSet: %v", err)
		}

		// Validate content
		if result.ItemId != record.ItemId || result.BatchId != record.BatchId {
			t.Errorf("Expected %v, got %v", record, result)
		}
	})

	t.Run("SadPath", func(t *testing.T) {
		invalidJSON := []byte(`{"ItemId": "invalid", "BatchId": 100, "Hit": {"key": "value"}}`)

		var result RecordSet
		err := json.Unmarshal(invalidJSON, &result)

		// Expect an error due to invalid data type
		if err == nil {
			t.Error("Expected error but got none")
		}
	})
}
