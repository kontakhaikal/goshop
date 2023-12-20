package model

import "time"

type OrderItem struct {
	ID              string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Order           Order
	OrderID         string `gorm:"size:36;index"`
	Product         Product
	ProductID       string `gorm:"size:36;index"`
	Qty             int
	BasePrice       float64 `gorm:"type:decimal(16,2)"`
	BaseTotal       float64 `gorm:"type:decimal(16,2)"`
	TaxAmount       float64 `gorm:"type:decimal(16,2)"`
	TaxPercent      float64 `gorm:"type:decimal(10,2)"`
	DiscountAmount  float64 `gorm:"type:decimal(16,2)"`
	DiscountPercent float64 `gorm:"type:decimal(10,2)"`
	SubTotal        float64 `gorm:"type:decimal(16,2)"`
	Sku             string  `gorm:"size:36;index"`
	Name            string  `gorm:"size:225"`
	Weight          float64 `gorm:"type:decimal(10,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
