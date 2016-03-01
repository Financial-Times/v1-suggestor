package service

import (
	"fmt"
	"testing"

	"github.com/Financial-Times/v1-suggestor/model"
	"github.com/stretchr/testify/assert"
)

func TestSubjectServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := SubjectService{"subjects"}
	tests := []struct {
		name        string
		contentRef  model.ContentRef
		suggestions []model.Suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 subject tag",
			buildContentRefWithSubjects(1),
			buildConceptSuggestionsWithSubjects(1),
		},
		{"Build concept suggestion from a contentRef with no subject tags",
			buildContentRefWithSubjects(0),
			buildConceptSuggestionsWithSubjects(0),
		},
		{"Build concept suggestion from a contentRef with multiple subject tags",
			buildContentRefWithSubjects(2),
			buildConceptSuggestionsWithSubjects(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.BuildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func TestSectionServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := SectionService{"sections"}
	tests := []struct {
		name        string
		contentRef  model.ContentRef
		suggestions []model.Suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 section tag",
			buildContentRefWithSections(1),
			buildConceptSuggestionsWithSections(1),
		},
		{"Build concept suggestion from a contentRef with no section tags",
			buildContentRefWithSections(0),
			buildConceptSuggestionsWithSections(0),
		},
		{"Build concept suggestion from a contentRef with multiple section tags",
			buildContentRefWithSections(2),
			buildConceptSuggestionsWithSections(2),
		},
		{"Build concept suggestion from a contentRef with a primary section",
			buildContentRefWithPrimarySection(2),
			buildConceptSuggestionsWithPrimarySection(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.BuildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func TestTopicServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := TopicService{"topics"}
	tests := []struct {
		name        string
		contentRef  model.ContentRef
		suggestions []model.Suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 topic tag",
			buildContentRefWithTopics(1),
			buildConceptSuggestionsWithTopics(1),
		},
		{"Build concept suggestion from a contentRef with no topic tags",
			buildContentRefWithTopics(0),
			buildConceptSuggestionsWithTopics(0),
		},
		{"Build concept suggestion from a contentRef with multiple topic tags",
			buildContentRefWithTopics(2),
			buildConceptSuggestionsWithTopics(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.BuildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func buildContentRefWithTopics(topicCount int) model.ContentRef {
	return buildContentRef(0, 0, topicCount, false)
}

func buildContentRefWithSubjects(subjectCount int) model.ContentRef {
	return buildContentRef(subjectCount, 0, 0, false)
}

func buildContentRefWithSections(sectionCount int) model.ContentRef {
	return buildContentRef(0, sectionCount, 0, false)
}

func buildContentRefWithPrimarySection(sectionCount int) model.ContentRef {
	return buildContentRef(0, sectionCount, 0, true)
}

func buildContentRef(subjectCount int, sectionCount int, topicCount int, hasPrimarySection bool) model.ContentRef {
	tags := []model.Tag{}
	for i := 0; i < subjectCount; i++ {
		subjectTerm := model.Term{CanonicalName: subjectNames[i], Taxonomy: "Subjects", ID: subjectTMEIDs[i]}
		tags = append(tags, model.Tag{Term: subjectTerm, TagScore: score})
	}

	for i := 0; i < sectionCount; i++ {
		sectionsTerm := model.Term{CanonicalName: sectionNames[i], Taxonomy: "Sections", ID: sectionTMEIDs[i]}
		sectionsTag := model.Tag{Term: sectionsTerm, TagScore: score}
		tags = append(tags, sectionsTag)
	}

	for i := 0; i < topicCount; i++ {
		topicTerm := model.Term{CanonicalName: topicNames[i], Taxonomy: "Topics", ID: topicTMEIDs[i]}
		topicTag := model.Tag{Term: topicTerm, TagScore: score}
		tags = append(tags, topicTag)
	}

	tagHolder := model.Tags{Tags: tags}

	var primarySection model.Term
	if hasPrimarySection {
		primarySection = model.Term{CanonicalName: sectionNames[0], Taxonomy: "Sections", ID: sectionTMEIDs[0]}
	}
	return model.ContentRef{TagHolder: tagHolder, PrimarySection: primarySection}
}

func buildConceptSuggestionsWithTopics(topicCount int) []model.Suggestion {
	return buildConceptSuggestions(0, 0, topicCount, false)
}

func buildConceptSuggestionsWithSections(sectionCount int) []model.Suggestion {
	return buildConceptSuggestions(0, sectionCount, 0, false)
}

func buildConceptSuggestionsWithPrimarySection(sectionCount int) []model.Suggestion {
	return buildConceptSuggestions(0, sectionCount, 0, true)
}

func buildConceptSuggestionsWithSubjects(subjectCount int) []model.Suggestion {
	return buildConceptSuggestions(subjectCount, 0, 0, false)
}

func buildConceptSuggestions(subjectCount int, sectionCount int, topicCount int, hasPrimarySection bool) []model.Suggestion {
	suggestions := []model.Suggestion{}

	relevance := model.Score{ScoringSystem: relevanceURI, Value: 0.65}
	confidence := model.Score{ScoringSystem: confidenceURI, Value: 0.93}
	provenance := model.Provenance{Scores: []model.Score{relevance, confidence}}

	for i := 0; i < subjectCount; i++ {
		thing := model.Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(subjectTMEIDs[i])).String(),
			PrefLabel: subjectNames[i],
			Predicate: predicate,
			Types:     []string{subjectURI},
		}
		subjectSuggestion := model.Suggestion{Thing: thing, Provenance: []model.Provenance{provenance}}
		suggestions = append(suggestions, subjectSuggestion)
	}

	for i := 0; i < sectionCount; i++ {
		thing := model.Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[i])).String(),
			PrefLabel: sectionNames[i],
			Predicate: predicate,
			Types:     []string{sectionURI},
		}
		sectionSuggestion := model.Suggestion{Thing: thing, Provenance: []model.Provenance{provenance}}
		suggestions = append(suggestions, sectionSuggestion)
	}
	if sectionCount > 0 && hasPrimarySection {
		thing := model.Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[0])).String(),
			PrefLabel: sectionNames[0],
			Predicate: primaryPredicate,
			Types:     []string{sectionURI},
		}
		sectionSuggestion := model.Suggestion{Thing: thing}
		suggestions = append(suggestions, sectionSuggestion)
	}

	for i := 0; i < topicCount; i++ {
		thing := model.Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(topicTMEIDs[i])).String(),
			PrefLabel: topicNames[i],
			Predicate: mentionsPredicate,
			Types:     []string{topicURI},
		}
		sectionSuggestion := model.Suggestion{Thing: thing, Provenance: []model.Provenance{provenance}}

		suggestions = append(suggestions, sectionSuggestion)

		thing = model.Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(topicTMEIDs[i])).String(),
			PrefLabel: topicNames[i],
			Predicate: aboutPredicate,
			Types:     []string{topicURI},
		}
		sectionSuggestion = model.Suggestion{Thing: thing, Provenance: []model.Provenance{provenance}}

		suggestions = append(suggestions, sectionSuggestion)
	}

	return suggestions
}

var score = model.TagScore{Confidence: 93, Relevance: 65}
var subjectNames = [...]string{"Mining Industry", "Oil Extraction Subsidies"}
var subjectTMEIDs = [...]string{"Mjk=-U2VjdGlvbnM=", "Nw==-R2VucmVz"}
var sectionNames = [...]string{"Companies", "Emerging Markets"}
var sectionTMEIDs = [...]string{"Nw==-R2Bucm3z", "Nw==-U2VjdGlvbnM="}
var topicNames = [...]string{"Big Data", "BP trial"}
var topicTMEIDs = [...]string{"M2YyN2I0NGEtZGZjMi00MDVjLTlkNjAtNGRlNTNhM2EwYjlm-VG9waWNz", "ZWE3YzNhNmQtNGU4MS00MzE0LWIxZWMtYWQxY2M4Y2ZjZDFk-VG9waWNz"}
