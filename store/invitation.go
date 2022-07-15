package store

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/database"
	"sophie-server/model/workspaces"
	globals "sophie-server/traits"
	"sophie-server/util"
)

// GenerateClaims generates claims for JWT invitation.
func GenerateClaims(create *workspaces.InvitationCreate) workspaces.InvitationGenerate {
	return workspaces.InvitationGenerate{
		WorkspaceID:    create.WorkspaceID,
		Host:           create.Host,
		StandardClaims: jwt.StandardClaims{},
	}
}

func CreateInvitationLink(generate *workspaces.InvitationGenerate) string {
	rawInvite := jwt.NewWithClaims(jwt.SigningMethodHS256, generate)
	inviteLink, err := rawInvite.SignedString([]byte(util.GetSecretKey()))
	if err != nil {
		return ""
	}
	return globals.HostUrl + "invitation?code=" + inviteLink
}

// UseInvitation executes when user uses invite link.
// It increases the number of usages and updates
// invitation in database.
func UseInvitation(c *gin.Context) {
	rawInvitation, _ := c.GetQuery("code")
	if _, err := util.ValidateToken(rawInvitation); err == nil {
		metadata, err := util.ParseToken(rawInvitation)
		if err == nil {
			invitation, success := GetInvitationById(metadata["id"].(int))
			if success {
				if invitation.IsAvailable() {
					invitation.Use()
					UpdateInvitation(&invitation)
					c.JSON(http.StatusOK, gin.H{"message": "invite link was successfully used"})
				}
			}
		}
	}
	c.JSON(http.StatusForbidden, gin.H{"error": "invalid invite link"})
}

func GetInvitationById(id int) (workspaces.Invitation, bool) {
	statement, _ := database.GetWorkspacesDB().Prepare(database.GetInvitationById)
	row, _ := statement.Query(id)
	for row.Next() {
		invitation := workspaces.Invitation{}
		_ = row.Scan(&invitation.ID, &invitation.WorkspaceID, &invitation.Host, &invitation.ExpirationDate, &invitation.Usages, &invitation.MaxUsages)
		return invitation, true
	}
	return workspaces.Invitation{}, false
}

// UpdateInvitation updates an invitation in database.
func UpdateInvitation(invitation *workspaces.Invitation) {
	statement, _ := database.GetWorkspacesDB().Prepare(database.UpdateInvitation)
	_, _ = statement.Exec(invitation.ExpirationDate, invitation.Usages, invitation.MaxUsages)
}
