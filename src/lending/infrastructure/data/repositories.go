package data

import (
	"context"
	"lending_service/catalog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type catalogRepository struct {
	database Database
}

func CreateCatalogRepository(database Database) catalogRepository {
	return catalogRepository{database: database}
}

func (r catalogRepository) GetCategories() ([]catalog.Category, error) {

	cur, err := r.database.List("categories", bson.D{{}}, nil)

	if err != nil {
		return nil, err
	}

	var categories []catalog.Category
	err = cur.All(context.Background(), &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r catalogRepository) GetBooks(categoryId string) ([]catalog.Book, error) {

	cur, err := r.database.List("books", bson.D{{"categoryId", categoryId}}, nil)

	if err != nil {
		return nil, err
	}

	var books []catalog.Book
	err = cur.All(context.Background(), &books)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return make([]catalog.Book, 0), nil
		}

		return nil, err
	}

	return books, nil
}

func (r catalogRepository) GetBookDetail(bookId string) (*catalog.Book, error) {

	var book *catalog.Book
	if err := r.database.Get("books", bson.D{{"id", bookId}}, nil).
		Decode(&book); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return book, nil
}
