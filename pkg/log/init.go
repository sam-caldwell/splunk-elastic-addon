package log

import (
	logger "log"
	"os"
)

// init - initialize logging to stderr
func init() {
	logger.SetOutput(os.Stderr)
}
