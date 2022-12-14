package database

import (
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

const defaultTimeout = 3 * time.Second

type DB struct {
	*sqlx.DB
}

func New(dsn string) (*DB, error) {
	db, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	return &DB{db}, nil
}
