package main

// GenreService extracts and transforms the genre taxonomy into a suggestion
type GenreService struct {
	HandledTaxonomy string
}

const genreURI = "http://www.ft.com/ontology/Genre"

// BuildSuggestions builds a list of genre suggestions from a ContentRef.
// Returns an empty array in case no genre annotations are found
func (genreService GenreService) buildSuggestions(contentRef ContentRef) []suggestion {
	genres := extractTags(genreService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}

	for _, value := range genres {
		suggestions = append(suggestions, buildSuggestion(value, genreURI, classification))
	}

	return suggestions
}
