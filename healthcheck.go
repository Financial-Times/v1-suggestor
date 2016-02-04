package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Financial-Times/go-fthealth"
)

// Healthcheck offers methods to measure application health.
type Healthcheck struct {
	client http.Client
	config AppConfig
}

func (h *Healthcheck) checkHealth() func(w http.ResponseWriter, r *http.Request) {
	return fthealth.HandlerParallel("Dependent services healthcheck", "Checks if all the dependent services are reachable and healthy.", h.messageQueueProxyReachable())
}

func (h *Healthcheck) gtg(writer http.ResponseWriter, req *http.Request) {
	healthChecks := []func() error{h.checkAggregateMessageQueueProxiesReachable}

	for _, hCheck := range healthChecks {
		if err := hCheck(); err != nil {
			writer.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}
}

func (h *Healthcheck) messageQueueProxyReachable() fthealth.Check {
	return fthealth.Check{
		BusinessImpact:   "Content V1 Metadata is not suggested. This will negatively impact V1 metadata availability.",
		Name:             "MessageQueueProxyReachable",
		PanicGuide:       "https://sites.google.com/a/ft.com/technology/systems/dynamic-semantic-publishing/extra-publishing/v1-suggestor-runbook",
		Severity:         1,
		TechnicalSummary: "Message queue proxy is not reachable/healthy",
		Checker:          h.checkAggregateMessageQueueProxiesReachable,
	}

}

func (h *Healthcheck) checkAggregateMessageQueueProxiesReachable() error {

	addresses := h.config.SourceQueueConfig.Addrs
	errMsg := ""
	for i := 0; i < len(addresses); i++ {
		error := h.checkMessageQueueProxyReachable(addresses[i])
		if error == nil {
			return nil
		}
		errMsg = errMsg + fmt.Sprintf("For %s there is an error %v \n", addresses[i], error.Error())
	}

	return errors.New(errMsg)

}

func (h *Healthcheck) checkMessageQueueProxyReachable(address string) error {
	req, err := http.NewRequest("GET", address + "/topics", nil)
	if err != nil {
		warnLogger.Printf("Could not connect to proxy: %v", err.Error())
		return err
	}

	if len(h.config.SourceQueueConfig.AuthorizationKey) > 0 {
		req.Header.Add("Authorization", h.config.SourceQueueConfig.AuthorizationKey)
	}

	if len(h.config.SourceQueueConfig.Queue) > 0 {
		req.Host = h.config.SourceQueueConfig.Queue
	}

	resp, err := h.client.Do(req)
	if err != nil {
		warnLogger.Printf("Could not connect to proxy: %v", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Proxy returned status: %d", resp.StatusCode)
		return errors.New(errMsg)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return checkIfTopicIsPresent(body, h.config.SourceQueueConfig.Topic)

}

func checkIfTopicIsPresent(body []byte, searchedTopic string) error {
	var topics []string

	err := json.Unmarshal(body, &topics)
	if err != nil {
		return fmt.Errorf("Error occured and topic could not be found. %v", err.Error())
	}

	for _, topic := range topics {
		if topic == searchedTopic {
			return nil
		}
	}

	return errors.New("Topic was not found")
}
