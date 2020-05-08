package tests

import (
	"database/sql"
)

const createTable = `
DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
  ID serial PRIMARY KEY,
  TITLE TEXT NOT NULL,
  USER_ID INTEGER 
); 
`

type TestDB struct {
	db *sql.DB
}

func Setup() *sql.DB {
	db, err := connectPostgresForTests()
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(createTable); err != nil {
		panic(err)
	}

	return db
}

func connectPostgresForTests() (*sql.DB, error) {
	connStr := "postgres://postgres@db:5432/todo?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
