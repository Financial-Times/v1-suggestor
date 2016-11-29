package main

// LocationService extracts and transforms the location taxonomy into a suggestion
type LocationService struct {
	HandledTaxonomy string
}

const locationURI = "http://www.ft.com/ontology/Location"

// BuildSuggestions builds a list of location suggestions from a ContentRef.
// Returns an empty array in case no location annotations are found
func (locationService LocationService) buildSuggestions(contentRef ContentRef) []suggestion {
	locations := extractTags(locationService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}

	for _, value := range locations {
		suggestions = append(suggestions, buildSuggestion(value, locationURI, conceptMentions))
	}

	if contentRef.PrimaryTheme.CanonicalName != "" {
		thing := thing{
			ID:        generateID(contentRef.PrimaryTheme.ID),
			PrefLabel: contentRef.PrimaryTheme.CanonicalName,
			Predicate: about,
			Types:     []string{locationURI},
		}
		suggestions = append(suggestions, suggestion{Thing: thing})
	}

	return suggestions
}
