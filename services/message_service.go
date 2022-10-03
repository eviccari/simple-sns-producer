package services

import "context"

// MessageService - Describe MessageService interface
type MessageService interface {
	Send(ctx context.Context, message string, topicName string) (protocolID string, err error)
}
