package services

import (
	"fmt"

	"github.com/rd67/go-gin-mySQL/models"
	"github.com/rd67/go-gin-mySQL/utils"
	"gorm.io/gorm"
)

func TeamDetails(teamId uint) (models.Team, error) {
	var team models.Team

	err := utils.DBConn.Table("teams").Preload("TeamUsers").Preload("TeamUsers.User").Where(fmt.Sprintf("id = %d", teamId)).First(&team).Error

	return team, err
}

type Store struct {
	db *gorm.DB
}

//  Handles Incrementing Team Steps, Adding Steps History
func TeamStepsAddEffect(tx *gorm.DB, teamId uint, teamUserId uint, steps uint) error {

	//  Adding TeamUserStep Log
	if err := tx.Create(&models.TeamUserStep{
		TeamUserId: teamUserId,
		Steps:      steps,
	}).Error; err != nil {
		return err
	}

	//  Increment total in Teams Model
	if err := tx.Model(models.Team{}).Where(fmt.Sprintf("id = %d", teamId)).UpdateColumn("steps", gorm.Expr("steps + ?", steps)).Error; err != nil {
		return err
	}

	return nil
}
