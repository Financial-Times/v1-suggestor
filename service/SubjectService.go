package service

import (
	"strings"

	"github.com/Financial-Times/v1-suggestor/model"
)

// SubjectService extracts and transforms the subject taxonomy into a suggestion
type SubjectService struct {
	HandledTaxonomy string
}

const subjectURI = "http://www.ft.com/ontology/thing/Subject"

// BuildSuggestions builds a list of subject suggestions from a ContentRef.
// Returns an empty array in case no subject annotations are found
func (subjectService SubjectService) BuildSuggestions(contentRef model.ContentRef) []model.Suggestion {
	subjects := extractTags(subjectService.HandledTaxonomy, contentRef)

	suggestions := []model.Suggestion{}
	for _, value := range subjects {
		relevance := model.Score{
			ScoringSystem: RelevanceURI,
			Value:         transformScore(value.TagScore.Relevance),
		}
		confidence := model.Score{
			ScoringSystem: ConfidenceURI,
			Value:         transformScore(value.TagScore.Confidence),
		}

		provenances := []model.Provenance{
			model.Provenance{
				Scores: []model.Score{relevance, confidence},
			},
		}
		thing := model.Thing{
			ID:        generateID(value.Term.ID),
			PrefLabel: value.Term.CanonicalName,
			Types:     []string{subjectURI},
		}

		suggestions = append(suggestions, model.Suggestion{Thing: thing, Provenance: provenances})
	}
	return suggestions
}

func transformScore(score int) float32 {
	return float32(score) / float32(100.0)
}

func generateID(cmrTermID string) string {
	return NewNameUUIDFromBytes([]byte(cmrTermID)).String()
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
