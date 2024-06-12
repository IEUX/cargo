package internal

import (
	"fmt"
	"log"
	"os"
)

var Info, Error, Debug *log.Logger

func init() {
	lf, err := os.OpenFile(C.Logger.Dir+C.Logger.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf(ErrorColor, fmt.Sprintf("[LOGGER INIT FAILED]: %s", err))
	}
	Info = log.New(lf, "[INFO] - ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lf, "[ERROR] - ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(lf, "[DEBUG] - ", log.Ldate|log.Ltime|log.Lshortfile)
}
