package model

type Transaction struct {
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	//Type         string `json:"transactionType"` //TODO make some sort of ENUM for debit/credit
	MainCategory string `json:"mainCategory"`
	SubCategory  string `json:"subCategory"`
	//TransactionSource Source
	Year  string `json:"year"`
	Month string `json:"month"`
}

type Source struct {
	Type string // Credit card, Venmo, Checking
	Name string
}
