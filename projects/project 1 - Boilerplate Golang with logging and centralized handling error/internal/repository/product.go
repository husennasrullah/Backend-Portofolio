package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/models"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/constants"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/logger"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/queryparam"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/utils/errorutil"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAllProduct(ctx context.Context, param queryparam.Param) (listOrders []models.Products, count int, err error) {
	queryBuilder := sq.Select(
		"products.id", "products.product_sku", "products.product_name",
		"products.quantity", "products.price", "products.created_at", "products.updated_at").
		From("products").
		Limit(uint64(param.Limit)).
		Offset(uint64(param.Offset))

	for key, value := range param.Filter {
		queryBuilder = queryBuilder.Where(sq.Eq{key: value}).
			PlaceholderFormat(sq.Dollar)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build sql query", logrus.Fields{
			"error": err, "query": query, "args": args,
		})
		return nil, 0, errorutil.Wrap(err, constants.FAILED_BUILD_QUERY)

	}

	var products []models.Products
	err = r.db.Select(&products, query, args...)
	if err != nil {
		return nil, 0, errorutil.Wrap(err, constants.FAILED_EXECUTE_QUERY)
	}

	queryCount := sq.Select("COUNT(*)").From("products")

	for key, value := range param.Filter {
		queryCount = queryCount.Where(sq.Eq{key: value}).
			PlaceholderFormat(sq.Dollar)
	}

	queryCountData, args, err := queryCount.ToSql()
	if err != nil {
		return nil, 0, errorutil.Wrap(err, constants.FAILED_BUILD_QUERY)
	}

	if err = r.db.QueryRowxContext(ctx, queryCountData, args...).Scan(&count); err != nil {
		logger.Error("Error get total count", logrus.Fields{
			"error": err,
		})
		return nil, 0, errorutil.Wrap(err, constants.FAILED_EXECUTE_QUERY)
	}

	return products, count, nil
}
