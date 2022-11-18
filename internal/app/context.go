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

	purchaseMapper := mpr.PurchaseMapper{}

	purchaseRepository := strg.PurchaseRepository{DB: db}
	productRepository := strg.ProductRepository{DB: db}

	purchaseService := srv.PurchaseService{PurchaseRepository: purchaseRepository}
	productService := srv.ProductService{ProductRepository: productRepository}

	indexController := cntr.IndexController{ProductService: productService}
	purchaseController := cntr.PurchaseController{
		PurchaseService: purchaseService,
		Mapper:          purchaseMapper}

	serverApp := server.Server{
		IndexController:    indexController,
		PurchaseController: purchaseController}

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
