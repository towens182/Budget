package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/towens182/budget/model"
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

func AddTransaction(ctx *gin.Context) {
	var newTransactions []model.Transaction

	if err := ctx.BindJSON(&newTransactions); err != nil {
		ctx.JSON(400, "invalid json")
		return
	}

	transactionService.AddTransaction(newTransactions)
	ctx.JSON(200, newTransactions)
}
