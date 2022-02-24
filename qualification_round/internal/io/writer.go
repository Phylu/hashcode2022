package io

import (
	"os"
	"simulation/internal/models"
	"strconv"
	"strings"
)

func WriteFile(filename string, output models.OutputDataset) error {
	var lines []string

	lines = append(lines, strconv.Itoa(output.ProjectCount))

	for _, assignment := range output.Assigments {
		lines = append(lines, assignment.Project.Name)

		var roles []string
		for _, role := range assignment.Roles {
			roles = append(roles, role.Name)
		}
		lines = append(lines, strings.Join(roles, " "))
	}

	return os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
}
