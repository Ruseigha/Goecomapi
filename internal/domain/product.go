package domain

import "time"

type Product struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	SKU         string    `bson:"sku" json:"sku"`
	Price       float64   `bson:"price" json:"price"`
	Stock       int       `bson:"stock" json:"stock"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
