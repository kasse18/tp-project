package postgres

import (
	"context"
	"database/sql"
	"fmt"

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

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		panic(err)
	}

	queries := []string{
		queryInitUsers,
		queryInitTattoos,
		queryInitCart,
		queryInitOrders,
		queryInitOrderItems,
		queryInitTags,
		queryInitTattooTags,
	}

	for _, query := range queries {
		_, err := tx.ExecContext(ctx, query)
		if err != nil {
			tx.Rollback()
			panic(fmt.Errorf("failed to execute query: %v\nQuery: %s", err, query))
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	return &DB{
		DB: db,
	}
}
