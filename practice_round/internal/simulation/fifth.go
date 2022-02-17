package simulation

import "simulation/internal/models"

type DoubleLikes struct{}

func (s DoubleLikes) Run(customers models.CustomerDataset) models.PerfectPizzaDataset {
	output := models.PerfectPizzaDataset{}
	likes := map[string]int{}
	dislikes := map[string]int{}

	for _, customer := range customers.Customers {
		for _, like := range customer.Likes {
			count, ok := likes[like]
			if !ok {
				count = 1
			} else {
				count++
			}

			likes[like] = count
		}

		for _, dislike := range customer.Dislikes {
			count, ok := dislikes[dislike]
			if !ok {
				count = 1
			} else {
				count++
			}

			dislikes[dislike] = count
		}
	}

	for like, likeCount := range likes {
		dislikeCount, ok := dislikes[like]

		if !ok {
			// no dislike, it is good to proceed
			output.Ingredients = append(output.Ingredients, like)
			continue
		}

		if likeCount >= 2*dislikeCount {
			// acceptable
			output.Ingredients = append(output.Ingredients, like)
			continue
		}
	}

	return output
}
