package adapters

import (
	"sync"

	"pubSub/app/repository"
)

var _ repository.TopicRepository = (*TopicStore)(nil)

// TopicStore is in-memory implementation of TopicRepository
type TopicStore struct {
	topics []*Topic
}

type Topic struct {
	name                string
	numberOfPartitions  int
	consumerGroupOffset map[string]int
	totalMessagesCount  int
	sync.RWMutex
}

func NewTopicStore() *TopicStore {
	return &TopicStore{}
}

func NewTopic(name string, partitionCount int) *Topic {
	return &Topic{name: name, numberOfPartitions: partitionCount, consumerGroupOffset: map[string]int{}}
}

func (t *TopicStore) IncrementMessageCount(topicName string, msgCount int) {
	for _, topic := range t.topics {
		if topic.GetName() == topicName {
			topic.totalMessagesCount += msgCount
		}
	}
}

func (t *TopicStore) CreateTopic(name string, partitionCount int) {
	t.topics = append(t.topics, NewTopic(name, partitionCount))
}

func (t *TopicStore) GetCGOffset(topicName string, consumerGroup string) int {
	for _, topic := range t.topics {
		if topic.GetName() == topicName {
			topic.RLock()
			result := topic.consumerGroupOffset[consumerGroup]
			topic.RUnlock()
			return result
		}
	}
	return 0
}

func (t *TopicStore) GetTotalMessageCount(topicName string) int {
	for _, topic := range t.topics {
		if topic.GetName() == topicName {
			return topic.totalMessagesCount
		}
	}
	return 0
}

func (t *TopicStore) GetNumberOfPartitions(topicName string) int {
	for _, topic := range t.topics {
		if topic.GetName() == topicName {
			return topic.GetNumberOfPartitions()
		}
	}
	return 0
}

func (t *TopicStore) IncrementConsumerGroupOffset(topicName string, consumerGroup string, offset int) int {
	for _, topic := range t.topics {
		if topic.GetName() == topicName {
			return topic.IncrementConsumerGroupOffset(consumerGroup, offset)
		}
	}
	return -1
}

func (t *Topic) GetName() string {
	return t.name
}

func (t *Topic) GetNumberOfPartitions() int {
	return t.numberOfPartitions
}

func (t *Topic) IncrementConsumerGroupOffset(consumerGroup string, offset int) int {
	t.Lock()
	defer t.Unlock()
	if t.consumerGroupOffset[consumerGroup]+offset > t.totalMessagesCount {
		return -1
	}
	t.consumerGroupOffset[consumerGroup] += offset
	return t.consumerGroupOffset[consumerGroup] - offset
}
