package simulation

import "simulation/internal/models"

type AddEverything struct{}

func (s AddEverything) Run(customers models.CustomerDataset) models.PerfectPizzaDataset {
	output := models.PerfectPizzaDataset{}

	for _, customer := range customers.Customers {
		output.Ingredients = append(output.Ingredients, customer.Likes...)
	}

	output.Ingredients = unique(output.Ingredients)

	return output
}
