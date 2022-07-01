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
	Workspaces    string // Workspaces are "workspaces" field in database. (Contains ids of workspaces declared in workspaces database)
}

// Workspace is a struct which
// describes a workspace instance for database.
type Workspace struct {
	id      int    // id is "id" field in database.
	members string // members are "members" field in database. (Contains json of members with their ids and permissions)
	pages   string // pages is "pages" field in database. (Contains json of pages with their metadata and types)
}

// ServerMessage is a struct which
// describes a server message instance.
// It is used for building server responses to API requests.
type ServerMessage struct {
	Title   string
	Message string
}
