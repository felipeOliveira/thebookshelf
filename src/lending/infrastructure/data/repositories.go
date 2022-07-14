package data

import (
	"lending_service/catalog"

	"go.mongodb.org/mongo-driver/bson"
)

type catalogRepository struct {
	database Database
}

func CreateCatalogRepository(database Database) catalogRepository {
	return catalogRepository{database: database}
}

func (r catalogRepository) GetCategories() ([]catalog.Category, error) {

	categories := make([]catalog.Category, 0)
	err := r.database.List("categories", bson.D{{}}, nil, categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r catalogRepository) GetBooks(categoryId string) ([]catalog.Book, error) {

	books := make([]catalog.Book, 0)
	err := r.database.List("books", bson.D{{"categoryId", categoryId}}, nil, &books)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r catalogRepository) GetBookDetail(bookId string) (*catalog.Book, error) {

	book := &catalog.Book{}
	err := r.database.Get("books", bson.D{{"id", bookId}}, nil, book)

	if err != nil {
		return nil, err
	}

	return book, nil
}
