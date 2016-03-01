package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubjectServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := SubjectService{"subjects"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []Suggestion
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
		contentRef  ContentRef
		suggestions []Suggestion
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
		contentRef  ContentRef
		suggestions []Suggestion
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

func buildContentRefWithTopics(topicCount int) ContentRef {
	return buildContentRef(0, 0, topicCount, false)
}

func buildContentRefWithSubjects(subjectCount int) ContentRef {
	return buildContentRef(subjectCount, 0, 0, false)
}

func buildContentRefWithSections(sectionCount int) ContentRef {
	return buildContentRef(0, sectionCount, 0, false)
}

func buildContentRefWithPrimarySection(sectionCount int) ContentRef {
	return buildContentRef(0, sectionCount, 0, true)
}

func buildContentRef(subjectCount int, sectionCount int, topicCount int, hasPrimarySection bool) ContentRef {
	tags := []Tag{}
	for i := 0; i < subjectCount; i++ {
		subjectTerm := Term{CanonicalName: subjectNames[i], Taxonomy: "Subjects", ID: subjectTMEIDs[i]}
		tags = append(tags, Tag{Term: subjectTerm, TagScore: score})
	}

	for i := 0; i < sectionCount; i++ {
		sectionsTerm := Term{CanonicalName: sectionNames[i], Taxonomy: "Sections", ID: sectionTMEIDs[i]}
		sectionsTag := Tag{Term: sectionsTerm, TagScore: score}
		tags = append(tags, sectionsTag)
	}

	for i := 0; i < topicCount; i++ {
		topicTerm := Term{CanonicalName: topicNames[i], Taxonomy: "Topics", ID: topicTMEIDs[i]}
		topicTag := Tag{Term: topicTerm, TagScore: score}
		tags = append(tags, topicTag)
	}

	tagHolder := Tags{Tags: tags}

	var primarySection Term
	if hasPrimarySection {
		primarySection = Term{CanonicalName: sectionNames[0], Taxonomy: "Sections", ID: sectionTMEIDs[0]}
	}
	return ContentRef{TagHolder: tagHolder, PrimarySection: primarySection}
}

func buildConceptSuggestionsWithTopics(topicCount int) []Suggestion {
	return buildConceptSuggestions(0, 0, topicCount, false)
}

func buildConceptSuggestionsWithSections(sectionCount int) []Suggestion {
	return buildConceptSuggestions(0, sectionCount, 0, false)
}

func buildConceptSuggestionsWithPrimarySection(sectionCount int) []Suggestion {
	return buildConceptSuggestions(0, sectionCount, 0, true)
}

func buildConceptSuggestionsWithSubjects(subjectCount int) []Suggestion {
	return buildConceptSuggestions(subjectCount, 0, 0, false)
}

func buildConceptSuggestions(subjectCount int, sectionCount int, topicCount int, hasPrimarySection bool) []Suggestion {
	suggestions := []Suggestion{}

	relevance := Score{ScoringSystem: relevanceURI, Value: 0.65}
	confidence := Score{ScoringSystem: confidenceURI, Value: 0.93}
	provenance := Provenance{Scores: []Score{relevance, confidence}}

	for i := 0; i < subjectCount; i++ {
		thing := Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(subjectTMEIDs[i])).String(),
			PrefLabel: subjectNames[i],
			Predicate: predicate,
			Types:     []string{subjectURI},
		}
		subjectSuggestion := Suggestion{Thing: thing, Provenance: []Provenance{provenance}}
		suggestions = append(suggestions, subjectSuggestion)
	}

	for i := 0; i < sectionCount; i++ {
		thing := Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[i])).String(),
			PrefLabel: sectionNames[i],
			Predicate: predicate,
			Types:     []string{sectionURI},
		}
		sectionSuggestion := Suggestion{Thing: thing, Provenance: []Provenance{provenance}}
		suggestions = append(suggestions, sectionSuggestion)
	}
	if sectionCount > 0 && hasPrimarySection {
		thing := Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[0])).String(),
			PrefLabel: sectionNames[0],
			Predicate: primaryPredicate,
			Types:     []string{sectionURI},
		}
		sectionSuggestion := Suggestion{Thing: thing}
		suggestions = append(suggestions, sectionSuggestion)
	}

	for i := 0; i < topicCount; i++ {
		thing := Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(topicTMEIDs[i])).String(),
			PrefLabel: topicNames[i],
			Predicate: mentionsPredicate,
			Types:     []string{topicURI},
		}
		sectionSuggestion := Suggestion{Thing: thing, Provenance: []Provenance{provenance}}

		suggestions = append(suggestions, sectionSuggestion)

		thing = Thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(topicTMEIDs[i])).String(),
			PrefLabel: topicNames[i],
			Predicate: aboutPredicate,
			Types:     []string{topicURI},
		}
		sectionSuggestion = Suggestion{Thing: thing, Provenance: []Provenance{provenance}}

		suggestions = append(suggestions, sectionSuggestion)
	}

	return suggestions
}

var score = TagScore{Confidence: 93, Relevance: 65}
var subjectNames = [...]string{"Mining Industry", "Oil Extraction Subsidies"}
var subjectTMEIDs = [...]string{"Mjk=-U2VjdGlvbnM=", "Nw==-R2VucmVz"}
var sectionNames = [...]string{"Companies", "Emerging Markets"}
var sectionTMEIDs = [...]string{"Nw==-R2Bucm3z", "Nw==-U2VjdGlvbnM="}
var topicNames = [...]string{"Big Data", "BP trial"}
var topicTMEIDs = [...]string{"M2YyN2I0NGEtZGZjMi00MDVjLTlkNjAtNGRlNTNhM2EwYjlm-VG9waWNz", "ZWE3YzNhNmQtNGU4MS00MzE0LWIxZWMtYWQxY2M4Y2ZjZDFk-VG9waWNz"}
