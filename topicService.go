package service

import "github.com/Financial-Times/v1-suggestor/model"

// TopicService extracts and transforms the topic taxonomy into a suggestion
type TopicService struct {
	HandledTaxonomy string
}

const topicURI = "http://www.ft.com/ontology/thing/Topic"
const mentionsPredicate = "mentions"
const aboutPredicate = "about"

// BuildSuggestions builds a list of topic suggestions from a ContentRef.
// Returns an empty array in case no topic annotations are found
func (topicService TopicService) BuildSuggestions(contentRef model.ContentRef) []model.Suggestion {
	topics := extractTags(topicService.HandledTaxonomy, contentRef)
	suggestions := []model.Suggestion{}

	for _, value := range topics {
		suggestions = append(suggestions, buildSuggestion(value, topicURI, mentionsPredicate))
		suggestions = append(suggestions, buildSuggestion(value, topicURI, aboutPredicate))
	}

	return suggestions
}
