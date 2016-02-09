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

func (subjectService SubjectService) BuildSuggestions(contentRef model.ContentRef) []model.Suggestion {
	subjects := extractTags(subjectService.HandledTaxonomy, contentRef)

	var suggestions []model.Suggestion
	for _, value := range subjects {
		relevance := model.Score{RelevanceURI, transformScore(value.TagScore.Relevance)}
		confidence := model.Score{ConfidenceURI, transformScore(value.TagScore.Confidence)}

		provenances := []model.Provenance{model.Provenance{[]model.Score{relevance, confidence}}}
		thing := model.Thing{generateID(value.Term.ID), value.Term.CanonicalName, []string{subjectURI}}

		suggestions = append(suggestions, model.Suggestion{thing, provenances})
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
