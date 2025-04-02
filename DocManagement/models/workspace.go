package models

import "time"

type Workspace struct {
	Id        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"not null; column:name" json:"name"`
	OwnerId   uint      `gorm:"column:owner_id" json:"owner_id"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
}
