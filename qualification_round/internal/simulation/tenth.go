package simulation

import (
	"simulation/internal/models"
	"sort"
)

type FreeContributors struct{}

func (s FreeContributors) Run(dataset models.InputDataset) models.OutputDataset {
	output := models.OutputDataset{}
	var assignments []models.Assignment
	contributorsProjects := map[*models.Contributor]int{}

	sortedProjects := dataset.Projects
	sort.SliceStable(sortedProjects, func(i, j int) bool {
		return sortedProjects[i].Score < sortedProjects[j].Score
	})

	for _, project := range sortedProjects {
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

			var bestContributor models.Contributor
			for _, c := range acceptableContributors {
				if _, ok := contributorsProjects[&c]; !ok {
					contributorsProjects[&c] = 1
					bestContributor = c
					break
				} else {
					if contributorsProjects[&c] < contributorsProjects[&bestContributor] {
						bestContributor = c
					}
				}
			}

			var contributorProjectsCount int
			if c, ok := contributorsProjects[&bestContributor]; ok {
				contributorProjectsCount = c
			} else {
				contributorProjectsCount = 0
			}
			contributorsProjects[&bestContributor] = contributorProjectsCount + 1
			assignment.Roles = append(assignment.Roles, bestContributor)
		}

		if len(assignment.Roles) == project.RolesCount {
			assignments = append(assignments, assignment)
		}

	}

	output.Assigments = assignments
	output.ProjectCount = len(assignments)

	return output
}
