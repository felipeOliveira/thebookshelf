package handlers

import (
	"lending_service/infrastructure/errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controllerFn func(c *gin.Context) (interface{}, error)

func ResponseHanler(fn controllerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := fn(c)

		if err != nil {
			log.Println("handling response error")
			handleError(c, err)
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func handleError(c *gin.Context, err error) {
	switch err.(type) {
	case *errors.NotFoundError:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing request"})
	}
}
