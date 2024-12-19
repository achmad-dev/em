package utils

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSqlDB(connStr string) (*sqlx.DB, error) {
	// init sql db
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
