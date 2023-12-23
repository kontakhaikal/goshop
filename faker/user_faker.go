package faker

import (
	"time"

	f "github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/kontakhaikal/goshop/model"
	"gorm.io/gorm"
)

func UserFaker() *model.User {
	return &model.User{
		ID:        uuid.NewString(),
		FirstName: f.FirstName(),
		LastName:  f.LastName(),
		Email:     f.Email(),
		Password:  "",
		CreatedAt: time.Time{},
		UpdateAt:  time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
