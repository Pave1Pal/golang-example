package entity

import "github.com/google/uuid"

type Product struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Price int64     `json:"price" db:"price"`
	Name  string    `json:"name" db:"name"`
}
