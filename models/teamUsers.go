package models

import (
	"time"

	"gorm.io/gorm"
)

type TeamUser struct {
	// gorm.Model

	ID uint `json:"id" gorm:"primarykey"`

	TeamId uint  `json:"teamId" gorm:"not null;index:idx_team_user_deleted;"`
	Team   *Team `json:"team" gorm:"foreignKey:team_id"`

	UserId uint  `json:"userId" gorm:"not null;index:idx_team_user_deleted;"`
	User   *User `json:"user" gorm:"foreignKey:user_id"`

	Steps uint `json:"steps" gorm:"default:0;"`

	CreatedAt time.Time      `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index:idx_team_user_deleted;"`

	TeamUserSteps []TeamUserStep `gorm:"foreignKey:team_user_id"`
}

// type Tabler interface {
// 	TableName() string
// }

// func (TeamUser) TableName() string {
// 	return "teamUsers"
// }
