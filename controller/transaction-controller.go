package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/towens182/budget/service"
)

var (
	transactionService service.TransactionService = service.New()
)

func GetAll(ctx *gin.Context) {
	result := transactionService.GetAll()
	if len(result) == 0 {
		ctx.JSON(404, nil)
	} else {
		ctx.JSON(200, result)
	}
}

// Add (return the new video also)
// ctx.BindJSON
