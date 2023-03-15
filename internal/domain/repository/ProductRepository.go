package repository

import "go-project/internal/domain/entity"

type ProductRepository interface {
	FindById(id string) (entity.Product, error)

	FindAll() (*[]entity.Product, error)

	Create(product *entity.Product) error

	Update(product *entity.Product) error

	Delete(id string) error
}
