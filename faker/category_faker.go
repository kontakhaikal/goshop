package faker

import (
	"github.com/google/uuid"
	"github.com/kontakhaikal/goshop/model"
)

func CategoriesFaker() []model.Category {
	electronicsId := uuid.NewString()
	return []model.Category{
		{
			ID:   uuid.NewString(),
			Name: "Fashions",
		},
		{
			ID:   electronicsId,
			Name: "Electronics",
		},
		{
			ID:       uuid.NewString(),
			Name:     "Handphones",
			ParentID: electronicsId,
		},
		{
			ID:       uuid.NewString(),
			Name:     "Laptops",
			ParentID: electronicsId,
		},
		{
			ID:   uuid.NewString(),
			Name: "Baby and Toys",
		},
		{
			ID:   uuid.NewString(),
			Name: "Baby and Toys",
		},
		{
			ID:   uuid.NewString(),
			Name: "Digital Goods",
		},
	}
}
