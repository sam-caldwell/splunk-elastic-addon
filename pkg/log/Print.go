package log

import logger "log"

// Print - print a message to logs
func Print(msg ...interface{}) {
	logger.Print(msg...)
}
