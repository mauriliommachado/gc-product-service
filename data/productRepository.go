package data

import (
	"context"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/mauriliommachado/go-commerce/product-service/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ProductRepository type
type ProductRepository struct {
	C *mongo.Collection
}

//Create product
func (r *ProductRepository) Create(product *models.Product) error {
	insertResult, err := r.C.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)
	return err
}

//Get product
func (r *ProductRepository) Get(product *models.Product) error {
	// create a value into which the result can be decoded
	var result models.Product
	filter := bson.D{{"name", "Ash"}}

	err := r.C.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//GetAll products
func (r *ProductRepository) GetAll() []models.Product {
	var products []models.Product
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*models.Product

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := r.C.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Product
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return products
}

//Delete Product
func (r *ProductRepository) Delete(id string) error {
	_, err := r.C.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
