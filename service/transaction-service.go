package service

import "github.com/towens182/budget/model"

type TransactionService interface {
	GetAll() []model.Transaction
}

type transactionService struct {
	transactions []model.Transaction
}

func New() TransactionService {

	testList := []model.Transaction{
		{
			Description:  "Oil Change",
			Cost:         20.00,
			MainCategory: "Automotive",
			SubCategory:  "Car Maintenence",
			Year:         "2022",
			Month:        "November",
		},
	}

	return &transactionService{
		transactions: testList,
	}
}

func (service *transactionService) GetAll() []model.Transaction {

	// TODO gather the transactions from csv
	return service.transactions
}
