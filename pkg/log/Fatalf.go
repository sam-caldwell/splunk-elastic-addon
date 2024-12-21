package log

import logger "log"

// Fatalf - Format log printing
func Fatalf(fmt string, msg ...interface{}) {
	logger.Fatalf(fmt, msg...)
}
