package controller

import (
	"github.com/towens182/budget/model"
	"github.com/towens182/budget/service"
)

type TransactionController interface {
	// Add
	GetAll() []model.Transaction
}

type controller struct {
	service service.TransactionService
}

func New(service service.TransactionService) TransactionController {
	return controller{
		service: service,
	}
}

func (c controller) GetAll() []model.Transaction {
	return c.service.GetAll()
}

// Add (return the new video also)
// ctx.BindJSON
