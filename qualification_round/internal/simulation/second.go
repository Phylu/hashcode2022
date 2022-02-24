package simulation

import (
	"simulation/internal/models"
)

type FirstContributorWithSkills struct{}

func (s FirstContributorWithSkills) Run(dataset models.InputDataset) models.OutputDataset {
	output := models.OutputDataset{}
	var assignments []models.Assignment

	for _, project := range dataset.Projects {
		//fmt.Println(`Working on Project ` + project.Name)

		assignment := models.Assignment{
			Project: project,
		}
		for _, role := range project.Roles {

			// fmt.Println(`- Working on Role ` + role.Name + ` which needs SkillLevel ` + strconv.Itoa(role.Level))
		OUTER:
			for _, contributor := range dataset.Contributors {
				for _, assignmentRole := range assignment.Roles {
					if assignmentRole.Name == contributor.Name {
						continue OUTER
					}
				}

				contributorSkillLevel, ok := contributor.Skills[role.Name]
				//if ok {
				// fmt.Println(`-- Contributor ` + contributor.Name + ` has SkillLevel ` + strconv.Itoa(contributorSkillLevel))
				//}
				if ok && contributorSkillLevel >= role.Level {
					assignment.Roles = append(assignment.Roles, contributor)
					//fmt.Println(`--- Assigning Contributor ` + contributor.Name + ` for Role ` + role.Name)
					break
				}
			}

		}

		if len(assignment.Roles) == project.RolesCount {
			assignments = append(assignments, assignment)
		}

	}

	output.Assigments = assignments
	output.ProjectCount = len(assignments)

	return output
}
