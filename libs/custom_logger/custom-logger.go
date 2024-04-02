package CustomLogger

import (
	"fmt"
	"log"
	"time"
)

// CustomLogger defines the interface for our custom logger
type CustomLogger interface {
	Info(msg string)
	Error(err error)
	Message(msg string)
}

type MyLogger struct {
	prefix string
}

func NewMyLogger(prefix string) CustomLogger {
	return &MyLogger{prefix: prefix}
}

// Info prints the message in yellow to the console with a timestamp
func (l *MyLogger) Info(msg string) {
	log.Printf("[%s] [INFO]  %s\n", l.prefix, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\033[33m%s\033[0m\n", msg)
}

// Error prints the error message in red to the console with a timestamp
func (l *MyLogger) Error(err error) {
	log.Printf("[%s] [ERROR] %s\n", l.prefix, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\033[31mERROR: %s\033[0m\n", err.Error())
}

// Message prints the message in green to the console with a timestamp
func (l *MyLogger) Message(msg string) {
	log.Printf("[%s] [MESSAGE] %s\n", l.prefix, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\033[32m%s\033[0m\n", msg)
}
