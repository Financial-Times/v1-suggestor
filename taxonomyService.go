package main

import (
	"strings"

)

// TaxonomyService defines the operations used to process taxonomies
type TaxonomyService interface {
	BuildSuggestions(ContentRef) []Suggestion
}

const predicate = "isClassifiedBy"
const relevanceURI = "http://api.ft.com/scoringsystem/FT-RELEVANCE-SYSTEM"
const confidenceURI = "http://api.ft.com/scoringsystem/FT-CONFIDENCE-SYSTEM"
const primaryPredicate = "isPrimarilyClassifiedBy"

func transformScore(score int) float32 {
	return float32(score) / float32(100.0)
}

func generateID(cmrTermID string) string {
	return "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(cmrTermID)).String()
}

func extractTags(wantedTagName string, contentRef ContentRef) []Tag {
	var wantedTags []Tag
	for _, tag := range contentRef.TagHolder.Tags {
		if strings.EqualFold(tag.Term.Taxonomy, wantedTagName) {
			wantedTags = append(wantedTags, tag)
		}
	}
	return wantedTags
}

func buildSuggestion(tag Tag, thingType string, predicate string) Suggestion {
	relevance := Score{
		ScoringSystem: relevanceURI,
		Value:         transformScore(tag.TagScore.Relevance),
	}
	confidence := Score{
		ScoringSystem: confidenceURI,
		Value:         transformScore(tag.TagScore.Confidence),
	}

	provenances := []Provenance{
		Provenance{
			Scores: []Score{relevance, confidence},
		},
	}
	thing := Thing{
		ID:        generateID(tag.Term.ID),
		PrefLabel: tag.Term.CanonicalName,
		Predicate: predicate,
		Types:     []string{thingType},
	}

	return Suggestion{Thing: thing, Provenance: provenances}
}
