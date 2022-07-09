package workspaces

// Workspace is a struct which
// describes a workspaces instance for database.
type Workspace struct {
	ID           int    // id is "id" field in database.
	Title        string // title is "title" field in database.
	CreationDate string // CreationDate is "creationDate" field in database.
	Members      string // members are "members" field in database. (Contains json of members with their ids and permissions)
	Pages        string // pages is "pages" field in database. (Contains json of pages with their metadata and types)
}

// Comment is a struct which
// describes a comment instance for database.
type Comment struct {
	ID           int    // id is "id" field in database.
	CreationDate string // CreationDate is "creationDate" field in database.
	Text         string // Text is "text" field in database.
	PageID       int    // PageID is "pageID" field in database.
	UserID       int    // UserID is "userID" field in database.
}
