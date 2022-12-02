package cntr

import (
	"encoding/json"
	"example.com/internal/domain/dto"
	"example.com/internal/mpr"
	"example.com/internal/srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseController struct {
	purchaseService srv.IPurchaseService
	mapper          mpr.IPurchaseMapper
}

func (p PurchaseController) GetBuyForm(ctx *gin.Context) {
	productId := ctx.Param("productId")
	ctx.HTML(http.StatusOK, "buy-form.html", gin.H{"productId": productId})
}

func (p PurchaseController) CreateFromForm(ctx *gin.Context) {

	var createDto dto.PurchaseCreate

	if err := ctx.ShouldBind(&createDto); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	purchase := p.mapper.FromDto(&createDto)

	_, err := p.purchaseService.Create(purchase)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}

func (p PurchaseController) Create(ctx *gin.Context) {
	var purchaseDto dto.PurchaseCreate
	requestBody := ctx.Request.Body

	if err := json.NewDecoder(requestBody).Decode(&purchaseDto); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}
	purchase := p.mapper.FromDto(&purchaseDto)

	created, err := p.purchaseService.Create(purchase)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"id": created.Id})
}

func NewPurchaseController(
	purchaseService *srv.IPurchaseService,
	purchaseMapper *mpr.IPurchaseMapper) *PurchaseController {

	return &PurchaseController{
		purchaseService: *purchaseService,
		mapper:          *purchaseMapper}
}
