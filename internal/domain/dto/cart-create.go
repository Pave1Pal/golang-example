package dto

import "example.com/internal/domain/entity"

type CartCreate struct {
	Products []entity.Product `json:"products"`
}
