package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"strings"
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
			buildContentRefWithPrimarySection("sections", 2),
			buildConceptSuggestionsWithPrimarySection("sections", 2),
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

func TestLocationServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := LocationService{"gl"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 location tag",
			buildContentRefWithLocations(1),
			buildConceptSuggestionsWithLocations(1),
		},
		{"Build concept suggestion from a contentRef with no location tags",
			buildContentRefWithLocations(0),
			buildConceptSuggestionsWithLocations(0),
		},
		{"Build concept suggestion from a contentRef with multiple location tags",
			buildContentRefWithLocations(2),
			buildConceptSuggestionsWithLocations(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func TestGenreServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := GenreService{"genres"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 genre tag",
			buildContentRefWithGenres(1),
			buildConceptSuggestionsWithGenres(1),
		},
		{"Build concept suggestion from a contentRef with no genre tags",
			buildContentRefWithGenres(0),
			buildConceptSuggestionsWithGenres(0),
		},
		{"Build concept suggestion from a contentRef with multiple genre tags",
			buildContentRefWithGenres(2),
			buildConceptSuggestionsWithGenres(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func TestSpecialReportServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := SpecialReportService{"specialReports"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 specialReports tag",
			buildContentRefWithSpecialReports(1),
			buildConceptSuggestionsWithSpecialReports(1),
		},
		{"Build concept suggestion from a contentRef with no specialReports tags",
			buildContentRefWithSpecialReports(0),
			buildConceptSuggestionsWithSpecialReports(0),
		},
		{"Build concept suggestion from a contentRef with multiple specialReports tags",
			buildContentRefWithSpecialReports(2),
			buildConceptSuggestionsWithSpecialReports(2),
		},
		{"Build concept suggestion from a contentRef with specialReports as a primary section",
			buildContentRefWithPrimarySection("specialReports", 2),
			buildConceptSuggestionsWithPrimarySection("specialReports", 2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect", test.name))
	}
}

func TestAlphavilleSeriesServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := AlphavilleSeriesService{"alphavilleSeries"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 alphavilleSeries tag",
			buildContentRefWithAlphavilleSeries(1),
			buildConceptSuggestionsWithAlphavilleSeries(1),
		},
		{"Build concept suggestion from a contentRef with no alphavilleSeries tags",
			buildContentRefWithAlphavilleSeries(0),
			buildConceptSuggestionsWithAlphavilleSeries(0),
		},
		{"Build concept suggestion from a contentRef with multiple alphavilleSeries tags",
			buildContentRefWithAlphavilleSeries(2),
			buildConceptSuggestionsWithAlphavilleSeries(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect.", test.name))
	}
}

func TestOrganisationsServiceBuildSuggestions(t *testing.T) {
	assert := assert.New(t)
	service := OrganisationService{"on"}
	tests := []struct {
		name        string
		contentRef  ContentRef
		suggestions []suggestion
	}{
		{"Build concept suggestion from a contentRef with 1 organisation tag",
			buildContentRefWithOrganisations(1),
			buildConceptSuggestionsWithOrganisations(1),
		},
		{"Build concept suggestion from a contentRef with no organisation tags",
			buildContentRefWithOrganisations(0),
			buildConceptSuggestionsWithOrganisations(0),
		},
		{"Build concept suggestion from a contentRef with multiple organisation tags",
			buildContentRefWithOrganisations(2),
			buildConceptSuggestionsWithOrganisations(2),
		},
	}

	for _, test := range tests {
		actualConceptSuggestions := service.buildSuggestions(test.contentRef)
		assert.Equal(test.suggestions, actualConceptSuggestions, fmt.Sprintf("%s: Actual concept suggestions incorrect: ACTUAL: %s  TEST: %s ", test.name, actualConceptSuggestions, test.suggestions))
	}
}

func buildContentRefWithLocations(locationCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["locations"] = locationCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRefWithTopics(topicCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["topics"] = topicCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRefWithSubjects(subjectCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["subjects"] = subjectCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRefWithSections(sectionCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["sections"] = sectionCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRefWithPrimarySection(taxonomyName string, sectionCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount[taxonomyName] = sectionCount
	return buildContentRef(taxonomyAndCount, true)
}

func buildContentRefWithGenres(genreCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["genres"] = genreCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRefWithSpecialReports(specialReportCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["specialReports"] = specialReportCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRefWithAlphavilleSeries(seriesCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["alphavilleSeries"] = seriesCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRefWithOrganisations(organisationCount int) ContentRef {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["organisations"] = organisationCount
	return buildContentRef(taxonomyAndCount, false)
}

func buildContentRef(taxonomyAndCount map[string]int, hasPrimarySection bool) ContentRef {
	metadataTags := []tag{}
	var primarySection term
	for key, count := range taxonomyAndCount {
		if strings.EqualFold("subjects", key) {
			for i := 0; i < count; i++ {
				subjectTerm := term{CanonicalName: subjectNames[i], Taxonomy: "Subjects", ID: subjectTMEIDs[i]}
				metadataTags = append(metadataTags, tag{Term: subjectTerm, TagScore: testScore})
			}
		}
		if strings.EqualFold("sections", key) {
			for i := 0; i < count; i++ {
				sectionsTerm := term{CanonicalName: sectionNames[i], Taxonomy: "Sections", ID: sectionTMEIDs[i]}
				sectionsTag := tag{Term: sectionsTerm, TagScore: testScore}
				metadataTags = append(metadataTags, sectionsTag)
			}

			if hasPrimarySection {
				primarySection = term{CanonicalName: sectionNames[0], Taxonomy: "Sections", ID: sectionTMEIDs[0]}
			}
		}
		if strings.EqualFold("topics", key) {
			for i := 0; i < count; i++ {
				topicTerm := term{CanonicalName: topicNames[i], Taxonomy: "Topics", ID: topicTMEIDs[i]}
				topicTag := tag{Term: topicTerm, TagScore: testScore}
				metadataTags = append(metadataTags, topicTag)
			}
		}
		if strings.EqualFold("locations", key) {
			for i := 0; i < count; i++ {
				locationTerm := term{CanonicalName: locationNames[i], Taxonomy: "GL", ID: locationTMEIDs[i]}
				locationTag := tag{Term: locationTerm, TagScore: testScore}
				metadataTags = append(metadataTags, locationTag)
			}
		}
		if strings.EqualFold("genres", key) {
			for i := 0; i < count; i++ {
				genreTerm := term{CanonicalName: genreNames[i], Taxonomy: "Genres", ID: genreTMEIDs[i]}
				metadataTags = append(metadataTags, tag{Term: genreTerm, TagScore: testScore})
			}
		}
		if strings.EqualFold("specialReports", key) {
			for i := 0; i < count; i++ {
				specialReportsTerm := term{CanonicalName: specialReportNames[i], Taxonomy: "SpecialReports", ID: specialReportTMEIDs[i]}
				sectionsTag := tag{Term: specialReportsTerm, TagScore: testScore}
				metadataTags = append(metadataTags, sectionsTag)
			}

			if hasPrimarySection {
				primarySection = term{CanonicalName: specialReportNames[0], Taxonomy: "SpecialReports", ID: specialReportTMEIDs[0]}
			}
		}
		if strings.EqualFold("alphavilleSeries", key) {
			for i := 0; i < count; i++ {
				alphavilleSeriesTerm := term{CanonicalName: alphavilleSeriesNames[i], Taxonomy: "AlphavilleSeries", ID: alphavilleSeriesTMEIDs[i]}
				alphavilleSeriesTag := tag{Term: alphavilleSeriesTerm, TagScore: testScore}
				metadataTags = append(metadataTags, alphavilleSeriesTag)
			}
		}
		if strings.EqualFold("organisations", key) {
			for i := 0; i < count; i++ {
				organisationTerm := term{CanonicalName: organisationNames[i], Taxonomy: "ON", ID: organisationTMEIDs[i]}
				organisationTag := tag{Term: organisationTerm, TagScore: testScore}
				metadataTags = append(metadataTags, organisationTag)
			}
		}
	}
	tagHolder := tags{Tags: metadataTags}

	return ContentRef{TagHolder: tagHolder, PrimarySection: primarySection}
}

func buildConceptSuggestionsWithLocations(locationCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["locations"] = locationCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestionsWithTopics(topicCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["topics"] = topicCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestionsWithSections(sectionCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["sections"] = sectionCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestionsWithPrimarySection(taxonomyName string, sectionCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount[taxonomyName] = sectionCount
	return buildConceptSuggestions(taxonomyAndCount, true)
}

func buildConceptSuggestionsWithSubjects(subjectCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["subjects"] = subjectCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestionsWithGenres(genreCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["genres"] = genreCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestionsWithSpecialReports(reportCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["specialReports"] = reportCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestionsWithAlphavilleSeries(seriesCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["alphavilleSeries"] = seriesCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestionsWithOrganisations(seriesCount int) []suggestion {
	taxonomyAndCount := make(map[string]int)
	taxonomyAndCount["organisations"] = seriesCount
	return buildConceptSuggestions(taxonomyAndCount, false)
}

func buildConceptSuggestions(taxonomyAndCount map[string]int, hasPrimarySection bool) []suggestion {
	suggestions := []suggestion{}

	relevance := score{ScoringSystem: relevanceURI, Value: 0.65}
	confidence := score{ScoringSystem: confidenceURI, Value: 0.93}
	metadataProvenance := provenance{Scores: []score{relevance, confidence}}
	for key, count := range taxonomyAndCount {
		if strings.EqualFold("subjects", key) {
			for i := 0; i < count; i++ {
				thing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(subjectTMEIDs[i])).String(),
					PrefLabel: subjectNames[i],
					Predicate: classification,
					Types:     []string{subjectURI},
				}
				subjectSuggestion := suggestion{Thing: thing, Provenance: []provenance{metadataProvenance}}
				suggestions = append(suggestions, subjectSuggestion)
			}
		}
		if strings.EqualFold("sections", key) {
			for i := 0; i < count; i++ {
				thing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[i])).String(),
					PrefLabel: sectionNames[i],
					Predicate: classification,
					Types:     []string{sectionURI},
				}
				sectionSuggestion := suggestion{Thing: thing, Provenance: []provenance{metadataProvenance}}
				suggestions = append(suggestions, sectionSuggestion)
			}

			if count > 0 && hasPrimarySection {
				thing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(sectionTMEIDs[0])).String(),
					PrefLabel: sectionNames[0],
					Predicate: primaryClassification,
					Types:     []string{sectionURI},
				}
				sectionSuggestion := suggestion{Thing: thing}
				suggestions = append(suggestions, sectionSuggestion)
			}
		}
		if strings.EqualFold("topics", key) {
			for i := 0; i < count; i++ {
				oneThing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(topicTMEIDs[i])).String(),
					PrefLabel: topicNames[i],
					Predicate: conceptMentions,
					Types:     []string{topicURI},
				}
				sectionSuggestion := suggestion{Thing: oneThing, Provenance: []provenance{metadataProvenance}}

				suggestions = append(suggestions, sectionSuggestion)
			}
		}
		if strings.EqualFold("locations", key) {
			for i := 0; i < count; i++ {
				oneThing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(locationTMEIDs[i])).String(),
					PrefLabel: locationNames[i],
					Predicate: conceptMentions,
					Types:     []string{locationURI},
				}
				sectionSuggestion := suggestion{Thing: oneThing, Provenance: []provenance{metadataProvenance}}

				suggestions = append(suggestions, sectionSuggestion)
			}
		}
		if strings.EqualFold("genres", key) {
			for i := 0; i < count; i++ {
				thing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(genreTMEIDs[i])).String(),
					PrefLabel: genreNames[i],
					Predicate: classification,
					Types:     []string{genreURI},
				}
				genreSuggestion := suggestion{Thing: thing, Provenance: []provenance{metadataProvenance}}
				suggestions = append(suggestions, genreSuggestion)
			}
		}
		if strings.EqualFold("specialReports", key) {
			for i := 0; i < count; i++ {
				thing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(specialReportTMEIDs[i])).String(),
					PrefLabel: specialReportNames[i],
					Predicate: classification,
					Types:     []string{specialReportURI},
				}
				specialReportSuggestion := suggestion{Thing: thing, Provenance: []provenance{metadataProvenance}}
				suggestions = append(suggestions, specialReportSuggestion)
			}

			if count > 0 && hasPrimarySection {
				thing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(specialReportTMEIDs[0])).String(),
					PrefLabel: specialReportNames[0],
					Predicate: primaryClassification,
					Types:     []string{specialReportURI},
				}
				specialReportSuggestion := suggestion{Thing: thing}
				suggestions = append(suggestions, specialReportSuggestion)
			}
		}
		if strings.EqualFold("alphavilleSeries", key) {
			for i := 0; i < count; i++ {
				oneThing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(alphavilleSeriesTMEIDs[i])).String(),
					PrefLabel: alphavilleSeriesNames[i],
					Predicate: classification,
					Types:     []string{alphavilleSeriesURI},
				}
				alphavilleSeriesSuggestion := suggestion{Thing: oneThing, Provenance: []provenance{metadataProvenance}}

				suggestions = append(suggestions, alphavilleSeriesSuggestion)
			}
		}
		if strings.EqualFold("organisations", key) {

			for i := 0; i < count; i++ {
				oneThing := thing{
					ID:        "http://api.ft.com/things/" + NewNameUUIDFromBytes([]byte(organisationTMEIDs[i])).String(),
					PrefLabel: organisationNames[i],
					Predicate: conceptMajorMentions,
					Types:     []string{organisationURI},
				}
				organisationSuggestion := suggestion{Thing: oneThing, Provenance: []provenance{metadataProvenance}}

				suggestions = append(suggestions, organisationSuggestion)

			}
		}
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
var locationNames = [...]string{"New York", "Rio"}
var locationTMEIDs = [...]string{"TmV3IFlvcms=-R0w=", "Umlv-R0w="}
var genreNames = [...]string{"News", "Letter"}
var genreTMEIDs = [...]string{"TmV3cw==-R2VucmVz", "TGV0dGVy-R2VucmVz"}
var specialReportNames = [...]string{"Business", "Investment"}
var specialReportTMEIDs = [...]string{"U3BlY2lhbFJlcG9ydHM=-R2Bucm3z", "U3BlY2lhbFJlcG9ydHM=-U2VjdGlvbnM="}
var alphavilleSeriesNames = [...]string{"AV Series 1", "AV Series 2"}
var alphavilleSeriesTMEIDs = [...]string{"series1-AV", "series2-AV"}
var organisationNames = [...]string{"Organisation 1", "Organisation 2"}
var organisationTMEIDs = [...]string{"Organisation-1-TME", "Organisation-2-TME"}
