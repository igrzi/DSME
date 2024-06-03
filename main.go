package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igrzi/DSME/controllers"
	"github.com/igrzi/DSME/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// Access routes
	router.POST("/acesso/entry", controllers.AccessEntry) // Registra um acesso na cancela
	router.POST("/acesso/leave", controllers.AccessLeave) // Registra uma saída na cancela

	// Cadastro routes
	router.POST("/usuario/create", controllers.UserCreate)   // Cria um usuário
	router.GET("/usuario", controllers.UserShow)             // Mostra os usuários
	router.PUT("/usuario/update", controllers.UserUpdate)    // Atualiza um usuário
	router.DELETE("/usuario/delete", controllers.UserDelete) // Deleta um usuário

	// Cancela routes
	router.POST("/cancela/open", controllers.BarrierOpen)   // Abre a cancela
	router.POST("/cancela/close", controllers.BarrierClose) // Fecha a cancela

	// Creditos routes
	router.POST("/creditos/add", controllers.AddCredits) // Adiciona créditos
	router.POST("/creditos/use", controllers.UseCredits) // Usa créditos

	// Vagas routes
	router.POST("/vagas/ocuppy", controllers.OcuppySpot) // Ocupa uma vaga
	router.POST("/vagas/vacate", controllers.VacateSpot) // Libera uma vaga

	router.Run()
}
