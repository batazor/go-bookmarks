package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionBook holds the name of the articles collection
	CollectionBook = "books"
)

// Book model
type Book struct{
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omintempty"`
	Title     string        `json:"body" form:"body" binding:"requered" bson:"body"`
	Author    string        `json:"author" form:"author" binding:"requered" bson:"author"`
	Price     float32       `json:"price" form:"price" bson:"price"`
	CreatedOn int64         `json:"created_on" bson:"created_od"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}