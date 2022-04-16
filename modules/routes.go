package modules

import (
	"github.com/gin-gonic/gin"
)

func SetupModuleRoutes(app *gin.Engine) {

	apiV1 := app.Group("/v1")

	//	Users Routes
	userV1Routes := apiV1.Group("/users")
	{
		userV1Routes.GET("/", UsersList)
		userV1Routes.POST("/", UserCreate)
	}

	//	Teams Routes
	teamV1Routes := apiV1.Group("/teams")
	{
		teamV1Routes.GET("/", TeamsList)
		teamV1Routes.POST("/", TeamCreate)
		teamV1Routes.GET("/:teamId", TeamDetails)
	}

	//	Team Users Routes
	teamUserV1Routes := apiV1.Group("/teamUsers")
	{
		teamUserV1Routes.POST("/", TeamUsersAdd)
		teamUserV1Routes.PUT("/:teamUserId", TeamUserStepsAdd)
	}

}
