package elastic

import (
	"io"
	"testing"
)

// TestCloseReader tests the closeReader function with subtests.
func TestCloseReader(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		reader := io.NopCloser(nil)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Unexpected panic: %v", r)
			}
		}()
		closeReader(reader)
	})

	t.Run("ErrorPath", func(t *testing.T) {
		reader := errorCloser{}
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Unexpected panic: %v", r)
			}
		}()
		closeReader(reader)
	})
}

// errorCloser is a mock that always returns an error on Close.
type errorCloser struct{}

func (errorCloser) Read(p []byte) (n int, err error) { return 0, nil }
func (errorCloser) Close() error                     { return io.ErrUnexpectedEOF }
