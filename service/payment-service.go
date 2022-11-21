package service

import (
	"github.com/goccy/go-json"
	"github.com/towens182/budget/data"
	"github.com/towens182/budget/model"
)

var paymentsFilePath string = "data/savedPaymentSources.json"

type PaymentService interface {
	AddPayment([]model.Payment) []model.Payment
	GetAll() []model.Payment
}

type paymentService struct {
	payments []model.Payment
}

func NewPaymentService() PaymentService {

	var payments []model.Payment

	data := data.LoadJsonFromDataFile(paymentsFilePath)
	err := json.Unmarshal(data, &payments)
	if err != nil {
		// TODO handle this properly
		panic(err)
	}

	return &paymentService{
		payments: payments,
	}
}

func (service *paymentService) AddPayment(newPayments []model.Payment) []model.Payment {
	service.payments = append(service.payments, newPayments...)

	// Save it to file
	data.AddJsonToDataFile(paymentsFilePath, service.payments)
	// TODO handle failure

	return newPayments
}

func (service *paymentService) GetAll() []model.Payment {
	return service.payments
}
