package logging

import (
	"log"
	"os"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
}

type StdLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewStdLogger() *StdLogger {
	return &StdLogger{
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *StdLogger) Debug(msg string, a ...interface{}) {
	l.debugLogger.Printf(msg, a...)
}

func (l *StdLogger) Info(msg string) {
	l.infoLogger.Println(msg)
}

func (l *StdLogger) Error(msg string, a ...interface{}) {
    l.errorLogger.Printf(msg, a...)
}