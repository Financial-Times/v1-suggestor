package main


// SubjectService extracts and transforms the subject taxonomy into a suggestion
type SubjectService struct {
	HandledTaxonomy string
}

const subjectURI = "http://www.ft.com/ontology/thing/Subject"

// BuildSuggestions builds a list of subject suggestions from a ContentRef.
// Returns an empty array in case no subject annotations are found
func (subjectService SubjectService) BuildSuggestions(contentRef ContentRef) []Suggestion {
	subjects := extractTags(subjectService.HandledTaxonomy, contentRef)
	suggestions := []Suggestion{}

	for _, value := range subjects {
		suggestions = append(suggestions, buildSuggestion(value, subjectURI, predicate))
	}

	return suggestions
}
