package utils

import (
	//	"reflect"
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
		t.Errorf("expected no args, got %v", args...)
	}
}

func TestQueryBuilder_WhereClause(t *testing.T) {
	var object = map[string]interface{}{
		"id":    124,
		"email": "test@example.com",
	}
	qb := &QueryBuilder{}

	query_builder, args := qb.
		SelectFields([]string{"id", "email"}).
		FromTable("public.user").
		AndWhereFields(object).
		Build()

	expectedQuery := "SELECT id, email FROM public.user WHERE id = $1 AND email = $2"
	if query_builder != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, query_builder)
	}

	if len(args) != 2 {
		t.Errorf("expected two args, got %v", args...)
	}

	if args[0] != 124 || args[1] != "test@example.com" {
		t.Errorf("expected args [124, 'test@example.com'], got %v", args)
	}
}
