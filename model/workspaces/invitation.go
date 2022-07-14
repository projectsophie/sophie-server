package workspaces

import "github.com/dgrijalva/jwt-go"

// Invitation is a struct which
// describes a workspace invite.
type Invitation struct {
	ID             int    `json:"id" form:"id"`
	WorkspaceID    int    `json:"workspaceId" form:"workspaceId"`
	Host           int    `json:"host" form:"host"`
	ExpirationDate string `json:"expirationDate" form:"expirationDate"`
	Usages         int    `json:"usages" form:"usages"`
	MaxUsages      int    `json:"maxUsages" form:"maxUsages"`
}

// InvitationGenerate is a struct which
// describes fields that are used to generate a
// JWT-key for invite link.
type InvitationGenerate struct {
	WorkspaceID int `json:"workspaceId" form:"workspaceId"`
	Host        int `json:"host" form:"host"`
	jwt.StandardClaims
}

// InvitationUse is a struct which
// describes an invitation when it used by a user.
type InvitationUse struct {
	InviteLink string `json:"inviteLink" form:"inviteLink"`
}
