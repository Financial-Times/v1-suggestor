package main

import (
	"strings"
)

// TaxonomyService defines the operations used to process taxonomies
type TaxonomyService interface {
	buildSuggestions(ContentRef) []suggestion
}

const conceptMentions = "mentions"
const conceptMajorMentions = "majorMentions"
const classification = "isClassifiedBy"
const primaryClassification = "isPrimarilyClassifiedBy"
const about = "about"
const hasAuthor = "hasAuthor"

const relevanceURI = "http://api.ft.com/scoringsystem/FT-RELEVANCE-SYSTEM"
const confidenceURI = "http://api.ft.com/scoringsystem/FT-CONFIDENCE-SYSTEM"

func transformScore(score int) float32 {
	return float32(score) / float32(100.0)
}

func generateID(cmrTermID string) string {
	return "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(cmrTermID)).String()
}

func extractTags(wantedTagName string, contentRef ContentRef) []tag {
	var wantedTags []tag
	for _, tag := range contentRef.TagHolder.Tags {
		if strings.EqualFold(tag.Term.Taxonomy, wantedTagName) {
			wantedTags = append(wantedTags, tag)
		}
	}
	return wantedTags
}

func buildSuggestion(tag tag, thingType string, predicate string) suggestion {
	relevance := score{
		ScoringSystem: relevanceURI,
		Value:         transformScore(tag.TagScore.Relevance),
	}
	confidence := score{
		ScoringSystem: confidenceURI,
		Value:         transformScore(tag.TagScore.Confidence),
	}

	provenances := []provenance{
		provenance{
			Scores: []score{relevance, confidence},
		},
	}
	thing := thing{
		ID:        generateID(tag.Term.ID),
		PrefLabel: tag.Term.CanonicalName,
		Predicate: predicate,
		Types:     []string{thingType},
	}

	return suggestion{Thing: thing, Provenance: provenances}
}
