package logging

import (
	"log"
	"os"
)

type Logger struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		InfoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
