package model

import "time"

type Section struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name       string `gorm:"size:100"`
	Slug       string `gorm:"size:100"`
	UpdateAt   time.Time
	Categories []Category
}