package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID                  string
	UserID              string
	User                User
	OrderItems          []OrderItem
	OrderCustomer       *OrderCustomer
	Code                string
	Status              int
	OrderDate           time.Time
	PaymentDue          time.Time
	PaymentStatus       string
	PaymentToken        string
	BaseTotalPrice      float64
	TaxAmount           float64
	DiscountAmount      float64
	DiscountPercent     float64
	ShippingCost        float64
	GrandTotal          float64
	Note                string
	ShippingCourier     string
	ShippingServiceName string
	ApprovedBy          string
	ApprovedAt          time.Time
	CancelledBy         string
	CancelledAt         time.Time
	CancellationNote    string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}
