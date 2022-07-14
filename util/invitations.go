package util

import (
	"github.com/dgrijalva/jwt-go"
	"sophie-server/model/workspaces"
)

// GenerateClaims generates claims for JWT invitation.
func GenerateClaims(workspace int, host int) workspaces.InvitationGenerate {
	return workspaces.InvitationGenerate{
		WorkspaceID:    workspace,
		Host:           host,
		StandardClaims: jwt.StandardClaims{},
	}
}

func GenerateLink(invite *workspaces.InvitationGenerate) string {
	rawInvite := jwt.NewWithClaims(jwt.SigningMethodHS256, invite)
	inviteLink, err := rawInvite.SignedString([]byte(GetSecretKey()))
	if err != nil {
		return ""
	}
	return inviteLink
}

func UseLink(link string) {

}
