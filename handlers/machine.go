package handlers

import (
	"log"
	"machine-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MachineHandler handles GET requests for /machine
func MachineHandler(c *gin.Context) {
	machineID := c.Query("machine_id")

	// Fetch the machine by ID from the service layer
	machine, err := services.GetMachineByID(machineID)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Machine not found"})
		} else {
			log.Printf("Error fetching machines: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading machine data"})
		}
		return
	}

	// Return the machine as JSON response
	c.JSON(http.StatusOK, machine)
}
