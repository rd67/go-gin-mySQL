package modules

import (
	"fmt"
	"net/http"

	"github.com/rd67/go-gin-mySQL/configs"
	"github.com/rd67/go-gin-mySQL/helpers"
	"github.com/rd67/go-gin-mySQL/models"
	"github.com/rd67/go-gin-mySQL/services"
	"github.com/rd67/go-gin-mySQL/utils"

	"github.com/gin-gonic/gin"
)

/////////////////////
//	Teams List
/////////////////////
type TeamListDataResponseStruct struct {
	Count int64         `json:"count"`
	Rows  []models.Team `json:"rows"`
}
type TeamsListResponseStruct struct {
	configs.CommonResponseStruct
	Data TeamListDataResponseStruct `json:"data"`
}

func TeamsList(c *gin.Context) {

	var data = struct {
		Limit  int    `form:"limit" binding:"min=10,max=100"`
		Offset int    `form:"offset" binding:"min=0"`
		Search string `form:"search"`
	}{
		Limit:  configs.DefaultPageSize,
		Offset: 0,
		Search: "",
	}

	//	Validation
	if err := c.ShouldBindQuery(&data); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	baseQuery := utils.DBConn.Model(models.Team{})

	//	If Searching
	if data.Search != "" {
		baseQuery.Where("name LIKE ?", "%"+data.Search+"%")
	}

	var count int64
	var rows []models.Team

	//	Getting Count of Rows
	if err := baseQuery.Count(&count).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	//	Getting Rows of table
	if err := baseQuery.Limit(data.Limit).Offset(data.Offset).Preload("TeamUsers").Preload("TeamUsers.User").Find(&rows).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	response := TeamsListResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusOK,
			Message:    "Teams listing",
		},
		Data: TeamListDataResponseStruct{
			Count: count,
			Rows:  rows,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}

/////////////////////
//	Team Create
/////////////////////
type TeamCreateDataResponseStruct struct {
	Team models.Team `json:"team"`
}
type TeamCreateResponseStruct struct {
	configs.CommonResponseStruct
	Data TeamCreateDataResponseStruct `json:"data"`
}

func TeamCreate(c *gin.Context) {

	var data struct {
		Name string `form:"name" binding:"required,min=3,max=50"`
	}

	//	Validation
	if err := c.ShouldBind(&data); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	var teamNameCheck int64
	utils.DBConn.Model(models.Team{}).Where(fmt.Sprintf("name = '%s'", data.Name)).Count(&teamNameCheck)
	if teamNameCheck > 0 {
		helpers.ActionFailedErrorResponse(c, "Sorry, this team name is already taken")
		return
	}

	var team = models.Team{
		Name: data.Name,
		// Users: []*models.User{
		// 	{Name: "Test User", Username: "TestUser"},
		// },
	}

	if err := utils.DBConn.Create(&team).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	teamDetails, _ := services.TeamDetails(team.ID)

	//	Creating Response
	response := TeamCreateResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusCreated,
			Message:    "Team created successfully",
		},
		Data: TeamCreateDataResponseStruct{
			Team: teamDetails,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}

/////////////////////
//	Team Create
/////////////////////
type TeamDetailsDataResponseStruct struct {
	Team models.Team `json:"team"`
}

type TeamDetailsResponseStruct struct {
	configs.CommonResponseStruct
	Data TeamDetailsDataResponseStruct `json:"data"`
}

func TeamDetails(c *gin.Context) {

	var data struct {
		TeamId uint64 `uri:"teamId" binding:"required"`
	}

	//	Validation
	if err := c.ShouldBindUri(&data); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	//	Db Query
	teamDetails, err := services.TeamDetails(uint(data.TeamId))
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	//	Creating Response
	response := TeamDetailsResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusOK,
			Message:    "Team Details",
		},
		Data: TeamDetailsDataResponseStruct{
			Team: teamDetails,
		},
	}

	c.JSON(response.StatusCode, response)
	return

}
