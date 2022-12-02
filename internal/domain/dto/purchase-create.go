package dto

type PurchaseCreate struct {
	Cart    CartCreate `json:"cart" binding:"required"`
	Person  string     `json:"person" binding:"required"`
	Address string     `json:"address" binding:"required"`
}
