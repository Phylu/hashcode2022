package simulation

import "simulation/internal/customer"

type AddEverything struct{}

func (s AddEverything) Run(customers customer.CustomerDataset) customer.PerfectPizzaDataset {
	output := customer.PerfectPizzaDataset{}

	for _, customer := range customers.Customers {
		output.Ingredients = append(output.Ingredients, customer.Likes...)
	}

	output.Ingredients = unique(output.Ingredients)

	return output
}
