package data

import (
	"context"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/mauriliommachado/go-commerce/product-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	product.ID = insertResult.InsertedID.(primitive.ObjectID)
	log.Println("Inserted a single document: ", insertResult.InsertedID.(primitive.ObjectID).Hex())
	return err
}

//Update product
func (r *ProductRepository) Update(product *models.Product) error {
	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": bson.M{"price": product.Price,
		"name": product.Name}}
	_, err := r.C.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Update a single document ")
	return err
}

//Get product
func (r *ProductRepository) Get(id string) (models.Product, error) {
	// create a value into which the result can be decoded
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	var result models.Product
	err := r.C.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Print(err)
	}
	return result, err
}

//GetAll products
func (r *ProductRepository) GetAll() []models.Product {
	// Pass these options to the Find method
	findOptions := options.Find()

	// Here's an array in which you can store the decoded documents
	var results []models.Product

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := r.C.Find(context.TODO(), bson.M{}, findOptions)
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

		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return results
}

//Delete Product
func (r *ProductRepository) Delete(id string) error {
	_, err := r.C.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
