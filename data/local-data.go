package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func LoadJsonFromDataFile(filePath string) []byte {

	// Check if the file exists, if not create it
	data, err := os.ReadFile(filePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Printf("Creating data file " + filePath)
		os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), 0644)
	}
	fmt.Println(data)
	return data
}

func AddJsonToDataFile(filePath string, T any) {
	file, _ := json.MarshalIndent(T, "", " ")
	_ = os.WriteFile(filePath, file, 0644)
	// TODO Handle error
}
