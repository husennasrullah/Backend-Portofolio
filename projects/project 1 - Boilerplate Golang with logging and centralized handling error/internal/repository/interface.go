package repository

import (
	"context"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/models"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/queryparam"
)

type ProductRepository interface {
	GetAllProduct(context.Context, queryparam.Param) ([]models.Products, int, error)
}
