package models

type Contributor struct {
	Name        string
	SkillsCount int
	Skills      map[string]int
}

type Role struct {
	Name  string
	Level int
}

type Project struct {
	Name       string
	Roles      []Role
	Score      int
	Duration   int
	BestBefore int
	RolesCount int
}

type InputDataset struct {
	ContributorsCount int
	ProjectsCount     int
	Contributors      []Contributor
	Projects          []Project
}
