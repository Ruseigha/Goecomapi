package domain

import "time"

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderCompleted OrderStatus = "completed"
	OrderCanceled  OrderStatus = "canceled"
)

type OrderItem struct {
	ProductID string  `bson:"product_id" json:"product_id"`
	Name      string  `bson:"name" json:"name"`
	Price     float64 `bson:"price" json:"price"`
	Quantity  int     `bson:"quantity" json:"quantity"`
}

type Order struct {
	ID        string      `bson:"_id,omitempty" json:"id"`
	UserID    string      `bson:"user_id" json:"user_id"`
	Items     []OrderItem `bson:"items" json:"items"`
	Total     float64     `bson:"total" json:"total"`
	Status    OrderStatus `bson:"status" json:"status"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at" json:"updated_at"`
}
