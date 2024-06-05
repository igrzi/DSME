package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/DSME/initializers"
	"github.com/igrzi/DSME/models"
)

// This is used to adjust the amount of spots available, or, if there isn't a record on the database, it will create a new spot
func AdjustAmountSpot(c *gin.Context) {
	quantitySpots, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		c.JSON(400, gin.H{"error": "quantity parameter can't be empty and must be a number!"})
		return
	}

	// Check if there's a record on the database
	var spot models.Spots
	result := initializers.DB.First(&spot)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			// If no record is found, create a new record
			newSpot := models.Spots{Quantity: quantitySpots}
			initializers.DB.Create(&newSpot)
			c.JSON(200, gin.H{"message": "New spot record created successfully!", "quantity": newSpot.Quantity})
		} else {
			// Handle other potential errors
			c.JSON(500, gin.H{"Internal error": result.Error.Error()})
		}
		return
	}

	// If a record is found, update the quantity of spots
	spot.Quantity = quantitySpots
	initializers.DB.Save(&spot)
	c.JSON(200, gin.H{"message": "Spot record updated successfully!", "quantity": spot.Quantity})
}

func OcupySpot(c *gin.Context) {
	// Find the spots record
	var spots models.Spots
	result := initializers.DB.First(&spots)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Spots record not found"})
		return
	}

	// Decrease the Quantity by 1 if available
	if spots.Quantity > 0 {
		spots.Quantity -= 1
	} else {
		c.JSON(400, gin.H{"error": "All spots are currently occupied"})
		return
	}

	// Save the updated spots back to the database
	initializers.DB.Save(&spots)

	c.JSON(200, gin.H{"message": "One spot used!", "available_spots": spots.Quantity})
}

func VacateSpot(c *gin.Context) {
	// Find the spots record
	var spots models.Spots
	result := initializers.DB.First(&spots)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Spots record not found"})
		return
	}

	// Decrease the Quantity by 1 if available
	if spots.Quantity > 0 {
		spots.Quantity -= 1
	} else {
		c.JSON(400, gin.H{"error": "All spots are currently occupied"})
		return
	}

	// Save the updated spots back to the database
	initializers.DB.Save(&spots)

	c.JSON(200, gin.H{"message": "One spot used!", "available_spots": spots.Quantity})
}
