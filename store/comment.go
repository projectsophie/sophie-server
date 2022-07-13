package store

import (
	"sophie-server/database"
	"sophie-server/model/pages"
	"time"
)

// GetCommentsByPage returns all comments from provided page id.
func GetCommentsByPage(page int) []pages.Comment {
	statement, _ := database.GetWorkspacesDB().Prepare(database.GetCommentsByPage)
	rows, _ := statement.Query(page)
	var array []pages.Comment
	for rows.Next() {
		var comment pages.Comment
		err := rows.Scan(&comment.ID, &comment.AuthorID, &comment.PageID, &comment.Text, &comment.CreationDate)
		if err == nil {
			array = append(array, comment)
		}
	}
	return array
}

// DeleteCommentById deletes comment from provided id.
func DeleteCommentById(id int) {
	statement, _ := database.GetWorkspacesDB().Prepare(database.DeleteCommentById)
	_, _ = statement.Query(id)
}

// CreateComment creates new comment in database.
func CreateComment(comment *pages.Comment) {
	// TODO: Validate comment page if user has access to it.
	statement, _ := database.GetWorkspacesDB().Prepare(database.CreateComment)
	_, _ = statement.Query(comment.AuthorID, comment.PageID, comment.Text, time.Now().Format("2006-01-02 15:04:05"))
}
