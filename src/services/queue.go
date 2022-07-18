package services

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type QueueService struct {
	client   *sqs.Client
	queueUrl string
}

func NewQueueService() *QueueService {
	cfg, err := initAwsConfig()
	if err != nil {
		log.Panicln("Failed To Load Configuration ", err)
	}
	c := sqs.NewFromConfig(*cfg)
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

func initAwsConfig() (*aws.Config, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if os.Getenv("ENV") != "production" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           os.Getenv("QUEUE_URL"),
				SigningRegion: os.Getenv("AWS_REGION"),
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
