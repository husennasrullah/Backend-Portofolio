package service

import (
	"crudproduct/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

func readGetListData(c *gin.Context) (inputStruct dto.GetlistDTO) {
	inputStruct.Page, _ = strconv.Atoi(c.Query("page"))
	inputStruct.Limit, _ = strconv.Atoi(c.Query("limit"))
	inputStruct.OrderBy = c.Query("orderby")
	return
}

func CountOffset(page int, limit int) int {
	return (page - 1) * limit
}
