package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func exec(db string, SQL string, args ...any) (sql.Result, error) {
	sqliteDatabase, _ := sql.Open("sqlite3", db)
	defer sqliteDatabase.Close()
	return sqliteDatabase.Exec(SQL, args...)
}

func query(db string, SQL string, args ...any) (*sql.Rows, error) {
	sqliteDatabase, _ := sql.Open("sqlite3", db)
	defer sqliteDatabase.Close()
	return sqliteDatabase.Query(SQL, args...)
}

func InitDatabase() {
	CreateUserDatabase()
}
