package main

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"encoding/base64"
	"encoding/json"
	"encoding/xml"

	"github.com/Financial-Times/message-queue-go-producer/producer"
	"github.com/Financial-Times/message-queue-gonsumer/consumer"
	"github.com/Financial-Times/v1-suggestor/model"
	"github.com/Financial-Times/v1-suggestor/service"
	"github.com/gorilla/mux"
	"github.com/kr/pretty"
)

var appConfig AppConfig
var messageProducer producer.MessageProducer
var taxonomyHandlers = make(map[string]service.TaxonomyService)

func main() {
	initLogs(os.Stdout, os.Stdout, os.Stderr)
	appConfig = buildConfig()
	infoLogger.Printf("Using configuration: %# v", pretty.Formatter(appConfig))

	setupTaxonomyHandlers()

	go enableHealthChecks()
	initializeProducer()
	readMessages()
}

func setupTaxonomyHandlers() {
	taxonomyHandlers["subjects"] = service.SubjectService{"subjects"}
}

func enableHealthChecks() {
	healthCheck := &Healthcheck{http.Client{}, appConfig}
	router := mux.NewRouter()
	router.HandleFunc("/__health", healthCheck.checkHealth())
	router.HandleFunc("/__gtg", healthCheck.gtg)
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
	var metadataPublishEvent model.MetadataPublishEvent
	err := json.Unmarshal([]byte(msg.Body), &metadataPublishEvent)
	if err != nil {
		errorLogger.Printf("Cannot unmarshal message body: %s", err.Error())
		return
	}

	metadataXML, err := base64.StdEncoding.DecodeString(metadataPublishEvent.Value)
	if err != nil {
		errorLogger.Printf("Error decoding body for uuid:  %s", err.Error())
		return
	}

	metadata := model.ContentRef{}
	err = xml.Unmarshal(metadataXML, &metadata)
	if err != nil {
		errorLogger.Printf("Error unmarshalling metadata XML: %s", err.Error())
	}

	var suggestions []model.Suggestion
	for key, value := range taxonomyHandlers {
		infoLogger.Printf("Processing taxonomy %s", key)
		suggestions = append(suggestions, value.BuildSuggestions(metadata)...)
	}

	conceptSuggestion := model.ConceptSuggestion{
		UUID:        metadataPublishEvent.UUID,
		Suggestions: suggestions,
	}

	marshalledSuggestions, err := json.Marshal(conceptSuggestion)
	if err != nil {
		panic(err)
	}

	infoLogger.Printf("Suggestions: \n %s", string(marshalledSuggestions))
}
