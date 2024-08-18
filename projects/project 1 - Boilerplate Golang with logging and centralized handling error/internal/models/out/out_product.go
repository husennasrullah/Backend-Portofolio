package out

import "github.com/husennasrullah/Backend-Portofolio/project-1/internal/models"

type ResponseListProduct struct {
	Count int           `json:"count"`
	Data  []ListProduct `json:"data"`
}

type ListProduct struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Sku        string `json:"sku"`
	Price      string `json:"price"`
	Quantity   string `json:"quantity"`
	CreateDate string `json:"create_date"`
	UpdateDate string `json:"update_date"`
}

func ConvertToResponseListProduct(count int, products []models.Products) ResponseListProduct {
	var response []ListProduct

	for _, product := range products {
		response = append(response, ListProduct{
			Id:         product.Id,
			Name:       product.ProductName,
			Sku:        product.ProductSku,
			Price:      product.Price,
			Quantity:   product.Quantity,
			CreateDate: product.CreatedAt,
			UpdateDate: product.UpdatedAt,
		})
	}

	return ResponseListProduct{
		Count: count,
		Data:  response,
	}
}
