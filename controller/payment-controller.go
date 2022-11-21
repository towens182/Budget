package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/towens182/budget/model"
	"github.com/towens182/budget/service"
)

type paymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController() Controller {
	var s service.PaymentService = service.NewPaymentService()

	return &paymentController{
		paymentService: s,
	}
}

func (c *paymentController) Add(ctx *gin.Context) {
	var newPayments []model.Payment

	if err := ctx.BindJSON(&newPayments); err != nil {
		ctx.JSON(400, "invalid json")
		return
	}

	c.paymentService.AddPayment(newPayments)
	ctx.JSON(200, newPayments)
}

func (c *paymentController) GetAll(ctx *gin.Context) {
	result := c.paymentService.GetAll()
	if len(result) == 0 {
		ctx.JSON(404, nil)
	} else {
		ctx.JSON(200, result)
	}
}
