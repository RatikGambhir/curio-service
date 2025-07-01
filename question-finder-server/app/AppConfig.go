package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	DB *pgxpool.Pool
}
