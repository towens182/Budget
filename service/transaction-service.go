package service

import (
	"github.com/goccy/go-json"
	"github.com/towens182/budget/data"
	"github.com/towens182/budget/model"
)

type TransactionService interface {
	AddTransaction([]model.Transaction) []model.Transaction
	GetAll() []model.Transaction
}

type transactionService struct {
	transactions []model.Transaction
}

func New() TransactionService {

	var transactions []model.Transaction

	data := data.LoadJsonFromLocalFileSystem()
	err := json.Unmarshal(data, &transactions)
	if err != nil {
		// TODO handle this properly
		panic(err)
	}

	return &transactionService{
		transactions: transactions,
	}
}

func (service *transactionService) AddTransaction(newTransactions []model.Transaction) []model.Transaction {
	service.transactions = append(service.transactions, newTransactions...)

	// Save it to file
	data.AddJsonToDataFile(service.transactions)
	// TODO handle failure

	return newTransactions
}

func (service *transactionService) GetAll() []model.Transaction {
	return service.transactions
}
