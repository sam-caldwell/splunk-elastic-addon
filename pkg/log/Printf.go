package log

import logger "log"

// Printf - Format log printing
func Printf(fmt string, msg ...interface{}) {
	logger.Printf(fmt, msg...)
}
