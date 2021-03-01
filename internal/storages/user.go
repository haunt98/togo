package storages

import (
	"context"
	"database/sql"
)

type UserStorage interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
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
func (udb *UserDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := udb.db.QueryRowContext(ctx, stmt, userID, pwd)
	u := &User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
