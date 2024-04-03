package customLogger

import (
	"fmt"
	"log"
	"time"
)

type stackedCustomLogger struct {
	prefix string
}

func NewStackedCustomLogger(prefix string) CustomLogger {
	return &stackedCustomLogger{prefix: prefix}
}

// Info prints the message in yellow to the console with a timestamp
func (l *stackedCustomLogger) Info(msg string) {
	log.Printf("%s [%s] [INFO] \033[33m%s\033[0m\n", time.Now().Format("2006-01-02 15:04:05"), l.prefix, msg)
}

// Error prints the error message in red to the console with a timestamp
func (l *stackedCustomLogger) Error(err error) {
	panic(fmt.Sprintf("%s [%s] [ERROR] \033[31mERROR: %s\033[0m\n", time.Now().Format("2006-01-02 15:04:05"), l.prefix, err.Error()))
}

// Message prints the message in green to the console with a timestamp
func (l *stackedCustomLogger) Message(msg string) {
	log.Printf("%s [%s] [MESSAGE] \033[32m%s\033[0m\n", time.Now().Format("2006-01-02 15:04:05"), l.prefix, msg)
}
