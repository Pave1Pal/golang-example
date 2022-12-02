package main

import (
	"example.com/internal/app"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	log.Println(dir)
	configPath := flag.String("config", "", "path to configuration file")
	flag.Parse()

	log.Println(*configPath)
	app.Run(*configPath)
}
