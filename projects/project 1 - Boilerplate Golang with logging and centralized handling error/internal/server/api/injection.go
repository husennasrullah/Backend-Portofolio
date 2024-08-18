package api

import (
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/controller"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/repository"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/service"
	"github.com/jmoiron/sqlx"
)

func getProductController(db *sqlx.DB) *controller.ProductController {
	productRepo := repository.NewProductRepository(db)
	productSrv := service.NewProductService(productRepo)
	productCtrl := controller.NewProductController(productSrv)
	return productCtrl
}
