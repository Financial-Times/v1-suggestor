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
		suggestions []suggestion
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
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func TestSectionServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := SectionService{"sections"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []suggestion
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
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func TestTopicServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := TopicService{"topics"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []suggestion
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
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
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
	metadataTags := []tag{}
	for i := 0; i < subjectCount; i++ {
		subjectTerm := term{CanonicalName: subjectNames[i], Taxonomy: "Subjects", ID: subjectTMEIDs[i]}
		metadataTags = append(metadataTags, tag{Term: subjectTerm, TagScore: testScore})
	}

	for i := 0; i < sectionCount; i++ {
		sectionsTerm := term{CanonicalName: sectionNames[i], Taxonomy: "Sections", ID: sectionTMEIDs[i]}
		sectionsTag := tag{Term: sectionsTerm, TagScore: testScore}
		metadataTags = append(metadataTags, sectionsTag)
	}

	for i := 0; i < topicCount; i++ {
		topicTerm := term{CanonicalName: topicNames[i], Taxonomy: "Topics", ID: topicTMEIDs[i]}
		topicTag := tag{Term: topicTerm, TagScore: testScore}
		metadataTags = append(metadataTags, topicTag)
	}

	tagHolder := tags{Tags: metadataTags}

	var primarySection term
	if hasPrimarySection {
		primarySection = term{CanonicalName: sectionNames[0], Taxonomy: "Sections", ID: sectionTMEIDs[0]}
	}
	return ContentRef{TagHolder: tagHolder, PrimarySection: primarySection}
}

func buildConceptSuggestionsWithTopics(topicCount int) []suggestion {
	return buildConceptSuggestions(0, 0, topicCount, false)
}

func buildConceptSuggestionsWithSections(sectionCount int) []suggestion {
	return buildConceptSuggestions(0, sectionCount, 0, false)
}

func buildConceptSuggestionsWithPrimarySection(sectionCount int) []suggestion {
	return buildConceptSuggestions(0, sectionCount, 0, true)
}

func buildConceptSuggestionsWithSubjects(subjectCount int) []suggestion {
	return buildConceptSuggestions(subjectCount, 0, 0, false)
}

func buildConceptSuggestions(subjectCount int, sectionCount int, topicCount int, hasPrimarySection bool) []suggestion {
	suggestions := []suggestion{}

	relevance := score{ScoringSystem: relevanceURI, Value: 0.65}
	confidence := score{ScoringSystem: confidenceURI, Value: 0.93}
	metadataProvenance := provenance{Scores: []score{relevance, confidence}}

	for i := 0; i < subjectCount; i++ {
		thing := thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(subjectTMEIDs[i])).String(),
			PrefLabel: subjectNames[i],
			Predicate: classification,
			Types:     []string{subjectURI},
		}
		subjectSuggestion := suggestion{Thing: thing, Provenance: []provenance{metadataProvenance}}
		suggestions = append(suggestions, subjectSuggestion)
	}

	for i := 0; i < sectionCount; i++ {
		thing := thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[i])).String(),
			PrefLabel: sectionNames[i],
			Predicate: classification,
			Types:     []string{sectionURI},
		}
		sectionSuggestion := suggestion{Thing: thing, Provenance: []provenance{metadataProvenance}}
		suggestions = append(suggestions, sectionSuggestion)
	}
	if sectionCount > 0 && hasPrimarySection {
		thing := thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[0])).String(),
			PrefLabel: sectionNames[0],
			Predicate: primaryClassification,
			Types:     []string{sectionURI},
		}
		sectionSuggestion := suggestion{Thing: thing}
		suggestions = append(suggestions, sectionSuggestion)
	}

	for i := 0; i < topicCount; i++ {
		oneThing := thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(topicTMEIDs[i])).String(),
			PrefLabel: topicNames[i],
			Predicate: conceptMentions,
			Types:     []string{topicURI},
		}
		sectionSuggestion := suggestion{Thing: oneThing, Provenance: []provenance{metadataProvenance}}

		suggestions = append(suggestions, sectionSuggestion)

		oneThing = thing{
			ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(topicTMEIDs[i])).String(),
			PrefLabel: topicNames[i],
			Predicate: conceptAbout,
			Types:     []string{topicURI},
		}
		sectionSuggestion = suggestion{Thing: oneThing, Provenance: []provenance{metadataProvenance}}

		suggestions = append(suggestions, sectionSuggestion)
	}

	return suggestions
}

var testScore = tagScore{Confidence: 93, Relevance: 65}
var subjectNames = [...]string{"Mining Industry", "Oil Extraction Subsidies"}
var subjectTMEIDs = [...]string{"Mjk=-U2VjdGlvbnM=", "Nw==-R2VucmVz"}
var sectionNames = [...]string{"Companies", "Emerging Markets"}
var sectionTMEIDs = [...]string{"Nw==-R2Bucm3z", "Nw==-U2VjdGlvbnM="}
var topicNames = [...]string{"Big Data", "BP trial"}
var topicTMEIDs = [...]string{"M2YyN2I0NGEtZGZjMi00MDVjLTlkNjAtNGRlNTNhM2EwYjlm-VG9waWNz", "ZWE3YzNhNmQtNGU4MS00MzE0LWIxZWMtYWQxY2M4Y2ZjZDFk-VG9waWNz"}
