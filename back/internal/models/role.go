package models

type Role struct {
	ID          int    `db:"id"`
	RoleName    string `db:"role_name"`
	Description string `db:"description"`
}

type RolePermission struct {
	ID             int    `db:"id"`
	RoleID         int    `db:"role_id"`
	PermissionName string `db:"permission_name"`
	CanView        bool   `db:"can_view"`
	CanEdit        bool   `db:"can_edit"`
	CanDelete      bool   `db:"can_delete"`
}
