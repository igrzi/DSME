package main

import (
	"github.com/igrzi/DSME/initializers"
	"github.com/igrzi/DSME/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Spots{}, &models.Access{}, &models.Credits{})
}
