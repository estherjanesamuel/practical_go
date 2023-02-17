package config

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ConnectMySql(connStr string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		log.Printf("error connecting to mysql: %v", err)
		return nil
	}

	return db
}