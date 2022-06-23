package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Database variables.
// usersDatabase is a database with user data.
// workspacesDatabase is a database with workspaces data.
var (
	usersDatabase      *sql.DB
	workspacesDatabase *sql.DB
)

// InitDB Creates databases if they are not exist.
// You can initialize every DB separately.
func InitDB(users bool, workspaces bool) {
	if users {
		usersDatabase, _ = sql.Open("sqlite3", "./users.db")
		statement, _ := usersDatabase.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, nickname TEXT, password TEXT, email TEXT, emailVerified BOOLEAN, workspaces INTEGER []);")
		_, _ = statement.Exec()
		log.Println("Users database was initialized successfully.")
	}
	if workspaces {
		workspacesDatabase, _ = sql.Open("sqlite3", "./workspaces.db")
		statement, _ := workspacesDatabase.Prepare("CREATE TABLE IF NOT EXISTS workspaces (id INTEGER PRIMARY KEY, members TEXT [], pages TEXT []);")
		_, _ = statement.Exec()
		log.Println("Workspaces database was initialized successfully.")
	}
}

// GetUsersDB Returns a usersDatabase
func GetUsersDB() *sql.DB {
	if usersDatabase != nil {
		return usersDatabase
	}
	panic("Users database not initialized. Seems, that InitDB() hasn't been executed or usersDatabase hasn't been initialized separately.")
}

// GetWorkspacesDB Returns a workspacesDatabase
func GetWorkspacesDB() *sql.DB {
	if workspacesDatabase != nil {
		return workspacesDatabase
	}
	panic("Workspaces database not initialized. Seems, that InitDB() hasn't been executed or workspacesDatabase hasn't been initialized separately.")
}
