package cntr

import (
	"example.com/internal/srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	productService srv.IProductService
}

func (i IndexController) IndexPage(ctx *gin.Context) {
	all, err := i.productService.FindAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.HTML(http.StatusOK, "index.html", all)
}

func NewIndexController(productService *srv.IProductService) *IndexController {
	return &IndexController{productService: *productService}
}
