package model

import "time"

type Product struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Price       int64     `json:"price"`
	Description string    `json:"description"`
	Quantity    int64     `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}
