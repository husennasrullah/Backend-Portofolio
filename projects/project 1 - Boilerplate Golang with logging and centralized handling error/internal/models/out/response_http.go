package out

import (
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/utils/errorutil"
	"net/http"
)

// DefaultHTTPResponse struct for all responses
type BaseResponse struct {
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

// SuccessResponse creates a successful response
func SuccessResponse(c *gin.Context, data interface{}) {
	response := BaseResponse{
		Status: http.StatusOK,
		Data:   data,
	}
	c.JSON(http.StatusOK, response)
}

// ErrorResponse creates an error response
func ErrorResponse(c *gin.Context, err error) {
	status, message := errorutil.HandleError(err)
	response := BaseResponse{
		Status:       status,
		ErrorMessage: message,
	}
	_ = c.Error(err)
	c.JSON(status, response)
}
