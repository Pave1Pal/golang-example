package srv

import (
	"example.com/internal/domain/entity"
	"example.com/internal/strg"
	"github.com/google/uuid"
)

type IPurchaseService interface {
	// FindAll find all purchase
	FindAll() ([]entity.Purchase, error)

	// FindById find purchase by id
	FindById(id uuid.UUID) (*entity.Purchase, error)

	//Create purchase
	Create(purchase *entity.Purchase) (*entity.Purchase, error)

	// Update purchase with "id", take purchase fileds for update
	Update(id uuid.UUID, product *entity.Purchase) (*entity.Purchase, error)

	// Delete product by "id"
	Delete(id uuid.UUID) (*uuid.UUID, error)
}

type PurchaseService struct {
	cartService        ICartService
	purchaseRepository strg.IPurchaseRepository
}

func (p PurchaseService) FindAll() ([]entity.Purchase, error) {
	all, err := p.purchaseRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (p PurchaseService) FindById(id uuid.UUID) (*entity.Purchase, error) {
	purchase, err := p.purchaseRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return purchase, err
}

func (p PurchaseService) Create(purchase *entity.Purchase) (*entity.Purchase, error) {
	createdCart, err := p.cartService.Create(&purchase.Cart)
	if err != nil {
		return nil, err
	}

	purchase.Cart.Id = createdCart.Id
	created, err := p.purchaseRepository.Create(purchase)
	if err != nil {
		return nil, err
	}

	return created, err
}

func (p PurchaseService) Update(id uuid.UUID, product *entity.Purchase) (*entity.Purchase, error) {
	product.Id = id
	updated, err := p.purchaseRepository.Update(product)
	if err != nil {
		return nil, err
	}
	return updated, err
}

func (p PurchaseService) Delete(id uuid.UUID) (*uuid.UUID, error) {
	deleted, err := p.purchaseRepository.Delete(id)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}

func NewPurchaseService(
	cartService *ICartService,
	purchaseRepository *strg.IPurchaseRepository) IPurchaseService {

	return &PurchaseService{
		cartService:        *cartService,
		purchaseRepository: *purchaseRepository}
}
