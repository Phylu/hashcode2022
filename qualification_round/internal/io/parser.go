package io

import (
	"simulation/internal/models"
	"strconv"
	"strings"
)

func parseFile(data string) (models.InputDataset, error) {
	rows := strings.Split(data, "\n")
	dataset := models.InputDataset{}
	var err error

	contributorsCount := 0
	skillCount := 0
	rolesCount := 0

	var contributor *models.Contributor
	var project *models.Project

	for index, row := range rows {
		if index == 0 {
			counters := parseLine(row)

			dataset.ContributorsCount, err = strconv.Atoi(counters[0])
			if err != nil {
				return models.InputDataset{}, err
			}

			dataset.ProjectsCount, err = strconv.Atoi(counters[1])
			if err != nil {
				return models.InputDataset{}, err
			}

			continue
		}

		if row == "" {
			continue
		}

		if contributorsCount < dataset.ContributorsCount {
			if contributor == nil {
				skillCount = 0
				line := parseLine(row)
				contributorSkillCount, _ := strconv.Atoi(line[1])
				contributor = &models.Contributor{
					Name:        line[0],
					SkillsCount: contributorSkillCount,
					Skills:      map[string]int{},
				}
				if contributorSkillCount == 0 {
					dataset.Contributors = append(dataset.Contributors, *contributor)
					contributor = nil
					contributorsCount += 1
				}
			} else {
				line := parseLine(row)
				skillName := line[0]
				skillLevel, _ := strconv.Atoi(line[1])
				contributor.Skills[skillName] = skillLevel
				skillCount += 1

				if skillCount == contributor.SkillsCount {
					dataset.Contributors = append(dataset.Contributors, *contributor)
					contributor = nil
					contributorsCount += 1
				}
			}
			continue
		}

		// Reading In Projects now
		if contributorsCount >= dataset.ContributorsCount {
			if project == nil {
				rolesCount = 0

				line := parseLine(row)

				duration, _ := strconv.Atoi(line[1])
				score, _ := strconv.Atoi(line[2])
				bestBefore, _ := strconv.Atoi(line[3])
				projectRolesCount, _ := strconv.Atoi(line[4])

				project = &models.Project{
					Name:       line[0],
					Duration:   duration,
					Score:      score,
					BestBefore: bestBefore,
					RolesCount: projectRolesCount,
				}

				if projectRolesCount == 0 {
					dataset.Projects = append(dataset.Projects, *project)
					project = nil
				}

			} else {
				line := parseLine(row)
				skillName := line[0]
				level, _ := strconv.Atoi(line[1])

				role := models.Role{
					Name:  skillName,
					Level: level,
				}

				project.Roles = append(project.Roles, role)
				rolesCount += 1

				if rolesCount == project.RolesCount {
					dataset.Projects = append(dataset.Projects, *project)
					project = nil
					rolesCount += 1
				}
			}

			continue
		}

	}

	return dataset, nil
}

func parseLine(line string) []string {
	return strings.Split(line, " ")
}
