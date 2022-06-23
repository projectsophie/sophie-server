package database

// Constant SQL statements
const (
	CREATE_USER = "INSERT INTO users (firstname, lastname, nickname, password, email) VALUES (?, ?, ?, ?, ?)"
	DELETE_USER = ""
)
