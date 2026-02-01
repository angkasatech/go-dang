package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("POSTGRES_URL")
	}
	if dsn == "" {
		dsn = os.Getenv("POSTGRES_URI")
	}
	if dsn == "" {
		dsn = os.Getenv("POSTGRES_CONNECTION_STRING")
	}

	if dsn == "" {
		log.Fatal("Database connection string not found. Please set DATABASE_URL (or POSTGRES_URL, POSTGRES_URI, or POSTGRES_CONNECTION_STRING). For Supabase, copy the project's Connection String (URI) and set it in your hosting provider (e.g. Railway Project Variables).")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create tables if they don't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Printf("Failed to create table: %v", err)
	} else {
		log.Println("Database schema verified")
	}

	log.Println("Connected to PostgreSQL")
	return db
}
