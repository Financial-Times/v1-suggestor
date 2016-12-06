package main

const brandURI = "http://www.ft.com/ontology/Brand"

type BrandService struct {
	HandledTaxonomy string
}

func (brandService BrandService) buildSuggestions(contentRef ContentRef) []suggestion {
	authors := extractTags(brandService.HandledTaxonomy, contentRef)
	suggestions := []suggestion{}
	for _, value := range authors {
		suggestions = append(suggestions, buildSuggestion(value, brandURI, classification))
	}
	return suggestions
}