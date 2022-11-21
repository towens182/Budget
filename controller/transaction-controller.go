package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/towens182/budget/model"
	"github.com/towens182/budget/service"
)

type transactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController() Controller {
	var s service.TransactionService = service.NewTransactionService()

	return &transactionController{
		transactionService: s,
	}
}

func (c *transactionController) Add(ctx *gin.Context) {
	var newTransactions []model.Transaction

	if err := ctx.BindJSON(&newTransactions); err != nil {
		ctx.JSON(400, "invalid json")
		return
	}

	c.transactionService.AddTransaction(newTransactions)
	ctx.JSON(200, newTransactions)
}

func (c *transactionController) GetAll(ctx *gin.Context) {
	result := c.transactionService.GetAll()
	if len(result) == 0 {
		ctx.JSON(404, nil)
	} else {
		ctx.JSON(200, result)
	}
}
