package workspaces

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
