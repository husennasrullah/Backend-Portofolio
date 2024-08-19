package productEndpoint

import (
	"crudproduct/endpoint"
	"crudproduct/service"
	"github.com/gin-gonic/gin"
)

type productEndpoint struct {
	endpoint.AbstractEndpoint
}

var ProductEndpoint = productEndpoint{}.New()

func (input productEndpoint) New() (output productEndpoint) {
	output.FileName = "endpoint//product.go"
	return
}

func (input productEndpoint) Insert(c *gin.Context) {
	input.ServeEndpoint(c, service.ProductService.InsertProduct)
}

func (input productEndpoint) View(c *gin.Context) {
	input.ServeEndpoint(c, service.ProductService.ViewProduct)
}

func (input productEndpoint) GetList(c *gin.Context) {
	input.ServeEndpoint(c, service.ProductService.ListProduct)
}
