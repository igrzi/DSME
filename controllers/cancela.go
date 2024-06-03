package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BarrierOpen(c *gin.Context) {

	fmt.Println("Open cancela")
}

func BarrierClose(c *gin.Context) {

	fmt.Println("Close cancela")
}
