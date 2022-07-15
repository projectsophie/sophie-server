package database

// Constant SQL statements
const (
	UsersPrepare       = "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT NOT NULL, lastname TEXT NOT NULL, nickname TEXT NOT NULL UNIQUE, password TEXT NOT NULL, email TEXT NOT NULL UNIQUE, register_date TEXT NOT NULL, email_verified BOOLEAN DEFAULT FALSE, sessions TEXT NOT NULL DEFAULT '[]', workspaces INTEGER [] DEFAULT []);"
	WorkspacesPrepare  = "CREATE TABLE IF NOT EXISTS workspaces (id INTEGER PRIMARY KEY, title TEXT NOT NULL, creation_date TEXT NOT NULL, members TEXT, pages TEXT DEFAULT []);"
	CommentsPrepare    = "CREATE TABLE IF NOT EXISTS comments (id INTEGER PRIMARY KEY, author INTEGER NOT NULL, page INTEGER NOT NULL, content TEXT NOT NULL, creation_date TEXT NOT NULL);"
	MediaPrepare       = "CREATE TABLE IF NOT EXISTS media (id INTEGER PRIMARY KEY, page INTEGER NOT NULL, workspace INTEGER NOT NULL, type TEXT NOT NULL, content TEXT NOT NULL);"
	EtcPrepare         = "CREATE TABLE IF NOT EXISTS etc (id INTEGER PRIMARY KEY, workspace INTEGER NOT NULL, type TEXT NOT NULL, metadata TEXT NOT NULL);"
	InvitationsPrepare = "CREATE TABLE IF NOT EXISTS invitations (id INTEGER PRIMARY KEY, workspace INTEGER NOT NULL, host INTEGER NOT NULL, expiration_date TEXT NOT NULL, usages INTEGER NOT NULL DEFAULT 0, max_usages INTEGER NOT NULL DEFAULT 1);"

	CreateUser           = "INSERT INTO users (firstname, lastname, nickname, password, email, register_date) VALUES (?, ?, ?, ?, ?, ?)"
	DeleteUserById       = "DELETE FROM users WHERE id = @id"
	DeleteUserByNickname = "DELETE FROM users WHERE nickname = @nickname"
	GetUserById          = "SELECT * FROM users WHERE id = @id"
	GetUserByNickname    = "SELECT * FROM users WHERE nickname = @nickname"
	GetUserByEmail       = "SELECT * FROM users WHERE email = @email"
	UpdateUser           = "UPDATE users SET firstname = @firstname, lastname = @lastname, nickname = @nickname, password = @password, email = @email, email_verified = @email_verified, sessions = @sessions, workspaces = @workspaces WHERE id = @id"

	CreateWorkspace         = "INSERT INTO workspaces (title, creation_date, members) VALUES (?, ?, ?)"
	UpdateWorkspace         = "UPDATE workspaces SET title = @title, creation_date = @creation_date, members = @members, pages = @pages WHERE id = @id"
	DeleteWorkspaceById     = "DELETE FROM workspaces WHERE id = @id"
	GetWorkspaceMembersById = "SELECT members FROM workspaces WHERE id = @id"

	GetCommentsByPage = "SELECT * FROM comments WHERE page = @page"
	DeleteCommentById = "DELETE FROM comments WHERE id = @id"
	CreateComment     = "INSERT INTO comments (author, page, content, creation_date) VALUES (?, ?, ?, ?)"

	CreateInvitation  = "INSERT INTO invitations (workspace, host, expiration_date, usages, max_usages) VALUES (?, ?, ?, ?, ?)"
	GetInvitationById = "SELECT * FROM invitations WHERE id = @id"
	UpdateInvitation  = "UPDATE invitations SET expiration_date = @expiration_date, usages = @usages, max_usages = @max_usages WHERE id = @id"
)
