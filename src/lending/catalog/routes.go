package catalog

import (
	"lending_service/infrastructure/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, repo catalogRepository) {
	r.GET("/categories", getCategories(repo))
	r.GET("/categories/:category/books", getBooks(repo))
	r.GET("/book/:booksid", getBookDetail(repo))
}

func getCategories(repo catalogRepository) gin.HandlerFunc {

	handlerFn := func(c *gin.Context) error {
		categories, err := repo.GetCategories()
		if err != nil {

			return err
		}

		c.JSON(http.StatusOK, categories)
		return nil
	}

	return handlers.ResponseHanler(handlerFn)
}

func getBooks(repo catalogRepository) gin.HandlerFunc {

	handlerFn := func(c *gin.Context) error {
		books, err := repo.GetBooks(c.Param("category"))

		if err != nil {
			return err
		}

		c.JSON(http.StatusOK, books)
		return nil
	}

	return handlers.ResponseHanler(handlerFn)
}

func getBookDetail(repo catalogRepository) gin.HandlerFunc {

	handlerFn := func(c *gin.Context) error {
		book, err := repo.GetBookDetail(c.Param("booksid"))

		if err != nil {
			return err
		}

		if book == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return nil
		}

		c.JSON(http.StatusOK, book)
		return nil
	}

	return handlers.ResponseHanler(handlerFn)
}
