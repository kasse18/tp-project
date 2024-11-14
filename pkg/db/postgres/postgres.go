package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func New(ctx context.Context, connStr string) *DB {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		panic(err)
	}

	_, err = conn.ExecContext(ctx, queryInitUsers)
	if err != nil {
		panic(err)
	}

	_, err = conn.ExecContext(ctx, queryInitTattoos)
	if err != nil {
		panic(err)
	}

	return &DB{
		DB: db,
	}
}
