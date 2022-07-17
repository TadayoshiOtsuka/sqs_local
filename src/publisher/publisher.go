package publisher

import (
	"context"
	"log"

	"github.com/TadayoshiOtsuka/sqs_local/src/services"
)

type Publisher struct {
	queueService services.QueueService
}

func NewPublisher(queue services.QueueService) *Publisher {
	return &Publisher{queueService: queue}
}

func (s *Publisher) SendMessages(ctx context.Context, messages []string) {
	for _, v := range messages {
		_, err := s.queueService.Send(ctx, v)
		if err != nil {
			log.Panicln("Send Message Error: ", err)
			break
		}
	}
}
