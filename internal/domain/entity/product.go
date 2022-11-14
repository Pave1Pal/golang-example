package entity

import "github.com/google/uuid"

type Product struct {
	Id    uuid.UUID
	Price int64
	Name  string
}
