package simulation

import (
	"simulation/internal/models"
)

type Simulation interface {
	Run(models.CustomerDataset) models.PerfectPizzaDataset
}

func unique(input []string) []string {
	keys := make(map[string]bool)

	list := []string{}
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
