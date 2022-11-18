package server

import (
	"example.com/internal/cntr"
	"example.com/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Server struct {
	IndexController    cntr.IndexController
	PurchaseController cntr.PurchaseController
}

func (s Server) Run(conf *config.ServerConfig) {
	router := gin.Default()
	router.LoadHTMLGlob(conf.Templates + "*")
	s.matchPaths(router)

	server := initServer(conf, router)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

func initServer(conf *config.ServerConfig, router *gin.Engine) *http.Server {
	s := &http.Server{
		Addr:         conf.Host + ":" + conf.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s
}

func (s Server) matchPaths(router *gin.Engine) {
	router.GET("/", s.IndexController.IndexPage)
	router.POST("/", s.IndexController.IndexPage)
	router.GET("/buy/:productId", s.PurchaseController.GetBuyForm)
	router.POST("/buy/:productId", s.PurchaseController.CreateFromForm)
}
