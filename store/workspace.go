package store

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"sophie-server/database"
	"sophie-server/model/workspaces"
	"sophie-server/util"
	"time"
)

// CreateWorkspace creates a new workspace in database.
func CreateWorkspace(workspace *workspaces.WorkspaceCreate, member *[]workspaces.WorkspaceMember) gin.H {
	statement, _ := database.GetWorkspacesDB().Prepare(database.CreateWorkspace)
	_, err := statement.Exec(workspace.Title, time.Now().Format("2006-01-02 15:04:05"), util.ToJson(member))
	if err != nil {
		return gin.H{
			"error": "an error occurred while creating workspace",
		}
	}
	return gin.H{
		"message": "workspace created successfully",
	}
}

// GetMemberList returns an array of workspace members taken from database.
func GetMemberList(workspace int) []workspaces.WorkspaceMember {
	statement, _ := database.GetWorkspacesDB().Prepare(database.GetWorkspaceMembersById)
	var metadata string
	_ = statement.QueryRow(workspace).Scan(&metadata)
	var members []workspaces.WorkspaceMember
	_ = json.Unmarshal([]byte(metadata), &members)
	return members
}
