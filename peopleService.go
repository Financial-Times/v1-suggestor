package main

// SubjectService extracts and transforms the subject taxonomy into a suggestion
type PeopleService struct {
	HandledTaxonomy string
}

const personURI = "http://www.ft.com/ontology/person/Person"

// BuildSuggestions builds a list of subject suggestions from a ContentRef.
// Returns an empty array in case no subject annotations are found
func (peopleService PeopleService) buildSuggestions(contentRef ContentRef) []suggestion {
	people := extractTags(peopleService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}
	for _, value := range people {
		suggestions = append(suggestions, buildSuggestion(value, personURI, conceptMajorMentions))
	}

	return suggestions
}
