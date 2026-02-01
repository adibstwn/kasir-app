package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	cfg, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	cfg.StatementCacheCapacity = 0
	cfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db := stdlib.OpenDB(*cfg)

	// Best practice pool config
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Second)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully")
	return db, nil
}
