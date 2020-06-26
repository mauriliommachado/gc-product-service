package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mauriliommachado/go-commerce/product-service/data"
	"github.com/mauriliommachado/go-commerce/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	router.PUT("/product/:id", protectMiddleware(update))
	router.POST("/product", protectMiddleware(add))
	router.DELETE("/product/:id", protectMiddleware(delete))
	log.Fatal(http.ListenAndServe(":8081", router))
	app.Router = router
	return nil
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "pong\n")
}

func get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p, err := pr.Get(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, _ := json.Marshal(p)
	w.Write(response)
}

func delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := pr.Delete(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func getAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, _ := json.Marshal(pr.GetAll())
	w.Write(response)
}

func update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	product.ID, _ = primitive.ObjectIDFromHex(ps.ByName("id"))
	pr.Update(&product)
}

func add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	pr.Create(&product)
	w.Header().Add("Location", product.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}
