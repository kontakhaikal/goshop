package model

import "time"

type OrderCustomer struct {
	ID         string
	User       User
	UserID     string
	Order      Order
	OrderID    string
	FirstName  string
	LastName   string
	CityID     string
	ProvinceID string
	Address1   string
	Address2   string
	Phone      string
	Email      string
	PostCode   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
