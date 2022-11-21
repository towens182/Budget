package main

import (
	"github.com/gin-gonic/gin"
	"github.com/towens182/budget/controller"
	initializers "github.com/towens182/budget/initializers"
)

func init() {
	initializers.LoadEnvVaraiables()
}

func main() {
	server := gin.Default()

	server.GET("/transactions", controller.GetAll)

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
