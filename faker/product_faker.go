package faker

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gorm.io/gorm"

	"github.com/kontakhaikal/goshop/model"
)

func ProductFaker(user *model.User) *model.Product {
	name := faker.Name()
	slug := slug.Make(name)
	return &model.Product{
		ID:               uuid.NewString(),
		UserID:           user.ID,
		Sku:              slug,
		Name:             faker.Name(),
		Slug:             slug,
		Price:            fakePrice(),
		Stock:            rand.Intn(100),
		Weight:           rand.Float64(),
		ShortDescription: faker.Paragraph(),
		Description:      faker.Paragraph(),
		Status:           1,
		UpdateAt:         time.Time{},
		CreatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

func fakePrice() float64 {
	pre := rand.Intn(2) + 1
	val := rand.Float64()
	return math.Round((val * float64(pre) / float64(pre)) * 100)
}
