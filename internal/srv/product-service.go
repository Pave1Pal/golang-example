package srv

import (
	"example.com/internal/domain/entity"
	"example.com/internal/strg"
	"github.com/google/uuid"
)

type IProductService interface {
	//find all products
	FindAll() ([]entity.Product, error)

	//find product by id
	FindById(id uuid.UUID) (*entity.Product, error)

	//create product
	Create(product *entity.Product) (*entity.Product, error)

	//update product with "id", take product fileds for update
	Update(targetId uuid.UUID, product *entity.Product) (*entity.Product, error)

	//deleate product by "id"
	Delete(id uuid.UUID) (*uuid.UUID, error)
}

type ProductService struct {
	ProductRepository strg.IProductRepository
}

func (p ProductService) FindAll() ([]entity.Product, error) {
	all, err := p.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (p ProductService) FindById(id uuid.UUID) (*entity.Product, error) {
	product, err := p.ProductRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p ProductService) Create(product *entity.Product) (*entity.Product, error) {
	created, err := p.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (p ProductService) Update(targetId uuid.UUID, product *entity.Product) (*entity.Product, error) {
	product.Id = targetId
	updated, err := p.ProductRepository.Update(product)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (p ProductService) Delete(id uuid.UUID) (*uuid.UUID, error) {
	deletedID, err := p.ProductRepository.Delete(id)
	if err != nil {
		return nil, err
	}
	return deletedID, nil
}
