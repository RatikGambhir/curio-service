package utils

import (
	"testing"
)

func TestQueryBuilder_SelectFrom(t *testing.T) {
	qb := &QueryBuilder{}

	query_builder, args := qb.
		SelectFields([]string{"id", "email"}).
		FromTable("public.user").
		Build()

	expectedQuery := "SELECT id, email FROM public.user"
	if query_builder != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, query_builder)
	}

	if len(args) != 0 {
		t.Errorf("expected no args, got %v", args)
	}
}

func TestQueryBuilder_SelectFieldsSingleField(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"*"})

	expectedQuery := "SELECT *"
	if qb.query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, qb.query)
	}
}

func TestQueryBuilder_WhereConditionsSingle(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"id", "email"})
	qb.FromTable("public.user")

	conditions := []Conditions{
		{
			conditions: []Condition{
				{
					field:    "id",
					operator: "=",
					value:    "$1",
				},
			},
		},
	}

	result, err := qb.WhereConditions(conditions)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedQuery := "WHERE   id = $1"
	if qb.query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, qb.query)
	}

	if result != qb {
		t.Error("WhereConditions should return the QueryBuilder instance for chaining")
	}
}

func TestQueryBuilder_WhereConditionsMultipleWithAnd(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"id", "email"})
	qb.FromTable("public.user")

	andOperator := "AND"
	conditions := []Conditions{
		{
			conditions: []Condition{
				{
					field:    "id",
					operator: "=",
					value:    "$1",
				},
				{
					field:             "email",
					operator:          "=",
					value:             "$2",
					inclusiveOperator: &andOperator,
				},
			},
		},
	}

	result, err := qb.WhereConditions(conditions)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedQuery := "WHERE   id = $1 AND email = $2"
	if qb.query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, qb.query)
	}

	if result != qb {
		t.Error("WhereConditions should return the QueryBuilder instance for chaining")
	}
}

func TestQueryBuilder_WhereConditionsMissingInclusiveOperator(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"id", "name"})
	qb.FromTable("users")

	conditions := []Conditions{
		{
			conditions: []Condition{
				{
					field:    "id",
					operator: "=",
					value:    "$1",
				},
				{
					field:             "name",
					operator:          "=",
					value:             "$2",
					inclusiveOperator: nil,
				},
			},
		},
	}

	result, err := qb.WhereConditions(conditions)
	if err == nil {
		t.Error("expected error for missing inclusive operator, got nil")
	}

	if err.Error() != "please make sure to include the inclusive operator" {
		t.Errorf("expected specific error message, got %q", err.Error())
	}

	if result != nil {
		t.Error("expected nil result when error occurs")
	}
}

func TestQueryBuilder_WhereConditionsWithOr(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"*"})
	qb.FromTable("products")

	orOperator := "OR"
	conditions := []Conditions{
		{
			conditions: []Condition{
				{
					field:    "price",
					operator: ">",
					value:    "$1",
				},
				{
					field:             "category",
					operator:          "=",
					value:             "$2",
					inclusiveOperator: &orOperator,
				},
			},
		},
	}

	result, err := qb.WhereConditions(conditions)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedQuery := "WHERE   price > $1 OR category = $2"
	if qb.query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, qb.query)
	}

	if result != qb {
		t.Error("WhereConditions should return the QueryBuilder instance for chaining")
	}
}

func TestQueryBuilder_OrderByAscending(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"id", "name"})
	qb.FromTable("users")
	result := qb.OrderByAscending("created_at")

	expectedQuery := "SELECT id, name FROM users ORDER BY created_at ASCENDING"
	if qb.query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, qb.query)
	}

	if result != qb {
		t.Error("OrderByAscending should return the QueryBuilder instance for chaining")
	}
}

func TestQueryBuilder_OrderByDescending(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"id", "name"})
	qb.FromTable("users")
	result := qb.OrderByDescending("created_at")

	expectedQuery := "SELECT id, name FROM users ORDER BY created_at DESCENDING"
	if qb.query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, qb.query)
	}

	if result != qb {
		t.Error("OrderByDescending should return the QueryBuilder instance for chaining")
	}
}

func TestQueryBuilder_CompleteQueryWithWhereAndOrderBy(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"id", "name", "email"})
	qb.FromTable("users")

	andOperator := "AND"
	conditions := []Conditions{
		{
			conditions: []Condition{
				{
					field:    "active",
					operator: "=",
					value:    "$1",
				},
				{
					field:             "role",
					operator:          "=",
					value:             "$2",
					inclusiveOperator: &andOperator,
				},
			},
		},
	}

	_, err := qb.WhereConditions(conditions)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	qb.OrderByDescending("created_at")

	query, args := qb.Build()

	expectedQuery := "WHERE   active = $1 AND role = $2 ORDER BY created_at DESCENDING"
	if query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, query)
	}

	if len(args) != 0 {
		t.Errorf("expected no args, got %v", args)
	}
}

func TestQueryBuilder_MethodChaining(t *testing.T) {
	qb := &QueryBuilder{}

	andOperator := "AND"
	conditions := []Conditions{
		{
			conditions: []Condition{
				{
					field:    "status",
					operator: "=",
					value:    "$1",
				},
				{
					field:             "verified",
					operator:          "=",
					value:             "$2",
					inclusiveOperator: &andOperator,
				},
			},
		},
	}

	result, err := qb.SelectFields([]string{"id", "username", "email"}).
		FromTable("users").
		WhereConditions(conditions)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	result.OrderByAscending("username")

	query, args := result.Build()

	expectedQuery := "WHERE   status = $1 AND verified = $2 ORDER BY username ASCENDING"
	if query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, query)
	}

	if len(args) != 0 {
		t.Errorf("expected no args, got %v", args)
	}
}

func TestQueryBuilder_ThreeConditions(t *testing.T) {
	qb := &QueryBuilder{}
	qb.SelectFields([]string{"*"})
	qb.FromTable("orders")

	andOperator := "AND"
	orOperator := "OR"
	conditions := []Conditions{
		{
			conditions: []Condition{
				{
					field:    "status",
					operator: "=",
					value:    "$1",
				},
				{
					field:             "total",
					operator:          ">",
					value:             "$2",
					inclusiveOperator: &andOperator,
				},
				{
					field:             "user_id",
					operator:          "=",
					value:             "$3",
					inclusiveOperator: &orOperator,
				},
			},
		},
	}

	_, err := qb.WhereConditions(conditions)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	query, _ := qb.Build()

	expectedQuery := "WHERE   status = $1 AND total > $2 OR user_id = $3"
	if query != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, query)
	}
}
