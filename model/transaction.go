package model

type Transaction struct {
	Id                   string  `json:"id"`
	Name                 string  `json:"name"`
	Amount               float32 `json:"amount"`
	Type                 string  `json:"transactionType"`
	MainCategory         string  `json:"category"`
	SubCategory          string  `json:"subCategory"`
	FinancialInstitution string  `json:"financialInsitution"`
	//PaymentSource Payment `json:"paymentSource"`
	PostedDate string `json:"postedDate"`
	Year       string `json:"year"`
	Month      string `json:"month"`
}
