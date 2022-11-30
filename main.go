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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	server := gin.Default()

	server.Use(CORSMiddleware())
	server.GET("/transactions", transactionController.GetAll)
	server.POST("/transactions", transactionController.Add)

	server.GET("/payments", paymentController.GetAll)
	server.POST("/payments", paymentController.Add)

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
