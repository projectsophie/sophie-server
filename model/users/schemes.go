package users

// UserCreate is a struct which
// describes a users instance while registration.
// Other fields like workspaces array are filled via SQL automatically.
type UserCreate struct {
	Firstname string `json:"firstname" form:"firstname" binding:"required"`
	Lastname  string `json:"lastname" form:"lastname" binding:"required"`
	Nickname  string `json:"nickname" form:"nickname" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
}

// UserAuth is a struct which
// describes a users instance while authentication.
type UserAuth struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// UserGet is a struct which
// describes a users instance while getting.
type UserGet struct {
	Firstname  string `json:"firstname" form:"firstname"`
	Lastname   string `json:"lastname" form:"lastname"`
	Nickname   string `json:"nickname" form:"nickname"`
	Email      string `json:"email" form:"email"`
	Workspaces string `json:"workspaces" form:"workspaces"`
}

// SessionCreate is a struct which
// describes a session instance while creation.
type SessionCreate struct {
	IP          string `json:"ip" form:"ip"`
	AccessToken string `json:"accessToken" form:"accessToken"`
	UserAgent   string `json:"userAgent" form:"userAgent"`
}
