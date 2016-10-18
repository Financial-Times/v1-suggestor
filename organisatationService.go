package main

// SubjectService extracts and transforms the subject taxonomy into a suggestion
type OrganisationService struct {
	HandledTaxonomy string
}

const organisationURI = "http://www.ft.com/ontology/organisation/Organisation"

// BuildSuggestions builds a list of subject suggestions from a ContentRef.
// Returns an empty array in case no subject annotations are found
func (organisationService OrganisationService) buildSuggestions(contentRef ContentRef) []suggestion {
	subjects := extractTags(organisationService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}
	for _, value := range subjects {
		suggestions = append(suggestions, buildSuggestion(value, organisationURI, conceptMajorMentions))
	}

	return suggestions
}
