package users

// User is a struct which
// describes a users instance for database.
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

// Session is a struct that represents a session.
type Session struct {
	ID            int    // ID is id of session.
	CreationDate  string // CreationDate is a creation date of session.
	LastUsageDate string // LastUsageDate is a last usage date of session.
	IP            string // IP is an IP-Address which was used for session creation.
	UserAgent     string // UserAgent is a user's agent which was used for session creation.
	AccessToken   string // AccessToken is an access token of session.
}
