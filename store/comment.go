package store

import (
	"sophie-server/database"
	"sophie-server/model/pages"
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

func CreateComment(comment *pages.Comment) {
	statement, _ := database.GetWorkspacesDB().Prepare(database.CreateComment)
	_, _ = statement.Query(comment.AuthorID, comment.PageID, comment.Text)
}
