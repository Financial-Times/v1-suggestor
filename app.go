package main

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Financial-Times/message-queue-go-producer/producer"
	"github.com/Financial-Times/message-queue-gonsumer/consumer"
	"github.com/gorilla/mux"
	"github.com/kr/pretty"
)

var appConfig AppConfig
var messageProducer producer.MessageProducer

func main() {
	initLogs(os.Stdout, os.Stdout, os.Stderr)
	appConfig = buildConfig()
	infoLogger.Printf("Using configuration: %# v", pretty.Formatter(appConfig))

	go enableHealthchecks()
	initializeProducer()
	readMessages()
}

func enableHealthchecks() {
	healthcheck := &Healthcheck{http.Client{}, appConfig}
	router := mux.NewRouter()
	router.HandleFunc("/__health", healthcheck.checkHealth())
	router.HandleFunc("/__gtg", healthcheck.gtg)
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		errorLogger.Panicf("Couldn't set up HTTP listener: %v\n", err)
	}
}

func initializeProducer() {
	messageProducer = producer.NewMessageProducer(appConfig.DestinationQueueConfig)
	infoLogger.Printf("Producer: %# v", pretty.Formatter(messageProducer))
}

func readMessages() {
	messageConsumer := consumer.NewConsumer(appConfig.SourceQueueConfig, handleMessage, http.Client{})
	infoLogger.Printf("Consumer: %# v", pretty.Formatter(messageConsumer))

	var consumerWaitGroup sync.WaitGroup
	consumerWaitGroup.Add(1)

	go func() {
		messageConsumer.Start()
		consumerWaitGroup.Done()
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	messageConsumer.Stop()
	consumerWaitGroup.Wait()
}

func handleMessage(msg consumer.Message) {
	//TODO
}
