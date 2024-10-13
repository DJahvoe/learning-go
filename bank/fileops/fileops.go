package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// in Go, you export function to other package by naming the first alphabet into Uppercase
func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	// Create if file doesn't exist
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		valueText := fmt.Sprint(defaultValue)
		os.WriteFile(fileName, []byte(valueText), 0644)
	}

	if err != nil {

		return defaultValue, errors.New("failed to read file")
	}

	valueText := string(data)

	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
