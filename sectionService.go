package main

// SectionService extracts and transforms the section taxonomy into a suggestion
type SectionService struct {
	HandledTaxonomy string
}

const sectionURI = "http://www.ft.com/ontology/Section"

// BuildSuggestions builds a list of section suggestions from a ContentRef.
// Returns an empty array in case no section annotations are found
func (sectionService SectionService) buildSuggestions(contentRef ContentRef) []suggestion {
	sections := extractTags(sectionService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}

	for _, value := range sections {
		suggestions = append(suggestions, buildSuggestion(value, sectionURI, classification))
	}

	if contentRef.PrimarySection.CanonicalName != "" {
		thing := thing{
			ID:        generateID(contentRef.PrimarySection.ID),
			PrefLabel: contentRef.PrimarySection.CanonicalName,
			Predicate: primaryClassification,
			Types:     []string{sectionURI},
		}
		suggestions = append(suggestions, suggestion{Thing: thing})
	}

	return suggestions
}
