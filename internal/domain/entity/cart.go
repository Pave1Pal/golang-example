package entity

import "github.com/google/uuid"

type Cart struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Price    int64     `json:"price" db:"price"`
	Products []Product `json:"products"`
}
