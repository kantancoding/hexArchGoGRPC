package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	// this blank import not being in main is necessary
	_ "github.com/go-sql-driver/mysql"
)

// Adapter is compatible with DbPort
type Adapter struct {
	db *sql.DB
}

// NewAdapter creates a new db Adapter
func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connect failure: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

// CloseDbConnection closes the db connection pool
func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

// AddToHistory creates a record in the arith_history table for the
// arithmetic operation
func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.
		Insert("arith_history").Columns(
		"date",
		"answer",
		"operation",
	).Values(
		time.Now(),
		answer,
		operation,
	).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
