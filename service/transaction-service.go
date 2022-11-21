package service

import (
	"github.com/goccy/go-json"
	"github.com/towens182/budget/data"
	"github.com/towens182/budget/model"
)

// TODO: load this from env, connect to a DB
var transactionsFilePath = "data/transactions.json"

type TransactionService interface {
	AddTransaction([]model.Transaction) []model.Transaction
	GetAll() []model.Transaction
}

type transactionService struct {
	transactions []model.Transaction
}

func NewTransactionService() TransactionService {

	var transactions []model.Transaction

	data := data.LoadJsonFromDataFile(transactionsFilePath)
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
	data.AddJsonToDataFile(transactionsFilePath, service.transactions)
	// TODO handle failure

	return newTransactions
}

func (service *transactionService) GetAll() []model.Transaction {
	return service.transactions
}
