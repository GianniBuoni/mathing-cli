package store

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed sql/schema.sql
var ddl string

func NewStore() (*Queries, error) {
	ctx := context.Background()
	conn, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, fmt.Errorf("issue connecting to db: %w", err)
	}

	if _, err = conn.ExecContext(ctx, ddl); err != nil {
		return nil, fmt.Errorf("issue creating tables: %w", err)
	}
	return New(conn), nil

}
