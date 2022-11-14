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
	mapper          mpr.IPurchaseMapper
}

//TODO протестировать этот контроллер
func (p PurchaseController) CreateFromForm(ctx *gin.Context) {
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var createDto dto.PurchaseCreate
	if err := ctx.Bind(&createDto); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	purchase := p.mapper.FromCreateDto(&createDto)
	_, err := p.PurchaseService.Create(purchase)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	ctx.Redirect(http.StatusCreated, "created-purchase.html")
}
