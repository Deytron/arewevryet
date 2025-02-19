package utils

import (
	"log"
	"runtime"
)

// Colors for logging
var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"

func NonFatal(err interface{}, m string) {
	// Get file name for Debug
	_, file, line, _ := runtime.Caller(1)
	file = file[5:]

	if err != nil {
		log.Printf(red+"[ERROR] - %s - Trace: %v, - %s, line %d"+reset, m, err, file, line)
	} else {
		log.Printf(green+"[OK] - %s - %s, line %d"+reset, m, file, line)
	}
}

func Fatal(err interface{}, m string) {
	// Get file name for Debug
	_, file, line, _ := runtime.Caller(1)
	file = file[5:]

	if err != nil {
		log.Fatalf(red+"[ERROR] - %s - Trace: %v, - %s, line %d"+reset, m, err, file, line)
	} else {
		log.Printf(green+"[OK] - %s - %s, line %d"+reset, m, file, line)
	}
}
