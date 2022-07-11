package workspaces

// Workspace is a struct which
// describes a workspace instance for database.
type Workspace struct {
	ID           int    // id is "id" field in database.
	Title        string // title is "title" field in database.
	CreationDate string // CreationDate is "creationDate" field in database.
	Members      string // members are "members" field in database. (Contains json of members with their ids and permissions)
	Pages        string // pages is "pages" field in database. (Contains json of pages with their metadata and types)
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
