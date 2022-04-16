package modules

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rd67/go-gin-mySQL/configs"
	"github.com/rd67/go-gin-mySQL/helpers"
	"github.com/rd67/go-gin-mySQL/models"
	"github.com/rd67/go-gin-mySQL/services"
	"github.com/rd67/go-gin-mySQL/utils"
	"gorm.io/gorm"
)

/////////////////////
// Team User Add
/////////////////////
func TeamUsersAdd(c *gin.Context) {

	var data = struct {
		TeamId uint `form:"teamId" binding:"required,gt=0"`
		UserId uint `form:"userId" binding:"required,gt=0"`
		Steps  uint `form:"steps" binding:"gte=0"`
	}{
		Steps: 0,
	}

	//	Validation
	if err := c.ShouldBind(&data); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	//  Checking the validness of User
	var user models.User
	if err := utils.DBConn.Model(models.User{}).Where(fmt.Sprintf("id = '%d'", data.UserId)).First(&user).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	// // //  Checking the validness of User
	var team models.Team
	if err := utils.DBConn.Where(fmt.Sprintf("id = %d", data.TeamId)).Preload("TeamUsers", "user_id = ?", data.UserId).First(&team).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	if len(team.TeamUsers) > 0 {
		helpers.ActionFailedErrorResponse(c, "Sorry, this user is already added to this team")
		return
	}

	err := utils.DBConn.Transaction(func(tx *gorm.DB) error {

		//	Creating Team User
		var teamUser = models.TeamUser{UserId: data.UserId, TeamId: data.TeamId, Steps: data.Steps}
		if err := tx.Model(models.TeamUser{}).Create(&teamUser).Error; err != nil {
			return err
		}

		if data.Steps == 0 {
			return nil
		}

		//If steps are there than Updating Team, Adding Steps History
		err := services.TeamStepsAddEffect(tx, data.TeamId, teamUser.ID, data.Steps)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	//	Team Details
	teamDetails, err := services.TeamDetails(data.TeamId)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	response := TeamDetailsResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusCreated,
			Message:    "Team user added successfully",
		},
		Data: TeamDetailsDataResponseStruct{
			Team: teamDetails,
		},
	}

	c.JSON(response.StatusCode, response)

	return
}

//	TeamUser Steps Add
func TeamUserStepsAdd(c *gin.Context) {

	var queryData struct {
		TeamUserId uint `uri:"teamUserId" binding:"required,gt=0"`
	}
	if err := c.BindUri(&queryData); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	var data struct {
		Steps uint `form:"steps" binding:"required,gt=0"`
	}
	if err := c.Bind(&data); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	var teamUser models.TeamUser
	err := utils.DBConn.Transaction(func(tx *gorm.DB) error {

		//.UpdateColumn("steps", gorm.Expr("steps + ?", steps))
		if err := tx.Model(models.TeamUser{}).Where(fmt.Sprintf("id = %d", queryData.TeamUserId)).First(&teamUser).UpdateColumn("steps", gorm.Expr("steps + ?", data.Steps)).Error; err != nil {
			return err
		}

		if err := services.TeamStepsAddEffect(tx, teamUser.TeamId, teamUser.ID, data.Steps); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	// var teamDetails models.Team
	teamDetails, err := services.TeamDetails(teamUser.TeamId)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	response := TeamDetailsResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusOK,
			Message:    "User Steps added successfully",
		},
		Data: TeamDetailsDataResponseStruct{
			Team: teamDetails,
		},
	}

	c.JSON(response.StatusCode, response)
	return

}
