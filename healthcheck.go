package main

import (
	"net/http"

	fthealth "github.com/Financial-Times/go-fthealth/v1_1"
	"github.com/Financial-Times/message-queue-go-producer/producer"
	"github.com/Financial-Times/message-queue-gonsumer/consumer"
	"github.com/Financial-Times/service-status-go/gtg"
)

type HealthCheck struct {
	consumer consumer.MessageConsumer
	producer producer.MessageProducer
}

func NewHealthCheck(p producer.MessageProducer, c consumer.MessageConsumer) *HealthCheck {
	return &HealthCheck{
		consumer: c,
		producer: p,
	}
}

func (h *HealthCheck) Health() func(w http.ResponseWriter, r *http.Request) {
	checks := []fthealth.Check{h.readQueueCheck(), h.writeQueueCheck()}
	hc := fthealth.HealthCheck{
		SystemCode:  "v1-suggestor",
		Name:        "V1 Suggestor",
		Description: "Checks if all the dependent services are reachable and healthy.",
		Checks:      checks,
	}
	return fthealth.Handler(hc)
}

func (h *HealthCheck) readQueueCheck() fthealth.Check {
	return fthealth.Check{
		ID:               "read-message-queue-proxy-reachable",
		Name:             "Read Message Queue Proxy Reachable",
		Severity:         1,
		BusinessImpact:   "Content V1 Metadata can't be read from queue. This will negatively impact V1 metadata availability.",
		TechnicalSummary: "Read message queue proxy is not reachable/healthy",
		PanicGuide:       "https://dewey.ft.com/",
		Checker:          h.consumer.ConnectivityCheck,
	}
}

func (h *HealthCheck) writeQueueCheck() fthealth.Check {
	return fthealth.Check{
		ID:               "write-message-queue-proxy-reachable",
		Name:             "Write Message Queue Proxy Reachable",
		Severity:         1,
		BusinessImpact:   "Content V1 Metadata can't be propagated through the stack. V1 metadata won't be available for new content.",
		TechnicalSummary: "Write message queue proxy is not reachable/healthy",
		PanicGuide:       "https://dewey.ft.com/",
		Checker:          h.producer.ConnectivityCheck,
	}
}

func (h *HealthCheck) GTG() gtg.Status {
	consumerCheck := func() gtg.Status {
		return gtgCheck(h.consumer.ConnectivityCheck)
	}
	producerCheck := func() gtg.Status {
		return gtgCheck(h.producer.ConnectivityCheck)
	}

	return gtg.FailFastParallelCheck([]gtg.StatusChecker{
		consumerCheck,
		producerCheck,
	})()
}

func gtgCheck(handler func() (string, error)) gtg.Status {
	if _, err := handler(); err != nil {
		return gtg.Status{GoodToGo: false, Message: err.Error()}
	}
	return gtg.Status{GoodToGo: true}
}
