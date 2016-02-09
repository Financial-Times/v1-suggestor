package main

import (
	"os"

	"errors"
	"fmt"

	"github.com/Financial-Times/message-queue-go-producer/producer"
	"github.com/Financial-Times/message-queue-gonsumer/consumer"
	"github.com/golang/go/src/pkg/strconv"
	"github.com/golang/go/src/pkg/strings"
)

// buildConfig reads the configuration from the system environment and build the app config
func buildConfig() AppConfig {
	concurrentProcessingString := os.Getenv("SRC_CONCURRENT_PROCESSING")
	srcConcurrentProcessing, err := strconv.ParseBool(concurrentProcessingString)
	if err != nil {
		errorMessage := fmt.Sprintf("Cannot parse value %v as boolean", concurrentProcessingString)
		panic(errors.New(errorMessage))
	}

	srcConf := consumer.QueueConfig{
		Addrs:                strings.Split(os.Getenv("SRC_ADDR"), ","),
		Group:                os.Getenv("SRC_GROUP"),
		Topic:                os.Getenv("SRC_TOPIC"),
		Queue:                os.Getenv("SRC_QUEUE"),
		ConcurrentProcessing: srcConcurrentProcessing,
	}

	destConf := producer.MessageProducerConfig{
		Addr:  os.Getenv("DEST_ADDRESS"),
		Topic: os.Getenv("DEST_TOPIC"),
		Queue: os.Getenv("DEST_QUEUE"),
	}

	return AppConfig{
		srcConf, destConf,
	}
}
