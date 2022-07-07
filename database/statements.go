package database

// Constant SQL statements
const (
	USERS_PREPARE      = "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT NOT NULL, lastname TEXT NOT NULL, nickname TEXT NOT NULL UNIQUE, password TEXT NOT NULL, email TEXT NOT NULL UNIQUE, email_verified BOOLEAN DEFAULT FALSE, register_date TEXT NOT NULL, sessions TEXT, workspaces INTEGER [] DEFAULT []);"
	WORKSPACES_PREPARE = "CREATE TABLE IF NOT EXISTS workspaces (id INTEGER PRIMARY KEY, creation_date TEXT NOT NULL, members TEXT, pages TEXT [] DEFAULT []);"

	CREATE_USER             = "INSERT INTO users (firstname, lastname, nickname, password, email) VALUES (?, ?, ?, ?, ?)"
	DELETE_USER_BY_ID       = "DELETE FROM users WHERE id = @id"
	DELETE_USER_BY_NICKNAME = "DELETE FROM users WHERE nickname = @nickname"
	GET_USER_BY_ID          = "SELECT * FROM users WHERE id = @id"
	GET_USER_BY_NICKNAME    = "SELECT * FROM users WHERE nickname = @nickname"
	GET_USER_BY_EMAIL       = "SELECT * FROM users WHERE email = @email"
	APPLY_SESSION_TO_USER   = "UPDATE users SET sessions = @sessions WHERE id = @id"

	CREATE_WORKSPACE       = "INSERT INTO workspaces (owner) VALUES (?)"
	DELETE_WORKSPACE_BY_ID = "DELETE FROM workspaces WHERE id = @id"
)
