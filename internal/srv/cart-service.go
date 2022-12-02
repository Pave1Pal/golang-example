package srv

import (
	"example.com/internal/domain/entity"
	"example.com/internal/strg"
)

type ICartService interface {
	Create(cart *entity.Cart) (*entity.Cart, error)
}

type CartService struct {
	cartRepository strg.ICartRepository
}

func (c CartService) Create(cart *entity.Cart) (*entity.Cart, error) {
	cartPrice := c.calculateCartPrice(cart)
	cart.Price = *cartPrice

	created, err := c.cartRepository.Creat(cart)
	if err != nil {
		return nil, err
	}
	return created, err
}

func (c CartService) calculateCartPrice(cart *entity.Cart) *int64 {
	cartPrice := int64(0)

	for _, product := range cart.Products {
		prodPrice := product.Price
		cartPrice += prodPrice
	}

	return &cartPrice
}

func NewCartService(cartRepository *strg.ICartRepository) ICartService {
	return &CartService{cartRepository: *cartRepository}
}
