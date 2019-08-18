package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	dataSource string
}

func NewDB(dataSource string) *DB {
	return &DB{
		dataSource: dataSource,
	}
}

func (db *DB) Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", db.dataSource)
}
