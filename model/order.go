package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID                  string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID              string
	User                User
	OrderItems          []OrderItem
	OrderCustomer       *OrderCustomer
	Code                string
	Status              int
	OrderDate           time.Time
	PaymentDue          time.Time
	PaymentStatus       string  `gorm:"size:50;index"`
	PaymentToken        string  `gorm:"size:100;index"`
	BaseTotalPrice      float64 `gorm:"type:decimal(16,2)"`
	TaxAmount           float64 `gorm:"type:decimal(16,2)"`
	DiscountAmount      float64 `gorm:"type:decimal(10,2)"`
	DiscountPercent     float64 `gorm:"type:decimal(16,2)"`
	ShippingCost        float64 `gorm:"type:decimal(10,2)"`
	GrandTotal          float64 `gorm:"type:decimal(16,2)"`
	Note                string  `gorm:"type:text"`
	ShippingCourier     string  `gorm:"size:100"`
	ShippingServiceName string  `gorm:"size:100"`
	ApprovedBy          string  `gorm:"size:36"`
	ApprovedAt          time.Time
	CancelledBy         string `gorm:"size:36"`
	CancelledAt         time.Time
	CancellationNote    string `gorm:"size:225"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}
