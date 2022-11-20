package main

import (
	"github.com/gin-gonic/gin"
	"github.com/towens182/budget/controller"
	initilizers "github.com/towens182/budget/initializers"
	"github.com/towens182/budget/service"
)

var (
	transactionService service.TransactionService       = service.New()
	videoController    controller.TransactionController = controller.New(transactionService)
)

func init() {
	initilizers.LoadEnvVaraiables()
}

func main() {
	server := gin.Default()

	server.GET("/transactions", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.GetAll())
	})

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
