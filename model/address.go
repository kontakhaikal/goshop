package model

import "time"

type Address struct {
	Id string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User
	UserId    string `gorm:"size:36;index"`
	Name      string `gorm:"size:100,not mull"`
	IsPrimary bool
	CityId    string `gorm:"size100"`
	Address1  string `gorm:"size:225"`
	Address2  string `gorm:"size:225"`
	Phone     string `gorm:"size100"`
	Email     string `gorm:"size100"`
	PostCode  string `gorm:"size100"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
