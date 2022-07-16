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

func (s *Subscriber) Start(ctx context.Context) {
	for {
		res, err := s.queueService.Receive(ctx)
		if err != nil {
			log.Println("Receive Message Error: ", err)
			break
		}
		if len(res.Messages) <= 0 {
			log.Println("No Message Contains")
			continue
		}
		log.Println("Receive Message Body is:", *res.Messages[0].Body)
		s.queueService.Delete(ctx, res.Messages[0].ReceiptHandle)
	}
}
