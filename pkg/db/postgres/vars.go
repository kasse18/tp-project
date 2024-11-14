package postgres

const (
	queryInitUsers = `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		login text NOT NULL,
		password text NOT NULL,
		email text NOT NULL,
		created_at timestamp NOT NULL
	  )`

	queryInitTattoos = `CREATE TABLE IF NOT EXISTS tattoos (
		id serial PRIMARY KEY,
		name text NOT NULL,
		artist text NOT NULL,
		price int NOT NULL,
		image text NOT NULL,
	  )`
)
