package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionBook holds the name of the articles collection
	CollectionBook = "books"
)

// Book model
type Book struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Isbn      string        `json:"isbn" bson:"isbn"`
	Author    string        `json:"author" form:"author" binding:"requered" bson:"author"`
	Title     string        `json:"title" form:"title" binding:"requered" bson:"title"`
	Price     int32         `json:"price" form:"price" bson:"price"`
	CreatedOn int64         `json:"created_on" bson:"created_od"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
}
