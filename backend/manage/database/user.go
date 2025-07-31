package database

import "database/sql"

func CreateUserDatabase() (sql.Result, error) {
	return exec("./user.db", `
	CREATE TABLE IF NOT EXISTS users (
		id char(36) NOT NULL PRIMARY KEY,
		name TEXT NOT NULL
	);
	`)
}

func CreateUser(token string, name string) (sql.Result, error) {
	return exec("./user.db", `
	INSERT INTO users (id, name) VALUES (?, ?)
	`, token, name)
}

func GetUser(token string) (*sql.Rows, error) {
	return query("./user.db", `
		SELECT id, name FROM users WHERE id = ?
	`, token)
}
