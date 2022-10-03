package adapters

import "github.com/sirupsen/logrus"

// BasicLoggerAdapter - Describe BasicLoggerAdapter struct
type BasicLoggerAdapter struct {
}

// NewBasicLogger - Create new BasicLoggerAdapter instance
func NewBasicLogger() LoggerAdapter {
	return &BasicLoggerAdapter{}
}

func (b *BasicLoggerAdapter) Info(i interface{}) {
	logrus.Info(i)
}

func (b *BasicLoggerAdapter) Debug(i interface{}) {
	logrus.Debug(i)
}

func (b *BasicLoggerAdapter) Error(i interface{}) {
	logrus.Error(i)
}

func (b *BasicLoggerAdapter) Warn(i interface{}) {
	logrus.Warn(i)
}

// init - Load logger configuration
func init() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableColors: false})
}
