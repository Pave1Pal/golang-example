package mpr

import (
	"example.com/internal/domain/dto"
	"example.com/internal/domain/entity"
	"github.com/google/uuid"
)

type IPurchaseMapper interface {
	FromCreateDto(dto *dto.PurchaseCreate) *entity.Purchase
}

type PurchaseMapper struct {
}

func (p PurchaseMapper) FromCreateDto(dto *dto.PurchaseCreate) *entity.Purchase {
	purchase := entity.Purchase{
		Product: entity.Product{Id: uuid.MustParse(*dto.ProductId)},
		Person:  *dto.Person,
		Address: *dto.Address,
	}
	return &purchase
}
