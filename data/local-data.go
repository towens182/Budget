package data

import (
	"errors"
	"fmt"
	"log"
	"os"
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
