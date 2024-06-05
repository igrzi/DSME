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

	// localhost:8010, this will need to talk to localhost:8030 and localhost:8040
	// Access routes
	router.POST("/acesso/entry", controllers.AccessEntry) // Registra um acesso na cancela
	// This will talk to localhost:8040, if there's a spot available, subtract one from the total spots
	router.POST("/acesso/leave", controllers.AccessLeave) // Registra uma saída na cancela
	// This will talk to localhost:8030, and subtract a credit from the CPF

	// localhost:8020 THIS IS ISOLATED FROM THE REST
	// Cadastro routes 				== 				DONE
	router.POST("/usuario/create", controllers.UserCreate)   // Cria um usuário
	router.GET("/usuario", controllers.UserShow)             // Mostra os usuários
	router.PUT("/usuario/update", controllers.UserUpdate)    // Atualiza um usuário
	router.DELETE("/usuario/delete", controllers.UserDelete) // Deleta um usuário

	// localhost:8030
	// Creditos routes 				== 				DONE
	router.POST("/creditos/add", controllers.AddCredits) // Adiciona créditos
	router.POST("/creditos/use", controllers.UseCredits) // Usa créditos

	// localhost:8040
	// Vagas routes
	router.POST("/vagas/adjust", controllers.AdjustAmountSpot) // Cria uma vaga
	router.POST("/vagas/ocuppy", controllers.OcupySpot)        // Ocupa uma vaga
	router.POST("/vagas/vacate", controllers.VacateSpot)       // Libera uma vaga

	// Cancela routes 				== 				DONE
	router.POST("/cancela/open", controllers.BarrierOpen)   // Abre a cancela
	router.POST("/cancela/close", controllers.BarrierClose) // Fecha a cancela

	router.Run()
}
