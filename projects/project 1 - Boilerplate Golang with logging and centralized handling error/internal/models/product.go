package models

type Products struct {
	Id          int    `json:"id" db:"id"`
	ProductName string `json:"product_name" db:"product_name"`
	ProductSku  string `json:"product_sku" db:"product_sku"`
	Price       string `json:"price" db:"price"`
	Quantity    string `json:"quantity" db:"quantity"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}
