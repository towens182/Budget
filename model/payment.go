package model

type Payment struct {
	Name        string `json:"name"`
	Method      string `json:"method"`
	Institution string `json:"institution"`
}
