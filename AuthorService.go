package main

// AuthorService extracts and transforms the author taxonomy into a suggestion
type AuthorService struct {
	HandledTaxonomy string
}

const authorURI = "http://www.ft.com/ontology/person/Person"

// BuildSuggestions builds a list of author suggestions from a ContentRef.
// Returns an empty array in case no author annotations are found
func (authorService AuthorService) buildSuggestions(contentRef ContentRef) []suggestion {
	authors := extractTags(authorService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}

	for _, value := range authors {
		suggestions = append(suggestions, buildSuggestion(value, authorURI, hasAuthor))
	}

	return suggestions
}
