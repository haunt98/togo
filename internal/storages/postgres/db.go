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
var _ storages.UserStorage = (*PostgresDB)(nil)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *PostgresDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	query := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := p.db.QueryContext(ctx, query,
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
func (p *PostgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	query := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	if _, err := p.db.ExecContext(ctx, query, &t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
		return err
	}

	return nil
}

// Get user by userID
func (p *PostgresDB) GetUser(ctx context.Context, userID string) (*storages.User, error) {
	query := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := p.db.QueryRowContext(ctx, query,
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
