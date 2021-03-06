package simulation

import "simulation/internal/models"

type NoDislikes struct{}

func (s NoDislikes) Run(customers models.CustomerDataset) models.PerfectPizzaDataset {
	output := models.PerfectPizzaDataset{}
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
