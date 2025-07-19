package repository

// MessageRepository stores all message data.
type MessageRepository interface {
	AddMessage(topic string, data string)
	GetMessages(topic string, offset int, count int) []string
}
