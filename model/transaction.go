package model

type Transaction struct {
	Description   string  `json:"description"`
	Cost          int     `json:"cost"`
	Type          string  `json:"transactionType"`
	MainCategory  string  `json:"mainCategory"`
	SubCategory   string  `json:"subCategory"`
	PaymentSource Payment `json:"paymentSource"`
	Year          string  `json:"year"`
	Month         string  `json:"month"`
}
