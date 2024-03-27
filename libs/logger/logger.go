package logger

import (
	"fmt"
	"log"
	"time"
)

// Logger defines the interface for our custom logger
type Logger interface {
	Info(msg string)
	Error(err error)
	Message(msg string)
}

// customLogger is a concrete implementation of the Logger interface
type customLogger struct {
	prefix string
}

// NewLogger creates a new custom logger with the provided prefix
func NewLogger(prefix string) Logger {
	return &customLogger{prefix: prefix}
}

// Info prints the message in yellow to the console with a timestamp
func (l *customLogger) Info(msg string) {
	log.Printf("[%s] [INFO]  %s\n", l.prefix, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\033[33m%s\033[0m\n", msg)
}

// Error prints the error message in red to the console with a timestamp
func (l *customLogger) Error(err error) {
	log.Printf("[%s] [ERROR] %s\n", l.prefix, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\033[31mERROR: %s\033[0m\n", err.Error())
}

// Message prints the message in green to the console with a timestamp
func (l *customLogger) Message(msg string) {
	log.Printf("[%s] [MESSAGE] %s\n", l.prefix, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\033[32m%s\033[0m\n", msg)
}
