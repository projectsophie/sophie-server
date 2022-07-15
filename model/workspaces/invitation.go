package workspaces

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

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
	ID          int `json:"id" form:"id"`
	WorkspaceID int `json:"workspaceId" form:"workspaceId"`
	Host        int `json:"host" form:"host"`
	jwt.StandardClaims
}

// InvitationCreate is a struct which
// describes fields that are used to generate an
// invitation for the database.
// As default, invitations can be created only by owner (OWNER_PERMISSION)
// or by administrator (ADMIN_PERMISSION).
type InvitationCreate struct {
	WorkspaceID    int    `json:"workspaceId" form:"workspaceId"`
	Host           int    `json:"host" form:"host"`
	ExpirationDate string `json:"expirationDate" form:"expirationDate"`
	MaxUsages      int    `json:"maxUsages" form:"maxUsages"`
}

// InvitationUse is a struct which
// describes an invitation when it used by a user.
type InvitationUse struct {
	InviteLink string `json:"inviteLink" form:"inviteLink"`
}

// Use is a function which
// uses an invitation and increases count of usages.
func (invitation *Invitation) Use() {
	invitation.Usages++
}

// IsAvailable is a function which
// checks if an invitation is available.
// It compares usages and maxUsages.
func (invitation *Invitation) IsAvailable() bool {
	exp, _ := time.Parse("2006-01-02 15:04:05", invitation.ExpirationDate)
	return invitation.Usages < invitation.MaxUsages && exp.Before(time.Now())
}
