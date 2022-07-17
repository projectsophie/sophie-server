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

// CommentCreate is a struct which
// describes a comment instance while create request.
type CommentCreate struct {
	PageID int    `json:"pageID" form:"pageID" binding:"required"`
	Text   string `json:"text" form:"text" binding:"required"`
}

// CommentDelete is a struct which
// describes comment instance in delete request.
type CommentDelete struct {
	ID int `json:"id" form:"id" binding:"required"`
}
