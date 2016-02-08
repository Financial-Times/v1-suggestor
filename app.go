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
	"fmt"
	"strings"

	"github.com/Financial-Times/message-queue-go-producer/producer"
	"github.com/Financial-Times/message-queue-gonsumer/consumer"
	"github.com/Financial-Times/v1-suggestor/model"
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
	metadataXML, err := base64.StdEncoding.DecodeString("base64bla")
	if err != nil {
		fmt.Printf("Error decoding body for uuid  %s", err.Error())
	}

	metadata := model.ContentRef{}
	err = xml.Unmarshal([]byte(metadataXML), &metadata)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Tags: %# v", pretty.Formatter(metadata))

	var subjects []model.Tag
	for _, value := range metadata.TagHolder.Tags {
		if strings.EqualFold(value.Term.Taxonomy, "subjects") {
			subjects = append(subjects, value)
		}
	}

	var conceptSuggestions []model.ConceptSuggestion
	for _, value := range subjects {
		var scores []model.Score
		relevance := model.Score{
			ScoringSystem: "http://api.ft.com/scoringsystem/FT-RELEVANCE-SYSTEM",
			Value:         float32(value.TagScore.Relevance) / float32(100.0),
		}
		confidence := model.Score{
			ScoringSystem: "http://api.ft.com/scoringsystem/FT-CONFIDENCE-SYSTEM",
			Value:         float32(value.TagScore.Confidence) / float32(100.0),
		}
		scores = append(scores, relevance)
		scores = append(scores, confidence)

		provenance := model.Provenance{Scores: scores}

		var provenances []model.Provenance
		provenances = append(provenances, provenance)

		var types []string
		types = append(types, "http://www.ft.com/ontology/thing/Subject")

		thing := model.Thing{
			Id:        "generate from CMR term ID",
			PrefLabel: value.Term.CanonicalName,
			Types:     types,
		}

		suggestion := model.Suggestion{
			Thing:      thing,
			Provenance: provenances,
		}

		var suggestions []model.Suggestion
		suggestions = append(suggestions, suggestion)

		conceptsuggestion := model.ConceptSuggestion{
			UUID:        "uuid",
			Suggestions: suggestions,
		}

		conceptSuggestions = append(conceptSuggestions, conceptsuggestion)
	}

	fmt.Printf("ConceptSuggestions: %# v", pretty.Formatter(conceptSuggestions))

	marshalledSuggestions, err := json.Marshal(conceptSuggestions)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Suggestions: \n %s", string(marshalledSuggestions))
}
