package cntr

import (
	"example.com/internal/srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	productService srv.IProductService
}

func (p ProductController) GetAllProducts(ctx *gin.Context) {
	all, err := p.productService.FindAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, all)
}

func NewProductController(productService *srv.IProductService) *ProductController {
	return &ProductController{productService: *productService}
}
