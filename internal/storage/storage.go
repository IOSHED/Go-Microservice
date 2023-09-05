package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	Db *sql.DB
}

func New(driverName string, storagePath string) (*Storage, error) {
	const op = "storage.New"

	db, err := sql.Open(driverName, storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db}, nil
}
