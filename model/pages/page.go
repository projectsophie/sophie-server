package pages

// Page is an interface which
// describes basic methods of various pages.
type Page interface {
	GetMetadata() string         // GetMetadata returns a JSON string of metadata of page.
	GetAsJsonModel() interface{} // GetAsJsonModel returns a JSON string of page.
}
