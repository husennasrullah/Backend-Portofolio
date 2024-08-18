package service

import (
	"context"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/models/out"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/queryparam"
)

type ProductService interface {
	GetListProducts(context.Context, queryparam.Param) (out.ResponseListProduct, error)
}
