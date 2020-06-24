package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product model
type Product struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Price float64            `json:"price"`
	Name  string             `json:"name"`
}
