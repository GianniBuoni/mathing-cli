package store

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed sqlc/schema.sql
var ddl string

type Store struct {
	Ctx     *context.Context
	Queries *Queries
}

func NewStore(ctx *context.Context) (Store, error) {
	conn, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return Store{}, fmt.Errorf("issue connecting to db: %w", err)
	}

	if _, err = conn.ExecContext(*ctx, ddl); err != nil {
		return Store{}, fmt.Errorf("issue creating tables: %w", err)
	}

	queries := New(conn)

	return Store{
		Ctx:     ctx,
		Queries: queries,
	}, nil
}
