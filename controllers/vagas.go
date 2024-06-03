package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func OcuppySpot(c *gin.Context) {

	fmt.Println("Ocupa vaga")
}

func VacateSpot(c *gin.Context) {

	fmt.Println("Libera vaga")
}
