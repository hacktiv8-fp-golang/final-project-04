package model

import "time"

type Category struct {
	ID int `gorm:"primaryKey" json:"id,omitempty"`
	Type string `gorm:"not null" json:"type" valid:"required~Product type is required"`
	SoldProductAmount int `json:"sold_product_amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Products []Product
}

type CategoryUpdate struct {
	Type string `gorm:"not null" json:"type" valid:"required~Product type is required"`
}
