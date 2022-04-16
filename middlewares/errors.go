package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rd67/go-gin-mySQL/configs"
)

type ErrorHandlerResponseStruct struct {
	configs.CommonResponseStruct
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		fmt.Printf(err.Error())
		// log, handle, etc.
	}

	response := ErrorHandlerResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong, please try again",
		},
	}

	c.JSON(int(response.StatusCode), response)
	return
}
