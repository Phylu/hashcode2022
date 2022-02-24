package simulation

import (
	"simulation/internal/models"
	"sort"
)

type LowerContributorWithSkills struct{}

func (s LowerContributorWithSkills) Run(dataset models.InputDataset) models.OutputDataset {
	output := models.OutputDataset{}
	var assignments []models.Assignment

	for _, project := range dataset.Projects {
		//fmt.Println(`Working on Project ` + project.Name)

		assignment := models.Assignment{
			Project: project,
		}
		for _, role := range project.Roles {
			acceptableContributors := []models.Contributor{}

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
					acceptableContributors = append(acceptableContributors, contributor)
					//fmt.Println(`--- Assigning Contributor ` + contributor.Name + ` for Role ` + role.Name)
				}
			}

			sort.SliceStable(acceptableContributors, func(i, j int) bool {
				return acceptableContributors[i].Skills[role.Name] < acceptableContributors[j].Skills[role.Name]
			})

			if len(acceptableContributors) == 0 {
				continue
			}

			accessibleContributor := acceptableContributors[0]

			assignment.Roles = append(assignment.Roles, accessibleContributor)
		}

		if len(assignment.Roles) == project.RolesCount {
			assignments = append(assignments, assignment)
		}

	}

	output.Assigments = assignments
	output.ProjectCount = len(assignments)

	return output
}
