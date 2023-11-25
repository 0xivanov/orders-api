package model

import "github.com/google/uuid"

type OrderStatus string

const (
	// Define the possible order statuses
	Processing OrderStatus = "Processing"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
	Canceled   OrderStatus = "Canceled"
)

type Order struct {
	OrderId    uint64      `json:"orderId"`
	CustomerId uuid.UUID   `json:"customerId"`
	Items      []Item      `json:"items"`
	Status     OrderStatus `json:"status"`
}

type Item struct {
	ItemId   uuid.UUID `json:"itemId"`
	Quantity uint      `json:"quantity"`
	Price    float32   `json:"price"`
}
