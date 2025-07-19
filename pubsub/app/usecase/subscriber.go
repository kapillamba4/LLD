package usecase

import (
	"fmt"
)

// Subscriber consumes from single topic and is a part of a single consumer group.
type Subscriber struct {
	topic string
}

func NewSubscriber(topic string) *Subscriber {
	return &Subscriber{topic: topic}
}

func (s *Subscriber) ConsumeMessage(msg string, consumerGroup string, consumerId int) {
	// All business logic should reside here.
	// time.Sleep(time.Duration(math.Pow(float64(consumerId), 2)) * time.Second)
	fmt.Printf("Message %s, consumerGroupId: %s, consumerId: %d\n", msg, consumerGroup, consumerId)
}

func (s *Subscriber) GetTopic() string {
	return s.topic
}
