package usecases

import "clean-go/src/domain/models"

type Subtract interface {
	Calculate(data models.Subtract) float32
}
