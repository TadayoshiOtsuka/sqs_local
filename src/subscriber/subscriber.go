package subscriber

import (
	"context"
	"log"

	"github.com/TadayoshiOtsuka/sqs_local/src/services"
)

type Subscriber struct {
	queueService services.QueueService
}

func NewSubscriber(queue services.QueueService) *Subscriber {
	return &Subscriber{queueService: queue}
}

func (s *Subscriber) SendMessages(ctx context.Context, messages []string) {
	for _, v := range messages {
		_, err := s.queueService.Send(ctx, v)
		if err != nil {
			log.Println("Send Message Error: ", err)
			break
		}
	}
}

func (s *Subscriber) Start(ctx context.Context) {
	for {
		res, err := s.queueService.Receive(ctx)
		if err != nil {
			log.Println("Receive Message Error: ", err)
			break
		}
		if len(res.Messages) > 0 {
			log.Println("Receive Message Body is:", *res.Messages[0].Body)
		}

	}
}
