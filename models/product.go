package models

import "github.com/globalsign/mgo/bson"

// Product model
type Product struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Price float64       `json:"price"`
	Name  string        `json:"name"`
}
