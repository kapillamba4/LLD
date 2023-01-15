# Low Level System Design - PubSub Messaging Queue

### Problem Statement
Design a messaging service based on publisher-subscriber model.
Important points to note:
- There should be a way to create a topic with N number of partitions.
- Number of consumers of a consumer group should not be allowed to be more than number of partitions defined for the topic.
- Each consumer of same consumer group (having same group ID) should read from one partition only.
- Subscribers/Consumers should be able to subscribe to a topic.
- All consumers should run in parallel.
- Gracefully close the subscribers when SIGINT or SIGTERM is received.