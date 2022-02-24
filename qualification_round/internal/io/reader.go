package io

import (
	"os"
	"simulation/internal/models"
)

func ReadFile(filename string) (models.InputDataset, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return models.InputDataset{}, err
	}

	return parseFile(string(data))
}
