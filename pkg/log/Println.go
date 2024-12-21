package log

import logger "log"

// Println - Print log message with new line/carriage return at the end
func Println(msg ...interface{}) {
	logger.Println(msg...)
}
