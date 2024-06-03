package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AccessEntry(c *gin.Context) {

	fmt.Println("AccessEntry")
}

func AccessLeave(c *gin.Context) {

	fmt.Println("AccessLeave")
}
