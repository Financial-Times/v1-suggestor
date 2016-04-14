package main

// SpecialReportService extracts and transforms the special report taxonomy into a suggestion
type SpecialReportService struct {
	HandledTaxonomy string
}

const specialReportURI = "http://www.ft.com/ontology/SpecialReport"

// BuildSuggestions builds a list of specialReport suggestions from a ContentRef.
// Returns an empty array in case no specialReport annotations are found
func (specialReportService SpecialReportService) buildSuggestions(contentRef ContentRef) []suggestion {
	specialReports := extractTags(specialReportService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}

	for _, value := range specialReports {
		suggestions = append(suggestions, buildSuggestion(value, specialReportURI, classification))
	}

	if contentRef.PrimarySection.CanonicalName != "" {
		thing := thing{
			ID:        generateID(contentRef.PrimarySection.ID),
			PrefLabel: contentRef.PrimarySection.CanonicalName,
			Predicate: primaryClassification,
			Types:     []string{specialReportURI},
		}
		suggestions = append(suggestions, suggestion{Thing: thing})
	}

	return suggestions
}
