package utils

type QueryBuilder interface {
	//Still deciding how I want to implement this
	FromTable(table string) *QueryBuilder
}
