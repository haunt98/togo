package storages

import (
	"context"
	"database/sql"
)

type UserStorage interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) (bool, error)
}

var _ UserStorage = (*UserDB)(nil)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

// ValidateUser returns tasks if match userID AND password
func (udb *UserDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) (bool, error) {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := udb.db.QueryRowContext(ctx, stmt, userID, pwd)
	if row.Err() != nil {
		return false, row.Err()
	}

	u := &User{}
	if err := row.Scan(&u.ID); err != nil {
		return false, err
	}

	return true, nil
}
