package cntr

import (
	"example.com/internal/domain/dto"
	"example.com/internal/mpr"
	"example.com/internal/srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseController struct {
	PurchaseService srv.IPurchaseService
	Mapper          mpr.IPurchaseMapper
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
	purchase := p.Mapper.FromCreateDto(&createDto)

	_, err := p.PurchaseService.Create(purchase)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
