package elastic

import "testing"

// TestCreateClient tests the CreateClient function.
func TestCreateClient(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		client, err := CreateClient("http://localhost:9200", "user", "pass", "", "")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if client == nil {
			t.Fatal("Expected client, got nil")
		}
	})

	t.Run("InvalidCACertPath", func(t *testing.T) {
		_, err := CreateClient("http://localhost:9200", "", "", "", "invalid/path/to/ca.cert")
		if err == nil {
			t.Fatal("Expected error for invalid CA cert path, got none")
		}
	})

	t.Run("MissingHost", func(t *testing.T) {
		_, err := CreateClient("", "", "", "", "")
		if err == nil {
			t.Fatal("Expected error for missing host, got none")
		}
	})
}
