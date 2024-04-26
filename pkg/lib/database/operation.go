package database

import (
	"context"
	"database/sql"
)

type DB interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	Begin() (*sql.Tx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	ExecContextTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (sql.Result, error)
}

func (db database) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.conn.GetContext(ctx, dest, query, args...)
}

func (db database) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.conn.SelectContext(ctx, dest, query, args...)
}

func (db database) Begin() (*sql.Tx, error) {
	return db.conn.Begin()
}

func (db database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.conn.ExecContext(ctx, query, args...)
}

func (db database) ExecContextTx(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return tx.ExecContext(ctx, query, args...)
}

func ExecuteTx(tx *sql.Tx, f func() error) error {
	err := f()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
