package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/DSME/initializers"
	"github.com/igrzi/DSME/models"
)

func AddCredits(c *gin.Context) {
	userCPF, err := strconv.Atoi(c.Query("cpf"))
	DealWithError(err, c)

	userCredits, err := strconv.Atoi(c.Query("credits"))
	DealWithError(err, c)

	credits := models.Credits{Cpf: userCPF, CreditAmount: userCredits}
	initializers.DB.Create(&credits)

	c.JSON(200, "Credits added successfully!")
}

func UseCredits(c *gin.Context) {
	userCPF, err := strconv.Atoi(c.Query("cpf"))
	DealWithError(err, c)

	// Find the credits for the given CPF
	var credits models.Credits
	result := initializers.DB.Where("cpf = ?", userCPF).First(&credits)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Decrease the CreditAmount by 1
	if credits.CreditAmount > 0 {
		credits.CreditAmount -= 1
	} else {
		c.JSON(409, gin.H{"error": "Not enough credits"})
		return
	}

	// Save the updated credits back to the database
	initializers.DB.Save(&credits)

	c.JSON(200, gin.H{"message": "Credit used successfully!", "remaining_credits": credits.CreditAmount})
}

func DealWithError(err error, c *gin.Context) {
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
	}
}
