package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DBI interface {
	Querier
	Conn() ConnI
	Begin(ctx context.Context) (*Tx, error)
}

type ConnI interface {
	DBTX
	Begin(ctx context.Context) (pgx.Tx, error)
}

type DB struct {
	conn ConnI
	*Queries
}

func NewDB(conn ConnI) *DB {
	return &DB{conn: conn, Queries: New(conn)}
}

func (d *DB) Conn() ConnI { return d.conn }

func (d *DB) Begin(ctx context.Context) (*Tx, error) {
	tx, err := d.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{tx: tx, Queries: New(tx)}, nil
}

type Tx struct {
	tx pgx.Tx
	*Queries
}

func (t *Tx) Conn() ConnI { return t.tx }

func (t *Tx) Begin(ctx context.Context) (*Tx, error) {
	panic("recursive transactions are not supported")
}

func (t *Tx) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *Tx) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

// CommitOrRollback can be used in a defer statement to commit or rollback a
// transaction depending on whether the enclosing function returned an error.
//
//	func myFunc() (err error) {
//	  tx, err := db.Begin(ctx)
//	  if err != nil { return err }
//	  defer tx.CommitOrRollback(ctx, &err)
//	  ...
//	}
func (t *Tx) CommitOrRollback(ctx context.Context, err *error) {
	if *err != nil {
		_ = t.Rollback(ctx)
	} else {
		*err = t.Commit(ctx)
	}
}
