package service

import "github.com/Financial-Times/v1-suggestor/model"

// TaxonomyService defines the operations used to process taxonomies
type TaxonomyService interface {
	BuildSuggestions(model.ContentRef) []model.Suggestion
}

// RelevanceURI used as scoring system identifier
const RelevanceURI = "http://api.ft.com/scoringsystem/FT-RELEVANCE-SYSTEM"

// ConfidenceURI used as scoring system identifier
const ConfidenceURI = "http://api.ft.com/scoringsystem/FT-CONFIDENCE-SYSTEM"
