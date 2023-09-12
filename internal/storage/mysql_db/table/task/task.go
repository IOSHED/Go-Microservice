package task

import (
	"fmt"
	"micriservice/internal/storage"
	err_mysql "micriservice/internal/storage/mysql_db"
	"micriservice/internal/storage/mysql_db/table"

	"github.com/go-sql-driver/mysql"
)

type TaskStorage struct {
	storage.Storage
}

func New(storage *storage.Storage) *TaskStorage {
	return &TaskStorage{
		Storage: *storage,
	}
}

// Saving the task. The name of the task and its description are passed.
func (t *TaskStorage) SaveTask(name string, description string) (int64, error) {
	const op = "storage.mysql.table.task.Save"

	stmt, err := t.Db.Prepare("INSERT INTO task(name, description) VALUES(?, ?)")
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	res, err := stmt.Exec(name, description)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == err_mysql.ErrNonUniqueValue {
			return 0, fmt.Errorf("%s: %w", op, table.ErrNameTaskExists)
		}

		return -1, fmt.Errorf("%s : %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	return id, nil
}

// Get a description of the task by name.
func (t *TaskStorage) GetDescriptionByName(name string) (string, error) {
	const op = "storage.mysql.table.task.GetDescriptionByName"

	stmt, err := t.Db.Prepare("SELECT description FROM task WHERE name = ?")
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

// Get the task id by name.
func (t *TaskStorage) GetByName(name string) (int64, error) {
	const op = "storage.mysql.table.task.GetByName"

	stmt, err := t.Db.Prepare("SELECT id FROM task WHERE name = ?")
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	var id int64

	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	return id, nil
}

// Вelete a task by name.
func (t *TaskStorage) DeleteByName(name string) (int64, error) {
	const op = "storage.mysql.table.task.DeleteByName"

	id, err := t.GetByName(name)
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	_, err = t.Db.Exec("DELETE FROM task WHERE name = ?", name)
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	return id, nil
}
