package searchquery

import (
	"curio/app"
)

type SearchQuery struct {
	Query []string
	Scope string
}

func NewSearchQuery(appConfig *app.AppConfig) *SearchQuery {
	return &SearchQuery{
		Query: []string{"id", "name", "email"},
	}
}
