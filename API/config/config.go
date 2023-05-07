package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDatabase() (*sql.DB, error) {
	database, err := sql.Open("postgres", "postgres://postgres:123589@localhost/laboraProject1?sslmode=disable")

	return database, err
}
