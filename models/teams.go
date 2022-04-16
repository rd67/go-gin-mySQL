package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	// gorm.Model

	ID uint `json:"id" gorm:"primarykey, AUTO_INCREMENT"`

	Name  string `json:"name" gorm:"size:50;not null;index:idx_name_deleted;"`
	Steps uint   `json:"steps" gorm:"default:0;"`

	Users []*User `json:"users" gorm:"many2many:team_users;joinForeignKey:TeamId;"`

	TeamUsers []*TeamUser `json:"teamUsers" gorm:"joinForeignKey:team_id;"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index:idx_name_deleted;"`
}
