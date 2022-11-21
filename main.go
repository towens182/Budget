package main

import (
	"github.com/gin-gonic/gin"
	"github.com/towens182/budget/controller"
	initializers "github.com/towens182/budget/initializers"
)

var (
	transactionController controller.Controller = controller.NewTransactionController()
	paymentController     controller.Controller = controller.NewPaymentController()
)

func init() {
	initializers.LoadEnvVaraiables()
}

func main() {
	server := gin.Default()

	server.GET("/transactions", transactionController.GetAll)
	server.POST("/transactions", transactionController.Add)

	server.GET("/payments", paymentController.GetAll)
	server.POST("/payments", paymentController.Add)

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
