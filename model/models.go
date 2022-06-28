package model

// User is a struct which
// describes a user instance for database.
type User struct {
	id            int    // id is "id" field in database.
	firstname     string // firstname is "firstname" field in database.
	lastname      string // lastname is "lastname" field in database.
	nickname      string // nickname is "nickname" field in database.
	password      string // password is "password" field in database. (Hashed)
	email         string // email is "email" field in database.
	emailVerified bool   // emailVerified is "emailVerified" field in database.
	workspaces    []int  // workspaces are "workspaces" field in database. (Contains ids of workspaces declared in workspaces database)
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
