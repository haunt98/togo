package sqlite

import (
	"context"
	"database/sql"

	"github.com/haunt98/togo/internal/storages"
)

const (
	defaultTaskLen = 10
)

var _ storages.TaskStorage = (*SQLiteDB)(nil)
var _ storages.UserStorage = (*SQLiteDB)(nil)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB(db *sql.DB) *SQLiteDB {
	return &SQLiteDB{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *SQLiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	query := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.db.QueryContext(ctx, query,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: createdDate,
			Valid:  true,
		},
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*storages.Task, 0, defaultTaskLen)
	for rows.Next() {
		t := &storages.Task{}
		if err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
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
func (l *SQLiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	query := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	if _, err := l.db.ExecContext(ctx, query, &t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
		return err
	}

	return nil
}

// Get user by userID
func (l *SQLiteDB) GetUser(ctx context.Context, userID string) (*storages.User, error) {
	query := `SELECT id, password, max_todo FROM users WHERE id = ?`
	row := l.db.QueryRowContext(ctx, query,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
	)

	user := &storages.User{}
	if err := row.Scan(&user.ID, &user.Password, &user.MaxTodo); err != nil {
		return nil, err
	}

	return user, nil
}
