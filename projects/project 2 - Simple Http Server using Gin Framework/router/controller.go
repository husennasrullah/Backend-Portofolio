package router

import (
	"crudproduct/config"
	"crudproduct/endpoint/productEndpoint"
	"github.com/gin-gonic/gin"
	"strconv"
)

func APIController(routes *gin.Engine) {
	//list api url here
	//routes.Use(a.AuthRequired)

	//insert product
	routes.POST("/product", productEndpoint.ProductEndpoint.Insert)
	//get list product
	routes.GET("product", productEndpoint.ProductEndpoint.GetList)
	//search product
	routes.GET("product/:id", productEndpoint.ProductEndpoint.View)

	routes.Run(config.ApplicationConfiguration.GetServerHost() + ":" + strconv.Itoa(config.ApplicationConfiguration.GetServerPort()))
}
