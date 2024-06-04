package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BarrierOpen(c *gin.Context) {

	fmt.Println("Cancela abrindo")
}

func BarrierClose(c *gin.Context) {

	fmt.Println("Cancela fechando")
}
