package main

import (
	"context"
	"log"
	"os"

	"github.com/TadayoshiOtsuka/sqs_local/src/publisher"
	"github.com/TadayoshiOtsuka/sqs_local/src/services"
	"github.com/TadayoshiOtsuka/sqs_local/src/subscriber"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	ctx := context.Background()
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
	cfg, err := config.LoadDefaultConfig(ctx, config.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		log.Fatalf("Failed To Load Configuration %v", err)
	}

	subscriber := subscriber.NewSubscriber(*services.NewQueueService(cfg))
	publisher := publisher.NewPublisher(*services.NewQueueService(cfg))
	publisher.SendMessages(ctx, []string{"hello", "world"})
	subscriber.Start(ctx)
}
