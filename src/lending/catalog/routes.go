package catalog

import (
	"lending_service/infrastructure/errors"
	"lending_service/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, repo catalogRepository) {
	r.GET("/categories", getCategories(repo))
	r.GET("/categories/:category/books", getBooks(repo))
	r.GET("/book/:booksid", getBookDetail(repo))
}

func getCategories(repo catalogRepository) gin.HandlerFunc {

	handlerFn := func(c *gin.Context) (interface{}, error) {
		categories, err := repo.GetCategories()
		if err != nil {

			return nil, err
		}

		return categories, nil
	}

	return handlers.ResponseHanler(handlerFn)
}

func getBooks(repo catalogRepository) gin.HandlerFunc {

	handlerFn := func(c *gin.Context) (interface{}, error) {
		books, err := repo.GetBooks(c.Param("category"))

		if err != nil {
			return nil, err
		}

		return books, nil
	}

	return handlers.ResponseHanler(handlerFn)
}

func getBookDetail(repo catalogRepository) gin.HandlerFunc {

	handlerFn := func(c *gin.Context) (interface{}, error) {
		book, err := repo.GetBookDetail(c.Param("booksid"))

		if err != nil {
			return nil, err
		}

		if book == nil {
			return nil, errors.NewNotFoundError()
		}

		return book, nil
	}

	return handlers.ResponseHanler(handlerFn)
}
