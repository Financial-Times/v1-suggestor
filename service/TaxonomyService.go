package service

import "github.com/Financial-Times/v1-suggestor/model"

type TaxonomyService interface {
	BuildSuggestions(model.ContentRef) []model.Suggestion
}

const RelevanceUri = "http://api.ft.com/scoringsystem/FT-RELEVANCE-SYSTEM"
const ConfidenceUri = "http://api.ft.com/scoringsystem/FT-CONFIDENCE-SYSTEM"
