package utils

import (
	"fmt"
	"strings"
)

type IQueryBuilder interface {
	//Still deciding how I want to implement this
	SelectFields(fields []string) *QueryBuilder
	FromTable(table string) *QueryBuilder
	AndWhereFields(fields map[string]interface{}) *QueryBuilder
}

type QueryBuilder struct {
	fields        []string
	table         string
	conditions    map[string]interface{}
	condition_val []any
	query         string
}

func (qb *QueryBuilder) SelectFields(fields []string) *QueryBuilder {
	qb.fields = fields
	var selectQuery = fmt.Sprintf(`SELECT %s`, strings.Join(fields, ", "))
	qb.query = selectQuery
	return qb
}

func (qb *QueryBuilder) FromTable(table string) *QueryBuilder {
	qb.table = table
	qb.query = fmt.Sprintf("%s FROM %s", qb.query, qb.table)
	return qb
}

func (qb *QueryBuilder) AndWhereFields(fields map[string]interface{}) *QueryBuilder {
	//TODO: Need to fix ordering with fields entered to match placeholder
	for key, value := range fields {
		var placeHolderIdx = len(qb.condition_val) + 1
		var slice = append(qb.condition_val, value)
		qb.condition_val = slice
		var placeholder = fmt.Sprintf(`$%d`, placeHolderIdx)
		if placeHolderIdx == 1 {
			qb.query = fmt.Sprintf(`%s WHERE %s = %s`, qb.query, key, placeholder)
		} else {
			qb.query = fmt.Sprintf(`%s AND %s = %s`, qb.query, key, placeholder)
		}
	}
	return qb
}

func (qb *QueryBuilder) OrWhereFields(fields map[string]interface{}) *QueryBuilder {
	//TODO: Need to fix ordering with fields entered to match placeholder

	for key, value := range fields {
		var placeHolderIdx = len(qb.condition_val) + 1
		var slice = append(qb.condition_val, value)
		qb.condition_val = slice
		var placeHolder = fmt.Sprintf(`$%d`, placeHolderIdx)
		if placeHolderIdx == 1 {
			qb.query = fmt.Sprintf(`%s WHERE %s = %s`, qb.query, key, placeHolder)
		} else {
			qb.query = fmt.Sprintf(`%s OR %s = %s`, qb.query, key, placeHolder)
		}
	}
	return qb
}

func (qb *QueryBuilder) OrderByAscending(field string) *QueryBuilder {
	qb.query = fmt.Sprintf(`%s ORDER BY %s ASCENDING`, qb.query, field)
	return qb
}

func (qb *QueryBuilder) OrderByDescending(field string) *QueryBuilder {
	qb.query = fmt.Sprintf(`%s ORDER BY %s DESCENDING`, qb.query, field)
	return qb
}

func (qb *QueryBuilder) Build() (string, []any) {
	return qb.query, qb.condition_val
}

// func (qb *QueryBuilder) JoinTable(target_table string, target_table_field string, source_table_field string) *QueryBuilder {
// 	// Work on join later
// }
