package storages

import (
	"context"
	"database/sql"
)

type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
}

var _ TaskStorage = (*TaskDB)(nil)

type TaskDB struct {
	db *sql.DB
}

func NewTaskDB(db *sql.DB) *TaskDB {
	return &TaskDB{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (tdb *TaskDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := tdb.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		t := &Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (tdb *TaskDB) AddTask(ctx context.Context, t *Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := tdb.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}
