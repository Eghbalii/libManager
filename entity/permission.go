package entity

type Permission struct {
	ID    uint
	Title PermissionTitle
}

type PermissionTitle string

const (
	BookUpdatePermission PermissionTitle = "book_update"
	BookDeletePermission PermissionTitle = "book_delete"
)
