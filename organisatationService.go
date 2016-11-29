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

	if contentRef.PrimaryTheme.CanonicalName != "" {
		thing := thing{
			ID:        generateID(contentRef.PrimaryTheme.ID),
			PrefLabel: contentRef.PrimaryTheme.CanonicalName,
			Predicate: about,
			Types:     []string{organisationURI},
		}
		suggestions = append(suggestions, suggestion{Thing: thing})
	}

	return suggestions
}
