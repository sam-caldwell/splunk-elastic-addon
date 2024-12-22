package input

import (
	"fmt"
	"io"
	"os"
)

// ReadStdin - Read bytes from stdin with a hard limit to avoid overload
func ReadStdin(maxInputSize int64) (data []byte, err error) {

	if data, err = io.ReadAll(io.LimitReader(os.Stdin, maxInputSize)); err == io.EOF {

		err = fmt.Errorf("input size exceeds allowed limit of %d bytes", maxInputSize)

	}

	return data, err

}
