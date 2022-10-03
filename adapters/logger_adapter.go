package adapters

// LoggerAdapter - Describe basic interface to be implemented by service loggers
type LoggerAdapter interface {
	Info(interface{})
	Debug(interface{})
	Error(interface{})
	Warn(interface{})
}
