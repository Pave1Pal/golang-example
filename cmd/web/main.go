package main

import (
	"example.com/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	app.Run("/home/pavelpal/projects/go-projects/TP-lab-2/resources/config/app-config.yaml")
}
