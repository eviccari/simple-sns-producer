package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// SNSMessageProducerAdapter - Describe SNSMessageProducerAdapter
type SNSMessageProducerAdapter struct {
	sess *session.Session
}

// NewSNSMessageProducerAdapter - Create new SNSMessageProducerAdapter instance
func NewSNSMessageProducerAdapter(sess *session.Session) *SNSMessageProducerAdapter {
	return &SNSMessageProducerAdapter{
		sess: sess,
	}
}

// Send - Send message to AWS SNS topic
func (mp *SNSMessageProducerAdapter) Send(ctx context.Context, message string, topicName string) (protocolID string, err error) {
	svc := sns.New(mp.sess)
	result, err := svc.PublishWithContext(ctx, &sns.PublishInput{Message: &message, TopicArn: &topicName})
	if err != nil {
		return "", err
	}

	return *result.MessageId, nil
}
