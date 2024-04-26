package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type database struct {
	conn *sqlx.DB
}

func InitDatabase() (DB, error) {
	conn, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=123456789 dbname=paperid sslmode=disable")
	if err != nil {
		return nil, err
	}

	return database{
		conn: conn,
	}, nil
}
