package database

import "database/sql"

func CreateMap(token string, name string, description string) (sql.Result, error) {
	return exec("./user.db", `
	CREATE TABLE IF NOT EXISTS users (
		id char(36) NOT NULL PRIMARY KEY,
		name TEXT NOT NULL
	);
	`)
}
