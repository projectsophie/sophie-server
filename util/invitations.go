package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/database"
	"sophie-server/model/workspaces"
	globals "sophie-server/traits"
)

// GenerateClaims generates claims for JWT invitation.
func GenerateClaims(workspace int, host int) workspaces.InvitationGenerate {
	return workspaces.InvitationGenerate{
		WorkspaceID:    workspace,
		Host:           host,
		StandardClaims: jwt.StandardClaims{},
	}
}

// GenerateLink generates a link for invitation.
func GenerateLink(invite *workspaces.InvitationGenerate) string {
	rawInvite := jwt.NewWithClaims(jwt.SigningMethodHS256, invite)
	inviteLink, err := rawInvite.SignedString([]byte(GetSecretKey()))
	if err != nil {
		return ""
	}
	return globals.HostUrl + "invitation?code=" + inviteLink
}

// UseLink executes when user uses invite link.
// It increases the number of usages and updates
// invitation in database.
func UseLink(c *gin.Context) {
	rawInvitation, _ := c.GetQuery("code")
	if _, err := ValidateToken(rawInvitation); err == nil {
		metadata, err := ParseToken(rawInvitation)
		if err == nil {
			invitation := GetInvitationById(metadata["id"].(int))
			if invitation != nil {
				if invitation.IsAvailable() {
					invitation.Use()
					UpdateInvitation(invitation)
					c.JSON(http.StatusOK, gin.H{"message": "invite link was successfully used"})
				}
			}
		}
	}
	c.JSON(http.StatusForbidden, gin.H{"error": "invalid invite link"})
}

func UpdateInvitation(invitation *workspaces.Invitation) {
	statement, _ := database.GetWorkspacesDB().Prepare(database.UpdateInvitation)
	_, _ = statement.Exec(invitation.ExpirationDate, invitation.Usages, invitation.MaxUsages)
}

func GetInvitationById(id int) *workspaces.Invitation {
	return &workspaces.Invitation{}
}
