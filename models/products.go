package models

import "time"

type Product struct {
	ID           uint64            `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Price        float64           `json:"price"`
	Tags         int               `json:"tags"`
	CategoriesID int               `json:"categories_id"`
	Categories   ProductCategories `json:"categories"`
}

type Products []Product

type ProductCategories struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ProductGalleries struct {
	ID        uint64    `json:"id"`
	ProductID Product   `json:"product_id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
