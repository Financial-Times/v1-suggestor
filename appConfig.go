package main

import (
	"github.com/Financial-Times/message-queue-go-producer/producer"
	"github.com/Financial-Times/message-queue-gonsumer/consumer"
)

// AppConfig is the application configuration
type AppConfig struct {
	SourceQueueConfig      consumer.QueueConfig           `json:"sourceQueueConfig"`
	DestinationQueueConfig producer.MessageProducerConfig `json:"destinationQueueConfig"`
}
