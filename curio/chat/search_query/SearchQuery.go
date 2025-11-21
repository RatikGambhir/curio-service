package searchquery

import (
	"curio/app"
	"fmt"
)

type SearchQuery struct {
	Query string
	Scope string
}

func NewSearchQuery(appConfig *app.AppConfig) *SearchQuery {
	return &SearchQuery{
		Query: genQueryFields([]string{"id", "name", "email"}),
	}

}

func genQueryFields(fields []string) string {
	var queryFields string
	for _, field := range fields {
		queryFields += field
		queryFields += ", "
	}

	return fmt.Sprintf("SELECT %s ", queryFields)

}

func (sq *SearchQuery) FromTable(table string) *SearchQuery {
	sq.Query = fmt.Sprintf("%s FROM %s", sq.Query, table)
	return sq
}

func genQuestionSource(questionSearchSource string) string {
	if questionSearchSource == "question" || questionSearchSource == "response" {
		return questionSearchSource
	}
	return "(question || response)"
}

func (sq *SearchQuery) WherePowerSearchParams(param []string, inclusive bool, questionSearchSource string) *SearchQuery {
	if len(param) == 0 {
		return sq
	}
	var searchSource = genQuestionSource(questionSearchSource)
	searchSource = genQuestionSource(questionSearchSource)

	var whereClause string
	for _, value := range param {
		//Appen tsvector params here, need some logic on the question source
		//	whereClause += fmt.Sprintf(" AND %s = '%s'", "id", value)
	}

	sq.Query = fmt.Sprintf("%s WHERE %s", sq.Query, whereClause)
	return sq
}

//TODO: Implement regular search params, we can pass in a map
