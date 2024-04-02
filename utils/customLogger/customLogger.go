package customLogger

// CustomLogger defines the interface for our custom logger
type CustomLogger interface {
	Info(msg string)
	Error(err error)
	Message(msg string)
}
