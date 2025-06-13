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
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
}
