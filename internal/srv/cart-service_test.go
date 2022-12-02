package srv

import (
	"example.com/internal/domain/entity"
	"testing"
)

func TestCalculateCartPrice(t *testing.T) {
	c := CartService{}

	prodPrice1 := int64(3000)
	prodPrice2 := int64(4000)

	totalPrice := prodPrice1 + prodPrice2

	cart := entity.Cart{
		Products: []entity.Product{
			{Price: prodPrice1},
			{Price: prodPrice2},
		}}

	price := c.calculateCartPrice(&cart)

	if totalPrice != *price {
		t.Error("expected: ", totalPrice, "actual: ", *price)
	}
}

func TestCalculateEmptyCartPrice(t *testing.T) {
	c := CartService{}
	cart := entity.Cart{}

	expectedPrice := int64(0)
	actualPrice := c.calculateCartPrice(&cart)

	if *actualPrice != expectedPrice {
		t.Error("expected: ", expectedPrice, "actual: ", actualPrice)
	}
}
