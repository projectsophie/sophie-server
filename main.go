package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	usersDatabase, _ = sql.Open("sqlite3", "./users.db")
	statement, _ := usersDatabase.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, nickname TEXT, password TEXT, email TEXT, email_verified BOOLEAN, workspaces INTEGER []);")
	statement.Exec()
}

func main() {
	InitDB()
}
