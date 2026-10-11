package main

import (
	"log"

	"test-api/src/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	log.Println("Server running at http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
