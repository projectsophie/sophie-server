package database

// Constant SQL statements
const (
	USERS_PREPARE      = "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT NOT NULL, lastname TEXT NOT NULL, nickname TEXT NOT NULL UNIQUE, password TEXT NOT NULL, email TEXT NOT NULL UNIQUE, register_date TEXT NOT NULL, email_verified BOOLEAN DEFAULT FALSE, sessions TEXT NOT NULL DEFAULT '[]', workspaces INTEGER [] DEFAULT []);"
	WORKSPACES_PREPARE = "CREATE TABLE IF NOT EXISTS workspaces (id INTEGER PRIMARY KEY, title TEXT NOT NULL, creation_date TEXT NOT NULL, members TEXT, pages TEXT [] DEFAULT []);"

	CREATE_USER             = "INSERT INTO users (firstname, lastname, nickname, password, email, register_date) VALUES (?, ?, ?, ?, ?, ?)"
	DELETE_USER_BY_ID       = "DELETE FROM users WHERE id = @id"
	DELETE_USER_BY_NICKNAME = "DELETE FROM users WHERE nickname = @nickname"
	GET_USER_BY_ID          = "SELECT * FROM users WHERE id = @id"
	GET_USER_BY_NICKNAME    = "SELECT * FROM users WHERE nickname = @nickname"
	GET_USER_BY_EMAIL       = "SELECT * FROM users WHERE email = @email"
	UPDATE_USER             = "UPDATE users SET firstname = @firstname, lastname = @lastname, nickname = @nickname, password = @password, email = @email, email_verified = @email_verified, sessions = @sessions, workspaces = @workspaces WHERE id = @id"

	CREATE_WORKSPACE       = "INSERT INTO workspaces (title) VALUES (?)"
	UPDATE_WORKSPACE       = "UPDATE workspaces SET title = @title, creation_date = @creation_date, members = @members, pages = @pages WHERE id = @id"
	DELETE_WORKSPACE_BY_ID = "DELETE FROM workspaces WHERE id = @id"
)
