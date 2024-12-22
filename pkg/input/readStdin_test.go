package input

import (
	"os"
	"testing"
)

// TestReadStdin tests the ReadStdin function with subtests.
func TestReadStdin(t *testing.T) {
	var writer *os.File
	oldStdin := os.Stdin

	t.Cleanup(func() {
		_ = os.Stdin.Close()
		os.Stdin = oldStdin
		_ = writer.Close()
	})

	createMockStdin := func() {
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}
		writer = w
		os.Stdin = r
	}
	loadTestData := func(input string) {
		if _, err := writer.Write([]byte(input)); err != nil {
			t.Fatal(err)
		}
		_ = writer.Close()
	}

	t.Run("HappyPath", func(t *testing.T) {
		input := "test input"
		maxSize := int64(20)

		// Simulate stdin input using a pipe and write test data to it.
		createMockStdin()
		loadTestData(input)

		// Test the function
		if data, err := ReadStdin(maxSize); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		} else if string(data) != input {
			t.Fatalf("Expected %q, got %q", input, string(data))
		}
	})
	t.Run("ExceedsLimit", func(t *testing.T) {
		input := "test input"
		maxSize := int64(4)

		// Simulate stdin input using a pipe and write test data to it.
		createMockStdin()
		loadTestData(input)

		// Test the function
		if data, err := ReadStdin(maxSize); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		} else if string(data) != "test" {
			t.Fatalf("Expected %q, got %q", input, string(data))
		}
	})

	t.Run("EmptyInput", func(t *testing.T) {
		input := ""
		maxSize := int64(10)

		// Simulate stdin input using a pipe and write test data to it.
		createMockStdin()
		loadTestData(input)

		// Test the function
		data, err := ReadStdin(maxSize)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(data) != 0 {
			t.Errorf("Expected empty input, got %q", string(data))
		}
	})
}
