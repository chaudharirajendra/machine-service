package main

import (
	"fmt"
	"log"
	"machine-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define the route for machine handler
	router.GET("/machine", handlers.MachineHandler)

	port := 8081
	log.Printf("Machine Service running at http://localhost:%d\n", port)

	// Start the server
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
