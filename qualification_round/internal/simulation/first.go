package simulation

import "simulation/internal/models"

type FirstContributorsInOrder struct{}

func (s FirstContributorsInOrder) Run(dataset models.InputDataset) models.OutputDataset {
	output := models.OutputDataset{}
	var assignments []models.Assignment

	for _, project := range dataset.Projects {
		assignment := models.Assignment{
			Project: project,
		}
		for i := range project.Roles {
			assignment.Roles = append(assignment.Roles, dataset.Contributors[i])
		}

		assignments = append(assignments, assignment)

	}

	output.Assigments = assignments
	output.ProjectCount = len(assignments)

	return output
}
