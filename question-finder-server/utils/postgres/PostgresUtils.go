package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToPostgres() (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error
	connStr := "postgres://postgres:RjChicago23!@localhost:5432/postgres?sslmode=disable"
	pool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	return pool, nil
}
