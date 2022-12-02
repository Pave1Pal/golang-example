package mpr

import (
	"example.com/internal/domain/dto"
	"example.com/internal/domain/entity"
)

type IPurchaseMapper interface {
	FromDto(dto *dto.PurchaseCreate) *entity.Purchase
}

type PurchaseMapper struct {
}

func (p PurchaseMapper) FromDto(dto *dto.PurchaseCreate) *entity.Purchase {
	purchase := entity.Purchase{
		Person:  dto.Person,
		Address: dto.Address,
		Cart: entity.Cart{
			Products: dto.Cart.Products,
		},
	}
	return &purchase
}
