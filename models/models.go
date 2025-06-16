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
	OrderID   int
	Amount    float64
	Status    string
	CreatedAt time.Time
}

type OrderDetail struct {
	OrderID           int       `json:"order_id"`
	ProductName       string    `json:"product_name"`
	Quantity          int       `json:"quantity"`
	Price             float64   `json:"price"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
}

type PopularProduct struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	TotalSold   float64 `json:"total_sold"`
}
