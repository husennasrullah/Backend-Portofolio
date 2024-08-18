package service

import (
	"context"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/models/out"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/repository"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/queryparam"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (srv productService) GetListProducts(ctx context.Context, param queryparam.Param) (out.ResponseListProduct, error) {
	listProduct, count, err := srv.repo.GetAllProduct(ctx, param)
	if err != nil {
		return out.ResponseListProduct{}, err
	}
	return out.ConvertToResponseListProduct(count, listProduct), nil
}
