package models

import (
	"time"

	"gorm.io/gorm"
)

type TeamUserStep struct {
	ID uint `json:"id" gorm:"primarykey"`

	TeamUserId uint      `json:"teamUserId" gorm:"not null;index:idx_team_user"`
	TeamUser   *TeamUser `json:"teamUser"`

	Steps uint `json:"steps" gorm:"default:0;"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
