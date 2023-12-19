package model

import "time"

type OrderItem struct {
	ID              string
	Order           Order
	OrderID         string
	Product         Product
	ProductID       string
	Qty             int
	BasePrice       float64
	BaseTotal       float64
	TaxAmount       float64
	TaxPercent      float64
	DiscountAmount  float64
	DiscountPercent float64
	SubTotal        float64
	Sku             string
	Name            string
	Weight          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
