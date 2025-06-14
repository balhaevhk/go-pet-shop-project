package models

import "time"

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int // Количество на складе
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Order struct {
	ID         int
	UserID     int
	TotalPrice float64
	CreatedAt  time.Time
}

type OrderItem struct {
	ID        int `json:"id"`
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Transaction struct {
	ID        int
	orderID   int
	Amount    float64
	Status    string
	CreatedAt time.Time
}
