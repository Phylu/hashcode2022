package simulation

import (
	"simulation/internal/customer"
)

type Simulation interface {
	Run(customer.CustomerDataset) customer.PerfectPizzaDataset
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
