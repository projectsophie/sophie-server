package users

import (
	"fmt"
	"sophie-server/database"
)

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

// UpdateUser updates users instance in database
// via provided users instance (*users.User).
func (user *User) UpdateUser() {
	statement, _ := database.GetUsersDB().Prepare(database.UpdateUser)
	_, err := statement.Exec(user.Firstname, user.Lastname, user.Nickname, user.Password, user.Email, user.EmailVerified, user.Sessions, user.Workspaces, user.ID)
	if err != nil {
		fmt.Println(err)
	}
}

// UserToUserGet converts User to UserGet
// via getting fields from User model.
func (user *User) UserToUserGet() UserGet {
	return UserGet{
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Nickname:   user.Nickname,
		Email:      user.Email,
		Workspaces: user.Workspaces,
	}
}
