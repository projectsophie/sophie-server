package pages

// Comment is a struct which
// describes a comment instance for database.
type Comment struct {
	ID           int    // id is "id" field in database.
	CreationDate string // CreationDate is "creationDate" field in database.
	Text         string // Text is "text" field in database.
	PageID       int    // PageID is "pageID" field in database.
	AuthorID     int    // AuthorID is "authorID" field in database.
}
