package io

import (
	"os"
	"simulation/internal/models"
)

func ReadFile(filename string) (models.CustomerDataset, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return models.CustomerDataset{}, err
	}

	return parseFile(string(data))
}
