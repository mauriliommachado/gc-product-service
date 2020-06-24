package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mauriliommachado/go-commerce/product-service/data"
	"github.com/mauriliommachado/go-commerce/product-service/models"
)

var pr data.ProductRepository

// InitServer function
func InitServer(app *models.App) error {
	log.Println("Initialing server")
	InitMiddleware(app)
	pr.C = app.Collection
	router := httprouter.New()
	router.GET("/ping", ping)
	router.GET("/product/:id", get)
	router.GET("/product", getAll)
	router.PUT("/product/:id", update)
	router.POST("/product", add)
	log.Fatal(http.ListenAndServe(":8081", router))
	app.Router = router
	return nil
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "pong\n")
}

func get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "You have the key!\n")
}

func getAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "You have the key!\n")
}

func update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "You have the key!\n")
}

func add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pr.Create(&models.Product{})
	fmt.Fprint(w, "You have the key!\n")
}
