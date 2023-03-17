package database

import (
	"github.com/gabrielmaximo/go-project/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductRepositoryGormImpl struct {
	DB *gorm.DB
}

func NewProductRepositoryGormImpl(db *gorm.DB) *ProductRepositoryGormImpl {
	return &ProductRepositoryGormImpl{DB: db}
}

func (r *ProductRepositoryGormImpl) Create(product *entity.Product) error {
	err := r.DB.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryGormImpl) Update(product *entity.Product) error {
	err := r.DB.UpdateColumns(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryGormImpl) FindByID(id string) (*entity.Product, error) {
	var product *entity.Product
	err := r.DB.First(&product, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryGormImpl) FindAll() (*[]entity.Product, error) {
	var products *[]entity.Product
	err := r.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryGormImpl) Delete(id string) error {
	err := r.DB.Delete(&entity.Product{}, "id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}
