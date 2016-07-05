package main

// AlphavilleSeriesService extracts and transforms the series taxonomy into a suggestion
type AlphavilleSeriesService struct {
	HandledTaxonomy string
}

const alphavilleSeriesURI = "http://www.ft.com/ontology/AlphavilleSeries"

// BuildSuggestions builds a list of topic suggestions from a ContentRef.
// Returns an empty array in case no topic annotations are found
func (alphavilleSeriesService AlphavilleSeriesService) buildSuggestions(contentRef ContentRef) []suggestion {
	series := extractTags(alphavilleSeriesService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}

	for _, value := range series {
		suggestions = append(suggestions, buildSuggestion(value, alphavilleSeriesURI, classification))
	}

	return suggestions
}
