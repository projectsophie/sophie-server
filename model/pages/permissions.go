package pages

// Permission is a special type for declaring permissions
// which are based on strings.
type Permission string

// Available permissions list
const (
	// ReadPermission allows users to view and read workspaces materials,
	// but denies any editing action.
	ReadPermission Permission = "READ_PERMISSION"
	// EditPermission allows users to view, read and edit workspaces materials,
	// but denies any admin action (like changing permissions for a users).
	EditPermission Permission = "EDIT_PERMISSION"
	// AdministratePermission allows users to do anything he wants
	// like editing, reading and managing members permissions.
	AdministratePermission Permission = "ADMIN_PERMISSION"
)
