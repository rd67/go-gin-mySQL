package helpers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rd67/go-gin-mySQL/configs"
	"gorm.io/gorm"
)

func ErrorResponse(c *gin.Context, err error) {

	var statusCode int

	if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusInternalServerError
	}

	fmt.Println(err)

	c.AbortWithStatusJSON(statusCode, gin.H{
		"statusCode": statusCode,
		"error":      err.Error(),
	})
}

type ValidationErrorDataStruct struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}
type ValidationErrorResponseStruc struct {
	configs.CommonResponseStruct
	Data  []ValidationErrorDataStruct `json:"data"`
	Error string
}

func Descriptive(verr validator.ValidationErrors) []ValidationErrorDataStruct {
	errs := []ValidationErrorDataStruct{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationErrorDataStruct{Field: f.Field(), Reason: err})
	}

	return errs
}

func ValidationErrorResponse(c *gin.Context, err error) {
	const StatusCode = http.StatusBadRequest

	var ve validator.ValidationErrors

	var errorData = make([]ValidationErrorDataStruct, 0)

	if errors.As(err, &ve) {
		// Handle validator error formatting
		errorData = Descriptive(ve)
	}

	response := ValidationErrorResponseStruc{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: StatusCode,
			Message:    "Validation failed",
		},
		Data:  errorData,
		Error: err.Error(),
	}

	c.JSON(int(response.StatusCode), response)
}

type ActionFailedErrorResponseStruc struct {
	configs.CommonResponseStruct
	Data map[string]string `json:"data"`
}

func ActionFailedErrorResponse(c *gin.Context, message string) {
	const StatusCode = http.StatusBadRequest

	//TODO for Future
	data := make(map[string]string)

	response := ActionFailedErrorResponseStruc{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: StatusCode,
			Message:    message,
		},
		Data: data,
	}

	c.JSON(int(response.StatusCode), response)
}
