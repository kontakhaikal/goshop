package repository

import (
	"github.com/kontakhaikal/goshop/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]*model.Product, error)
	FindBySlug(slug string) (*model.Product, error)
}

type GormProductRepository struct {
	db *gorm.DB
}

func (g GormProductRepository) FindBySlug(slug string) (*model.Product, error) {
	var product = model.Product{}
	err := g.db.Model(&model.Product{}).First(&product, &model.Product{Slug: slug}).Error
	return &product, err
}

func (g GormProductRepository) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	err := g.db.Model(&model.Product{}).Find(&products).Error
	return products, err
}

func NewGormProductRepository(db *gorm.DB) ProductRepository {
	return GormProductRepository{
		db,
	}
}
