package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controllerFn func(c *gin.Context) error

func ResponseHanler(fn controllerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			log.Printf("Error while handling request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing request"})
		}
	}
}
