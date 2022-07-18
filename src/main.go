package main

import (
	"context"

	"github.com/TadayoshiOtsuka/sqs_local/src/publisher"
	"github.com/TadayoshiOtsuka/sqs_local/src/services"
	"github.com/TadayoshiOtsuka/sqs_local/src/subscriber"
)

func main() {
	ctx := context.Background()
	queueService := services.NewQueueService()
	subscriber := subscriber.NewSubscriber(*queueService)
	publisher := publisher.NewPublisher(*queueService)
	publisher.SendMessages(ctx, []string{"hello", "world"})
	subscriber.Start(ctx)
}
