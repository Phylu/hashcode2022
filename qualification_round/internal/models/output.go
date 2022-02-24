package models

type Assignment struct {
	Project Project
	Roles   []Contributor
}

type OutputDataset struct {
	ProjectCount int
	Assigments   []Assignment
}
