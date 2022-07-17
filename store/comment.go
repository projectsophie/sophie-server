package store

import (
	"sophie-server/database"
	"sophie-server/model/pages"
	"sophie-server/model/users"
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
func DeleteCommentById(comment *pages.CommentDelete, user *users.User) bool {
	model := GetCommentById(comment.ID)
	if HasAccessToPage(user, model.PageID) {
		statement, _ := database.GetWorkspacesDB().Prepare(database.DeleteCommentById)
		_, _ = statement.Query(comment.ID)
		return true
	}
	return false
}

// GetCommentById deletes pages.Comment with provided id.
func GetCommentById(id int) pages.Comment {
	statement, _ := database.GetWorkspacesDB().Prepare(database.GetCommentById)
	row := statement.QueryRow(id)
	var comment pages.Comment
	err := row.Scan(&comment.ID, &comment.CreationDate, &comment.Text, &comment.PageID, &comment.AuthorID)
	if err != nil {
		return pages.Comment{}
	}
	return comment
}

// CreateComment creates new comment in database.
func CreateComment(comment *pages.CommentCreate, user int) {
	// TODO: Validate comment page if user has access to it.
	statement, _ := database.GetWorkspacesDB().Prepare(database.CreateComment)
	_, _ = statement.Query(user, comment.PageID, comment.Text, time.Now().Format("2006-01-02 15:04:05"))
}
