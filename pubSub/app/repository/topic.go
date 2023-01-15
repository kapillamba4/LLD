package repository

// TopicRepository stores all topics data.
type TopicRepository interface {
	CreateTopic(name string, partitionCount int)
	GetNumberOfPartitions(topic string) int
	IncrementConsumerGroupOffset(topicName string, consumerGroup string, offset int) int
	GetCGOffset(topicName string, consumerGroup string) int
	GetTotalMessageCount(topicName string) int
	IncrementMessageCount(topicName string, msgCount int)
}
