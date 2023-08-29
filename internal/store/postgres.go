package store

import (
	"database/sql"
)

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Db: db,
	}
}
