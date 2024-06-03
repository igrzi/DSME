package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddCredits(c *gin.Context) {

	fmt.Println("Adiciona crédito")
}

func UseCredits(c *gin.Context) {

	fmt.Println("Consome crédito")
}
