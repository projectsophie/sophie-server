package store

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"sophie-server/database"
	"sophie-server/model/users"
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
	return GetWorkspaceMembersFromJson(metadata)
}

// GetWorkspaceMembersFromJson parses given string
// from database and converts it to []workspaces.Workspace.
func GetWorkspaceMembersFromJson(metadata string) []workspaces.WorkspaceMember {
	var members []workspaces.WorkspaceMember
	_ = json.Unmarshal([]byte(metadata), &members)
	return members
}

// GetWorkspaceById returns workspaces.Workspace
// model from database with given id.
func GetWorkspaceById(id int) workspaces.Workspace {
	statement, _ := database.GetWorkspacesDB().Prepare(database.GetWorkspaceById)
	row := statement.QueryRow(id)
	var workspaceModel workspaces.Workspace
	var __users string
	err := row.Scan(
		&workspaceModel.ID,
		&workspaceModel.Title,
		&workspaceModel.CreationDate,
		&__users,
		&workspaceModel.Pages)
	if err == nil {
		workspaceModel.Members = GetWorkspaceMembersFromJson(__users)
	} else {
		return workspaces.Workspace{}
	}
	return workspaceModel
}

// AddMember adds provided user model as workspaces.WorkspaceMember
// to workspaces.Workspace with provided id and updates
// it in database.
func AddMember(user *users.User, workspace int) {
	workspaceModel := GetWorkspaceById(workspace)
	workspaceModel.Members = append(
		workspaceModel.Members,
		user.AsWorkspaceMember())
	workspaceModel.UpdateWorkspace()
}

// GetWorkspaceByPage returns workspaces.Workspace
// model from database with given page id.
func GetWorkspaceByPage(page int) workspaces.Workspace {
	statement, _ := database.GetWorkspacesDB().Prepare(database.GetWorkspaceByPageId)
	var workspaceID int
	row := statement.QueryRow(page)
	err := row.Scan(&workspaceID)
	if err != nil {
		return workspaces.Workspace{}
	}
	return GetWorkspaceById(workspaceID)
}

// HasAccessToPage returns true if user has access to workspace.
func HasAccessToPage(user *users.User, page int) bool {
	workspace := GetWorkspaceByPage(page)
	fmt.Println(workspace)
	for _, member := range workspace.Members {
		if member.ID == user.ID {
			return true
		}
	}
	return false
}
