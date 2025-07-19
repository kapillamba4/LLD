package adapters

import (
	"sync"

	"pubSub/app/repository"
)

var _ repository.MessageRepository = (*MessageStore)(nil)

// MessageStore is in-memory implementation of MessageRepository
type MessageStore struct {
	topicToMessageStream map[string][]string
	sync.RWMutex
}

func NewMessageStore() *MessageStore {
	return &MessageStore{topicToMessageStream: map[string][]string{}}
}

func (m *MessageStore) AddMessage(topic string, data string) {
	m.Lock()
	defer m.Unlock()
	m.topicToMessageStream[topic] = append(m.topicToMessageStream[topic], data)
}

func (m *MessageStore) GetMessages(topic string, offset int, count int) []string {
	m.RLock()
	defer m.RUnlock()
	return m.topicToMessageStream[topic][offset : offset+count]
}
