package simulation

import "simulation/internal/customer"

type NoDislikes struct{}

func (s NoDislikes) Run(customers customer.CustomerDataset) customer.PerfectPizzaDataset {
	output := customer.PerfectPizzaDataset{}
	likes := []string{}
	dislikes := []string{}

	for _, customer := range customers.Customers {
		likes = append(likes, customer.Likes...)
		dislikes = append(dislikes, customer.Dislikes...)
	}

	likes = unique(likes)
	dislikes = unique(dislikes)

	for _, dislike := range dislikes {
		for likeIndex, like := range likes {
			if dislike == like {
				likes = remove(likes, likeIndex)
			}
		}
	}

	output.Ingredients = likes

	return output
}
