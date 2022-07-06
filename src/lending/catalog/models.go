package catalog

import (
	"time"
)

type Book struct {
	Id            string    `json:"id" bson:"id"`
	Pages         uint16    `json:"pages" bson:"pages"`
	Isbn          string    `json:"isbn" bson:"isbn"`
	PublishedDate time.Time `json:"publishedDate" bson:"publishedDate"`
	Description   string    `json:"description" bson:"description"`
	Title         string    `json:"title" bson:"title"`
	Authors       []Author  `json:"authors" bson:"authors"`
	Publisher     Publisher `json:"publisher" bson:"publisher"`
	CategoryId    string    `json:"categoryId" bson:"categoryId"`
}

type Author struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Publisher struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Category struct {
	Id   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
