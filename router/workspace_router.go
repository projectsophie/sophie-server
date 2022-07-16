package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/model/pages"
	"sophie-server/model/workspaces"
	"sophie-server/store"
	"sophie-server/util"
	"time"
)

func CreateWorkspace(c *gin.Context) {
	user, success := store.GetUserByToken(util.GetToken(c))
	if success {
		member := workspaces.WorkspaceMember{
			ID:          user.ID,
			Permissions: string(pages.OwnerPermission),
			JoinedAt:    time.Now().Format("2006-01-02 15:04:05"),
			LastSeen:    "NA",
		}
		var workspaceModel workspaces.WorkspaceCreate
		if err := c.BindJSON(&workspaceModel); err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			c.IndentedJSON(http.StatusOK, store.CreateWorkspace(&workspaceModel, &[]workspaces.WorkspaceMember{member}))
		}
	} else {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
	}
}

func UseInvitation(c *gin.Context) {
	store.UseInvitation(c)
}
