package main

import (
	"net/http"

	"github.com/rd67/go-gin-mySQL/configs"
	"github.com/rd67/go-gin-mySQL/modules"
	"github.com/rd67/go-gin-mySQL/utils"

	"github.com/gin-gonic/gin"
)

var PORT string

func init() {
	utils.ConnectDb()

	PORT = configs.GetConfig().App.Port
}

func main() {
	router := setupRouter()

	router.Run(":" + PORT)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// router.Use(middlewares.ErrorHandler)

	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	modules.SetupModuleRoutes(router)

	//	Handle 404 Error
	router.NoRoute(func(c *gin.Context) {

		var response = struct {
			configs.CommonResponseStruct
		}{
			CommonResponseStruct: configs.CommonResponseStruct{
				StatusCode: http.StatusNotFound,
				Message:    "Route Not Found",
			},
		}

		c.JSON(int(response.StatusCode), response)

		return

	})

	return router
}
