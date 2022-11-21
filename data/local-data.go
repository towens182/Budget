package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/towens182/budget/model"
)

// TODO: load this from env
var dataPath = "data/transactions.json"

func LoadJsonFromLocalFileSystem() []byte {

	// Check if the file exists, if not create it
	data, err := os.ReadFile(dataPath)
	if errors.Is(err, os.ErrNotExist) {
		log.Printf("Creating data file " + dataPath)
		os.Create(dataPath)
		os.WriteFile(dataPath, []byte("[]"), 0644)
	}
	fmt.Println(data)
	return data
}

func AddJsonToDataFile(newTransactions []model.Transaction) {
	file, _ := json.MarshalIndent(newTransactions, "", " ")
	_ = os.WriteFile(dataPath, file, 0644)
	// TODO Handle error
}
