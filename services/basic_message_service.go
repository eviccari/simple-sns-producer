package services

import (
	"context"
	"errors"
	"github.com/eviccari/simple-sns-producer/adapters"
)

// BasicMessageService - Describe BasicMessageService type.
type BasicMessageService struct {
	mpa adapters.MessageProducerAdapter
}

// NewBasicMessageService - Create new BasicMessageService instance
func NewBasicMessageService(mpa adapters.MessageProducerAdapter) *BasicMessageService {
	return &BasicMessageService{
		mpa: mpa,
	}
}

// Send - Send message to topic
func (ms *BasicMessageService) Send(ctx context.Context, message string, topicName string) (protocolID string, err error) {
	if message == "" || topicName == "" {
		return "", errors.New("message and topic name are required")
	}

	protocolID, err = ms.mpa.Send(ctx, message, topicName)
	if err != nil {
		return "", err
	}
	return
}
