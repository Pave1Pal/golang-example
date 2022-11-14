package entity

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	Id      uuid.UUID
	Person  string
	Address string
	Date    time.Time
	Product Product
}
