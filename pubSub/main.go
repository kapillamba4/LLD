package main

import (
	"fmt"
	"log"
	"sync"

	"pubSub/app/adapters"
	"pubSub/app/usecase"
)

func main() {
	topicStore := adapters.NewTopicStore()
	msgStore := adapters.NewMessageStore()
	topicName := "example_topic"
	consumerGroupName1 := "example_cg_1"
	consumerGroupName2 := "example_cg_2"

	// Create a topic with N partitions.
	topicStore.CreateTopic(topicName, 5)

	// Create a message stream processor that reads from a topic & is also a part of some consumer group.
	streamProcessor := usecase.NewStreamProcessor(topicName, topicStore, msgStore)

	wt := sync.WaitGroup{}
	wt.Add(1)
	msgProcessingStopped := false
	messageCounter := 0

	go func() {
		for !msgProcessingStopped {
			msg := fmt.Sprintf("data_%d", messageCounter)
			msgStore.AddMessage("example_topic", msg)
			topicStore.IncrementMessageCount("example_topic", 1)
			messageCounter += 1
		}
		wt.Done()
	}()

	// Spawn 2 consumers for given stream processor having same consumer group 1.
	streamProcessor.AddConsumerGroup(consumerGroupName1)
	streamProcessor.AddConsumerGroup(consumerGroupName1)

	// Spawn 1 more consumer for given stream processor having consumer group 2.
	streamProcessor.AddConsumerGroup(consumerGroupName2)

	// Process method is Blocking.
	log.Output(1, "Start processing")
	streamProcessor.Process()

	// End processing, close all routines.
	msgProcessingStopped = true
	wt.Wait()

	log.Output(1, "Closing program - END")
}
