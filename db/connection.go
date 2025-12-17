package db

import (
	"database/sql"
	_ "embed"
	"fmt"
)

//go:embed sqlc/schema.sql
var schemaSQL string

func CreateConnection(dsn string) (*Queries, error) {
	conn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	if _, err := conn.Exec(schemaSQL); err != nil {
		return nil, fmt.Errorf("failed database initialization: %w", err)
	}
	return New(conn), nil
}
