package models

type Permission string

const (
	Read Permission = "read"
	Edit Permission = "edit"
	Del  Permission = "delete"
)
