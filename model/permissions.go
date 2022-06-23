package model

// Permission is a special type for declaring permissions
// which are based on strings.
type Permission string

// Available permissions list
const (
	// ReadPermission allows user to view and read workspace materials,
	// but denies any editing action.
	ReadPermission Permission = "READ_PERMISSION"
	// EditPermission allows user to view, read and edit workspace materials,
	// but denies any admin action (like changing permissions for a user).
	EditPermission Permission = "EDIT_PERMISSION"
	// AdministratePermission allows user to do anything he wants
	// like editing, reading and managing members permissions.
	AdministratePermission Permission = "ADMIN_PERMISSION"
)
