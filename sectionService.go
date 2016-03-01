package main

// SectionService extracts and transforms the section taxonomy into a suggestion
type SectionService struct {
	HandledTaxonomy string
}

const sectionURI = "http://www.ft.com/ontology/thing/Section"

// BuildSuggestions builds a list of section suggestions from a ContentRef.
// Returns an empty array in case no section annotations are found
func (sectionService SectionService) BuildSuggestions(contentRef ContentRef) []Suggestion {
	sections := extractTags(sectionService.HandledTaxonomy, contentRef)
	suggestions := []Suggestion{}

	for _, value := range sections {
		suggestions = append(suggestions, buildSuggestion(value, sectionURI, predicate))
	}

	if contentRef.PrimarySection.CanonicalName != "" {
		thing := Thing{
			ID:        generateID(contentRef.PrimarySection.ID),
			PrefLabel: contentRef.PrimarySection.CanonicalName,
			Predicate: primaryPredicate,
			Types:     []string{sectionURI},
		}
		suggestions = append(suggestions, Suggestion{Thing: thing})
	}

	return suggestions
}
