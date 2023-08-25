package mysql

import (
	"database/sql"
	"fmt"
	"micriservice/internal/storage"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.mysql.New"

	db, err := sql.Open("mysql", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db}, nil
}

func (s *Storage) SaveTask(name string, description string) (int64, error) {
	const op = "storage.mysql.SaveTask"

	stmt, err := s.db.Prepare("INSERT INTO task(name, description) VALUES(?, ?)")
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	res, err := stmt.Exec(name, description)
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("%s : %s %w", op, storage.ErrGetInsertId, err)
	}

	return id, nil
}

func (s *Storage) GetTaskByName(name string) (string, error) {
	const op = "storage.mysql.GetTaskByName"

	stmt, err := s.db.Prepare("SELECT description FROM task WHERE name = ?")
	if err != nil {
		return "", fmt.Errorf("%s : %w", op, err)
	}

	var description string

	err = stmt.QueryRow(name).Scan(&description)
	if err != nil {
		return "", fmt.Errorf("%s : %w", op, err)
	}

	return description, nil
}
