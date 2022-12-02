package app

import (
	"example.com/internal/cntr"
	"example.com/internal/config"
	"example.com/internal/mpr"
	"example.com/internal/server"
	"example.com/internal/srv"
	"example.com/internal/strg"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func Run(pathToConfig string) {
	appConfig := getAppConfig(pathToConfig)
	db := strg.InitDB(appConfig.DBConfig)
	defer db.Close()

	var purchaseMapper mpr.IPurchaseMapper = mpr.PurchaseMapper{}

	cartRepository := strg.NewCartRepository(db)
	purchaseRepository := strg.NewPurchaseRepository(db)
	productRepository := strg.NewProductRepository(db)

	cartService := srv.NewCartService(&cartRepository)
	purchaseService := srv.NewPurchaseService(&cartService, &purchaseRepository)
	productService := srv.NewProductService(&productRepository)

	indexController := cntr.NewIndexController(&productService)
	productController := cntr.NewProductController(&productService)
	purchaseController := cntr.NewPurchaseController(&purchaseService, &purchaseMapper)

	serverApp := server.Server{
		IndexController:    *indexController,
		PurchaseController: *purchaseController,
		ProductController:  *productController}

	serverApp.Run(appConfig.ServerConfig)
}

func getAppConfig(pathToConfig string) *config.AppConfig {
	var appConfig config.AppConfig
	configFile, err := os.Open(pathToConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)

	if err := decoder.Decode(&appConfig); err != nil {
		log.Fatal(err.Error())
	}
	return &appConfig
}
