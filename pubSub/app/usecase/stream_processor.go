package usecase

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"

	"pubSub/app/repository"
)

// StreamProcessor is a message mediator between topic & multiple consumers of a consumer group.
type StreamProcessor struct {
	subscriber        *Subscriber
	topicRepository   repository.TopicRepository
	messageRepository repository.MessageRepository
	freeSubscribers   map[string]int32
	totalSubscribers  map[string]int32
	sync.WaitGroup
	sync.RWMutex
}

func NewStreamProcessor(topicName string, topicRepository repository.TopicRepository, messageRepository repository.MessageRepository) *StreamProcessor {
	return &StreamProcessor{
		subscriber:        NewSubscriber(topicName),
		topicRepository:   topicRepository,
		messageRepository: messageRepository,
		freeSubscribers:   map[string]int32{},
		totalSubscribers:  map[string]int32{},
	}
}

func addDelta(i *int32, delta int32) int32 {
	return atomic.AddInt32(i, delta)
}

func (s *StreamProcessor) updateSubsCount(subsCount map[string]int32, cg string, delta int32) {
	s.Lock()
	subsCount[cg] = subsCount[cg] + delta
	s.Unlock()
}

func (s *StreamProcessor) AddConsumerGroup(consumerGroup string) {
	s.Lock()
	defer s.Unlock()
	numOfPartitions := s.topicRepository.GetNumberOfPartitions(s.subscriber.GetTopic())
	if numOfPartitions >= int(s.totalSubscribers[consumerGroup])+1 {
		s.freeSubscribers[consumerGroup] += 1
		s.totalSubscribers[consumerGroup] += 1
	}
}

func (s *StreamProcessor) Process() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	numberOfMessagesProcessed := int32(0)
	subscriberCounter := 0
	for quit := false; !quit; {

		for cg := range s.freeSubscribers {
			freeSubs := s.freeSubscribers[cg]
			if freeSubs > 0 {
				// register new subscriber
				for i := 0; i < int(freeSubs); i += 1 {
					s.Add(1)
					s.updateSubsCount(s.freeSubscribers, cg, -1)
					subscriberCounter += 1
					go func(subscriberCount int, consumerGroup string) {
						log.Output(1, fmt.Sprintf("Started consumer %d, Consumer Group %s \n", subscriberCount, consumerGroup))
						for !quit {
							offset := s.topicRepository.IncrementConsumerGroupOffset(s.subscriber.GetTopic(), consumerGroup, 1)
							if offset == -1 {
								continue
							}

							messages := s.messageRepository.GetMessages(s.subscriber.GetTopic(), offset, 1)
							if len(messages) > 0 {
								s.subscriber.ConsumeMessage(messages[0], consumerGroup, subscriberCount)
								addDelta(&numberOfMessagesProcessed, 1)
							}
						}
						s.Done()
					}(subscriberCounter, cg)
				}
			}
		}
		sig := <-done
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			log.Output(1, "Received a signal: "+sig.String())
			quit = true
		}
	}

	s.Wait()
	log.Output(1, fmt.Sprintf("Messages processed: %d\n", numberOfMessagesProcessed))
}
