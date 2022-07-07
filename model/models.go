package model

// User is a struct which
// describes a user instance for database.
type User struct {
	ID            int    // ID is "id" field in database.
	Firstname     string // Firstname is "firstname" field in database.
	Lastname      string // Lastname is "lastname" field in database.
	Nickname      string // Nickname is "nickname" field in database.
	Password      string // Password is "password" field in database. (Hashed)
	Email         string // Email is "email" field in database.
	EmailVerified bool   // EmailVerified is "emailVerified" field in database.
	RegisterDate  string // RegisterDate is "registerDate" field in database.
	Sessions      string // Sessions is "sessions" field in database.
	Workspaces    string // Workspaces are "workspaces" field in database. (Contains ids of workspaces declared in workspaces database)
}

// Workspace is a struct which
// describes a workspace instance for database.
type Workspace struct {
	ID           int    // id is "id" field in database.
	CreationDate string // CreationDate is "creationDate" field in database.
	Members      string // members are "members" field in database. (Contains json of members with their ids and permissions)
	Pages        string // pages is "pages" field in database. (Contains json of pages with their metadata and types)
}
