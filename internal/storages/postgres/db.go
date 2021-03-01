package postgres

import (
	"context"
	"database/sql"

	"github.com/haunt98/togo/internal/storages"
)

const (
	defaultTaskLen = 10
)

var _ storages.TaskStorage = (*PostgresDB)(nil)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *PostgresDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	query := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := p.db.QueryContext(ctx, query, userID, createdDate)
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
func (p *PostgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	query := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	if _, err := p.db.ExecContext(ctx, query, &t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (p *PostgresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) (bool, error) {
	query := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := p.db.QueryRowContext(ctx, query, userID, pwd)

	u := &storages.User{}
	if err := row.Scan(&u.ID); err != nil {
		return false, err
	}

	return true, nil
}
