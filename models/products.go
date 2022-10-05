package models

import (
	"time"
)

type Product struct {
	ID          uint64           `json:"id" gorm:"primary_key"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	Tags        string           `json:"tags"`
	CategorieID int64            `json:"categorie_id"`
	Categorie   ProductCategorie `json:"categorie" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Galleries   ProductGalleries `json:"galleries"`
}

type Products []Product

type ProductCategorie struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ProductCategories []ProductCategorie

type ProductGallerie struct {
	ID        uint64     `json:"id"`
	ProductID int        `json:"product_id"`
	Url       string     `json:"url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ProductGalleries []ProductGallerie
