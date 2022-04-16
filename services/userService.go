package services

import (
	"fmt"

	"github.com/rd67/go-gin-mySQL/models"
	"github.com/rd67/go-gin-mySQL/utils"
)

func UserDetails(userId uint) (models.User, error) {
	var user models.User

	err := utils.DBConn.Model(models.User{}).Where(fmt.Sprintf("id = %d", userId)).First(&user).Error

	return user, err
}
