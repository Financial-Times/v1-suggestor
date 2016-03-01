package main

// ConceptSuggestion models the suggestion as it will be written on the queue
type ConceptSuggestion struct {
	UUID        string       `json:"uuid"`
	Suggestions []suggestion `json:"suggestions"`
}

type suggestion struct {
	Thing      thing        `json:"thing"`
	Provenance []provenance `json:"provenances,omitempty"`
}

type thing struct {
	ID        string   `json:"id"`
	PrefLabel string   `json:"prefLabel"`
	Predicate string   `json:"predicate"`
	Types     []string `json:"types"`
}

type provenance struct {
	Scores []score `json:"scores"`
}

type score struct {
	ScoringSystem string  `json:"scoringSystem"`
	Value         float32 `json:"value"`
}
