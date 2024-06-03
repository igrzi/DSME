package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/DSME/initializers"
	"github.com/igrzi/DSME/models"
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {

	fmt.Println("Batata doce!")

	// Get data of user from request body
	var userData struct {
		Cpf      int
		Name     string
		Category string
	}
	c.Bind(&userData)

	if !checkIfCPFisRegistered(userData.Cpf) {
		// Se o CPF não estiver registrado, cria o usuário
		user := models.User{Cpf: userData.Cpf, Name: userData.Name, Category: userData.Category}
		result := initializers.DB.Create(&user)

		if result.Error != nil {
			c.Status(400)
			return
		}

		// Return a 200 status

		c.Status(200)
	} else {
		fmt.Println()
		c.JSON(409, "User already on database")
	}
}

func UserShow(c *gin.Context) {

	fmt.Println("UserShow")
}

func UserUpdate(c *gin.Context) {

	fmt.Println("UserUpdate")
}

func UserDelete(c *gin.Context) {

	fmt.Println("UserDelete")
}

func checkIfCPFisRegistered(cpf int) bool {
	var user models.User

	result := initializers.DB.Table("users").Where("cpf = ?", cpf).First(&user)
	if result.Error != nil {
		// Log any errors
		fmt.Println("Error checking if CPF is registered:", result.Error)
		return false
	}

	// Log the user found (if any)
	fmt.Println("User found:", user)

	// If the user is found, return true; otherwise, return false
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
