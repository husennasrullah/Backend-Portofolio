package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/models/out"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/service"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/queryparam"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(svc service.ProductService) *ProductController {
	return &ProductController{service: svc}
}

// GetListProduct
// @Tags			Internal API - List Product
// @Summary			List Product
// @Description		"List Product"
// @Accept			json
// @Produce			json
// @Param			page				query		int					false	"filter page"
// @Param			limit				query		int					false	"set limit"
// @Param			product_name		query		int					false	"filter by product_name"
// @Param			product_sku			query		int					false	"filter by product_sku"
// @Success			200				 	{object}	out.BaseResponse
// @Router		    /v1/order [get]
func (ctrl *ProductController) GetListProduct(ctx *gin.Context) {
	param, _ := queryparam.RequestParam(ctx)

	orders, err := ctrl.service.GetListProducts(ctx, param)
	if err != nil {
		out.ErrorResponse(ctx, err)
		return
	}
	out.SuccessResponse(ctx, orders)
}
