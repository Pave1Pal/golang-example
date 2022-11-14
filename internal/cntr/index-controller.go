package cntr

import (
	"example.com/internal/srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	ProductService srv.IProductService
}

func (i IndexController) IndexPage(ctx *gin.Context) {
	all, err := i.ProductService.FindAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.HTML(http.StatusOK, "index.html", all)
}
