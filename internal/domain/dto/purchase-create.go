package dto

type PurchaseCreate struct {
	ProductId *string `form:"productId" binding:"required"`
	Person    *string `from:"person" binding:"required"`
	Address   *string `from:"address" binding:"required"`
}
