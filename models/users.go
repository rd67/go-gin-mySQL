package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model

	ID uint `json:"id" gorm:"primarykey"`

	Name     string `json:"name" gorm:"size:50;not null;"`
	Username string `json:"username" gorm:"size:50;not null;index:idx_username_deleted;"`

	//https://gorm.io/docs/many_to_many.html
	Teams []*Team `json:"teams" gorm:"many2many:team_users;joinForeignKey:UserId;"`

	TeamUsers []*TeamUser `json:"teamUsers" gorm:"hasMany:team_users;joinForeignKey:UserId;"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index:idx_username_deleted;"`
}
