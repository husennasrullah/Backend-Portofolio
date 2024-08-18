package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-1/config"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/controller"
	"github.com/jmoiron/sqlx"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
	"net/http"
)

func Register(ginEngine *gin.Engine, db *sqlx.DB, cfg *config.Config, ctx context.Context) {
	// swagger
	ginEngine.GET("/swagger/*any", ginSwag.WrapHandler(swagFiles.Handler))
	ginEngine.GET("/swagger", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })

	// Health check endpoint
	ginEngine.GET("/health", controller.HealthCheck(db, cfg))
	ginEngine.NoRoute(unknownPage)

	ApiV1(&ginEngine.RouterGroup, db, cfg)
}

func ApiV1(app *gin.RouterGroup, db *sqlx.DB, cfg *config.Config) {
	productCtrl := getProductController(db)

	dashboardV1 := app.Group("/v1")
	{
		orderGroup := dashboardV1.Group("product")
		{
			orderGroup.GET("/", productCtrl.GetListProduct)
		}
	}
}

func unknownPage(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":        404,
		"data":          nil,
		"error_message": "no route match",
	})
}
