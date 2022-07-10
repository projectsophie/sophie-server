package pages

// Table is a struct which
// describes a table instance for database.
type Table struct {
	ID        int
	Title     string
	Workspace int
	Type      string
	Columns   []string
	Strings   []string
}
