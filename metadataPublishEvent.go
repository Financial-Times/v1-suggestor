package main

// MetadataPublishEvent models the events we process from the queue
type MetadataPublishEvent struct {
	UUID  string `json:"uuid"`
	Value string `json:"value"`
}
