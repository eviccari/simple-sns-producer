package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

// GetAWSSession - Get AWS Credentials Session
func GetAWSSession() *session.Session {
	return sess
}

// init - Load AWS Credentials Session
func init() {
	ep := "http://localhost:4566"
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("test", "test", "TOKEN"),
		Endpoint:    aws.String(ep),
	})

	if err != nil {
		panic(err)
	}

	setSession(s)
}

func setSession(s *session.Session) {
	sess = s
}
