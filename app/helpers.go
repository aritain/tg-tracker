package app

import (
	"encoding/json"
	"os"
)

const FILEPATH = "/data/value.json"

func GetValue() (value Balance) {
	filepath := FILEPATH
	data, err := os.ReadFile(filepath)
	if err == nil {
		_ = json.Unmarshal(data, &value)
	}
	return
}

func WriteValue(value Balance) {
	filepath := FILEPATH
	file, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close()
	json.NewEncoder(file).Encode(value)
}
