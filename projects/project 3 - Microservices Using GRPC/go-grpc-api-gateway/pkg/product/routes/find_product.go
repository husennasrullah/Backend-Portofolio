package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-3/go-grpc-api-gateway/pkg/product/pb"
	"net/http"
	"strconv"
)

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
