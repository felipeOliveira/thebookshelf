package main

import (
	"lending_service/catalog"
	"lending_service/infrastructure/data"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db, err := data.NewDatabase(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	catalog.Register(router, data.CreateCatalogRepository(db))

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
