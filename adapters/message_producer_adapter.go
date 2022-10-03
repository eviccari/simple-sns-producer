package adapters

import "context"

// MessageProducerAdapter - Describe basic message producer interface
type MessageProducerAdapter interface {
	Send(ctx context.Context, message string, topicName string) (protocolID string, err error)
}
