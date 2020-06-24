package main

import (
	"log"

	"github.com/mauriliommachado/go-commerce/product-service/data"
	"github.com/mauriliommachado/go-commerce/product-service/models"
	"github.com/mauriliommachado/go-commerce/product-service/services"
)

func main() {
	log.Println("Initialing application")
	log.Println("Loading configurations")
	// Initialize server
	app := models.App{}
	app.ServiceName = "products"
	app.AuthService = "http://localhost:8080/validateToken"
	data.InitDb(&app)
	services.InitServer(&app)
}
