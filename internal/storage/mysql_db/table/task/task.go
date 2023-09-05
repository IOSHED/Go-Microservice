package task

import (
	"database/sql"
	"fmt"
	err_mysql "micriservice/internal/storage/mysql_db"
	"micriservice/internal/storage/mysql_db/table"

	"github.com/go-sql-driver/mysql"
)

type Task struct{}

// Saving the task. The name of the task and its description are passed.
func (t *Task) SaveTask(db *sql.DB, name string, description string) (int64, error) {
	const op = "storage.mysql.table.task.Save"

	stmt, err := db.Prepare("INSERT INTO task(name, description) VALUES(?, ?)")
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
func GetDescriptionByName(db *sql.DB, name string) (string, error) {
	const op = "storage.mysql.table.task.GetDescriptionByName"

	stmt, err := db.Prepare("SELECT description FROM task WHERE name = ?")
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
func GetByName(db *sql.DB, name string) (int64, error) {
	const op = "storage.mysql.table.task.GetByName"

	stmt, err := db.Prepare("SELECT id FROM task WHERE name = ?")
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
func DeleteByName(db *sql.DB, name string) (int64, error) {
	const op = "storage.mysql.table.task.DeleteByName"

	id, err := GetByName(db, name)
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	_, err = db.Exec("DELETE FROM task WHERE name = ?", name)
	if err != nil {
		return -1, fmt.Errorf("%s : %w", op, err)
	}

	return id, nil
}
