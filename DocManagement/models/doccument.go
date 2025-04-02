package models

import "time"

type Document struct {
	ID         uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name       string    `json:"name" gorm:"column:name"`
	FilePath   string    `json:"file_path" gorm:"column:file_path"`
	UploadedBy uint      `json:"uploaded_by" gorm:"column:uploaded_by"`
	Workspace  uint      `json:"workspace" gorm:"column:workspace_id"`
	UploadedAt time.Time `json:"updated_at" gorm:"column:uploaded_at"`
}
