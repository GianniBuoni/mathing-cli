package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed seed.sql
var ddl string

func main() {
	conn, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("❌ issue connecting to db: %v", err)
	}

	if _, err = conn.Exec(ddl); err != nil {
		log.Fatalf("❌ could not seed db data: %v", err)
	}

	fmt.Println("🪴 DB seeded! YAY!")

}
