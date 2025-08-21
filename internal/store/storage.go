package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("rsource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
	}

	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}

	Cells interface {
		Create(context.Context, *Cell) (*Cell, error)
		GetByID(context.Context, int64) (*Cell, error)
		GetCells(context.Context, *PaginatedCellQuery) (*[]Cell, error)
		Delete(context.Context, int64) error
		IsEmpty(context.Context) (bool, error)
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
		Roles: &RoleStore{db},
		Cells: &CellStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
