package entity

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	Id      uuid.UUID `json:"id"`
	Person  string    `json:"person"`
	Address string    `json:"address"`
	Date    time.Time `json:"date"`
	Cart    Cart      `json:"shopingCart"`
}
