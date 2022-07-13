package types

import "sophie-server/util"

// Table is a struct which
// describes a table instance for database.
type Table struct {
	ID        int
	Title     string
	Workspace int
	Type      string
	Columns   []map[string]interface{}
	Strings   [][]map[string]interface{}
}

// TableJSON is a struct which
// describes table while getting it via request.
type TableJSON struct {
	ID        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Workspace int    `json:"workspace" form:"workspace"`
	Type      string `json:"type" form:"type"`
	Metadata  string `json:"metadata" form:"metadata"`
}

// GetMetadata is an implementation of Page interface.
func (table *Table) GetMetadata() map[string]interface{} {
	return map[string]interface{}{
		"metadata": map[string]interface{}{
			"columns": table.Columns,
			"strings": table.Strings,
		},
	}
}

// GetAsJsonModel is an implementation of Page interface.
func (table *Table) GetAsJsonModel() TableJSON {
	return TableJSON{
		ID:        table.ID,
		Title:     table.Title,
		Workspace: table.Workspace,
		Type:      table.Type,
		Metadata:  util.ToJson(table.GetMetadata()),
	}
}
