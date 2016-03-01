package service

import (
	"strings"

	"github.com/Financial-Times/v1-suggestor/model"
)

// TaxonomyService defines the operations used to process taxonomies
type TaxonomyService interface {
	BuildSuggestions(model.ContentRef) []model.Suggestion
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

func extractTags(wantedTagName string, contentRef model.ContentRef) []model.Tag {
	var wantedTags []model.Tag
	for _, tag := range contentRef.TagHolder.Tags {
		if strings.EqualFold(tag.Term.Taxonomy, wantedTagName) {
			wantedTags = append(wantedTags, tag)
		}
	}
	return wantedTags
}

func buildSuggestion(tag model.Tag, thingType string) model.Suggestion {
	relevance := model.Score{
		ScoringSystem: relevanceURI,
		Value:         transformScore(tag.TagScore.Relevance),
	}
	confidence := model.Score{
		ScoringSystem: confidenceURI,
		Value:         transformScore(tag.TagScore.Confidence),
	}

	provenances := []model.Provenance{
		model.Provenance{
			Scores: []model.Score{relevance, confidence},
		},
	}
	thing := model.Thing{
		ID:        generateID(tag.Term.ID),
		PrefLabel: tag.Term.CanonicalName,
		Predicate: predicate,
		Types:     []string{thingType},
	}

	return model.Suggestion{Thing: thing, Provenance: provenances}
}
