package dto

type PurchaseCreate struct {
	ProductId string `form:"productId"`
	Person    string `from:"person"`
	Address   string `from:"address"`
}
