package io

import (
	"os"
	"simulation/internal/models"
	"strconv"
	"strings"
)

func WriteFile(filename string, output models.PerfectPizzaDataset) error {
	line := strconv.Itoa(len(output.Ingredients)) + " " + strings.Join(output.Ingredients, " ")

	return os.WriteFile(filename, []byte(line), 0644)
}
