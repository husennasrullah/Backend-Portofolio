package endpoint

import (
	"crudproduct/dto/out"
	"crudproduct/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"time"
)

type AbstractEndpoint struct {
	FileName  string
	TimeStart time.Time
	TimeEnd   time.Time
}

func (input AbstractEndpoint) ServeEndpoint(c *gin.Context, serveFunction func(ctx *gin.Context) (out.Payload, model.ErrorModel)) {
	serve(c, serveFunction)
}

func serve(c *gin.Context, serve func(c *gin.Context) (out.Payload, model.ErrorModel)) {
	var (
		err       model.ErrorModel
		output    out.Payload
		timeStart = time.Now()
		timeEnd   = time.Now()
	)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(string(debug.Stack()))
		} else {
			if err.Error != nil {
				fmt.Println(err.Error.Error())
			}
		}
		AbstractEndpoint{TimeStart: timeStart, TimeEnd: timeEnd}.finish(c, err, output)
	}()

	output, err = serve(c)
	if err.Error != nil {
		return
	}

	timeEnd = time.Now()
}

func (input AbstractEndpoint) finish(c *gin.Context, err model.ErrorModel, output out.Payload) {
	if err.Error != nil {
		writeErrorResponse(c, err)
	} else {
		writeSuccessResponse(c, output)
	}
}

func writeSuccessResponse(c *gin.Context, output out.Payload) {
	responseMessage := out.Payload{
		Status: out.StatusResponse{
			Success: true,
			Code:    "success",
			Message: output.Status.Message,
		},
		Data: output.Data,
	}
	c.JSON(http.StatusOK, responseMessage)
}

func writeErrorResponse(c *gin.Context, err model.ErrorModel) {
	responseMessage := out.Payload{
		Status: out.StatusResponse{
			Success: false,
			Code:    "failed",
			Message: err.Error.Error(),
		},
		Data: nil,
	}

	c.JSON(err.Code, responseMessage)
}
