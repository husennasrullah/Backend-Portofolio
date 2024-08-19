package service

import (
	"crudproduct/attribute"
	"crudproduct/dto/out"
	"crudproduct/model"
	"crudproduct/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type productService struct {
	FileName string
}

var ProductService = productService{}.New()

func (input productService) New() (output productService) {
	input.FileName = "productService"
	return
}

// =============================================
func (input productService) InsertProduct(c *gin.Context) (result out.Payload, err model.ErrorModel) {
	db := attribute.ServerAttribute.GormClient
	// Validate input
	var productModel model.Product
	errs := c.ShouldBindJSON(&productModel)
	if errs != nil {
		err = model.GenerateErrorModel(400, "invalid request", input.FileName, "InsertProduct", errors.New("invalid json"))
		return
	}

	// Create book
	prod := model.Product{
		Name:        productModel.Name,
		Price:       productModel.Price,
		Description: productModel.Description,
		Quantity:    productModel.Quantity,
		CreatedAt:   time.Now(),
	}
	db.Create(&prod)

	//create data in redis
	key := "product-" + fmt.Sprint(prod.Id)
	WriteToRedis(key, util.StructToJSON(prod))

	result.Status.Message = "success insert product"
	return
}

func (input productService) ListProduct(c *gin.Context) (result out.Payload, err model.ErrorModel) {
	funcName := "ListProduct"
	db := attribute.ServerAttribute.GormClient
	var prod []model.Product

	getListData := readGetListData(c)
	getListData, err = getListData.ValidatePageLimitAndOrderBy()
	if err.Error != nil {
		return
	}

	offset := CountOffset(getListData.Page, getListData.Limit)

	//gorm get list data
	resultDB := db.Offset(offset).Limit(getListData.Limit).Order(getListData.OrderBy).Find(&prod)
	if resultDB.Error != nil {
		err = model.GenerateErrorModel(400, resultDB.Error.Error(), input.FileName, funcName, nil)
		return
	}

	//write data to redis
	for i := 0; i < len(prod); i++ {
		key := "product-" + fmt.Sprint(prod[i].Id)
		WriteToRedis(key, util.StructToJSON(prod[i]))
	}

	result.Status.Message = "success get data"
	result.Data = prod
	return
}

func (input productService) ViewProduct(c *gin.Context) (result out.Payload, err model.ErrorModel) {
	db := attribute.ServerAttribute.GormClient
	var product model.Product

	id, errs := strconv.ParseInt(c.Param("id"), 10, 32)
	if errs != nil {
		err = model.GenerateErrorModel(400, "unknown error", input.FileName, "InsertProduct", nil)
		return
	}

	tempResult := db.First(&product, id)
	if tempResult.Error != nil {
		err = model.GenerateErrorModel(400, "unknown data with this id", input.FileName, "InsertProduct", errors.New("no record data"))
		return
	}

	result.Status.Success = true
	result.Data = product

	return
}
