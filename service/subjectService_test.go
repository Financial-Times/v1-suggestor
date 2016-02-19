package service

import (
	"fmt"
	"testing"

	"github.com/Financial-Times/v1-suggestor/model"
	"github.com/stretchr/testify/assert"
)

func TestBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := SubjectService{"subjects"}
	tests := []struct {
		name        string
		contentRef  model.ContentRef
		suggestions []model.Suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 subject tag",
			buildContentRef(1),
			buildConceptSuggestions(1),
		},
		{"Build concept suggestion from a contentRef with no subject tags",
			buildContentRef(0),
			buildConceptSuggestions(0),
		},
		{"Build concept suggestion from a contentRef with multiple subject tags",
			buildContentRef(2),
			buildConceptSuggestions(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.BuildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func buildContentRef(subjectCount int) model.ContentRef {
	tags := []model.Tag{}
	for i := 0; i < subjectCount; i++ {
		subjectTerm := model.Term{CanonicalName: subjectNames[i], Taxonomy: "Subjects", ID: subjectTMEIDS[i]}
		tags = append(tags, model.Tag{Term: subjectTerm, TagScore: score})
	}

	sectionsTerm := model.Term{CanonicalName: sectionName, Taxonomy: "Sections", ID: sectionTMEID}
	sectionsTag := model.Tag{Term: sectionsTerm, TagScore: score}
	tags = append(tags, sectionsTag)

	tagHolder := model.Tags{Tags: tags}
	return model.ContentRef{TagHolder: tagHolder}
}

func buildConceptSuggestions(subjectCount int) []model.Suggestion {
	suggestions := []model.Suggestion{}

	relevance := model.Score{ScoringSystem: RelevanceURI, Value: 0.65}
	confidence := model.Score{ScoringSystem: ConfidenceURI, Value: 0.93}
	provenance := model.Provenance{Scores: []model.Score{relevance, confidence}}

	for i := 0; i < subjectCount; i++ {
		thing := model.Thing{
			ID:        NewNameUUIDFromBytes([]byte(subjectTMEIDS[i])).String(),
			PrefLabel: subjectNames[i],
			Predicate: predicate,
			Types:     []string{subjectURI},
		}
		subjectSuggestion := model.Suggestion{Thing: thing, Provenance: []model.Provenance{provenance}}
		suggestions = append(suggestions, subjectSuggestion)
	}

	return suggestions
}

var score = model.TagScore{Confidence: 93, Relevance: 65}
var subjectNames = [...]string{"Mining Industry", "Oil Extraction Subsidies"}
var subjectTMEIDS = [...]string{"Mjk=-U2VjdGlvbnM=", "Nw==-R2VucmVz"}

const sectionName = "Companies"
const sectionTMEID = "Nw==-R2Bucm3z"
