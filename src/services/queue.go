package services

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type QueueService struct {
	client   *sqs.Client
	queueUrl string
}

func NewQueueService(cfg aws.Config) *QueueService {
	c := sqs.NewFromConfig(cfg)
	url := os.Getenv("QUEUE_URL")
	return &QueueService{client: c, queueUrl: url}
}

func (s *QueueService) Send(ctx context.Context, body string) (*string, error) {
	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(body),
		QueueUrl:     aws.String(s.queueUrl),
		DelaySeconds: 5,
	}
	res, err := s.client.SendMessage(ctx, params)
	if err != nil {
		return nil, err
	}

	return res.MessageId, nil
}

func (s *QueueService) Receive(ctx context.Context) (*sqs.ReceiveMessageOutput, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:        aws.String(s.queueUrl),
		WaitTimeSeconds: 20,
	}
	res, err := s.client.ReceiveMessage(ctx, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *QueueService) Delete(ctx context.Context, receiptHandle *string) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(s.queueUrl),
		ReceiptHandle: receiptHandle,
	}
	if _, err := s.client.DeleteMessage(ctx, params); err != nil {
		return err
	}

	return nil
}
