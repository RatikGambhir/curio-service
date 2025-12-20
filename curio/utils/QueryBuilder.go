package utils

import (
	"fmt"
	"strings"
)

type Conditions struct {
	conditions     []Condition
	querySubstring string
}

type Condition struct {
	field             string
	operator          string
	inclusiveOperator *string
	value             string
}

type QueryBuilder struct {
	fields        []string
	table         string
	conditions    map[string]interface{}
	condition_val []any
	condition     []Conditions
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

func (qb *QueryBuilder) WhereConditions(conditions []Conditions) (*QueryBuilder, error) {

	var querySubstring string = "WHERE "
	qb.condition = conditions

	for i, v := range conditions {
		var conditionArray = v.conditions
		for j, k := range conditionArray {
			var inclusiveOperator = ""
			if j > 0 {
				if k.inclusiveOperator == nil {
					return nil, fmt.Errorf("please make sure to include the inclusive operator")
				}
				inclusiveOperator = *conditionArray[j].inclusiveOperator
			}
			var field = k.field
			var operator = k.operator
			var value = k.value

			row := []string{
				inclusiveOperator,
				field,
				operator,
				value,
			}

			var whereSubstring = strings.Join(row, " ")
			var whereCondition = fmt.Sprintf(`%s %s`, querySubstring, whereSubstring)
			querySubstring = whereCondition
			conditions[i].querySubstring = querySubstring

		}

	}
	qb.query = querySubstring

	return qb, nil
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
