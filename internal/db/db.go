package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() error {
	url := os.Getenv("POSTGRES_URL")
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return fmt.Errorf("cannot connect to postgres: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("cannot ping postgres: %w", err)
	}
	Pool = pool
	fmt.Println("✅ Connected to Postgres")
	return nil
}
