package model

// UserCreate is a struct which
// describes a user instance while registration.
// Other fields like workspaces array are filled via SQL automatically.
type UserCreate struct {
	Firstname string `json:"firstname" form:"firstname" binding:"required"`
	Lastname  string `json:"lastname" form:"lastname" binding:"required"`
	Nickname  string `json:"nickname" form:"nickname" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
}
