package srv

import (
	"example.com/internal/domain/entity"
	"example.com/internal/strg"
	"github.com/google/uuid"
)

type IPurchaseService interface {
	//find all purchase
	FindAll() ([]entity.Purchase, error)

	//find purchase by id
	FindById(id uuid.UUID) (*entity.Purchase, error)

	//create purchase
	Create(purchase *entity.Purchase) (*entity.Purchase, error)

	//update purchase with "id", take purchase fileds for update
	Update(id uuid.UUID, product *entity.Purchase) (*entity.Purchase, error)

	//deleate product by "id"
	Delete(id uuid.UUID) (*uuid.UUID, error)
}

type PurchaseService struct {
	PurchaseRepository strg.IPurchaseRepository
}

func (p PurchaseService) FindAll() ([]entity.Purchase, error) {
	all, err := p.PurchaseRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (p PurchaseService) FindById(id uuid.UUID) (*entity.Purchase, error) {
	purchase, err := p.PurchaseRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return purchase, err
}

func (p PurchaseService) Create(purchase *entity.Purchase) (*entity.Purchase, error) {
	created, err := p.PurchaseRepository.Create(purchase)
	if err != nil {
		return nil, err
	}
	return created, err
}

func (p PurchaseService) Update(id uuid.UUID, product *entity.Purchase) (*entity.Purchase, error) {
	product.Id = id
	updated, err := p.PurchaseRepository.Update(product)
	if err != nil {
		return nil, err
	}
	return updated, err
}

func (p PurchaseService) Delete(id uuid.UUID) (*uuid.UUID, error) {
	deleted, err := p.PurchaseRepository.Delete(id)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}
