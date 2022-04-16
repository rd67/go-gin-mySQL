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
	// "github.com/go-playground/validator/v10"
)

/////////////////////
//	Users Listing
/////////////////////
type UsersListResponseDataStruct struct {
	Count int64         `json:"count"`
	Rows  []models.User `json:"rows"`
}
type UsersListResponseStruc struct {
	configs.CommonResponseStruct
	Data UsersListResponseDataStruct `json:"data"`
}

func UsersList(c *gin.Context) {

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

	baseQuery := utils.DBConn.Model(models.User{})

	//	If seaching
	if data.Search != "" {
		baseQuery.Where("name LIKE ?", "%"+data.Search+"%")
	}

	var count int64
	rows := []models.User{}

	//DB Queryies
	baseQuery.Count(&count)
	baseQuery.Limit(data.Limit).Offset(data.Offset).Find(&rows)

	//	Generating common Response parameters like statusCode, message, data
	response := UsersListResponseStruc{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusOK,
			Message:    "Success",
		},
		Data: UsersListResponseDataStruct{
			Count: count,
			Rows:  rows,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}

/////////////////////
//	User Create
/////////////////////
type UserCreateResponseDataStruct struct {
	User models.User `json:"user"`
}
type UserCreateResponseStruc struct {
	configs.CommonResponseStruct
	Data UserCreateResponseDataStruct `json:"data"`
}

func UserCreate(c *gin.Context) {

	var data struct {
		Name     string `form:"name" binding:"required,min=3,max=50"`
		Username string `form:"username" binding:"required,min=3,max=50"`
	}

	if err := c.ShouldBind(&data); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	//	User Name Check
	var usernameCheck int64
	utils.DBConn.Model(models.User{}).Where(fmt.Sprintf("username = '%s'", data.Username)).Count(&usernameCheck)
	if usernameCheck > 0 {
		helpers.ActionFailedErrorResponse(c, "Sorry, this username is already taken")
		return
	}

	user := models.User{
		Name:     data.Name,
		Username: data.Username,
	}

	utils.DBConn.Create(&user)

	userDetails, _ := services.UserDetails(user.ID)

	// Generating common Response parameters like statusCode, message, data
	response := UserCreateResponseStruc{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusCreated,
			Message:    "User created successfully",
		},
		Data: UserCreateResponseDataStruct{
			User: userDetails,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}
