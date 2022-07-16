package workspaces

import "sophie-server/database"

// Workspace is a struct which
// describes a workspace instance for database.
type Workspace struct {
	ID           int               // id is "id" field in database.
	Title        string            // title is "title" field in database.
	CreationDate string            // CreationDate is "creationDate" field in database.
	Members      []WorkspaceMember // members are "members" field in database. (Contains json of members with their ids and permissions)
	Pages        string            // pages is "pages" field in database. (Contains json of pages with their metadata and types)
}

// WorkspaceGet is a struct which
// describes a workspaces instance while getting.
type WorkspaceGet struct {
	ID           int    `json:"id" form:"id"`
	CreationDate string `json:"creationDate" form:"creationDate"`
	Members      string `json:"members" form:"members"`
	Pages        string `json:"pages" form:"pages"`
}

// WorkspaceCreate is a struct which
// describes a workspaces instance while creation.
type WorkspaceCreate struct {
	Title string `json:"title" form:"title" binding:"required"`
}

// WorkspaceMember is a struct which
// describes a workspace member.
type WorkspaceMember struct {
	ID          int    `json:"id" form:"id"`
	Permissions string `json:"permissions" form:"permissions"`
	JoinedAt    string `json:"joinedAt" form:"joinedAt"`
	LastSeen    string `json:"lastSeen" form:"lastSeen"`
}

// UpdateWorkspace updates provided Workspace
// instance in database.
func (workspace *Workspace) UpdateWorkspace() {
	statement, _ := database.GetWorkspacesDB().Prepare(database.UpdateWorkspace)
	_, _ = statement.Exec(workspace.Title, workspace.CreationDate, workspace.Members, workspace.Pages, workspace.ID)
}
