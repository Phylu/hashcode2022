package simulation

import (
	"simulation/internal/models"
)

type Simulation interface {
	Run(models.InputDataset) models.OutputDataset
}
