package service

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/towens182/budget/data"
	"github.com/towens182/budget/model"
)

type TransactionService interface {
	GetAll() []model.Transaction
}

type transactionService struct {
	transactions []model.Transaction
}

func New() TransactionService {

	var transactions []model.Transaction

	data := data.LoadJsonFromLocalFileSystem()
	fmt.Println(data)
	err := json.Unmarshal(data, &transactions)
	if err != nil {
		// TODO handle this properly
		panic(err)
	}

	return &transactionService{
		transactions: transactions,
	}
}

func (service *transactionService) GetAll() []model.Transaction {
	return service.transactions
}
